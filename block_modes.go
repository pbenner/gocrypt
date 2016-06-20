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

package gocrypto

/* -------------------------------------------------------------------------- */

import "fmt"

/* -------------------------------------------------------------------------- */

type ECBCipher struct {
  BlockCipher
}

func NewECBCipher(cipher BlockCipher) *ECBCipher {
  return &ECBCipher{cipher}
}

func (cipher ECBCipher) Encrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("ECBCipher.Encrypt(): invalid input length")
  }
  if len(input) != len(output) {
    return fmt.Errorf("ECBCipher.Encrypt(): invalid output length")
  }
  for i := 0; i < len(input)-bl+1; i += bl {
    cipher.BlockCipher.Encrypt(input[i:i+bl], output[i:i+bl])
  }
  return nil
}

func (cipher ECBCipher) Decrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("ECBCipher.Decrypt(): invalid input length")
  }
  if len(input) != len(output) {
    return fmt.Errorf("ECBCipher.Decrypt(): invalid output length")
  }
  for i := 0; i < len(input)-bl+1; i += bl {
    cipher.BlockCipher.Decrypt(input[i:i+bl], output[i:i+bl])
  }
  return nil
}

/* -------------------------------------------------------------------------- */

type CBCCipher struct {
  BlockCipher
  iv []byte
}

func NewCBCCipher(cipher BlockCipher, iv []byte) (*CBCCipher, error) {
  if len(iv) != cipher.GetBlockLength() {
    return nil, fmt.Errorf("NewCBCCipher(): iv has invalid length")
  }
  return &CBCCipher{cipher, iv}, nil
}

func (cipher CBCCipher) Encrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("CBCCipher.Encrypt(): invalid input length")
  }
  if len(input) != len(output) {
    return fmt.Errorf("CBCCipher.Encrypt(): invalid output length")
  }
  if len(input) == 0 {
    return nil
  }
  // encrypt first block using IV
  Bits(output[0:bl]).Xor(cipher.iv, input[0:0+bl])
  cipher.BlockCipher.Encrypt(output[0:0+bl], output[0:0+bl])
  // loop over remaining blocks
  for i := bl; i < len(input)-bl+1; i += bl {
    Bits(output[i:i+bl]).Xor(output[i-bl:i], input[i:i+bl])
    cipher.BlockCipher.Encrypt(output[i:i+bl], output[i:i+bl])
  }
  return nil
}

func (cipher CBCCipher) Decrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("CBCCipher.Decrypt(): invalid input length")
  }
  if len(input) != len(output) {
    return fmt.Errorf("CBCCipher.Decrypt(): invalid output length")
  }
  if len(input) == 0 {
    return nil
  }
  // decrypt first block using IV
  cipher.BlockCipher.Decrypt(input[0:0+bl], output[0:0+bl])
  Bits(output[0:0+bl]).Xor(output[0:0+bl], cipher.iv)
  // loop over remaining blocks
  for i := bl; i < len(input)-bl+1; i += bl {
    cipher.BlockCipher.Decrypt(input[i:i+bl], output[i:i+bl])
    Bits(output[i:i+bl]).Xor(output[i:i+bl], input[i-bl:i])
  }
  return nil
}
