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
import "testing"

/* -------------------------------------------------------------------------- */

func TestBinaryExtensionField1(t *testing.T) {

  pf := NewPrimeField(2)

  // irreducible polynomial
  p1 := NewPolynomial(pf)
  p1.AddTerm(1, 8)
  p1.AddTerm(1, 4)
  p1.AddTerm(1, 3)
  p1.AddTerm(1, 1)
  p1.AddTerm(1, 0)

  p2 := NewBinaryPolynomial(1)
  p2.AddTerm(1, 8)
  p2.AddTerm(1, 4)
  p2.AddTerm(1, 3)
  p2.AddTerm(1, 1)
  p2.AddTerm(1, 0)

  // fields
  f1 := NewExtensionField(p1)
  f2 := NewBinaryExtensionField(p2)

  // field elements
  a1 := NewPolynomial(pf)
  b1 := NewPolynomial(pf)
  a1.AddTerm(1, 0)

  a2 := NewBinaryPolynomial(1)
  b2 := NewBinaryPolynomial(1)
  r2 := NewBinaryPolynomial(1)
  a2.AddTerm(1, 0)

  for i := 0; i <= 0xFF; i++ {
    b2.Terms[0] = byte(i)
    b1.SetZero()
    for _, e := range b2.Exponents() {
      b1.AddTerm(1, e)
    }

    r1 := f1.Div(a1, b1)
    f2.Div(r2, a2, b2)

    exponents1 := r1.Exponents()
    exponents2 := r2.Exponents()

    if len(exponents1) != len(exponents2) {
      t.Error("binary field extension test failed")
    } else {
      for i := 0; i < len(exponents1); i++ {
        if exponents1[i] != exponents2[i] {
          t.Error("binary field extension test failed")
        }
      }
    }
  }
}
