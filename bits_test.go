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

package lib

/* -------------------------------------------------------------------------- */

//import "fmt"
import "testing"

/* -------------------------------------------------------------------------- */

func TestBitsRead(t *testing.T) {

  str1 := "01001011 00000001 01011000 11110001 11000001 01111000 00111001 01010011"
  str2 := Bits{}.Read(str1).String()

  if str1 != str2 {
    t.Error("reading bits from string failed")
  }
}

/* -------------------------------------------------------------------------- */

func TestRotateSlice1(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, -2)

  if y[0] != (x[0] >> 2) + (x[1] << 6) {
    t.Error("rotate test failed")
  }
  if y[1] != (x[1] >> 2) + (x[2] << 6) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[2] >> 2) + (x[3] << 6) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[3] >> 2) + (x[4] << 6) {
    t.Error("rotate test failed")
  }
  if y[4] != (x[4] >> 2) + (x[0] << 6) {
    t.Error("rotate test failed")
  }

}

func TestRotateSlice2(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, -10)

  if y[4] != (x[0] >> 2) + (x[1] << 6) {
    t.Error("rotate test failed")
  }
  if y[0] != (x[1] >> 2) + (x[2] << 6) {
    t.Error("rotate test failed")
  }
  if y[1] != (x[2] >> 2) + (x[3] << 6) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[3] >> 2) + (x[4] << 6) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[4] >> 2) + (x[0] << 6) {
    t.Error("rotate test failed")
  }

}

func TestRotateSlice3(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, 13)

  if y[1] != (x[0] << 5) + (x[4] >> 3) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[1] << 5) + (x[0] >> 3) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[2] << 5) + (x[1] >> 3) {
    t.Error("rotate test failed")
  }
  if y[4] != (x[3] << 5) + (x[2] >> 3) {
    t.Error("rotate test failed")
  }
  if y[0] != (x[4] << 5) + (x[3] >> 3) {
    t.Error("rotate test failed")
  }

}
