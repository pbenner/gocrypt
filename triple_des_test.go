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

func TestTripleDESencrypt(t *testing.T) {

	key    := []byte{
    0x3b, 0x38, 0x98, 0x37, 0x15, 0x20, 0xf7, 0x5e,
    0x92, 0x2f, 0xb5, 0x10, 0xc7, 0x1f, 0x43, 0x6e,
    0x3b, 0x38, 0x98, 0x37, 0x15, 0x20, 0xf7, 0x5e }
  des, _ := NewTripleDESCipher(key)
  ecb    := NewECBCipher(des)

  plaintext := []byte("Gott wuerfelt nicht... !")
  encrypted := make([]byte, len(plaintext))
  decrypted := make([]byte, len(plaintext))

  result := []byte{
    0x22, 0x86, 0x1e, 0x54, 0xd0, 0x8e, 0x48, 0xb0,
    0x2a, 0x35, 0x77, 0xb5, 0x45, 0xa4, 0xa9, 0x26,
    0xc2, 0x07, 0x45, 0xb4, 0x44, 0x2c, 0xf0, 0x79 }

  err1 := ecb.Encrypt(plaintext, encrypted)
  err2 := ecb.Decrypt(encrypted, decrypted)

  if err1 != nil {
    t.Error(err1)
  }
  if err2 != nil {
    t.Error(err2)
  }
  if !Bits(encrypted).Equals(result) {
    t.Error("triple DES encryption failed")
  }
  if !Bits(plaintext).Equals(decrypted) {
    t.Error("triple DES decryption failed")
  }
}
