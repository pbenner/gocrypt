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

type ExtensionField struct {
  IP *Polynomial
}

/* -------------------------------------------------------------------------- */

func NewExtensionField(ip *Polynomial) ExtensionField {
  return ExtensionField{ip}
}

/* -------------------------------------------------------------------------- */

func (f ExtensionField) Add(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Add(a, b)
  return r
}

func (f ExtensionField) Sub(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Sub(a, b)
  return r
}

func (f ExtensionField) Mul(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.Mul(a, b)
  r.Mod(r, f.IP)
  return r
}

func (f ExtensionField) Div(a, b *Polynomial) *Polynomial {
  r := NewPolynomial(f.IP.Field)
  _, _, t := PolynomialEEA(f.IP, b)
  r.Mul(a, t)
  return r
}

func (f ExtensionField) Zero() *Polynomial {
  r := NewPolynomial(f.IP.Field)
  return r
}

func (f ExtensionField) One() *Polynomial {
  r := NewPolynomial(f.IP.Field)
  r.AddTerm(1, 0)
  return r
}

func (f ExtensionField) IsZero(a *Polynomial) bool {
  return len(a.Terms) == 0
}

func (f ExtensionField) IsOne(a *Polynomial) bool {
  if len(a.Terms) == 1 {
    if v, ok := a.Terms[0]; ok {
      return f.IP.Field.IsOne(v)
    }
  }
  return false
}
