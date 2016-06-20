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

func TestECBCipher(t *testing.T) {

	key    := []byte{0x3b, 0x38, 0x98, 0x37, 0x15, 0x20, 0xf7, 0x5e}
  des, _ := NewDESCipher(key)
  ecb    := NewECBCipher(des)

  plaintext := []byte("Gott wuerfelt nicht... !")
  encrypted := make([]byte, len(plaintext))
  decrypted := make([]byte, len(plaintext))

  err1 := ecb.Encrypt(plaintext, encrypted)
  err2 := ecb.Decrypt(encrypted, decrypted)

  if err1 != nil {
    t.Error(err1)
  }
  if err2 != nil {
    t.Error(err2)
  }
  if !Bits(plaintext).Equals(decrypted) {
    t.Error("ECB encryption/decryption failed")
  }
}

func TestCBCCipher(t *testing.T) {

	key    := []byte{0x3b, 0x38, 0x98, 0x37, 0x15, 0x20, 0xf7, 0x5e}
  des, _ := NewDESCipher(key)
  cbc, _ := NewCBCCipher(des, []byte{1,2,3,4,5,6,7,8})

  plaintext := []byte("Gott wuerfelt nicht... !")
  encrypted := make([]byte, len(plaintext))
  decrypted := make([]byte, len(plaintext))

  err1 := cbc.Encrypt(plaintext, encrypted)
  err2 := cbc.Decrypt(encrypted, decrypted)

  if err1 != nil {
    t.Error(err1)
  }
  if err2 != nil {
    t.Error(err2)
  }
  if !Bits(plaintext).Equals(decrypted) {
    t.Error("CBC encryption/decryption failed")
  }
}
