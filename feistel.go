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

// key function returning the ith subkey
type Kfunc func(i int) []byte
// feistel function
type Ffunc func(key, input, output []byte)

type FeistelNetwork struct {
  Rounds      int
  BlockLength int // block length in bytes
  K           Kfunc
  F           Ffunc
}

/* -------------------------------------------------------------------------- */

func NewFeistelNetwork(round, blockLength int, k Kfunc, f Ffunc) FeistelNetwork {
  return FeistelNetwork{round, blockLength, k, f}
}

func (network FeistelNetwork) encryptBlock(input, output, fTmp []byte) {
  l := network.BlockLength
  // variables at the end of a round
  Li := output[0:l/2]
  Ri := output[l/2:l]
  // let j = i+1
  Lj := input[0:l/2]
  Rj := input[l/2:l]
  // apply encryption multiple times
  for i := 0; i < network.Rounds; i++ {
    // swap i and j
    Li, Ri, Lj, Rj = Lj, Rj, Li, Ri
    // copy Ri to Lj
    for k := 0; k < l/2; k++ {
      Lj[k] = Ri[k]
    }
    // get the ith key
    key := network.K(i)
    // call F function
    network.F(key, Ri, fTmp)
    // encrypte Li and store result in Rj
    xorSlice(Li, fTmp, Rj)
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
  }
  return output
}
