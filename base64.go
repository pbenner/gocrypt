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

package gocrypt

/* -------------------------------------------------------------------------- */

//import "fmt"
import "bufio"
import "encoding/base64"
import "io/ioutil"
import "os"

/* -------------------------------------------------------------------------- */

type Base64 struct {}

/* -------------------------------------------------------------------------- */

func (Base64) ReadString(str string) ([]byte, error) {
  tmp, err := base64.StdEncoding.DecodeString(str)
  if err != nil {
    return nil, err
  }

  return tmp, nil
}

func (Base64) WriteString(b []byte) string {

  s := base64.StdEncoding.EncodeToString(b)

  return s
}

func (Base64) ReadFile(filename string) ([]byte, error) {
  var str string

  f, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    str = str + scanner.Text()
  }
  return Base64{}.ReadString(str)
}

func (Base64) WriteFile(filename string, b []byte) error {

  if err := ioutil.WriteFile(filename, b, 0666); err != nil {
    return err
  }
  return nil
}
