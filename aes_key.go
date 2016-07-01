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

/* -------------------------------------------------------------------------- */

func (cipher AESCipher) subkeys128(key []byte) {
  
}
