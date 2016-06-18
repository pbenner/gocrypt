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
import "strings"
import "unicode"

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
      if bits[i] & (1 << uint(j)) != 0 {
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
  // strip all whitespace
  str = strings.Map(func(r rune) rune {
    if unicode.IsSpace(r) {
    return -1
  }
    return r
  }, str)
  // allocate memory
  r := make([]byte, divIntUp(len(str), 8))
  // loope over string and convert each bit
  for i := 0; i < len(str); i++ {
    if str[i] == '1' {
      r[i/8] |= 1 << uint(i%8)
    }
  }
  return r
}

/* -------------------------------------------------------------------------- */

func (y Bits) RotateLeft(x []byte, n uint) {
  var tmp1, tmp2 byte
  l := len(x)
  m := (n/8) % uint(l)
  k := n % 8
  for i := 0; i < l; i++ {
    y[(i+int(m))%l] = x[i]
  }
  for i := 0; i < l; i++ {
    tmp1 = (y[i] << k) | tmp2
    tmp2 = (y[i] >> (8-k))
    y[i] = tmp1
  }
  y[0] = y[0] + tmp2
}

func (y Bits) RotateRight(x []byte, n uint) {
  var tmp1, tmp2 byte
  l := len(x)
  m := int((n/8) % uint(l))
  k := n % 8
  for i := 0; i < l; i++ {
    y[(i+int(l-m))%l] = x[i]
  }
  for i := l-1; i >= 0; i-- {
    tmp1 = (y[i] >> k) | tmp2
    tmp2 = (y[i] << (8-k))
    y[i] = tmp1
  }
  y[l-1] = y[l-1] + tmp2
}

func (y Bits) Rotate(x []byte, n int) {
  if n < 0 {
    y.RotateRight(x, uint(-n))
  } else {
    y.RotateLeft(x, uint(n))
  }
}

/* -------------------------------------------------------------------------- */

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

func (x Bits) Clear() {
  for i := 0; i < len(x); i++ {
    x[i] = 0
  }
}
