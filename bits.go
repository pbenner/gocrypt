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

import "fmt"
import "bufio"
import "bytes"
import "errors"

/* -------------------------------------------------------------------------- */

type Bits []byte 

/* -------------------------------------------------------------------------- */

func (bits Bits) String() string {

  buffer := new(bytes.Buffer)
  writer := bufio.NewWriter(buffer)

  for i := 0; i < len(bits); i++ {
    if i != 0 {
      fmt.Fprintf(writer, " ")
    }
    for j := 0; j < 8; j++ {
      if bits[i] & (1 << uint(8-1-j)) != 0 {
        fmt.Fprintf(writer, "1")
      } else {
        fmt.Fprintf(writer, "0")
      }
    }
  }
  writer.Flush()

  return buffer.String()
}

func (Bits) Read(str string) Bits {
  r := []byte{}
  // loop over string and convert each bit
  for i := 0; i < len(str); i += 9 {
    b := byte(0)
    for j := 0; j < 8 && i+j < len(str); j++ {
      k := i+j
      if str[k] == '1' {
        b |= 1 << byte(8-1-j)
      } else if str[k] != '0' {
        // error, invalid character
        return nil
      }
    }
    if i+8 < len(str) && str[i+8] != ' ' {
      // error, bytes need to be separated by space
      return nil
    }
    r = append(r, b)
  }
  return r
}

/* -------------------------------------------------------------------------- */

func (y Bits) RotateLeft(x []byte, n uint) {
  var tmp1, tmp2 byte
  l := len(x)
  m := int((n/8) % uint(l))
  k := n % 8
  for i := 0; i < l; i++ {
    y[(i+int(l-m))%l] = x[i]
  }
  for i := l-1; i >= 0; i-- {
    tmp1 = (y[i] << k) | tmp2
    tmp2 = (y[i] >> (8-k))
    y[i] = tmp1
  }
  y[l-1] = y[l-1] + tmp2
}

func (y Bits) RotateRight(x []byte, n uint) {
  var tmp1, tmp2 byte
  l := len(x)
  m := int((n/8) % uint(l))
  k := n % 8
  for i := 0; i < l; i++ {
    y[(i+int(m))%l] = x[i]
  }
  for i := 0; i < l; i++ {
    tmp1 = (y[i] >> k) | tmp2
    tmp2 = (y[i] << (8-k))
    y[i] = tmp1
  }
  y[0] = y[0] + tmp2
}

func (y Bits) Rotate(x []byte, n int) {
  if n < 0 {
    y.RotateRight(x, uint(-n))
  } else {
    y.RotateLeft(x, uint(n))
  }
}

/* -------------------------------------------------------------------------- */

func (y Bits) Equals(x []byte) bool {
  for i := 0; i < len(x); i++ {
    if x[i] != y[i] {
      return false
    }
  }
  return true
}

// compute element-wise xor: z = x (+) y
func (z Bits) Xor(x, y []byte) {
  for i := 0; i < len(x); i++ {
    z[i] = x[i] ^ y[i]
  }
}

func (y Bits) Reverse(x []byte) {
  for i := 0; i < len(x); i++ {
    y[i] = x[i]
    y[i] = (y[i] & 0xF0) >> 4 | (y[i] & 0x0F) << 4
    y[i] = (y[i] & 0xCC) >> 2 | (y[i] & 0x33) << 2
    y[i] = (y[i] & 0xAA) >> 1 | (y[i] & 0x55) << 1
  }
  for i := 0; i < len(y)/2; i++ {
    y[i], y[len(y)-i-1] = y[len(y)-i-1], y[i]
  }
}

func (y Bits) ReverseEndian(x []byte) {
  for i := 0; i < len(y)/2; i++ {
    y[i], y[len(y)-i-1] = x[len(y)-i-1], x[i]
  }
}

func (x Bits) Clear() {
  for i := 0; i < len(x); i++ {
    x[i] = 0
  }
}

func (x Bits) Set(i int) {
  x[i/8] |= 1 << uint(7-i%8)
}

func (x Bits) Clr(i int) {
  x[i/8] &= ^(1 << uint(7-i%8))
}

func (x Bits) Swap(i, j int) {
  b1 := x[i/8] & (1 << byte(7-i%8))
  b2 := x[j/8] & (1 << byte(7-j%8))
  if b1 != 0 {
    x.Set(j)
  } else {
    x.Clr(j)
  }
  if b2 != 0 {
    x.Set(i)
  } else {
    x.Clr(i)
  }
}

/* -------------------------------------------------------------------------- */

func reduceTableInjective(table [][]int, result []int) error {
  for i := 0; i < len(result); i++ {
    result[i] = -1
  }
  for i := 0; i < len(table); i++ {
    // i: input bit
    // j: output bit
    for k := 0; k < len(table[i]); k++ {
      j := table[i][k]-1
      if result[j] == -1 {
        result[j] = i+1
      } else {
        return errors.New("table cannot be converted")
      }
    }
  }
  return nil
}

// Surjective mapping of input bits to output bits. The mapping
// is defined by the table. The ith bit in the input slice is
// mapped to position j = table[i] in the output.
func (output Bits) MapSurjective(input []byte, table []int) {
  // number of input bits
  n := 8*len(input)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over input bits
  for i := 0; i < n; i++ {
    // index of the output bit
    j := table[i]-1
    if input[i/8]  & byte(1 << byte(7 - (i % 8))) != 0 {
      output[j/8] |= byte(1 << byte(7 - (j % 8)))
    }
  }
}

// Injective mapping of input bits to output bits. The mapping
// is defined by the table. The jth bit in the output slice is
// copied from position i = table[j] in the input.
func (output Bits) MapInjective(input []byte, table []int) {
  // number of input bits
  n := 8*len(output)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over output bits
  for j := 0; j < n; j++ {
    // index of the input bit
    i := table[j]-1
    if input[i/8]  & byte(1 << byte(7 - (i % 8))) != 0 {
      output[j/8] |= byte(1 << byte(7 - (j % 8)))
    }
  }
}

func (output Bits) Map(input []byte, table [][]int) {
  // number of input bits
  n := 8*len(input)
  // check if table is long enough
  if len(table) != n {
    panic("table has invalid length")
  }
  // loop over input bits
  for i := 0; i < n; i++ {
    for k := 0; k < len(table[i]); k++ {
      // index of the output bit
      j := table[i][k]-1
      if input[i/8]  & byte(1 << byte(7 - (i % 8))) != 0 {
        output[j/8] |= byte(1 << byte(7 - (j % 8)))
      }
    }
  }
}
