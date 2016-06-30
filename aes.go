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

type AESCipher struct {
  BlockLength int // block length in bytes
  Keys        [][]byte
}

/* -------------------------------------------------------------------------- */

func (AESCipher) substitute(input, output []byte) {
  for i := 0; i < len(input); i++ {
    output[i] = aesSbox[input[i]]
  }
}

func (AESCipher) shiftRows(input, output []byte) {
  output[ 0] = input[ 0]
  output[ 5] = input[ 1]
  output[10] = input[ 2]
  output[15] = input[ 3]
  output[ 4] = input[ 4]
  output[ 9] = input[ 5]
  output[14] = input[ 6]
  output[ 3] = input[ 7]
  output[ 8] = input[ 8]
  output[13] = input[ 9]
  output[ 2] = input[10]
  output[ 7] = input[11]
  output[12] = input[12]
  output[ 1] = input[13]
  output[ 6] = input[14]
  output[11] = input[15]
}

func (AESCipher) mixColumn(input, output []byte) {
  add := aesMixColAdd
  mul := aesMixColMul
  for i := 0; i < 16; i += 4 {
    output[i+0] = add[add[add[mul[2][input[i+0]]][mul[3][input[i+1]]]][mul[1][input[i+2]]]][mul[1][input[i+3]]]
    output[i+1] = add[add[add[mul[1][input[i+0]]][mul[2][input[i+1]]]][mul[3][input[i+2]]]][mul[1][input[i+3]]]
    output[i+2] = add[add[add[mul[1][input[i+0]]][mul[1][input[i+1]]]][mul[2][input[i+2]]]][mul[3][input[i+3]]]
    output[i+3] = add[add[add[mul[3][input[i+0]]][mul[1][input[i+1]]]][mul[1][input[i+2]]]][mul[2][input[i+3]]]
  }
}
