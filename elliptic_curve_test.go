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
import "math/big"
import "testing"

/* -------------------------------------------------------------------------- */

func TestEllipticCurve(t *testing.T) {

  c := NewEllipticCurve(
    big.NewInt(1),
    big.NewInt(7),
    big.NewInt(17))

  p := NewECPoint(
    big.NewInt(2),
    big.NewInt(0))

  q := NewECPoint(
    big.NewInt(1),
    big.NewInt(3))

  r := c.Add(p, q)

  if r.x.Cmp(big.NewInt(6)) != 0 {
    t.Error("elliptic curve test failed")
  }
  if r.y.Cmp(big.NewInt(12)) != 0 {
    t.Error("elliptic curve test failed")
  }

}