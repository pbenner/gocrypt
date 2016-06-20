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

package gocrypto

/* -------------------------------------------------------------------------- */

//import "fmt"
import "math/rand"
import "time"
import "github.com/seehuhn/mt19937"

/* -------------------------------------------------------------------------- */

type OneTimePad struct
{
  privateKey Key
}

/* constructors
 * -------------------------------------------------------------------------- */

func NewOneTimePad(keyLength int) OneTimePad {
  cipher := OneTimePad{}
  cipher.Generate(keyLength)
  return cipher
}

/* -------------------------------------------------------------------------- */

func (cipher *OneTimePad) GetKey() Key {
  return cipher.privateKey
}

func (cipher *OneTimePad) Generate(args ...interface{}) {

  length := 8;

  if len(args) > 1 {
    panic("Generate(): Invalid argument")
  }
  if len(args) == 1 {
    length = args[0].(int)
  }

  cipher.privateKey= make(Key, length)

  rng := rand.New(mt19937.New())
  rng.Seed(time.Now().UnixNano())

  for i := range cipher.privateKey {
    cipher.privateKey[i] = byte(rng.Int())
  }
}

func (cipher *OneTimePad) Encrypt(m Message) Message {

  result := NullMessage(len(m))
  length := len(cipher.privateKey)

  if len(cipher.privateKey) != length {
    panic("No private key available!")
  }
  mt  := mt19937.New()
  mt.SeedFromSlice(cipher.privateKey.Uint64Slice())
  rng := rand.New(mt)

  for i := 0; i < len(m); i++ {
    result[i] = m[i] ^ byte(rng.Int())
  }

  return result
}

func (cipher *OneTimePad) Decrypt(m Message) Message {

  return cipher.Encrypt(m)
}
