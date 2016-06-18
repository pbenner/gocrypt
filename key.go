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

package lib

/* -------------------------------------------------------------------------- */

//import "fmt"
import "bufio"
import "encoding/binary"
import "encoding/base64"
import "io/ioutil"
import "math"
import "os"

/* type definition
 * -------------------------------------------------------------------------- */

type Key []byte

/* constructors
 * -------------------------------------------------------------------------- */

func NewKey(n int) Key {
  return make(Key, n)
}

/* type conversion
 * -------------------------------------------------------------------------- */

func (k Key) Uint64Slice() []uint64 {

  const step int = 8;

  length := int(math.Ceil(float64(len(k))/float64(step)))
  tmp    := make([]byte,   step)
  result := make([]uint64, length)

  for i := 0; i < len(k); i += step {
    for j := 0; j < step; j++ {
      if i+j < len(k) {
        tmp[j] = k[i+j]
      } else {
        tmp[j] = 0
      }
    }
    result[i/step] = binary.LittleEndian.Uint64(tmp)
  }
  return result
}

/* i/o
 * -------------------------------------------------------------------------- */

func (key *Key) Read(filename string) error {
  var keystr string

  f, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    keystr = keystr + scanner.Text()
  }
  tmp, err := base64.StdEncoding.DecodeString(string(keystr))

  if err != nil {
    return err
  }
  *key = tmp

  return nil
}

func (key Key) Write(filename string) error {

  if err := ioutil.WriteFile(filename, []byte(key.String()), 0666); err != nil {
    return err
  }
  return nil
}

func (k Key) String() string {

  s := base64.StdEncoding.EncodeToString(k)

  return s
}
