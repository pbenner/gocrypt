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

import "fmt"
import "testing"

/* -------------------------------------------------------------------------- */

func TestPrimeField1(t *testing.T) {

  p := NewPrimeField(2)

  fmt.Println(p.Add(1,0))

}

func TestPrimeField2(t *testing.T) {

  p := NewPrimeField(5)

  for i := 1; i < 5; i++ {
    fmt.Printf("1/%d: = %d\n", i, p.Div(1,i))
  }

}
