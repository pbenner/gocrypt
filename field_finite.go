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
import "math"
import "sort"

/* -------------------------------------------------------------------------- */

type FiniteField struct {
  P, N  int
  Terms map[int]float64
}

/* -------------------------------------------------------------------------- */

func NewFiniteField() *FiniteField {
  terms := make(map[int]float64)
  return &FiniteField{0, 0, terms}
}

/* -------------------------------------------------------------------------- */

func (r *FiniteField) Clone() *FiniteField {
  s := NewFiniteField()
  for k, v := range r.Terms {
    s.Terms[k] = v
  }
  return s
}

func (r *FiniteField) AddTerm(c float64, e int) {
  r.Terms[e] += c
  r.Clean()
}

func (r *FiniteField) Clear() {
  r.Terms = make(map[int]float64)
}

func (r *FiniteField) Exponents() []int {
  var keys []int
  for k := range r.Terms {
    keys = append(keys, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(keys)))
  return keys
}

func (r *FiniteField) Degree() int {
  if len(r.Terms) == 0 {
    return 0
  }
  return r.Exponents()[0]
}

func (r *FiniteField) Lead() (float64, int) {
  if len(r.Terms) == 0 {
    return 0, 0
  }
  k := r.Exponents()[0]

  return r.Terms[k], k
}

func (r *FiniteField) Clean() {
  for k, _ := range r.Terms {
    r.Terms[k] = math.Abs(float64(int(r.Terms[k]) % 2))
  }
  for k, v := range r.Terms {
    if v == 0 {
      delete(r.Terms, k)
    }
  }
}

/* -------------------------------------------------------------------------- */

func (r *FiniteField) Equals(a *FiniteField) bool {
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

func (r *FiniteField) neg() {
  for k, v := range r.Terms {
    r.Terms[k] = -v
  }
}

func (r *FiniteField) Neg(a *FiniteField) {
  if r == a {
    r.neg()
  } else {
    r.Clear()
    r.add(a)
    r.neg()
  }
}

/* -------------------------------------------------------------------------- */

func (r *FiniteField) add(a *FiniteField) {
  for k, v := range a.Terms {
    r.Terms[k] += v
  }
  r.Clean()
}

func (r *FiniteField) Add(a, b *FiniteField) {
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

func (r *FiniteField) sub(a *FiniteField) {
  for k, v := range a.Terms {
    r.Terms[k] -= v
  }
  r.Clean()
}

func (r *FiniteField) Sub(a, b *FiniteField) {
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

func (r *FiniteField) mul(a *FiniteField) {
  t := NewFiniteField()
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

func (r *FiniteField) Mul(a, b *FiniteField) {
  if r == a || r == b {
    r.mul(b)
  } else {
    r.Clear()
    r.add(a)
    r.mul(b)
  }
}

/* -------------------------------------------------------------------------- */

func (s *FiniteField) Div(a, b *FiniteField) *FiniteField {
  z := NewFiniteField()
  t := NewFiniteField()
  q := NewFiniteField()
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
  s.Terms = q.Terms

  return r
}

/* -------------------------------------------------------------------------- */

func (r *FiniteField) Mod(a, b *FiniteField) {
  s := r.Div(a, b)
  r.Terms = s.Terms
}

/* -------------------------------------------------------------------------- */

func (ri *FiniteField) EEA(rj *FiniteField) (*FiniteField, *FiniteField, *FiniteField) {

  z0 := NewFiniteField()
  si := NewFiniteField()
  si.AddTerm(1, 0)
  ti := NewFiniteField()
  // j = i+1
  sj := NewFiniteField()
  tj := NewFiniteField()
  tj.AddTerm(1, 0)
  qj := NewFiniteField()
  // k = j+1
  sk := NewFiniteField()
  tk := NewFiniteField()
  rk := NewFiniteField()

  for !rj.Equals(z0) {
    // r_i = r_i-2 mod r_i-1
    rk.Mod(ri, rj)
    // q_i-1 = (r_i-2 - r_i)/r_i-1
    qj.Sub(ri, rk)
    qj.Div(qj, rj)
    // s_i = s_i-2 - q_i-1*s_i-1  
    sk.Mul(qj, sj)
    sk.Sub(si, sk)
    // t_i = t_i-2 - q_i-1*t_i-1
    tk.Mul(qj, tj)
    tk.Sub(ti, tk)

    si, sj, sk = sj, sk, si
    ti, tj, tk = tj, tk, ti
    ri, rj, rk = rj, rk, ri
  }
  if _, e := ri.Lead(); e == 0 {
    si.Div(si, ri)
    ti.Div(ti, ri)
    ri.Div(ri, ri)
  }
  // gcd(r0, r1) = ri = s r_0 + t r_1
  return ri, si, ti
}

/* -------------------------------------------------------------------------- */

func (r *FiniteField) String() string {

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
