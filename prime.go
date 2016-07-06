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
import "math/rand"

/* -------------------------------------------------------------------------- */

func FermatTest(p *big.Int, s int, rnd *rand.Rand) bool {
  // copy p to q
  q := big.NewInt(0)
  q.Set(p)
  if q.Sign() == -1 {
    q.Mul(q, big.NewInt(-1))
  }
  // some constants
  c1 := big.NewInt(1)
  c2 := big.NewInt(2)
  c3 := big.NewInt(3)
  c4 := big.NewInt(4)
  // variables for the test
  a := big.NewInt(0)
  n := big.NewInt(0)
  n.Set(q)
  n.Sub(n, c4)
  if q.Cmp(c2) == 0 || q.Cmp(c3) == 0 {
    return true
  }
  if n.Sign() <= 0 {
    return false
  }
  // compute q = p-1
  q.Sub(q, c1)
  for i := 0; i < s; i++ {
    a.Rand(rnd, n)
    a.Add(a, c2)
    a.Exp(a, q, p)
    if a.Cmp(c1) != 0 {
      return false
    }
  }
  return true
}
