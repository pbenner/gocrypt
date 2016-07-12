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
import "math"

/* -------------------------------------------------------------------------- */

type RealField struct {
}

/* -------------------------------------------------------------------------- */

func NewRealField() RealField {
  return RealField{}
}

/* -------------------------------------------------------------------------- */

func (f RealField) Neg(a float64) float64 {
  return -a
}

func (f RealField) Add(a, b float64) float64 {
  return a+b
}

func (f RealField) Sub(a, b float64) float64 {
  return a-b
}

func (f RealField) Mul(a, b float64) float64 {
  return a*b
}

func (f RealField) Div(a, b float64) float64 {
  return a/b
}

func (f RealField) IsZero(a float64) bool {
  return math.Abs(a) < 1e-12
}

func (f RealField) IsOne(a float64) bool {
  return math.Abs(1.0-a) < 1e-12
}

func (f RealField) Zero() float64 {
  return 0.0
}

func (f RealField) One() float64 {
  return 1.0
}
