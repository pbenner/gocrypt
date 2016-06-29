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

func (f FiniteField) modp(r *Polynomial) *Polynomial {
  for k, v := range r.Terms {
    r.Terms[k] = float64(f.P.Modp(int(v)))
  }
  return r
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) Add(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.Add(a, b)
  return f.modp(r)
}

func (f FiniteField) Sub(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.Sub(a, b)
  return f.modp(r)
}

func (f FiniteField) Mul(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.Sub(a, b)
  return f.modp(r)
}
