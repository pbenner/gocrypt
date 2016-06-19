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

func TestDESsBox(t *testing.T) {
  input  := Bits{}.Read("11111100 10001011 00011011 10010101 01110001 00011101")
  output := make([]byte, 4)
  result := Bits{}.Read("11010110 00111010 11001110 00101001")

  DESCipher{}.Sbox(input, output)

  for i := 0; i < 4; i++ {
    if output[i] != result[i] {
      t.Errorf("DES s-box test %d failed", i)
    }
  }
}

func TestDESkeys(t *testing.T) {
  key := Key(Bits{}.Read("00111011 00111000 10011000 00110111 00010101 00100000 11110111 01011110"))

  result := [][]byte{
    Bits{}.Read("01011100 00001000 01001100 01010101 10001111 01001111"),
    Bits{}.Read("01010001 00101101 11110000 01100100 10010111 11001100"),
    Bits{}.Read("11010100 11100100 10000101 11011000 10110100 11101111"),
    Bits{}.Read("01010011 10000111 00000110 01101110 11011110 10101001"),
    Bits{}.Read("01101000 10010000 10100111 00011010 01111101 01111011"),
    Bits{}.Read("10110001 10000000 01101110 10101111 11011001 00110000"),
    Bits{}.Read("10100000 01000010 10110010 11000001 01101111 01110010"),
    Bits{}.Read("10110100 00011011 00110100 11111101 10001010 00011100"),
    Bits{}.Read("00100010 11011101 01000010 10010011 10000110 01111100"),
    Bits{}.Read("01101000 01100001 01010111 11011001 10111111 10000100"),
    Bits{}.Read("00100101 11000101 00011001 00111000 01100110 10111101"),
    Bits{}.Read("01000111 00000001 10110011 01111011 01111000 10000111"),
    Bits{}.Read("10111111 10001000 10010001 10100110 01100001 10111011"),
    Bits{}.Read("00011111 00100010 10001010 10100111 00111011 01000111"),
    Bits{}.Read("00111010 00010100 10011100 11110110 10000011 11110010"),
    Bits{}.Read("00010001 01111100 10000001 11010111 11100001 01001110")}

  des := NewDESCipher(key)

  for i := 0; i < len(result); i++ {
    if ! Bits(des.Keys[i]).Equals(result[i]) {
      t.Errorf("DES subkey %d is invalid", i+1)
      }
  }
}

func TestDESencrypt(t *testing.T) {
  key := Key(Bits{}.Read("00111011 00111000 10011000 00110111 00010101 00100000 11110111 01011110"))
  // simply use the key as message
  msg := key
  des := NewDESCipher(Key(key))

  encrypted := des.Encrypt(msg)
  decrypted := des.Decrypt(encrypted)

  result := Bits{}.Read("10001111 00000011 01000101 01101101 00111111 01111000 11100010 11000101")

  for i := 0; i < len(result); i++ {
    if result[i] != encrypted[i] {
      t.Error("DES encryption failed")
    }
  }
  for i := 0; i < len(result); i++ {
    if msg[i] != decrypted[i] {
      t.Error("DES decryption failed")
    }
  }
}
