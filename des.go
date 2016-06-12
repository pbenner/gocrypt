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

//import "fmt"

/* -------------------------------------------------------------------------- */

type DESCipher struct {
  FeistelNetwork
}

/* -------------------------------------------------------------------------- */

func init() {
  // convert s-boxes to simplify indexing
  desFsbox1 = convertSbox6to4(desFsbox1)
  desFsbox2 = convertSbox6to4(desFsbox2)
  desFsbox3 = convertSbox6to4(desFsbox3)
  desFsbox4 = convertSbox6to4(desFsbox4)
  desFsbox5 = convertSbox6to4(desFsbox5)
  desFsbox6 = convertSbox6to4(desFsbox6)
  desFsbox7 = convertSbox6to4(desFsbox7)
  desFsbox8 = convertSbox6to4(desFsbox8)
}

/* -------------------------------------------------------------------------- */

/* initial permutation IP */
var desIP = []int{
  58, 50, 42, 34, 26, 18, 10,  2, 60, 52, 44, 36, 28, 20, 12,  4,
  62, 54, 46, 38, 30, 22, 14,  6, 64, 56, 48, 40, 32, 24, 16,  8,
  57, 49, 41, 33, 25, 17,  9,  1, 59, 51, 43, 35, 27, 19, 11,  3,
  61, 53, 45, 37, 29, 21, 13,  5, 63, 55, 47, 39, 31, 23, 15,  7 }
 
/* final permutation IP^-1 */
var desFP = []int{
  40,  8, 48, 16, 56, 24, 64, 32, 39,  7, 47, 15, 55, 23, 63, 31,
  38,  6, 46, 14, 54, 22, 62, 30, 37,  5, 45, 13, 53, 21, 61, 29,
  36,  4, 44, 12, 52, 20, 60, 28, 35,  3, 43, 11, 51, 19, 59, 27,
  34,  2, 42, 10, 50, 18, 58, 26, 33,  1, 41,  9, 49, 17, 57, 25 }

/* expansion map */
var desFexpansion = []int{
  32,  1,  2,  3,  4,  5,  4,  5,  6,  7,  8,  9,  8,  9, 10, 11,
  12, 13, 12, 13, 14, 15, 16, 17, 16, 17, 18, 19, 20, 21, 20, 21,
  22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32,  1 }

var desFsbox1 = []byte{
  14,  4, 13,  1,  2, 15, 11,  8,  3, 10,  6, 12,  5,  9,  0,  7,
   0, 15,  7,  4, 14,  2, 13,  1, 10,  6, 12, 11,  9,  5,  3,  8,
   4,  1, 14,  8, 13,  6,  2, 11, 15, 12,  9,  7,  3, 10,  5,  0,
  15, 12,  8,  2,  4,  9,  1,  7,  5, 11,  3, 14, 10,  0,  6, 13}

var desFsbox2 = []byte{
  15,  1,  8, 14,  6, 11,  3,  4,  9,  7,  2, 13, 12,  0,  5, 10,
   3, 13,  4,  7, 15,  2,  8, 14, 12,  0,  1, 10,  6,  9, 11,  5,
   0, 14,  7, 11, 10,  4, 13,  1,  5,  8, 12,  6,  9,  3,  2, 15,
  13,  8, 10,  1,  3, 15,  4,  2, 11,  6,  7, 12,  0,  5, 14,  9 }

var desFsbox3 = []byte{
  10,  0,  9, 14,  6,  3, 15,  5,  1, 13, 12,  7, 11,  4,  2,  8,
  13,  7,  0,  9,  3,  4,  6, 10,  2,  8,  5, 14, 12, 11, 15,  1,
  13,  6,  4,  9,  8, 15,  3,  0, 11,  1,  2, 12,  5, 10, 14,  7,
   1, 10, 13,  0,  6,  9,  8,  7,  4, 15, 14,  3, 11,  5,  2, 12 }

var desFsbox4 = []byte{
   7, 13, 14,  3,  0,  6,  9, 10,  1,  2,  8,  5, 11, 12,  4, 15,
  13,  8, 11,  5,  6, 15,  0,  3,  4,  7,  2, 12,  1, 10, 14,  9,
  10,  6,  9,  0, 12, 11,  7, 13, 15,  1,  3, 14,  5,  2,  8,  4,
   3, 15,  0,  6, 10,  1, 13,  8,  9,  4,  5, 11, 12,  7,  2, 14 }

var desFsbox5 = []byte{
   2, 12,  4,  1,  7, 10, 11,  6,  8,  5,  3, 15, 13,  0, 14,  9,
  14, 11,  2, 12,  4,  7, 13,  1,  5,  0, 15, 10,  3,  9,  8,  6,
   4,  2,  1, 11, 10, 13,  7,  8, 15,  9, 12,  5,  6,  3,  0, 14,
  11,  8, 12,  7,  1, 14,  2, 13,  6, 15,  0,  9, 10,  4,  5,  3 }

var desFsbox6 = []byte{
  12,  1, 10, 15,  9,  2,  6,  8,  0, 13,  3,  4, 14,  7,  5, 11,
  10, 15,  4,  2,  7, 12,  9,  5,  6,  1, 13, 14,  0, 11,  3,  8,
   9, 14, 15,  5,  2,  8, 12,  3,  7,  0,  4, 10,  1, 13, 11,  6,
   4,  3,  2, 12,  9,  5, 15, 10, 11, 14,  1,  7,  6,  0,  8, 13 }

var desFsbox7 = []byte{
   4, 11,  2, 14, 15,  0,  8, 13,  3, 12,  9,  7,  5, 10,  6,  1,
  13,  0, 11,  7,  4,  9,  1, 10, 14,  3,  5, 12,  2, 15,  8,  6,
   1,  4, 11, 13, 12,  3,  7, 14, 10, 15,  6,  8,  0,  5,  9,  2,
   6, 11, 13,  8,  1,  4, 10,  7,  9,  5,  0, 15, 14,  2,  3, 12 }

var desFsbox8 = []byte{
  13,  2,  8,  4,  6, 15, 11,  1, 10,  9,  3, 14,  5,  0, 12,  7,
   1, 15, 13,  8, 10,  3,  7,  4, 12,  5,  6, 11,  0, 14,  9,  2,
   7, 11,  4,  1,  9, 12, 14,  2,  0,  6, 10, 13, 15,  3,  5,  8,
   2,  1, 14,  7,  4, 10,  8, 13, 15, 12,  9,  0,  3,  5,  6, 11 }

var desFsboxP = []int{
  16,  7, 20, 21, 29, 12, 28, 17,  1, 15, 23, 26,  5, 18, 31, 10,
   2,  8, 24, 14, 32, 27,  3,  9, 19, 13, 30,  6, 22, 11,  4, 25 }

/* -------------------------------------------------------------------------- */

// Shuffle Sbox entries such that the indexing with outer
// bits for rows and middle bits for columns is converted to a
// standard lookup table.
func convertSbox6to4(input []byte) []byte {
  output := make([]byte, len(input))

  for i := 0; i < 64; i++ {
    row := (i & 1) | ((i>>4) & 2)
    col := (i >> 1) & 0xF
    k := row*16 + col
    output[i] = input[k]
  }
  return output
}

func desFeistelFunctionSbox(input, output []byte) {
  i1 := input[0] & 0x3F
  i2 := (input[0] >> 6) | ((input[1] & 0xF) << 2)
  i3 := (input[1] >> 4) | ((input[2] & 0x3) << 0)
  i4 := (input[2] >> 2)
  i5 := input[3] & 0x3F
  i6 := (input[3] >> 6) | ((input[4] & 0xF) << 2)
  i7 := (input[4] >> 4) | ((input[5] & 0x3) << 0)
  i8 := (input[5] >> 2)
  o1 := desFsbox1[i1]
  o2 := desFsbox2[i2]
  o3 := desFsbox3[i3]
  o4 := desFsbox4[i4]
  o5 := desFsbox5[i5]
  o6 := desFsbox6[i6]
  o7 := desFsbox7[i7]
  o8 := desFsbox8[i8]
  output[0] = o1 + (o2 << 4)
  output[1] = o3 + (o4 << 4)
  output[2] = o5 + (o6 << 4)
  output[3] = o7 + (o8 << 4)
}

func desFeistelFunction(input, output, key []byte) {
  tmp1 := make([]byte, 48/8)
  tmp2 := make([]byte, 32/8)
  BitmapInjective(input, tmp1, desFexpansion)
  xorSlice(tmp1, key, tmp1)
  desFeistelFunctionSbox(tmp1, tmp2)
  BitmapInjective(tmp2, output, desFsboxP)
}

/* -------------------------------------------------------------------------- */

func NewDESCipher() {
}
