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

/* -------------------------------------------------------------------------- */

type FiniteField struct {
  P   PrimeField
  N   int
  IP *Polynomial
}

/* -------------------------------------------------------------------------- */

func NewFiniteField(p, n int, ip *Polynomial) FiniteField {
  return FiniteField{NewPrimeField(p), n, ip}
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) FieldAdd(a_, b_ FieldElement) FieldElement {
  a := a_.(*Polynomial)
  b := b_.(*Polynomial)
  return f.Add(a, b)
}

func (f FiniteField) FieldSub(a_, b_ FieldElement) FieldElement {
  a := a_.(*Polynomial)
  b := b_.(*Polynomial)
  return f.Sub(a, b)
}

func (f FiniteField) FieldMul(a_, b_ FieldElement) FieldElement {
  a := a_.(*Polynomial)
  b := b_.(*Polynomial)
  return f.Mul(a, b)
}

func (f FiniteField) FieldDiv(a_, b_ FieldElement) FieldElement {
  a := a_.(*Polynomial)
  b := b_.(*Polynomial)
  return f.Div(a, b)
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) Add(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.P)
  r.Add(a, b)
  return r
}

func (f FiniteField) Sub(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.P)
  r.Sub(a, b)
  return r
}

func (f FiniteField) Mul(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.P)
  r.Mul(a, b)
  r.Mod(r, f.IP)
  return r
}

func (f FiniteField) Div(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.P)
  _, _, t := PolynomialEEA(f.IP, b)
  r.Mul(a, t)
  return r
}

/* -------------------------------------------------------------------------- */

func (p *Polynomial) ReadByte(b byte) {
  p.Clear()
  for i := 0; i < 8; i ++ {
    if b & (1 << byte(i)) != 0 {
      p.AddTerm(p.Field.FieldOne(), i)
    }
  }
}

func (p *Polynomial) WriteByte() byte {
  b := byte(0)
  for k, _ := range p.Terms {
    if k < 8 {
      b |= (1 << byte(k))
    }
  }
  return b
}
