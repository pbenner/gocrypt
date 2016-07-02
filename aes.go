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

import "fmt"

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

func (AESCipher) substituteInv(input, output []byte) {
  for i := 0; i < len(input); i++ {
    output[i] = aesSboxInv[input[i]]
  }
}

func (AESCipher) shiftRows(input, output []byte) {
  output[ 0] = input[ 0]
  output[13] = input[ 1]
  output[10] = input[ 2]
  output[ 7] = input[ 3]
  output[ 4] = input[ 4]
  output[ 1] = input[ 5]
  output[14] = input[ 6]
  output[11] = input[ 7]
  output[ 8] = input[ 8]
  output[ 5] = input[ 9]
  output[ 2] = input[10]
  output[15] = input[11]
  output[12] = input[12]
  output[ 9] = input[13]
  output[ 6] = input[14]
  output[ 3] = input[15]
}

func (AESCipher) shiftRowsInv(input, output []byte) {
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
    output[i+0] = add[add[add[mul[0x02][input[i+0]]][mul[0x03][input[i+1]]]][mul[0x01][input[i+2]]]][mul[0x01][input[i+3]]]
    output[i+1] = add[add[add[mul[0x01][input[i+0]]][mul[0x02][input[i+1]]]][mul[0x03][input[i+2]]]][mul[0x01][input[i+3]]]
    output[i+2] = add[add[add[mul[0x01][input[i+0]]][mul[0x01][input[i+1]]]][mul[0x02][input[i+2]]]][mul[0x03][input[i+3]]]
    output[i+3] = add[add[add[mul[0x03][input[i+0]]][mul[0x01][input[i+1]]]][mul[0x01][input[i+2]]]][mul[0x02][input[i+3]]]
  }
}

func (AESCipher) mixColumnInv(input, output []byte) {
  add := aesMixColAdd
  mul := aesMixColMul
  for i := 0; i < 16; i += 4 {
    output[i+0] = add[add[add[mul[0x0E][input[i+0]]][mul[0x0B][input[i+1]]]][mul[0x0D][input[i+2]]]][mul[0x09][input[i+3]]]
    output[i+1] = add[add[add[mul[0x09][input[i+0]]][mul[0x0E][input[i+1]]]][mul[0x0B][input[i+2]]]][mul[0x0D][input[i+3]]]
    output[i+2] = add[add[add[mul[0x0D][input[i+0]]][mul[0x09][input[i+1]]]][mul[0x0E][input[i+2]]]][mul[0x0B][input[i+3]]]
    output[i+3] = add[add[add[mul[0x0B][input[i+0]]][mul[0x0D][input[i+1]]]][mul[0x09][input[i+2]]]][mul[0x0E][input[i+3]]]
  }
}

/* -------------------------------------------------------------------------- */

func NewAESCipher(key []byte) (*AESCipher, error) {
  cipher := AESCipher{}
  cipher.BlockLength = 16
  if err := cipher.GenerateSubkeys(key); err != nil {
    return nil, err
  }
  return &cipher, nil
}

/* -------------------------------------------------------------------------- */

func (cipher AESCipher) Encrypt(input, output []byte) error {
  if len(input) != cipher.BlockLength {
    return fmt.Errorf("AESCipher.Encrypt(): invalid input length")
  }
  if len(output) != cipher.BlockLength {
    return fmt.Errorf("AESCipher.Encrypt(): invalid output length")
  }
  tmp := make([]byte, cipher.BlockLength)
  // xor input
  Bits(output).Xor(input, cipher.Keys[0])
  // apply rounds
  for i := 1; i < len(cipher.Keys); i++ {
    cipher.substitute(output, tmp)
    cipher.shiftRows (tmp, output)
    if i == len(cipher.Keys)-1 {
      Bits(output).Xor(output, cipher.Keys[i])
    } else {
      cipher.mixColumn (output, tmp)
      Bits(output).Xor(tmp, cipher.Keys[i])
    }
  }
  return nil
}

func (cipher AESCipher) Decrypt(input, output []byte) error {
  if len(input) != cipher.BlockLength {
    return fmt.Errorf("AESCipher.Dencrypt(): invalid input length")
  }
  if len(output) != cipher.BlockLength {
    return fmt.Errorf("AESCipher.Dencrypt(): invalid output length")
  }
  tmp := make([]byte, cipher.BlockLength)
  // copy input to output
  copy(output, input)
  // apply rounds
  for i := 0; i < len(cipher.Keys)-1; i++ {
    if i == 0 {
      Bits(output).Xor(output, cipher.Keys[len(cipher.Keys)-i-1])
    } else {
      Bits(tmp).Xor(output, cipher.Keys[len(cipher.Keys)-i-1])
      cipher.mixColumnInv(tmp, output)
    }
    cipher.shiftRowsInv (output, tmp)
    cipher.substituteInv(tmp, output)
  }
  Bits(output).Xor(output, cipher.Keys[0])
  return nil
}

func (cipher *AESCipher) GenerateSubkeys(key []byte) error {
  switch len(key) {
  case 16: cipher.subkeys(key, 10, false)
  case 24: cipher.subkeys(key, 12, false)
  case 32: cipher.subkeys(key, 14, true)
  default:
    return fmt.Errorf("AESCipher.GenerateSubkeys(): invalid key length")
  }
  return nil
}
