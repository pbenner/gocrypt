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

package main

/* -------------------------------------------------------------------------- */

import   "fmt"
import   "math"
import   "math/rand"
import   "time"

import . "github.com/pbenner/autodiff"
import . "github.com/pbenner/0100101101"

/* -------------------------------------------------------------------------- */

func proposal(alphabet AsciiAlphabet, r *rand.Rand) (int, int) {
  i := alphabet.Geti()
  j := alphabet.Getj()
  n := j - i + 1
  j1 := i+r.Intn(n)
  j2 := i+r.Intn(n)
  return j1, j2
}

func sampler(n int, alphabet AsciiAlphabet, ciphertext Message, t Matrix, verbose bool) PermutationCipher {

  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)
  l := alphabet.Getj()-alphabet.Geti()+1

  cipher := NewAsciiPermutationCipher(alphabet)

  for i := 0; i < n*l; i++ {
    j1, j2 := proposal(alphabet, r)
    text1 := cipher.Decrypt(ciphertext)
    cipher.Swap(j1, j2)
    text2 := cipher.Decrypt(ciphertext)

    l1 := likelihood(text1, t)
    l2 := likelihood(text2, t)

    a := Div(l2, l1)
    b := math.Min(1, a.Value())
    c := r.Float64()

    if (c <= b) {
      // accept (do nothing)
    } else {
      // reject (reverse swap)
      cipher.Swap(j1, j2)
    }
    if verbose && i % int(math.Pow(2, 8)) == 0 {
      fmt.Println("message:", text1)
      fmt.Println("likelihood:", l1)
      fmt.Println()
    }
  }
  return cipher
}
