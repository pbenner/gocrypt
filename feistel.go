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

/* -------------------------------------------------------------------------- */

type FeistelNetwork struct {
  Rounds      int
  BlockLength int // block length in bytes
  Key         func(i int) []byte
  F           func(key, input, output []byte)
}

/* -------------------------------------------------------------------------- */

// compute element-wise xor: z = x (+) y
func xorSlice(x, y, z []byte) {
  for i := 0; i < len(x); i++ {
    z[i] = x[i] ^ y[i]
  }
}

func (network FeistelNetwork) EncryptBlock(input, output []byte) {
  n := len(input)
  // variables at the end of a round
  Li := output[0:n/2]
  Ri := output[n/2:n]
  // let j = i+1
  Rj := input[n/2:n]
  // result of the F function
  Fout := make([]byte, network.BlockLength/2)
  // apply encryption multiple times
  for i := 0; i < network.Rounds; i++ {
    // switch input and output
    input, output = output, input
    // get the ith key
    key := network.Key(i)
    // call F function
    network.F(key, Ri, Fout)
    // encrypte Li and store result in Rj
    xorSlice(Li, Fout, Rj)
  }
}

func (network FeistelNetwork) Encrypt() {
}
