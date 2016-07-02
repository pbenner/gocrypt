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

  k1 := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
  r1 := []byte{0xd0, 0x14, 0xf9, 0xa8, 0xc9, 0xee, 0x25, 0x89, 0xe1, 0x3f, 0x0c, 0xc8, 0xb6, 0x63, 0x0c, 0xa6}

  k2 := []byte{
    0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b,
    0x80, 0x90, 0x79, 0xe5, 0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b }
  r2 := []byte{0xe9, 0x8b, 0xa0, 0x6f, 0x44, 0x8c, 0x77, 0x3c, 0x8e, 0xcc, 0x72, 0x04, 0x01, 0x00, 0x22, 0x02}

  k3 := []byte{
    0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
    0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4 }
  r3 := []byte{0xfe, 0x48, 0x90, 0xd1, 0xe6, 0x18, 0x8d, 0x0b, 0x04, 0x6d, 0xf3, 0x44, 0x70, 0x6c, 0x63, 0x1e}

  c1, _ := NewAESCipher(k1)
  c2, _ := NewAESCipher(k2)
  c3, _ := NewAESCipher(k3)

  if !Bits(c1.Keys[10]).Equals(r1) {
    t.Error("aes keys test failed")
  }
  if !Bits(c2.Keys[12]).Equals(r2) {
    t.Error("aes keys test failed")
  }
  if !Bits(c3.Keys[14]).Equals(r3) {
    t.Error("aes keys test failed")
  }
}
