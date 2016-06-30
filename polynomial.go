/* Copyright (C) 2016 Philipp Benner
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package gocrypt

/* -------------------------------------------------------------------------- */

import "fmt"
import "bufio"
import "bytes"
import "sort"

/* -------------------------------------------------------------------------- */

type Polynomial struct {
  Terms map[int]FieldElement
  Field Field
}

/* -------------------------------------------------------------------------- */

func NewPolynomial(field Field) *Polynomial {
  terms := make(map[int]FieldElement)
  return &Polynomial{terms, field}
}

func NewRealPolynomial() *Polynomial {
  terms := make(map[int]FieldElement)
  return &Polynomial{terms, NewRealField()}
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) Clone() *Polynomial {
  s := NewPolynomial(r.Field)
  for k, v := range r.Terms {
    s.Terms[k] = v
  }
  s.Field = r.Field
  return s
}

func (r *Polynomial) CloneEmpty() *Polynomial {
  return NewPolynomial(r.Field)
}

func (r *Polynomial) AddTerm(c FieldElement, e int) {
  if v, ok := r.Terms[e]; ok {
    r.Terms[e] = r.Field.FieldAdd(v, c)
  } else {
    r.Terms[e] = c
  }
  r.Clean()
}

func (r *Polynomial) Clear() {
  r.Terms = make(map[int]FieldElement)
}

func (r *Polynomial) Clean() {
  for k, v := range r.Terms {
    if r.Field.FieldIsZero(v) {
      delete(r.Terms, k)
    }
  }
}

func (r *Polynomial) Exponents() []int {
  var keys []int
  for k := range r.Terms {
    keys = append(keys, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(keys)))
  return keys
}

func (r *Polynomial) Degree() int {
  if len(r.Terms) == 0 {
    return 0
  }
  return r.Exponents()[0]
}

func (r *Polynomial) Lead() (FieldElement, int) {
  if len(r.Terms) == 0 {
    return 0, 0
  }
  k := r.Exponents()[0]

  return r.Terms[k], k
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) Equals(a *Polynomial) bool {
  if r == a {
    return true
  } else {
    if len(r.Terms) != len(a.Terms) {
      return false
    } else {
      for k, v := range r.Terms {
        if a.Terms[k] != v {
          return false
        }
      }
      return true
    }
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) neg() {
  for k, v := range r.Terms {
    r.Terms[k] = r.Field.FieldNeg(v)
  }
}

func (r *Polynomial) Neg(a *Polynomial) {
  if r == a {
    r.neg()
  } else {
    r.Clear()
    r.add(a)
    r.neg()
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) add(a *Polynomial) {
  for k, v := range a.Terms {
    if w, ok := r.Terms[k]; ok {
      r.Terms[k] = r.Field.FieldAdd(w, v)
    } else {
      r.Terms[k] = v
    }
  }
  r.Clean()
}

func (r *Polynomial) Add(a, b *Polynomial) {
  if r == a {
    r.add(b)
  } else if r == b {
    r.add(a)
  } else {
    r.Clear()
    r.add(a)
    r.add(b)
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) sub(a *Polynomial) {
  for k, v := range a.Terms {
    if w, ok := r.Terms[k]; ok {
      r.Terms[k] = r.Field.FieldSub(w, v)
    } else {
      r.Terms[k] = r.Field.FieldNeg(v)
    }
  }
  r.Clean()
}

func (r *Polynomial) Sub(a, b *Polynomial) {
  if r == a {
    r.sub(b)
  } else if r == b {
    r.neg()
    r.add(a)
  } else {
    r.Clear()
    r.add(a)
    r.sub(b)
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) mul(a *Polynomial) {
  t := NewPolynomial(r.Field)
  for k1, v1 := range r.Terms {
    for k2, v2 := range a.Terms {
      k := k1+k2
      v := r.Field.FieldMul(v1, v2)
      t.AddTerm(v, k)
    }
  }
  r.Terms = t.Terms
  r.Clean()
}

func (r *Polynomial) Mul(a, b *Polynomial) {
  if r == a || r == b {
    r.mul(b)
  } else {
    r.Clear()
    r.add(a)
    r.mul(b)
  }
}

/* -------------------------------------------------------------------------- */

func (r1 *Polynomial) div(a, b, r2 *Polynomial) {
  z := NewPolynomial(r1.Field)
  t := NewPolynomial(r1.Field)
  q := NewPolynomial(r1.Field)
  r := a.Clone()
  if b.Equals(z) {
    panic("Div(): division by zero")
  }
  c2, e2 := b.Lead()
  for !r.Equals(z) && r.Degree() >= b.Degree() {
    c1, e1 := r.Lead()
    t.Clear()
    t.AddTerm(r1.Field.FieldDiv(c1, c2), e1-e2)
    q.AddTerm(r1.Field.FieldDiv(c1, c2), e1-e2)
    t.Mul(t, b)
    r.Sub(r, t)
  }
  r1.Terms = q.Terms
  // save remainder if s is given
  if r2 != nil {
    r2.Terms = r.Terms
  }
}

func (r *Polynomial) Div(a, b *Polynomial) {
  r.div(a, b, nil)
}

func (r *Polynomial) Mod(a, b *Polynomial) {
  r.div(a, b, r)
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) String() string {

  var buffer bytes.Buffer
  writer := bufio.NewWriter(&buffer)

  keys := r.Exponents()

  if len(keys) == 0 {
    fmt.Fprintf(writer, "0.0")
  }
  for i, k := range keys {
    v := 0.0
    switch x := r.Terms[k].(type) {
    case float64: v = float64(x)
    case int    : v = float64(x)
    default: panic("String(): field not supported")
    }
    if i != 0 {
      if v >= 0.0 {
        fmt.Fprintf(writer, " + ")
      } else {
        fmt.Fprintf(writer, " - ")
        v = -v
      }
    }
    if k == 0 {
      fmt.Fprintf(writer, "%f", v)
    } else {
      if v == 1.0 {
        fmt.Fprintf(writer, "x^%d", k)
      } else {
        fmt.Fprintf(writer, "%fx^%d", v, k)
      }
    }
  }
  writer.Flush()

  return buffer.String()
}
