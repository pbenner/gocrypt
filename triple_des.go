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

type TripleDESCipher struct {
  cipher1 DESCipher
  cipher2 DESCipher
  cipher3 DESCipher
}

/* -------------------------------------------------------------------------- */

func NewTripleDESCipher(key Key) (*TripleDESCipher, error) {
  if len(key) != 24 {
    return nil, fmt.Errorf("NewDESCipher(): invalid key length")
  }
  result := TripleDESCipher{}
  p1, _ := NewDESCipher(key[ 0: 8])
  p2, _ := NewDESCipher(key[ 8:16])
  p3, _ := NewDESCipher(key[16:24])
  result.cipher1 = *p1
  result.cipher2 = *p2
  result.cipher3 = *p3
  return &result, nil
}

/* -------------------------------------------------------------------------- */

func (cipher TripleDESCipher) GetBlockLength() int {
  return cipher.cipher1.GetBlockLength()
}

func (cipher TripleDESCipher) Encrypt(input, output []byte) error {
  err := cipher.cipher1.Encrypt(input,  output)
  if err != nil { return err }
  err  = cipher.cipher2.Decrypt(output, output)
  if err != nil { return err }
  err  = cipher.cipher3.Encrypt(output, output)
  return err
}

func (cipher TripleDESCipher) Decrypt(input, output []byte) error {
  err := cipher.cipher3.Decrypt(input,  output)
  if err != nil { return err }
  err  = cipher.cipher2.Encrypt(output, output)
  if err != nil { return err }
  err  = cipher.cipher1.Decrypt(output, output)
  return err
}
