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

package gocrypto

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

func NewAsciiPermutationCipher(alphabet AsciiAlphabet) PermutationCipher {
  cipher := PermutationCipher{}
  cipher.Generate(alphabet)
  return cipher
}

/* -------------------------------------------------------------------------- */

func (cipher *PermutationCipher) GetKey() Key {
  return cipher.forwardKey
}

func (cipher *PermutationCipher) Generate(args ...interface{}) {
  i := 0
  j := 255
  // permutation of the full byte range
  n := int(math.Pow(2, 8))
  // if an alphabet is given as an argument, restrict the
  // permutation range accordingly
  if len(args) == 1 {
    alphabet := args[0].(AsciiAlphabet)
    i = alphabet.i
    j = alphabet.j
    n = j - i + 1
  }
  cipher.forwardKey = NewKey(256)
  cipher.reverseKey = NewKey(256)
  p := rand.Perm(n)
  for k := 0; k < 256; k++ {
    if k >= i && k <= j {
      cipher.forwardKey[k]        = byte(i+p[k-i])
      cipher.reverseKey[i+p[k-i]] = byte(k)
    } else {
      cipher.forwardKey[k] = byte(k)
      cipher.reverseKey[k] = byte(k)
    }
  }
}

func (cipher *PermutationCipher) Encrypt(m Message) Message {

  result := NullMessage(len(m))
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

  result := NullMessage(len(m))
  length := len(cipher.forwardKey)

  if len(cipher.forwardKey) != length {
    panic("No private key available!")
  }

  for i := 0; i < len(m); i++ {
    result[i] = cipher.reverseKey[m[i]]
  }

  return result
}

/* utility
 * -------------------------------------------------------------------------- */

func (cipher *PermutationCipher) Swap(j1, j2 int) {
  k1, k2 := cipher.forwardKey[j1], cipher.forwardKey[j2]
  cipher.forwardKey[j1], cipher.forwardKey[j2] =
    cipher.forwardKey[j2], cipher.forwardKey[j1]
  cipher.reverseKey[k1], cipher.reverseKey[k2] =
    cipher.reverseKey[k2], cipher.reverseKey[k1]
}
