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

package gocrypt

/* -------------------------------------------------------------------------- */

//import "fmt"

/* -------------------------------------------------------------------------- */

type ByteMatrix []byte
type ByteVector   byte

/* -------------------------------------------------------------------------- */

func NewByteVector(a byte) ByteVector {
  return ByteVector(a)
}

func NewByteMatrix(a []byte) ByteMatrix {
  if len(a) != 8 {
    panic("NewByteMatrix(): invalid argument")
  }
  return ByteMatrix(a)
}

/* -------------------------------------------------------------------------- */

func ByteMmulV(a ByteMatrix, b ByteVector) byte {
  r := byte(0.0)
  for i := 0; i < 8; i++ {
    s := byte(0)
    for j := 0; j < 8; j++ {
      if a[i] & (1 << byte(j)) != 0 && b & (1 << byte(j)) != 0 {
        s = (s + 1) % 2
      }
    }
    r |= (s << byte(7-i))
  }
  return r
}

func ByteVaddV(a ByteVector, b ByteVector) byte {
  r := byte(0.0)
  for i := 0; i < 8; i++ {
    if a & (1 << byte(i)) != 0 && b & (1 << byte(i)) != 0 {
      r |= (1 << byte(7-i))
    }
  }
  return r
}
