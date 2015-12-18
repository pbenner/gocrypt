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

import "strings"
import "bytes"

/* -------------------------------------------------------------------------- */

func CodeBlock(s string) string {

  var buffer bytes.Buffer

  reader := strings.NewReader(s)

  for i := 0;; i++ {

    if i != 0 && i % 80 == 0 {
      buffer.WriteByte('\n')
    }
    b, err := reader.ReadByte()
    if err != nil {
      break
    }
    buffer.WriteByte(b)

  }

  return (buffer.String())
}
