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

func TestBitmap(t *testing.T) {

  table := []int{
    0, 1, 9, 3, 4, 5, 6, 7,
    8, 2,10,11,12,13,14,15}
  input  := []byte{4,0}
  output := remapBits(input, table)

  if output[0] != 0 {
    t.Error("bitmap test failed")
  }
  if output[1] != 2 {
    t.Error("bitmap test failed")
  }
}