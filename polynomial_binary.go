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

type BinaryPolynomial struct {
  Terms []byte
}

/* -------------------------------------------------------------------------- */

func NewBinaryPolynomial(n int) *BinaryPolynomial {
  r := BinaryPolynomial{}
  r.Terms = make([]byte, n, 2*n)
  r.SetZero()
  return &r
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Clone() *BinaryPolynomial {
  s := NewBinaryPolynomial(len(r.Terms))
  s.Set(r)
  return s
}

func (r *BinaryPolynomial) AddTerm(c, e int) {
  r.Terms[e/8] ^= byte(c%2) << byte(e%8)
}

func (r *BinaryPolynomial) Degree() int {

  for i := len(r.Terms); i > 0; i-- {
    for j := 8; j > 0; j-- {
      if r.Terms[i-1] & (1 << byte(j-1)) != 0 {
        return 8*(i-1)+(j-1)
      }
    }
  }
  return 0
}

func (r *BinaryPolynomial) Exponents() []int {
  exponents := []int{}

  for i := len(r.Terms); i > 0; i-- {
    for j := 8; j > 0; j-- {
      if r.Terms[i-1] & (1 << byte(j-1)) != 0 {
        exponents = append(exponents, 8*(i-1)+(j-1))
      }
    }
  }
  sort.Sort(sort.Reverse(sort.IntSlice(exponents)))
  return exponents
}

func (r *BinaryPolynomial) Lead() (int, int) {
  if len(r.Terms) == 0 {
    return 0, 0
  }
  k := r.Degree()

  return 1, k
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Realloc(n int) {
  if cap(r.Terms) < n {
    r.Terms = make([]byte, n, 2*n)
  } else {
    r.Terms = r.Terms[0:n]
  }
}

func (r *BinaryPolynomial) Resize(n int) {
  if cap(r.Terms) < n {
    t := make([]byte, n, 2*n)
    copy(t, r.Terms)
    r.Terms = t
  } else {
    r.Terms = r.Terms[0:n]
  }
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) SetZero() {
  for i := 0; i < len(r.Terms); i++ {
    r.Terms[i] = 0
  }
}

func (r *BinaryPolynomial) Set(a *BinaryPolynomial) {
  if len(r.Terms) != len(a.Terms) {
    r.Terms = make([]byte, len(a.Terms), 2*len(a.Terms))
  }
  copy(r.Terms, a.Terms)
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Equals(a *BinaryPolynomial) bool {
  n := len(r.Terms)
  m := len(a.Terms)
  if n == m {
    for i := 0; i < n; i++ {
      if r.Terms[i] != a.Terms[i] {
        return false
      }
    }
  } else {
    if n < m {
      r, a = a, r
      n, m = m, n
    }
    for i := 0; i < m; i++ {
      if r.Terms[i] != a.Terms[i] {
        return false
      }
    }
    for i := m; i < n; i++ {
      if r.Terms[i] != 0 {
        return false
      }
    }
  }
  return true
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Neg(a *BinaryPolynomial) {
  // nothing to do
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Add(a, b *BinaryPolynomial) {
  n := len(a.Terms)
  m := len(b.Terms)
  if n == m {
    if len(r.Terms) != n {
      r.Terms = make([]byte, n, 2*n)
    }
    for i := 0; i < n; i++ {
      r.Terms[i] = a.Terms[i] ^ b.Terms[i]
    }
  } else {
    if n < m {
      a, b = b, a
      n, m = m, n
    }
    if len(r.Terms) != n {
      r.Resize(n)
    }
    for i := 0; i < m; i++ {
      r.Terms[i] = a.Terms[i] ^ b.Terms[i]
    }
    for i := m; i < n; i++ {
      r.Terms[i] = a.Terms[i]
    }
  }
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Sub(a, b *BinaryPolynomial) {
  r.Add(a, b)
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Mul(a, b *BinaryPolynomial) {
  for i := len(a.Terms); i > 0; i-- {
    t := byte(0)
    for j := 8; j > 0; j-- {
      k := 8*(i-1) + (j-1)
      s := 0
      for k1 := 0; k1 <= k; k1++ {
        k2 := k-k1
        if a.Terms[k1/8] & (1 << byte(k1%8)) != 0 && a.Terms[k2/8] & (1 << byte(k2%8)) != 0 {
          s++
        }
      }
      t ^= byte(s%2) << byte(k%8)
    }
    r.Terms[i-1] = t
  }
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) Div(a, b *BinaryPolynomial) {
}

func (r *BinaryPolynomial) Mod(a, b *BinaryPolynomial) {
}

/* -------------------------------------------------------------------------- */

func (r *BinaryPolynomial) String() string {

  var buffer bytes.Buffer
  writer := bufio.NewWriter(&buffer)

  first := true

  for i := len(r.Terms); i > 0; i-- {
    for j := 8; j > 0; j-- {
      if r.Terms[i-1] & (1 << byte(j-1)) != 0 {
        if first {
          first = false
        } else {
          fmt.Fprintf(writer, " + ")
        }
        fmt.Fprintf(writer, "x^%d", 8*(i-1)+(j-1))
      }
    }
  }
  writer.Flush()

  return buffer.String()
}
