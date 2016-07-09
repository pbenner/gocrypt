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
import "math/big"
//import "math/rand"

/* -------------------------------------------------------------------------- */

type ECPoint struct {
  x, y *big.Int
}

func NewECPoint(x, y *big.Int) ECPoint {
  return ECPoint{x, y}
}


func NullECPoint() ECPoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  return ECPoint{x, y}
}

/* -------------------------------------------------------------------------- */

type EllipticCurve struct {
  a, b *big.Int
  f    BigPrimeField
}

/* -------------------------------------------------------------------------- */

func NewEllipticCurve(a, b, p *big.Int) EllipticCurve {
  f := NewBigPrimeField(p)
  return EllipticCurve{a, b, f}
}

/* -------------------------------------------------------------------------- */

func (ec EllipticCurve) Add(p, q ECPoint) ECPoint {
  r := NullECPoint()
  s := big.NewInt(0)
  t := big.NewInt(0)
  f := ec.f
  if p.x.Cmp(q.x) == 0 {
    // p == q
    s = f.Mul(p.x, p.x)
    s = f.Mul(s, big.NewInt(3))
    s = f.Sub(s, ec.a)
    s = f.Div(s, p.y)
    s = f.Div(s, big.NewInt(2))
  } else {
    // p != q
    s = f.Sub(p.y, q.y)
    t = f.Sub(p.x, q.x)
    s = f.Div(s, t)
  }
  r.x = f.Mul(s, s)
  r.x = f.Sub(r.x, p.x)
  r.x = f.Sub(r.x, q.x)
  r.y = f.Sub(r.x, p.x)
  r.y = f.Mul(r.y, s)
  r.y = f.Add(r.y, p.y)
  r.y = f.Neg(r.y)
  return r
}