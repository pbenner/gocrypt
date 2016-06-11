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

func TestFeistel(t *testing.T) {

  // Feistel network with just one round and a block length
  // of two bytes. Let input = [L0, R0], then the output =
  // [L1,R1] is given by L1 = R0 and R1 = L0 (+) R0.

  k := func(i int) []byte {
    return []byte{}
  }
  f := func(key, input, output []byte) {
    for i := 0; i < len(input); i++ {
      output[i] = input[i]
    }
  }
  network := NewFeistelNetwork(1, 2, k, f)

  input  := []byte{13,3}
  output := network.Encrypt(input)

  if output[0] != 3 {
    t.Error("bitmap test failed")
  }
  if output[1] != 14 {
    t.Error("bitmap test failed")
  }
}
