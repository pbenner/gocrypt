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

type PrimeField int

/* -------------------------------------------------------------------------- */

func NewPrimeField(p int) PrimeField {
  return PrimeField(p)
}

/* -------------------------------------------------------------------------- */

func (p PrimeField) Neg(a int) int {
  return p.Modp(-a)
}

func (p PrimeField) Add(a, b int) int {
  return p.Modp(a+b)
}

func (p PrimeField) Sub(a, b int) int {
  return p.Modp(a-b)
}

func (p PrimeField) Mul(a, b int) int {
  return p.Modp(a*b)
}

func (p PrimeField) Div(a, b int) int {
  r, _, t := EEA(int(p), b)
  if r != 1 {
    panic("divisor does not have an inverse")
  }
  return p.Mul(a, p.Modp(t))
}

func (p PrimeField) Modp(a int) int {
  r := a % int(p)
  if r < 0 {
    return int(p) + r
  } else {
    return r
  }
}

func (p PrimeField) Zero() int {
  return 0
}

func (p PrimeField) One() int {
  return 1
}

func (p PrimeField) IsZero(a int) bool {
  return a == 0
}

func (p PrimeField) IsOne(a int) bool {
  return a == 1
}
