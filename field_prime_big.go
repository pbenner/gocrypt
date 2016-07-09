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

func NewBigPrimeField(p *big.Int) BigPrimeField {
  return BigPrimeField{p}
}

/* -------------------------------------------------------------------------- */

func (f BigPrimeField) FieldNeg(a_ FieldElement) FieldElement {
  a := a_.(*big.Int)
  return f.Neg(a)
}

func (f BigPrimeField) FieldAdd(a_, b_ FieldElement) FieldElement {
  a := a_.(*big.Int)
  b := b_.(*big.Int)
  return f.Add(a, b)
}

func (f BigPrimeField) FieldSub(a_, b_ FieldElement) FieldElement {
  a := a_.(*big.Int)
  b := b_.(*big.Int)
  return f.Sub(a, b)
}

func (f BigPrimeField) FieldMul(a_, b_ FieldElement) FieldElement {
  a := a_.(*big.Int)
  b := b_.(*big.Int)
  return f.Mul(a, b)
}

func (f BigPrimeField) FieldDiv(a_, b_ FieldElement) FieldElement {
  a := a_.(*big.Int)
  b := b_.(*big.Int)
  return f.Div(a, b)
}

func (f BigPrimeField) FieldIsZero(a_ FieldElement) bool {
  a := a_.(*big.Int)
  return f.IsZero(a)
}

func (f BigPrimeField) FieldIsOne(a_ FieldElement) bool {
  a := a_.(*big.Int)
  return f.IsOne(a)
}

func (f BigPrimeField) FieldZero() FieldElement {
  return 0
}

func (f BigPrimeField) FieldOne() FieldElement {
  return 1
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
