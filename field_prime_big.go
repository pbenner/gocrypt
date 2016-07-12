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

type BigPrimeField struct {
  p *big.Int
}

/* -------------------------------------------------------------------------- */

func NewBigPrimeField(p_ *big.Int) BigPrimeField {
  p := big.NewInt(0)
  p.Set(p_)
  return BigPrimeField{p}
}

/* -------------------------------------------------------------------------- */

func (f BigPrimeField) Neg(a *big.Int) *big.Int {
  return f.Modp(a.Neg(a))
}

func (f BigPrimeField) Add(a, b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Add(a, b)
  return f.Modp(r)
}

func (f BigPrimeField) Sub(a, b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Sub(a, b)
  return f.Modp(r)
}

func (f BigPrimeField) Mul(a, b *big.Int) *big.Int {
  r := big.NewInt(0)
  r.Mul(a, b)
  return f.Modp(r)
}

func (f BigPrimeField) Div(a, b *big.Int) *big.Int {
  r, _, t := BigEEA(f.p, b)
  if r.Cmp(big.NewInt(1)) != 0 {
    panic("divisor does not have an inverse")
  }
  return f.Mul(a, f.Modp(t))
}

func (f BigPrimeField) IsZero(a *big.Int) bool {
  return a.Cmp(big.NewInt(0)) == 0
}

func (f BigPrimeField) IsOne(a *big.Int) bool {
  return a.Cmp(big.NewInt(1)) == 0
}

func (f BigPrimeField) Modp(a *big.Int) *big.Int {
  r := a.Mod(a, f.p)
  if r.Cmp(big.NewInt(0)) < 0 {
    r.Add(r, f.p)
  }
  return r
}
