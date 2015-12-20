/* Copyright (C) 2015 Philipp Benner
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

package main

/* -------------------------------------------------------------------------- */

import . "github.com/pbenner/autodiff"
import . "github.com/pbenner/0100101101"

/* -------------------------------------------------------------------------- */

func likelihood(text Message, t Matrix) Scalar {
  result := NewScalar(t.ElementType(), 1.0)

  for i := 1; i < len(text); i++ {
    j1 := int(text[i-1])
    j2 := int(text[i])
    result = Mul(result, t.At(j1, j2))
  }

  return result
}
