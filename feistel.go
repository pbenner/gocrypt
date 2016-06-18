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

import "fmt"

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

func (network FeistelNetwork) encryptBlock(input, output, fTmp []byte) {
  l := network.BlockLength
  // variables at the end of a round
  Lj := output[0:l/2]
  Rj := output[l/2:l]
  // let j = i+1
  Li := input[0:l/2]
  Ri := input[l/2:l]
  fmt.Println("L0:", Bits(Lj))
  fmt.Println("R0:", Bits(Rj))
  // apply encryption multiple times
  for i := 0; i < len(network.Keys); i++ {
    fmt.Println("Round:", i)
    // copy Ri to Lj
    for k := 0; k < l/2; k++ {
      Lj[k] = Ri[k]
    }
    // call F function
    network.F(network.Keys[i], Ri, fTmp)
    // encrypte Li and store result in Rj
    xorSlice(Li, fTmp, Rj)
    fmt.Println("Li:", Bits(Lj))
    fmt.Println("Ri:", Bits(Rj))
    fmt.Println()
    // swap i and j
    Li, Ri, Lj, Rj = Lj, Rj, Li, Ri
  }
}

func (network FeistelNetwork) Encrypt(input []byte) []byte {
  l := network.BlockLength
  // allocate some memory for holding temporary data
  iTmp := make([]byte, l)
  fTmp := make([]byte, l/2)
  // allocate memory for holding the output
  output := make([]byte, len(input))
  // make a copy the input since EncryptBlock 
  for i := 0; i < len(input); i += l {
    iBlock := input [i:i+l]
    oBlock := output[i:i+l]
    copy(iTmp, iBlock)
    network.encryptBlock(iTmp, oBlock, fTmp)
    fmt.Println("feistel result:", Bits(oBlock))
  }
  return output
}
