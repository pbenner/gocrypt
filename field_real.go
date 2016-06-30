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

func (f RealField) FieldNeg(a_ FieldElement) FieldElement {
  a := a_.(float64)
  return -a
}

func (f RealField) FieldAdd(a_, b_ FieldElement) FieldElement {
  a := a_.(float64)
  b := b_.(float64)
  return a+b
}

func (f RealField) FieldSub(a_, b_ FieldElement) FieldElement {
  a := a_.(float64)
  b := b_.(float64)
  return a-b
}

func (f RealField) FieldMul(a_, b_ FieldElement) FieldElement {
  a := a_.(float64)
  b := b_.(float64)
  return a*b
}

func (f RealField) FieldDiv(a_, b_ FieldElement) FieldElement {
  a := a_.(float64)
  b := b_.(float64)
  return a/b
}

func (f RealField) FieldIsZero(a_ FieldElement) bool {
  a := a_.(float64)
  return math.Abs(a) < 1e-12
}

func (f RealField) FieldIsOne(a_ FieldElement) bool {
  a := a_.(float64)
  return math.Abs(1.0-a) < 1e-12
}

func (f RealField) FieldZero() FieldElement {
  return 0.0
}

func (f RealField) FieldOne() FieldElement {
  return 1.0
}
