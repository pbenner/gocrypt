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
import "math/big"

/* -------------------------------------------------------------------------- */

func TestBigPrimeField(t *testing.T) {

  p := NewBigPrimeField(big.NewInt(29))
  a := big.NewInt(1)
  b := big.NewInt(17)
  r := big.NewInt(0)

  p.Div(r, a, b)

  if r.Cmp(big.NewInt(12)) != 0 {
    t.Error("big prime field test failed")
  }

}
