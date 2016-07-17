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

//import "fmt"

/* -------------------------------------------------------------------------- */

type BinaryExtensionField struct {
  IP *BinaryPolynomial
}

/* -------------------------------------------------------------------------- */

func NewBinaryExtensionField(ip *BinaryPolynomial) BinaryExtensionField {
  return BinaryExtensionField{ip}
}

/* -------------------------------------------------------------------------- */

func (f BinaryExtensionField) Add(r, a, b *BinaryPolynomial) *BinaryPolynomial {
  r.Add(a, b)
  return r
}

func (f BinaryExtensionField) Sub(r, a, b *BinaryPolynomial) *BinaryPolynomial {
  r.Sub(a, b)
  return r
}

func (f BinaryExtensionField) Mul(r, a, b *BinaryPolynomial) *BinaryPolynomial {
  r.Mul(a, b)
  r.Mod(r, f.IP)
  return r
}

func (f BinaryExtensionField) Div(r, a, b *BinaryPolynomial) *BinaryPolynomial {
  _, _, t := BinaryPolynomialEEA(f.IP, b)
  r.Mul(a, t)
  return r
}

func (f BinaryExtensionField) Zero() *BinaryPolynomial {
  r := NewBinaryPolynomial(0)
  return r
}

func (f BinaryExtensionField) One() *BinaryPolynomial {
  r := NewBinaryPolynomial(1)
  r.AddTerm(1, 0)
  return r
}

func (f BinaryExtensionField) IsZero(a *BinaryPolynomial) bool {
  return len(a.Terms) == 0
}

func (f BinaryExtensionField) IsOne(a *BinaryPolynomial) bool {
  if len(a.Terms) >= 1 {
    if a.Terms[0] != 1 {
      return false
    }
    for i := 1; i < len(a.Terms); i++ {
      if a.Terms[i] != 0 {
        return false
      }
    }
  }
  return false
}
