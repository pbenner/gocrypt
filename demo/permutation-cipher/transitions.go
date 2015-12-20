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

import   "math"

import . "github.com/pbenner/autodiff"

/* -------------------------------------------------------------------------- */

func newTransitionMatrix(sType ScalarType, n int, alpha float64) Matrix {
  t := NullMatrix(sType, n, n)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      t.Set(NewScalar(sType, alpha), i, j)
    }
  }
  return t
}

func normalizeMatrix(t Matrix) {
  rows, cols := t.Dims()
  for i := 0; i < rows; i++ {
    sum := NewScalar(t.ElementType(), 0.0)
    for j := 0; j < cols; j++ {
      sum = Add(sum, t.At(i, j))
    }
    for j := 0; j < cols; j++ {
      t.Set(Div(t.At(i, j), sum), i, j)
    }
  }
}

func estimateTransitions(sType ScalarType, text string) Matrix {
  n := int(math.Pow(2, 8))
  t := newTransitionMatrix(sType, n, 0.001)

  for i := 1; i < len(text); i++ {
    j1 := int(text[i-1])
    j2 := int(text[i])
    // increment count by one
    t.Set(Add(t.At(j1, j2), NewScalar(sType, 1.0)), j1, j2)
  }
  normalizeMatrix(t)

  return t
}
