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

/* affine algebra
 * -------------------------------------------------------------------------- */

func (ec EllipticCurve) Add(p, q AffinePoint) AffinePoint {

  r := NullAffinePoint()

  if p.IsZero() && q.IsZero() {
    return r
  }
  if p.IsZero() {
    r.Set(q)
    return r
  }
  if q.IsZero() {
    r.Set(p)
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

func (ec EllipticCurve) Neg(p AffinePoint) AffinePoint {

  r := NewAffinePoint(p.x, p.y)
  r.y.Neg(r.y)

  return r
}

func (ec EllipticCurve) MulInt(p AffinePoint, n *big.Int) AffinePoint {

  r := NullAffinePoint()

  for i := 0; i < n.BitLen(); i++ {
    j := n.BitLen() - i - 1
    r  = ec.Add(r, r)
    if n.Bit(j) != 0 {
      r = ec.Add(r, p)
    }
  }
  return r
}

/* projective algebra (cf. Hankerson et al. 2003, pp. 88)
 * -------------------------------------------------------------------------- */

func (ec EllipticCurve) DoubleProjective(p ProjectivePoint) ProjectivePoint {

  r := NullProjectivePoint()

  if p.IsZero() {
    return r
  }

  a := big.NewInt(0)
  ec.f.Mul(a, p.y, p.y)
  b := big.NewInt(4)
  ec.f.Mul(b, b, p.x)
  ec.f.Mul(b, b, a)
  c := big.NewInt(8)
  ec.f.Mul(c, c, a)
  ec.f.Mul(c, c, a)
  t := big.NewInt(0)
  ec.f.Mul(t, ec.a, p.z)
  ec.f.Mul(t,    t, p.z)
  ec.f.Mul(t,    t, p.z)
  ec.f.Mul(t,    t, p.z)
  d := big.NewInt(3)
  ec.f.Mul(d, d, p.x)
  ec.f.Mul(d, d, p.x)
  ec.f.Add(d, d, t)

  ec.f.Mul(r.x, d, d)
  ec.f.Sub(r.x, r.x, b)
  ec.f.Sub(r.x, r.x, b)

  ec.f.Sub(r.y, b, r.x)
  ec.f.Mul(r.y, r.y, d)
  ec.f.Sub(r.y, r.y, c)

  r.z.SetInt64(2)
  ec.f.Mul(r.z, r.z, p.y)
  ec.f.Mul(r.z, r.z, p.z)

  return r
}

func (ec EllipticCurve) AddMixed(p ProjectivePoint, q AffinePoint) ProjectivePoint {

  r := NullProjectivePoint()

  if p.IsZero() && q.IsZero() {
    return r
  }
  if p.IsZero() {
    r.SetAffine(q)
    return r
  }
  if q.IsZero() {
    r.Set(p)
    return r
  }

  a := big.NewInt(0)
  ec.f.Mul(a, p.z, p.z)
  b := big.NewInt(0)
  ec.f.Mul(b, p.z, a)
  c := big.NewInt(0)
  ec.f.Mul(c, q.x, a)
  d := big.NewInt(0)
  ec.f.Mul(d, q.y, b)
  e := big.NewInt(0)
  ec.f.Sub(e, c, p.x)
  f := big.NewInt(0)
  ec.f.Sub(f, d, p.y)
  g := big.NewInt(0)
  ec.f.Mul(g, e, e)
  h := big.NewInt(0)
  ec.f.Mul(h, g, e)
  i := big.NewInt(0)
  ec.f.Mul(i, p.x, g)

  ec.f.Mul(r.x, f, f)
  ec.f.Sub(r.x, r.x, h)
  ec.f.Sub(r.x, r.x, i)
  ec.f.Sub(r.x, r.x, i)

  ec.f.Sub(r.y, i, h)
  ec.f.Mul(r.y, r.y, f)
  ec.f.Mul(h, h, p.y)
  ec.f.Sub(r.y, r.y, h)

  ec.f.Mul(r.z, p.z, e)

  return r
}

func (ec EllipticCurve) NegProjective(p ProjectivePoint) ProjectivePoint {

  r := p.Clone()
  r.y.Neg(r.y)

  return r
}

func (ec EllipticCurve) MulIntProjective(q AffinePoint, n *big.Int) AffinePoint {

  p := NullProjectivePoint()
  r := NullAffinePoint()

  for i := 0; i < n.BitLen(); i++ {
    j := n.BitLen() - i - 1
    p  = ec.DoubleProjective(p)
    if n.Bit(j) != 0 {
      p = ec.AddMixed(p, q)
    }
  }
  // convert to affine point
  t := big.NewInt(0)
  ec.f.Mul(t, p.z, p.z)
  ec.f.Div(r.x, p.x, t)
  ec.f.Mul(t, t, p.z)
  ec.f.Div(r.y, p.y, t)

  return r
}
