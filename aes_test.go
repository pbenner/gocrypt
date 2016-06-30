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
import "testing"

/* -------------------------------------------------------------------------- */

func TestAES1(t *testing.T) {

  m := ByteMatrix(Bits{}.Read("10001111 11000111 11100011 11110001 11111000 01111100 00111110 00011111"))
  v := ByteVector(Bits{}.Read("01100011")[0])

  // irreducible polynomial
  p := NewPolynomial()
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)

  f := NewFiniteField(2, 8, p)

  a := NewPolynomial()
  a.AddTerm(1, 0)

  for i := 0; i <= 0xFF; i++ {
    b := NewPolynomial()
    b.ReadByte(byte(i))

    r := f.Div(a, b)
    y := ByteVaddV(ByteMmulV(m, ByteVector(r.WriteByte())), v)

    if byte(y) != aesSbox[i] {
      t.Error("finite field aes s-box test failed")
    }
  }
}
