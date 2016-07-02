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

package gocrypt

/* -------------------------------------------------------------------------- */

//import "fmt"

/* -------------------------------------------------------------------------- */

func (AESCipher) g(input, output []byte, i byte) {
  output[0] = aesSbox[input[1]] ^ aesRcon[i]
  output[1] = aesSbox[input[2]]
  output[2] = aesSbox[input[3]]
  output[3] = aesSbox[input[0]]
}

func (AESCipher) h(input, output []byte) {
  output[0] = aesSbox[input[0]]
  output[1] = aesSbox[input[1]]
  output[2] = aesSbox[input[2]]
  output[3] = aesSbox[input[3]]
}

/* -------------------------------------------------------------------------- */

func (cipher *AESCipher) subkeys(key []byte, rounds int, h bool) {
  subkeylen := cipher.BlockLength
  if 4*len(key) % subkeylen != 0 {
    panic("AESCipher.subkeys(): invalid key length")
  }
  subkeys := make([]byte, (rounds+1)*subkeylen)
  copy(subkeys[0:len(key)], key)
  for m, i := 4*len(key)/subkeylen, 1; 4*m*i < len(subkeys); i++ {
    i0 := 4*m*(i-0)
    i1 := 4*m*(i-1)
    cipher.g(subkeys[i0-4:i0], subkeys[i0:i0+4], byte(i))
    Bits(subkeys[i0:i0+1*4]).Xor(subkeys[i1:i1+4], subkeys[i0:i0+4])
    for j := 1; j < m && i0+4*(j+1) <= len(subkeys); j++ {
      j0 := 4*(j-0)
      j1 := 4*(j-1)
      j2 := 4*(j+1)
      if h && j % (m/2) == 0 {
        cipher.h(subkeys[i0+j1:i0+j0], subkeys[i0+j0:i0+j2])
        Bits(subkeys[i0+j0:i0+j2]).Xor(subkeys[i0+j0:i0+j2], subkeys[i1+j0:i1+j2])
      } else {
        Bits(subkeys[i0+j0:i0+j2]).Xor(subkeys[i1+j0:i1+j2], subkeys[i0+j1:i0+j0])
      }
    }
  }
  cipher.Keys = make([][]byte, rounds+1)
  for i := 0; i < len(subkeys)/subkeylen; i++ {
    cipher.Keys[i] = subkeys[i*subkeylen:(i+1)*subkeylen]
  }
}
