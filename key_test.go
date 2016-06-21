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

func TestKeyReadWrite(t *testing.T) {
  filename := "key_test.key"

  key1 := Key(Bits{}.Read("00111011 00111000 10011000 00110111 00010101 00100000 11110111 01011110"))
  key1.Write(filename)
  key2 := Key{}
  key2.Read(filename)

  for i := 0; i < len(key1); i++ {
    if key1[i] != key2[i] {
      t.Error("key read/write failed")
    }
  }

}
