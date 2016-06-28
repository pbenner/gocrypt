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

/* -------------------------------------------------------------------------- */

func EEA(ri, rj *big.Int) (*big.Int, *big.Int, *big.Int) {

  zero := big.NewInt(0)

  si := big.NewInt(1)
  ti := big.NewInt(0)
  // j = i+1
  sj := big.NewInt(0)
  tj := big.NewInt(1)
  qj := big.NewInt(0)
  // k = j+1
  sk := big.NewInt(0)
  tk := big.NewInt(0)
  rk := big.NewInt(0)

  for rj.Cmp(zero) != 0 {
    // r_i = r_i-2 mod r_i-1
    rk.Mod(ri, rj)
    // q_i-1 = (r_i-2 - r_i)/r_i-1
    qj.Sub(ri, rk)
    qj.Div(qj, rj)
    // s_i = s_i-2 - q_i-1*s_i-1  
    sk.Mul(qj, sj)
    sk.Sub(si, sk)
    // t_i = t_i-2 - q_i-1*t_i-1
    tk.Mul(qj, tj)
    tk.Sub(ti, tk)

    si, sj, sk = sj, sk, si
    ti, tj, tk = tj, tk, ti
    ri, rj, rk = rj, rk, ri
  }
  // gcd(r0, r1) = ri = s r_0 + t r_1
  return ri, si, ti
}
