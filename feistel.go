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

// compute element-wise xor: z = x (+) y
func xorSlice(x, y, z []byte) {
  for i := 0; i < len(x); i++ {
    z[i] = x[i] ^ y[i]
  }
}

/* -------------------------------------------------------------------------- */

// key function returning the ith subkey
type Kfunc func(i int) []byte
// encryption function
type Ffunc func(key, input, output []byte)

type FeistelNetwork struct {
  Rounds      int
  BlockLength int // block length in bytes
  K           Kfunc
  F           Ffunc
  Fout        []byte
}

/* -------------------------------------------------------------------------- */

func NewFeistelNetwork(round, blockLength int, k Kfunc, f Ffunc) FeistelNetwork {
  // result of the F function
  fout := make([]byte, blockLength/2)
  return FeistelNetwork{round, blockLength, k, f, fout}
}

func (network FeistelNetwork) EncryptBlock(input, output []byte) {
  n := len(input)
  // variables at the end of a round
  Li := output[0:n/2]
  Ri := output[n/2:n]
  // let j = i+1
  Rj := input[n/2:n]
  // apply encryption multiple times
  for i := 0; i < network.Rounds; i++ {
    // switch input and output
    input, output = output, input
    // get the ith key
    key := network.K(i)
    // call F function
    network.F(key, Ri, network.Fout)
    // encrypte Li and store result in Rj
    xorSlice(Li, network.Fout, Rj)
  }
}

func (network FeistelNetwork) Encrypt(input []byte) []byte {
  l := network.BlockLength
  output := make([]byte, len(input))
  for i := 0; i < len(input); i += l {
    iBlock := input [i:i+l]
    oBlock := output[i:i+l]
    network.EncryptBlock(iBlock, oBlock)
  }
  return output
}
