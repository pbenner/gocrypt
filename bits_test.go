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

//import "fmt"
import "testing"

/* -------------------------------------------------------------------------- */

func TestBitsString(t *testing.T) {

  str  := "00111011 00111000 10011000 00110111 00010101 00100000 11110111 01011110"
	bits1 := Bits{}.Read(str)
  bits2 := []byte{0x3b, 0x38, 0x98, 0x37, 0x15, 0x20, 0xf7, 0x5e}

  if !bits1.Equals(bits2) {
    t.Error("bits string failed")
  }
}

func TestBitsRead(t *testing.T) {

  str1 := "01001011 00000001 01011000 11110001 11000001 01111000 00111001 01010011"
  str2 := Bits{}.Read(str1).String()

  if str1 != str2 {
    t.Error("reading bits from string failed")
  }
}

func TestBitsSetClr(t *testing.T) {

  bits1 := Bits{}.Read("01001011 00000001 01011000 11110001 11000001 01111000 00111001 01010011")
  bits2 := Bits{}.Read("01001011 00100001 00011000 11110001 11000001 01111000 00111001 01010011")

  bits1.Set(11)
  bits1.Clr(18)

  for i := 0; i < len(bits1); i++ {
    if bits1[i] != bits2[i] {
      t.Error("set/clr bits failed")
    }
  }
}

func TestBitsSwap(t *testing.T) {

  bits1 := Bits{}.Read("01001011 00000001 01011000 11110001 11000001 01111000 00111001 01010011")
  bits2 := Bits{}.Read("01001011 00000001 01011000 11101001 11000001 01111000 00111001 01010011")

  bits1.Swap(28, 29)

  for i := 0; i < len(bits1); i++ {
    if bits1[i] != bits2[i] {
      t.Error("swapping bits failed")
    }
  }
}

/* -------------------------------------------------------------------------- */

func TestRotateSlice1(t *testing.T) {

  x := Bits{}.Read("01111010 00001010 00000001 00100001 00000100")
  y := make([]byte, len(x))
  z := Bits{}.Read("00011110 10000010 10000000 01001000 01000001")

  Bits(y).Rotate(x, -2)

  if !Bits(y).Equals(z) {
    t.Error("rotate test failed")
  }
}

func TestRotateSlice2(t *testing.T) {

  x := Bits{}.Read("01111010 00001010 00000001 00100001 00000100")
  y := make([]byte, len(x))
  z := Bits{}.Read("01000001 00011110 10000010 10000000 01001000")

  Bits(y).Rotate(x, -10)

  if !Bits(y).Equals(z) {
    t.Error("rotate test failed")
  }
}

func TestRotateSlice3(t *testing.T) {

  x := Bits{}.Read("01111010 00001010 00000001 00100001 00000100")
  y := make([]byte, len(x))
  z := Bits{}.Read("01000000 00100100 00100000 10001111 01000001")

  Bits(y).Rotate(x, 13)

  if !Bits(y).Equals(z) {
    t.Error("rotate test failed")
  }

}

func TestBitsReverse(t *testing.T) {
  input  := Bits{}.Read("11111100 10001011 00011011 10010101 01110001 00011101")
  output := Bits(make([]byte, len(input)))
  output.Reverse(input)

  result := Bits{}.Read("10111000 10001110 10101001 11011000 11010001 00111111")

  for i := 0; i < 4; i++ {
    if output[i] != result[i] {
      t.Errorf("reversing bits failed")
    }
  }
}

/* -------------------------------------------------------------------------- */

func TestMapSurjective(t *testing.T) {

  table := []int{
    1,  2, 3,  4,  5,  15,  7,  8,
    9,  6, 11, 12, 13, 14, 15, 16}
  input  := []byte{4,0}
  output := []byte{0,0}

  Bits(output).MapSurjective(input, table)

  if output[0] != 0 {
    t.Error("bitmap test failed")
  }
  if output[1] != 2 {
    t.Error("bitmap test failed")
  }
}

func TestMap(t *testing.T) {

  table := [][]int{
    { 2, 48},
    { 3},
    { 4},
    { 5,  7},
    { 8,  6},
    { 9},
    {10},
    {11, 13},
    {14, 12},
    {15},
    {16},
    {17, 19},
    {20, 18},
    {21},
    {22},
    {23, 25},
    {26, 24},
    {27},
    {28},
    {29, 31},
    {32, 30},
    {33},
    {34},
    {35, 37},
    {38, 36},
    {39},
    {40},
    {41, 43},
    {44, 42},
    {45},
    {46},
    {47,  1} }

  input  := Bits{}.Read("00000000 00000000 00000001 10000000")
  output := Bits{}.Read("00000000 00000000 00000000 00000000 00000000 00000000")
  result := Bits{}.Read("00000000 00000000 00000000 00000000 00111100 00000000")

  Bits(output).Map(input, table)

  if !Bits(output).Equals(result) {
    t.Error("bitmap test failed")
  }
}

func TestMapInjective1(t *testing.T) {
  table1 := [][]int{
    { 2, 48},
    { 3},
    { 4},
    { 5,  7},
    { 8,  6},
    { 9},
    {10},
    {11, 13},
    {14, 12},
    {15},
    {16},
    {17, 19},
    {20, 18},
    {21},
    {22},
    {23, 25},
    {26, 24},
    {27},
    {28},
    {29, 31},
    {32, 30},
    {33},
    {34},
    {35, 37},
    {38, 36},
    {39},
    {40},
    {41, 43},
    {44, 42},
    {45},
    {46},
    {47,  1} }
  table2 := make([]int, 48)
  reduceTableInjective(table1, table2)

  input   := []byte{0,0,1,1<<7}
  output1 := []byte{0,0,0,0,0,0}
  output2 := []byte{0,0,0,0,0,0}

  Bits(output1).Map         (input, table1)
  Bits(output2).MapInjective(input, table2)

  for i := 0; i < len(output1); i++ {
    if output1[i] != output2[i] {
      t.Error("bitmap test failed")
    }
  }
}

func TestMapInjective2(t *testing.T) {

  table := []int{
    2,  3, 4,  5,  6,  7,  8,  1,
    1,  2, 3,  4,  5,  6,  7,  8}
  input  := Bits{}.Read("10101010 00000000")
  output := Bits{}.Read("00000000 00000000")
  result := Bits{}.Read("01010101 10101010")

  Bits(output).MapInjective(input, table)

  if !Bits(output).Equals(result) {
      t.Error("injective bitmap test failed")
  }
}
