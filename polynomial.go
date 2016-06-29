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
  Terms map[int]float64
}

/* -------------------------------------------------------------------------- */

func NewPolynomial() *Polynomial {
  terms := make(map[int]float64)
  return &Polynomial{terms}
}

/* -------------------------------------------------------------------------- */

func (r *Polynomial) Clone() *Polynomial {
  s := NewPolynomial()
  for k, v := range r.Terms {
    s.Terms[k] = v
  }
  return s
}

func (r *Polynomial) AddTerm(c float64, e int) {
  r.Terms[e] += c
  r.Clean()
}

func (r *Polynomial) Clear() {
  r.Terms = make(map[int]float64)
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

func (r *Polynomial) Lead() (float64, int) {
  if len(r.Terms) == 0 {
    return 0, 0
  }
  k := r.Exponents()[0]

  return r.Terms[k], k
}

func (r *Polynomial) Clean() {
  for k, v := range r.Terms {
    if v == 0 {
      delete(r.Terms, k)
    }
  }
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
    r.Terms[k] = -v
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
    r.Terms[k] += v
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
    r.Terms[k] -= v
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
  t := NewPolynomial()
  for k1, v1 := range r.Terms {
    for k2, v2 := range a.Terms {
      k := k1+k2
      v := v1*v2
      t.Terms[k] += v
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
  z := NewPolynomial()
  t := NewPolynomial()
  q := NewPolynomial()
  r := a.Clone()
  if b.Equals(z) {
    panic("Div(): division by zero")
  }
  c2, e2 := b.Lead()
  for !r.Equals(z) && r.Degree() >= b.Degree() {
    c1, e1 := r.Lead()
    t.Clear()
    t.AddTerm(c1/c2, e1-e2)
    q.AddTerm(c1/c2, e1-e2)
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
      fmt.Fprintf(writer, "%f", v)
    } else {
      fmt.Fprintf(writer, "%fx^%d", v, k)
    }
  }
  writer.Flush()

  return buffer.String()
}
