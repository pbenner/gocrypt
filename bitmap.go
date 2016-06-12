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
import "errors"

/* -------------------------------------------------------------------------- */

func reduceTableInjective(table [][]int, result []int) error {
  for i := 0; i < len(result); i++ {
    result[i] = -1
  }
  for i := 0; i < len(table); i++ {
    // i: input bit
    // j: output bit
    for k := 0; k < len(table[i]); k++ {
      j := table[i][k]
      if result[j] == -1 {
        result[j] = i
      } else {
        return errors.New("table cannot be converted")
      }
    }
  }
  return nil
}

/* -------------------------------------------------------------------------- */

// Surjective mapping of input bits to output bits. The mapping
// is defined by the table. The ith bit in the input slice is
// mapped to position j = table[i] in the output.
func BitmapSurjective(input, output []byte, table []int) {
  // number of input bits
  n := 8*len(input)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over input bits
  for i := 0; i < n; i++ {
    // index of the output bit
    j := table[i]
    if input[i/8]  & byte(1 << byte(i % 8)) != 0 {
      output[j/8] |= byte(1 << byte(j % 8))
    }
  }
}

// Injective mapping of input bits to output bits. The mapping
// is defined by the table. The jth bit in the output slice is
// copied from position i = table[j] in the input.
func BitmapInjective(input, output []byte, table []int) {
  // number of input bits
  n := 8*len(output)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over output bits
  for j := 0; j < n; j++ {
    // index of the input bit
    i := table[j]
    if input[i/8]  & byte(1 << byte(i % 8)) != 0 {
      output[j/8] |= byte(1 << byte(j % 8))
    }
  }
}

func Bitmap(input, output []byte, table [][]int) {
  // number of input bits
  n := 8*len(input)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over input bits
  for i := 0; i < n; i++ {
    for k := 0; k < len(table[i]); k++ {
      // index of the output bit
      j := table[i][k]
      if input[i/8]  & byte(1 << byte(i % 8)) != 0 {
        output[j/8] |= byte(1 << byte(j % 8))
      }
    }
  }
}
