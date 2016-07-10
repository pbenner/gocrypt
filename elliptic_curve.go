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
import "math/big"

/* -------------------------------------------------------------------------- */

type ECPoint struct {
  x, y *big.Int
}

func NewECPoint(x, y *big.Int) ECPoint {
  return ECPoint{x, y}
}

func NilECPoint() ECPoint {
  return ECPoint{nil, nil}
}

func NullECPoint() ECPoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  return ECPoint{x, y}
}

func (p ECPoint) String() string {
  return fmt.Sprintf("(%v,%v)", p.x, p.y)
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

  if ec.IsZero(p) && ec.IsZero(q) {
    return ec.Zero()
  }
  if ec.IsZero(p) {
    r.x = q.x
    r.y = q.y
    return r
  }
  if ec.IsZero(q) {
    r.x = p.x
    r.y = p.y
    return r
  }

  s := big.NewInt(0)
  t := big.NewInt(0)
  f := ec.f
  if p.x.Cmp(q.x) == 0 {
    if p.y.Cmp(q.y) != 0 || p.y.Cmp(big.NewInt(0)) == 0 {
      // p must be the inverse of q, i.e. either
      // p.y != q.y or p.y == q.y == 0
      return ec.Zero()
    }
    // p == q
    s = f.Mul(p.x, p.x)
    s = f.Mul(s, big.NewInt(3))
    s = f.Add(s, ec.a)
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

func (ec EllipticCurve) Neg(p ECPoint) ECPoint {

  r := NullECPoint()
  r.x.Set(p.x)
  r.y.Neg(p.y)

  return r
}

func (ec EllipticCurve) MulInt(p ECPoint, n *big.Int) ECPoint {

  r := NilECPoint()

  for i := 0; i < n.BitLen(); i++ {
    j := n.BitLen() - i - 1
    r  = ec.Add(r, r)
    if n.Bit(j) != 0 {
      r = ec.Add(r, p)
    }
  }
  return r
}

func (ec EllipticCurve) Zero() ECPoint {
  return NilECPoint()
}

func (ec EllipticCurve) IsZero(p ECPoint) bool {
  if p.x == nil || p.y == nil {
    return true
  }
  return false
}
