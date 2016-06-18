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
import "testing"

/* -------------------------------------------------------------------------- */

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

  bits1.Set(10)
  bits1.Clr(17)

  for i := 0; i < len(bits1); i++ {
    if bits1[i] != bits2[i] {
      t.Error("set/clr bits failed")
    }
  }
}

func TestBitsSwap(t *testing.T) {

  bits1 := Bits{}.Read("01001011 00000001 01011000 11110001 11000001 01111000 00111001 01010011")
  bits2 := Bits{}.Read("01001011 00000001 01011000 11101001 11000001 01111000 00111001 01010011")

  bits1.Swap(27, 28)

  for i := 0; i < len(bits1); i++ {
    if bits1[i] != bits2[i] {
      t.Error("swapping bits failed")
    }
  }
}

/* -------------------------------------------------------------------------- */

func TestRotateSlice1(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, -2)

  if y[0] != (x[0] >> 2) + (x[1] << 6) {
    t.Error("rotate test failed")
  }
  if y[1] != (x[1] >> 2) + (x[2] << 6) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[2] >> 2) + (x[3] << 6) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[3] >> 2) + (x[4] << 6) {
    t.Error("rotate test failed")
  }
  if y[4] != (x[4] >> 2) + (x[0] << 6) {
    t.Error("rotate test failed")
  }

}

func TestRotateSlice2(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, -10)

  if y[4] != (x[0] >> 2) + (x[1] << 6) {
    t.Error("rotate test failed")
  }
  if y[0] != (x[1] >> 2) + (x[2] << 6) {
    t.Error("rotate test failed")
  }
  if y[1] != (x[2] >> 2) + (x[3] << 6) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[3] >> 2) + (x[4] << 6) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[4] >> 2) + (x[0] << 6) {
    t.Error("rotate test failed")
  }

}

func TestRotateSlice3(t *testing.T) {

  x := []byte{122,10,1,33,4}
  y := make([]byte, len(x))

  Bits(y).Rotate(x, 13)

  if y[1] != (x[0] << 5) + (x[4] >> 3) {
    t.Error("rotate test failed")
  }
  if y[2] != (x[1] << 5) + (x[0] >> 3) {
    t.Error("rotate test failed")
  }
  if y[3] != (x[2] << 5) + (x[1] >> 3) {
    t.Error("rotate test failed")
  }
  if y[4] != (x[3] << 5) + (x[2] >> 3) {
    t.Error("rotate test failed")
  }
  if y[0] != (x[4] << 5) + (x[3] >> 3) {
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
    1,  2, 10,  4,  5,  6,  7,  8,
    9,  3, 11, 12, 13, 14, 15, 16}
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

  input  := []byte{0,0,1,1<<7}
  output := []byte{0,0,0,0,0,0}

  Bits(output).Map(input, table)

  if output[0] != 1 {
    t.Error("bitmap test failed")
  }
  if output[1] != 0 {
    t.Error("bitmap test failed")
  }
  if output[2] != (1<<7) {
    t.Error("bitmap test failed")
  }
  if output[3] != (1<<1) {
    t.Error("bitmap test failed")
  }
  if output[4] != 0 {
    t.Error("bitmap test failed")
  }
  if output[5] != (1<<6) {
    t.Error("bitmap test failed")
  }
  // fmt.Println("input[0]:", strconv.FormatInt(int64(input[0]), 2))
  // fmt.Println("input[1]:", strconv.FormatInt(int64(input[1]), 2))
  // fmt.Println("input[2]:", strconv.FormatInt(int64(input[2]), 2))
  // fmt.Println("input[3]:", strconv.FormatInt(int64(input[3]), 2))

  // fmt.Println("output[0]:", strconv.FormatInt(int64(output[0]), 2))
  // fmt.Println("output[1]:", strconv.FormatInt(int64(output[1]), 2))
  // fmt.Println("output[2]:", strconv.FormatInt(int64(output[2]), 2))
  // fmt.Println("output[3]:", strconv.FormatInt(int64(output[3]), 2))
  // fmt.Println("output[4]:", strconv.FormatInt(int64(output[4]), 2))
  // fmt.Println("output[5]:", strconv.FormatInt(int64(output[5]), 2))
}

func TestMapInjective(t *testing.T) {
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
