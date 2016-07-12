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
  Terms  map[int]int
  Field  FieldInt
}

/* -------------------------------------------------------------------------- */

func NewPolynomial(field FieldInt) *Polynomial {
  r := Polynomial{}
  r.SetZero()
  r.Field  = field
  return &r
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) Clone() *Polynomial {
  s := NewPolynomial(r.Field)
  s.Set(r)
  return s
}

func (r *Polynomial) AddTerm(c, e int) {
  if r.Field.IsZero(c) {
    return
  }
  if v, ok := r.Terms[e]; ok {
    r.Terms[e] = r.Field.Add(v, c)
    // if term has zero coefficient, delete it
    if r.Field.IsZero(r.Terms[e]) {
      delete(r.Terms, e)
    }
  } else {
    r.Terms[e] = c
  }
}

func (r *Polynomial) Degree() int {
  deg := 0
  for k, _ := range r.Terms {
    if k > deg {
      deg = k
    }
  }
  return deg
}

func (r *Polynomial) Exponents() []int {
  var keys []int
  for k := range r.Terms {
    keys = append(keys, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(keys)))
  return keys
}

func (r *Polynomial) Lead() (int, int) {
  if len(r.Terms) == 0 {
    return 0, 0
  }
  k := r.Degree()

  return r.Terms[k], k
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) SetZero() {
  r.Terms  = make(map[int]int)
}

func (r *Polynomial) Set(a *Polynomial) {
  r.SetZero()
  for k, v := range a.Terms {
    r.Terms[k] = v
  }
  r.Field  = a.Field
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
    r.Terms[k] = r.Field.Neg(v)
  }
}

func (r *Polynomial) Neg(a *Polynomial) {
  if r != a {
    r.Set(a)
  }
  r.neg()
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) add(a *Polynomial) {
  for k, v := range a.Terms {
    r.AddTerm(v, k)
  }
}

func (r *Polynomial) Add(a, b *Polynomial) {
  if r == a {
    r.add(b)
  } else if r == b {
    r.add(a)
  } else {
    r.SetZero()
    r.add(a)
    r.add(b)
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) sub(a *Polynomial) {
  for k, v := range a.Terms {
    r.AddTerm(r.Field.Neg(v), k)
  }
}

func (r *Polynomial) Sub(a, b *Polynomial) {
  if r == a {
    r.sub(b)
  } else if r == b {
    r.neg()
    r.add(a)
  } else {
    r.SetZero()
    r.add(a)
    r.sub(b)
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) mul(a, b *Polynomial) {
  for k1, v1 := range a.Terms {
    for k2, v2 := range b.Terms {
      k := k1+k2
      v := r.Field.Mul(v1, v2)
      r.AddTerm(v, k)
    }
  }
}

func (r *Polynomial) Mul(a, b *Polynomial) {
  if r == a || r == b {
    t := NewPolynomial(r.Field)
    t.mul(a, b)
    r.Set(t)
  } else {
    r.SetZero()
    r.mul(a, b)
  }
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) div(a, b *Polynomial, remainder bool) {
  z := NewPolynomial(r.Field)
  t := NewPolynomial(r.Field)
  q := NewPolynomial(r.Field)
  s := a.Clone()
  if b.Equals(z) {
    panic("Div(): division by zero")
  }
  c2, e2 := b.Lead()
  for !s.Equals(z) && s.Degree() >= b.Degree() {
    c1, e1 := s.Lead()
    // new coefficient
    c := s.Field.Div(c1, c2)
    // new exponent
    e := e1-e2
    t.SetZero()
    t.AddTerm(c, e)
    q.AddTerm(c, e)
    t.Mul(t, b)
    s.Sub(s, t)
  }
  if remainder {
    r.Set(s)
  } else {
    r.Set(q)
  }
}

func (r *Polynomial) Div(a, b *Polynomial) {
  r.div(a, b, false)
}

func (r *Polynomial) Mod(a, b *Polynomial) {
  r.div(a, b, true)
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) String() string {

  var buffer bytes.Buffer
  writer := bufio.NewWriter(&buffer)

  keys := r.Exponents()

  if len(keys) == 0 {
    fmt.Fprintf(writer, "0")
  }
  for i, k := range keys {
    v := r.Terms[k]
    if i != 0 {
      if v >= 0.0 {
        fmt.Fprintf(writer, " + ")
      } else {
        fmt.Fprintf(writer, " - ")
        v = -v
      }
    }
    if k == 0 {
      fmt.Fprintf(writer, "%d", v)
    } else {
      if v == 1.0 {
        fmt.Fprintf(writer, "x^%d", k)
      } else {
        fmt.Fprintf(writer, "%dx^%d", v, k)
      }
    }
  }
  writer.Flush()

  return buffer.String()
}
