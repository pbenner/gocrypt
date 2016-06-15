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
//import "strconv"
import "testing"

/* -------------------------------------------------------------------------- */

func TestDESsBox(t *testing.T) {
  //  input  100111 010100 100110 100001 111010 011110 010001 011000
  //  input  10011101 01001001 10100001 11101001 11100100 01011000
  // output: 0111 1001 0101 1011 0010 1000 1100 0101
  input := []byte{
    0x58,  // 01011000
    0xE4,  // 11100100
    0xE9,  // 11101001
    0xA1,  // 10100001
    0x49,  // 01001001
    0x9D } // 10011101

  output := make([]byte, 32/8)

  desSbox(input, output)

  if output[0] != 0x5 | (0xC << 4) {
    t.Error("DES s-box test failed")
  }
  if output[1] != 0x8 | (0x2 << 4) {
    t.Error("DES s-box test failed")
  }
  if output[2] != 0xB | (0x5 << 4) {
    t.Error("DES s-box test failed")
  }
  if output[3] != 0x9 | ( 0x7 << 4) {
    t.Error("DES s-box test failed")
  }

}

func TestDESkeyShift(t *testing.T) {
  // 00111011 00111000 10011000 00110111 00010101 00100000 11110111 01011110
  key1 := []byte{0x5e,0xf7,0x20,0x15,0x37,0x98,0x38,0x3b}
  key2 := make([]byte, 8)
  // 01000100 11000000 01101011 11011100 10011101 10001000 01111111
  tmp1 := make([]byte, 56/8)


  ReverseBits(key1, key2)
  BitmapInjective(key2, tmp1, desKeyPC1)

  fmt.Println(Key(key2).BinarySequence())
  fmt.Println(Key(tmp1).BinarySequence())
  desSplitRotateKey(tmp1)
  fmt.Println(Key(tmp1).BinarySequence())
}
