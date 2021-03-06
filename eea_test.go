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

func TestEEA(t *testing.T) {

  r0 := 67
  r1 := 12

  ri, si, ti := EEA(r0, r1)

  if ri != 1 {
      t.Error("EEA failed")
  }
  if si != -5 {
      t.Error("EEA failed")
  }
  if ti != 28 {
      t.Error("EEA failed")
  }

}

func TestBigEEA(t *testing.T) {

  r0 := big.NewInt(67)
  r1 := big.NewInt(12)

  ri, si, ti := BigEEA(r0, r1)

  if ri.Cmp(big.NewInt(1)) != 0 {
      t.Error("BigEEA failed")
  }
  if si.Cmp(big.NewInt(-5)) != 0 {
      t.Error("BigEEA failed")
  }
  if ti.Cmp(big.NewInt(28)) != 0 {
      t.Error("BigEEA failed")
  }

}
