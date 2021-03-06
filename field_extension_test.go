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

func TestExtensionField1(t *testing.T) {

  pf := NewPrimeField(2)

  // irreducible polynomial
  p := NewPolynomial(pf)
  p.AddTerm(1, 4)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)

  f := NewExtensionField(p)

  a := NewPolynomial(pf)
  a.AddTerm(1, 3)
  a.AddTerm(1, 2)
  a.AddTerm(1, 0)
  b := NewPolynomial(pf)
  b.AddTerm(1, 2)
  b.AddTerm(1, 1)

  r1 := NewPolynomial(pf)
  r1.AddTerm(1, 3)

  r2 := f.Mul(a, b)

  if !r1.Equals(r2) {
    t.Error("finite field test failed")
  }
}

func TestExtensionField2(t *testing.T) {

  pf := NewPrimeField(2)

  // irreducible polynomial
  p := NewPolynomial(pf)
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)

  f := NewExtensionField(p)

  a := NewPolynomial(pf)
  a.AddTerm(1, 0)
  b := NewPolynomial(pf)
  b.AddTerm(1, 7)
  b.AddTerm(1, 6)
  b.AddTerm(1, 1)

  r1 := NewPolynomial(pf)
  r1.AddTerm(1, 5)
  r1.AddTerm(1, 3)
  r1.AddTerm(1, 2)
  r1.AddTerm(1, 1)
  r1.AddTerm(1, 0)

  r2 := f.Div(a, b)

  if !r1.Equals(r2) {
    t.Error("finite field test failed")
  }
}

func TestExtensionField3(t *testing.T) {

  pf := NewPrimeField(2)

  // irreducible polynomial
  p := NewPolynomial(pf)
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)

  f := NewExtensionField(p)

  zero := f.Zero()
  one  := f.One()

  if !f.IsZero(zero) {
    t.Error("finite field test failed")
  }
  if !f.IsOne(one) {
    t.Error("finite field test failed")
  }

}
