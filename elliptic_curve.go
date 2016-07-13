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

/* -------------------------------------------------------------------------- */

type EllipticCurve struct {
  a, b *big.Int
  f    BigPrimeField
}

/* -------------------------------------------------------------------------- */

func NewEllipticCurve(a_, b_, p *big.Int) EllipticCurve {
  a := big.NewInt(0)
  a.Set(a_)
  b := big.NewInt(0)
  b.Set(b_)
  f := NewBigPrimeField(p)
  return EllipticCurve{a, b, f}
}

/* -------------------------------------------------------------------------- */

func (ec EllipticCurve) Add(p, q ECPoint) ECPoint {

  r := NullECPoint()

  if p.IsZero() && q.IsZero() {
    return r
  }
  if p.IsZero() {
    r.SetZero(false)
    r.x = q.x
    r.y = q.y
    return r
  }
  if q.IsZero() {
    r.SetZero(false)
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
      return r
    }
    // p == q
    f.Mul(s, p.x, p.x)
    f.Mul(s, s, big.NewInt(3))
    f.Add(s, s, ec.a)
    f.Div(s, s, p.y)
    f.Div(s, s, big.NewInt(2))
  } else {
    // p != q
    f.Sub(s, p.y, q.y)
    f.Sub(t, p.x, q.x)
    f.Div(s, s, t)
  }
  f.Mul(r.x, s, s)
  f.Sub(r.x, r.x, p.x)
  f.Sub(r.x, r.x, q.x)
  f.Sub(r.y, r.x, p.x)
  f.Mul(r.y, r.y, s)
  f.Add(r.y, r.y, p.y)
  f.Neg(r.y, r.y)
  r.SetZero(false)
  return r
}

func (ec EllipticCurve) Neg(p ECPoint) ECPoint {

  r := NewECPoint(p.x, p.y)
  r.y.Neg(r.y)

  return r
}

func (ec EllipticCurve) MulInt(p ECPoint, n *big.Int) ECPoint {

  r := NewECPoint(nil, nil)

  for i := 0; i < n.BitLen(); i++ {
    j := n.BitLen() - i - 1
    r  = ec.Add(r, r)
    if n.Bit(j) != 0 {
      r = ec.Add(r, p)
    }
  }
  return r
}
