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

// round function
type RoundFunction func(key, input, output []byte)

type FeistelNetwork struct {
  BlockLength int // block length in bytes
  Keys        [][]byte
  F           RoundFunction
}

/* -------------------------------------------------------------------------- */

func NewFeistelNetwork(blockLength int, keys [][]byte, f RoundFunction) FeistelNetwork {
  return FeistelNetwork{blockLength, keys, f}
}

func (network FeistelNetwork) encryptBlock(input, output, iTmp, fTmp []byte) {
  l := network.BlockLength
  // copy input
  copy(iTmp, input)
  // variables at the end of a round
  Li := output[0:l/2]
  Ri := output[l/2:l]
  // let j = i+1
  Lj := iTmp[0:l/2]
  Rj := iTmp[l/2:l]
  // apply encryption multiple times
  for i := 0; i < len(network.Keys); i++ {
    // swap i and j
    Li, Ri, Lj, Rj = Lj, Rj, Li, Ri
    // copy Ri to Lj
    for k := 0; k < l/2; k++ {
      Lj[k] = Ri[k]
    }
    // call F function
    network.F(network.Keys[i], Ri, fTmp)
    // encrypte Li and store result in Rj
    Bits(Rj).Xor(Li, fTmp)
  }
  for i := 0; i < l/2; i++ {
    output[i], output[l/2+i] = Rj[i], Lj[i]
  }
}

func (network FeistelNetwork) Encrypt(input []byte) []byte {
  l := network.BlockLength
  // allocate some memory for holding temporary data
  iTmp := make([]byte, l)
  fTmp := make([]byte, l/2)
  // allocate memory for holding the output
  output := make([]byte, len(input))
  // loop over message and encrypt each block
  for i := 0; i < len(input); i += l {
    iBlock := input [i:i+l]
    oBlock := output[i:i+l]
    network.encryptBlock(iBlock, oBlock, iTmp, fTmp)
  }
  return output
}
