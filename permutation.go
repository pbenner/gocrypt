/* Copyright (C) 2015 Philipp Benner
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

import "math"
import "math/rand"

/* -------------------------------------------------------------------------- */

type PermutationCipher struct
{
  forwardKey Key
  reverseKey Key
}

func NewPermutationCipher() PermutationCipher {
  cipher := PermutationCipher{}
  cipher.Generate()
  return cipher
}

/* -------------------------------------------------------------------------- */

func (cipher *PermutationCipher) GetKey() Key {
  return cipher.forwardKey
}

func (cipher *PermutationCipher) Generate(_n ...int) {
  n := int(math.Pow(2, 8))
  cipher.forwardKey = NewKey(n)
  cipher.reverseKey = NewKey(n)
  p := rand.Perm(n)
  for i := 0; i < n; i++ {
    cipher.forwardKey[  i ] = byte(p[i])
    cipher.reverseKey[p[i]] = byte(i)
  }
}

func (cipher *PermutationCipher) Encrypt(m Message) Message {

  result := NewMessage(len(m))
  length := len(cipher.forwardKey)

  if len(cipher.forwardKey) != length {
    panic("No private key available!")
  }

  for i := 0; i < len(m); i++ {
    result[i] = cipher.forwardKey[m[i]]
  }

  return result
}

func (cipher *PermutationCipher) Decrypt(m Message) Message {

  result := NewMessage(len(m))
  length := len(cipher.forwardKey)

  if len(cipher.forwardKey) != length {
    panic("No private key available!")
  }

  for i := 0; i < len(m); i++ {
    result[i] = cipher.reverseKey[m[i]]
  }

  return result
}
