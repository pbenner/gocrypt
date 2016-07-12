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

type FieldInt interface {
  Neg(a    int) int
  Add(a, b int) int
  Sub(a, b int) int
  Mul(a, b int) int
  Div(a, b int) int
  IsZero(a int) bool
  IsOne (a int) bool
  Zero  () int
  One   () int
}

type FieldFloat interface {
  Neg(a    float64) float64
  Add(a, b float64) float64
  Sub(a, b float64) float64
  Mul(a, b float64) float64
  Div(a, b float64) float64
  IsZero(a float64) bool
  IsOne (a float64) bool
  Zero  () float64
  One   () float64
}

type FieldPolynomial interface {
  Neg(a    *Polynomial) *Polynomial
  Add(a, b *Polynomial) *Polynomial
  Sub(a, b *Polynomial) *Polynomial
  Mul(a, b *Polynomial) *Polynomial
  Div(a, b *Polynomial) *Polynomial
  IsZero(a *Polynomial) bool
  IsOne (a *Polynomial) bool
  Zero  () *Polynomial
  One   () *Polynomial
}
