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
