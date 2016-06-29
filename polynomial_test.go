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
import "math"
import "testing"

/* -------------------------------------------------------------------------- */

func TestPolynomial1(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm(20,30)
  p.AddTerm(2,3)
  q := p.Clone()
  q.AddTerm(1024, 256)

  r := NewPolynomial()
  r.Mul(p,q)
  // 20480x^286 + 2048x^259 + 400x^60 + 80x^33 + 4x^6
  s := NewPolynomial()
  s.AddTerm(20480, 286)
  s.AddTerm( 2048, 259)
  s.AddTerm(  400,  60)
  s.AddTerm(   80,  33)
  s.AddTerm(    4,   6)

  if !s.Equals(r) {
    t.Error("polynomial test failed")
  }

}

func TestPolynomial2(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm(20, 30)
  p.AddTerm( 2  ,3)
  q := p.Clone()

  r1 := NewPolynomial()
  r1.Add(p, q)

  r2 := NewPolynomial()
  r2.Add(p, p)

  r3 := p.Clone()
  r3.Add(r3, p)

  r4 := p.Clone()
  r4.Add(p, r4)

  r5 := p.Clone()
  r5.Add(r5, r5)

  if !r1.Equals(r2) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r3) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r4) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r5) {
    t.Error("polynomial test failed")
  }

}

func TestPolynomial3(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm(20, 30)
  p.AddTerm( 2  ,3)
  q := p.Clone()
  p.AddTerm(10, 35)

  r1 := NewPolynomial()
  r1.Sub(p, q)

  r2 := p.Clone()
  r2.Sub(r2, q)

  r3 := q.Clone()
  r3.Sub(p, r3)

  r4 := p.Clone()
  r4.Sub(r4, r4)

  if !r1.Equals(r2) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r3) {
    t.Error("polynomial test failed")
  }
  if !r4.Equals(NewPolynomial()) {
    t.Error("polynomial test failed")
  }

}

func TestPolynomial4(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm(20, 30)
  p.AddTerm( 2  ,3)
  q := p.Clone()

  r1 := NewPolynomial()
  r1.Mul(p, q)

  r2 := NewPolynomial()
  r2.Mul(p, p)

  r3 := p.Clone()
  r3.Mul(r3, p)

  r4 := p.Clone()
  r4.Mul(p, r4)

  r5 := p.Clone()
  r5.Mul(r5, r5)

  if !r1.Equals(r2) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r3) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r4) {
    t.Error("polynomial test failed")
  }
  if !r1.Equals(r5) {
    t.Error("polynomial test failed")
  }

}

func TestPolynomial5(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm( 1, 3)
  p.AddTerm(-2, 2)
  p.AddTerm(-4, 0)
  q := NewPolynomial()
  q.AddTerm( 1, 1)
  q.AddTerm(-3, 0)

  r1 := NewPolynomial()
  r1.AddTerm( 1, 2)
  r1.AddTerm( 1, 1)
  r1.AddTerm( 3, 0)

  r2 := NewPolynomial()
  r2.AddTerm( 5, 0)

  r3 := NewPolynomial()
  r4 := r3.Div(p, q)

  if !r3.Equals(r1) {
    t.Error("polynomial test failed")
  }
  if !r4.Equals(r2) {
    t.Error("polynomial test failed")
  }

}

func TestPolynomial6(t *testing.T) {

  p := NewPolynomial()
  p.AddTerm( 1, 8)
  p.AddTerm( 1, 4)
  p.AddTerm( 1, 3)
  p.AddTerm( 1, 1)
  p.AddTerm( 1, 0)
  a := NewPolynomial()
  a.AddTerm( 1, 7)
  a.AddTerm( 1, 6)
  a.AddTerm( 1, 1)

  r, si, ti := PolynomialEEA(a, p)

  if c, e := r.Lead(); math.Abs(c - 1.0) > 1e-12 || e != 0 {
    t.Error("polynomial test failed")
  }
  if c, e := si.Lead(); math.Abs(c - -0.796748) > 1e-4 || e != 7 {
    t.Error("polynomial test failed")
  }
  if c, e := ti.Lead(); math.Abs(c - 0.796748) > 1e-4 || e != 6 {
    t.Error("polynomial test failed")
  }
}
