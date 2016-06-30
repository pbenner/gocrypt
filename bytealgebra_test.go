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
import "testing"

/* -------------------------------------------------------------------------- */

func TestByteAlgebra(t *testing.T) {

  // y = m x + v

  m := ByteMatrix(Bits{}.Read("10001111 11000111 11100011 11110001 11111000 01111100 00111110 00011111"))
  v := ByteVector(Bits{}.Read("01100011")[0])

  x := ByteVector(Bits{}.Read("00101111")[0])
  y := ByteVaddV(ByteMmulV(m, x), v)

  r := ByteVector(Bits{}.Read("00100101")[0])

  if y != r {
    t.Error("byte algebra test failed")
  }
}
