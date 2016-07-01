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

  pf := NewPrimeField(2)

  // irreducible polynomial
  p := NewPolynomial(pf)
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)

  f := NewFiniteField(2, 8, p)

  a := NewPolynomial(pf)
  a.AddTerm(1, 0)

  for i := 0; i <= 0xFF; i++ {
    b := NewPolynomial(pf)
    b.ReadByte(byte(i))

    r := f.Div(a, b)
    y := ByteVaddV(ByteMmulV(m, ByteVector(r.WriteByte())), v)

    if byte(y) != aesSbox[i] {
      t.Error("finite field aes s-box test failed")
    }
  }
}

func TestAES2(t *testing.T) {

  key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
  r   := []byte{0xd0, 0x14, 0xf9, 0xa8, 0xc9, 0xee, 0x25, 0x89, 0xe1, 0x3f, 0x0c, 0xc8, 0xb6, 0x63, 0x0c, 0xa6}

  cipher, _ := NewAESCipher(key)

  if !Bits(cipher.Keys[10]).Equals(r) {
    t.Error("aes keys test failed")
  }
}
