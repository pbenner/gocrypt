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
import "math/rand"

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
}

func NewCBCCipher(cipher BlockCipher) *CBCCipher {
  return &CBCCipher{cipher}
}

func (cipher CBCCipher) Encrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("CBCCipher.Encrypt(): invalid input length")
  }
  if len(input)+bl != len(output) {
    return fmt.Errorf("CBCCipher.Encrypt(): invalid output length")
  }
  // the first part of the message contains the IV
  iv := output[0:bl]
  // draw a new IV
  for i := 0; i < bl; i++ {
    iv[i] = byte(rand.Int())
  }
  for i := 0; i < len(input); i += bl {
    Bits(output[i+bl:i+2*bl]).Xor(output[i:i+bl], input[i:i+bl])
    cipher.BlockCipher.Encrypt(output[i+bl:i+2*bl], output[i+bl:i+2*bl])
  }
  return nil
}

func (cipher CBCCipher) Decrypt(input, output []byte) error {
  bl := cipher.GetBlockLength()
  // check arguments
  if len(input) % bl != 0 {
    return fmt.Errorf("CBCCipher.Decrypt(): invalid input length")
  }
  if len(input)-bl != len(output) {
    return fmt.Errorf("CBCCipher.Decrypt(): invalid output length")
  }
  for i := bl; i < len(input)-bl+1; i += bl {
    cipher.BlockCipher.Decrypt(input[i:i+bl], output[i-bl:i])
    Bits(output[i-bl:i]).Xor(output[i-bl:i], input[i-bl:i])
  }
  return nil
}
