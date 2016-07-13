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
  isZero bool
}

/* -------------------------------------------------------------------------- */

func NewECPoint(x_, y_ *big.Int) ECPoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  if x_ != nil {
    x.Set(x_)
  }
  if y_ != nil {
    y.Set(y_)
  }
  return ECPoint{x, y, false}
}

func NullECPoint() ECPoint {
  x := big.NewInt(0)
  y := big.NewInt(0)
  return ECPoint{x, y, true}
}

/* -------------------------------------------------------------------------- */

func (p ECPoint) Clone() ECPoint {
  q := NullECPoint()
  q.x.Set(p.x)
  q.y.Set(p.y)
  return q
}

func (p ECPoint) IsZero() bool {
  return p.isZero
}

func (p ECPoint) GetX() *big.Int {
  return p.x
}

func (p ECPoint) GetY() *big.Int {
  return p.y
}

func (p *ECPoint) Set(q ECPoint) {
  p.isZero = q.isZero
  if !p.isZero {
    p.x.Set(q.x)
    p.y.Set(q.y)
  }
}

func (p *ECPoint) SetX(x *big.Int) {
  p.x.Set(x)
}

func (p *ECPoint) SetY(y *big.Int) {
  p.y.Set(y)
}

func (p *ECPoint) SetZero(v bool) {
  p.isZero = v
}

func (p ECPoint) String() string {
  if p.IsZero() {
    return fmt.Sprintf("(zero)")
  } else {
    return fmt.Sprintf("(%v,%v)", p.x, p.y)
  }
}
