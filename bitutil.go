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

// compute element-wise xor: z = x (+) y
func xorSlice(x, y, z []byte) {
  for i := 0; i < len(x); i++ {
    z[i] = x[i] ^ y[i]
  }
}

/* -------------------------------------------------------------------------- */

func rotateSliceLeft(x, y []byte, n uint) {
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

func rotateSliceRight(x, y []byte, n uint) {
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

func rotateSlice(x, y []byte, n int) {
  if n < 0 {
    rotateSliceRight(x, y, uint(-n))
  } else {
    rotateSliceLeft(x, y, uint(n))
  }
}
