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

type AffinePoint struct {
  x, y *big.Int
  isZero bool
}

/* -------------------------------------------------------------------------- */

func NewAffinePoint(x_, y_ *big.Int) AffinePoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  if x_ != nil {
    x.Set(x_)
  }
  if y_ != nil {
    y.Set(y_)
  }
  return AffinePoint{x, y, false}
}

func NullAffinePoint() AffinePoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  return AffinePoint{x, y, true}
}

/* -------------------------------------------------------------------------- */

func (p AffinePoint) Clone() AffinePoint {
  q := NullAffinePoint()
  q.x.Set(p.x)
  q.y.Set(p.y)
  return q
}

func (p AffinePoint) IsZero() bool {
  return p.isZero
}

func (p AffinePoint) GetX() *big.Int {
  return p.x
}

func (p AffinePoint) GetY() *big.Int {
  return p.y
}

func (p *AffinePoint) Set(q AffinePoint) {
  p.isZero = q.isZero
  if !p.isZero {
    p.x.Set(q.x)
    p.y.Set(q.y)
  }
}

func (p *AffinePoint) SetX(x *big.Int) {
  p.x.Set(x)
}

func (p *AffinePoint) SetY(y *big.Int) {
  p.y.Set(y)
}

func (p *AffinePoint) SetZero(v bool) {
  p.isZero = v
}

func (p AffinePoint) String() string {
  if p.IsZero() {
    return fmt.Sprintf("(zero)")
  } else {
    return fmt.Sprintf("(%v,%v)", p.x, p.y)
  }
}

/* -------------------------------------------------------------------------- */

type ProjectivePoint struct {
  x, y, z *big.Int
}

/* -------------------------------------------------------------------------- */

func NewProjectivePoint(x_, y_, z_ *big.Int) ProjectivePoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  z := big.NewInt(1)
  if x_ != nil {
    x.Set(x_)
  }
  if y_ != nil {
    y.Set(y_)
  }
  if y_ != nil {
    z.Set(z_)
  }
  return ProjectivePoint{x, y, z}
}

func NullProjectivePoint() ProjectivePoint {
  x := big.NewInt(0)
  y := big.NewInt(1)
  z := big.NewInt(0)
  return ProjectivePoint{x, y, z}
}

/* -------------------------------------------------------------------------- */

func (p ProjectivePoint) Clone() ProjectivePoint {
  q := NullProjectivePoint()
  q.x.Set(p.x)
  q.y.Set(p.y)
  q.z.Set(p.z)
  return q
}

func (p ProjectivePoint) IsZero() bool {
  return p.z.Cmp(big.NewInt(0)) == 0
}

func (p ProjectivePoint) GetX() *big.Int {
  return p.x
}

func (p ProjectivePoint) GetY() *big.Int {
  return p.y
}

func (p ProjectivePoint) GetZ() *big.Int {
  return p.z
}

func (p *ProjectivePoint) Set(q ProjectivePoint) {
  p.x.Set(q.x)
  p.y.Set(q.y)
  p.z.Set(q.z)
}

func (p *ProjectivePoint) SetAffine(q AffinePoint) {
  p.x.Set(q.x)
  p.y.Set(q.y)
  p.z.SetInt64(1)
}

func (p *ProjectivePoint) SetX(x *big.Int) {
  p.x.Set(x)
}

func (p *ProjectivePoint) SetY(y *big.Int) {
  p.y.Set(y)
}

func (p *ProjectivePoint) SetZ(z *big.Int) {
  p.z.Set(z)
}

func (p *ProjectivePoint) SetZero() {
  p.z.SetInt64(0)
}

func (p ProjectivePoint) String() string {
  if p.IsZero() {
    return fmt.Sprintf("(zero)")
  } else {
    return fmt.Sprintf("(%v:%v:%v)", p.x, p.y, p.z)
  }
}
