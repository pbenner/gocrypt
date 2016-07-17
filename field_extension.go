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

type FieldExtension struct {
  IP *Polynomial
}

/* -------------------------------------------------------------------------- */

func NewFieldExtension(ip *Polynomial) FieldExtension {
  return FieldExtension{ip}
}

/* -------------------------------------------------------------------------- */

func (f FieldExtension) Add(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Add(a, b)
  return r
}

func (f FieldExtension) Sub(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Sub(a, b)
  return r
}

func (f FieldExtension) Mul(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Mul(a, b)
  r.Mod(r, f.IP)
  return r
}

func (f FieldExtension) Div(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  _, _, t := PolynomialEEA(f.IP, b)
  r.Mul(a, t)
  return r
}

func (f FieldExtension) Zero() *Polynomial {
  r := NewPolynomial(f.IP.Field)
  return r
}

func (f FieldExtension) One() *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.AddTerm(1, 0)
  return r
}

func (f FieldExtension) IsZero(a *Polynomial) bool {
  return len(a.Terms) == 0
}

func (f FieldExtension) IsOne(a *Polynomial) bool {
  if len(a.Terms) == 1 {
    if v, ok := a.Terms[0]; ok {
      return f.IP.Field.IsOne(v)
    }
  }
  return false
}

/* -------------------------------------------------------------------------- */

func (p *Polynomial) ReadByte(b byte) {
  p.SetZero()
  for i := 0; i < 8; i ++ {
    if b & (1 << byte(i)) != 0 {
      p.AddTerm(p.Field.One(), i)
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
