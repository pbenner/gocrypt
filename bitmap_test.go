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
//import "strconv"
import "testing"

/* -------------------------------------------------------------------------- */

func TestBitPermuation(t *testing.T) {

  table := []int{
    0, 1, 9, 3, 4, 5, 6, 7,
    8, 2,10,11,12,13,14,15}
  input  := []byte{4,0}
  output := []byte{0,0}

  PermuteBits(input, output, table)

  if output[0] != 0 {
    t.Error("bitmap test failed")
  }
  if output[1] != 2 {
    t.Error("bitmap test failed")
  }
}

func TestBitMap(t *testing.T) {

  input  := []byte{0,0,1,1<<7}
  output := []byte{0,0,0,0,0,0}

  fExpansion(input, output)

  if output[0] != 1 {
    t.Error("bitmap test failed")
  }
  if output[1] != 0 {
    t.Error("bitmap test failed")
  }
  if output[2] != (1<<7) {
    t.Error("bitmap test failed")
  }
  if output[3] != (1<<1) {
    t.Error("bitmap test failed")
  }
  if output[4] != 0 {
    t.Error("bitmap test failed")
  }
  if output[5] != (1<<6) {
    t.Error("bitmap test failed")
  }
  // fmt.Println("input[0]:", strconv.FormatInt(int64(input[0]), 2))
  // fmt.Println("input[1]:", strconv.FormatInt(int64(input[1]), 2))
  // fmt.Println("input[2]:", strconv.FormatInt(int64(input[2]), 2))
  // fmt.Println("input[3]:", strconv.FormatInt(int64(input[3]), 2))

  // fmt.Println("output[0]:", strconv.FormatInt(int64(output[0]), 2))
  // fmt.Println("output[1]:", strconv.FormatInt(int64(output[1]), 2))
  // fmt.Println("output[2]:", strconv.FormatInt(int64(output[2]), 2))
  // fmt.Println("output[3]:", strconv.FormatInt(int64(output[3]), 2))
  // fmt.Println("output[4]:", strconv.FormatInt(int64(output[4]), 2))
  // fmt.Println("output[5]:", strconv.FormatInt(int64(output[5]), 2))
}
