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

package main

/* -------------------------------------------------------------------------- */

import   "fmt"

import . "github.com/pbenner/gocrypt"

/* -------------------------------------------------------------------------- */

func printTable(name string, table [][]byte) {

  fmt.Printf("var %s = [0x100][0x100]byte{\n", name)

  for i := 0; i < len(table); i++ {
    fmt.Printf("  { ")
    for j := 0; j < len(table); j++ {
      if j != 0 {
        fmt.Printf(", ")
      }
      if j != 0 && j % 16 == 0 {
        fmt.Printf("\n    ")
      }
      fmt.Printf("0x%02X", table[i][j])
    }
    if i != len(table)-1 {
      fmt.Printf(" },\n")
    } else {
      fmt.Printf(" }}\n")
    }
  }
  
}

/* -------------------------------------------------------------------------- */

func main() {

  // precompute addition and multiplication in GF(2^8)

  // irreducible polynomial
  p := NewBinaryPolynomial(1)
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)
  // create GF(2^8) with irriducible polynomial p
  f := NewBinaryExtensionField(p)
  // field elements
  a := NewBinaryPolynomial(1)
  b := NewBinaryPolynomial(1)
  r := NewBinaryPolynomial(1)
  // allocate memory and compute results
  aesMixColAdd := make([][]byte, 0xFF+1)
  aesMixColMul := make([][]byte, 0xFF+1)
  for i := 0; i <= 0xFF; i++ {
    aesMixColAdd[i] = make([]byte, 0xFF+1)
    aesMixColMul[i] = make([]byte, 0xFF+1)
    a.Terms[0] = byte(i)
    for j := 0; j <= 0xFF; j++ {
      b.Terms[0] = byte(j)
      f.Add(r, a, b)
      if len(r.Terms) == 0 {
        aesMixColAdd[i][j] = 0
      } else {
        aesMixColAdd[i][j] = r.Terms[0]
      }
      f.Mul(r, a, b)
      if len(r.Terms) == 0 {
        aesMixColMul[i][j] = 0
      } else {
        aesMixColMul[i][j] = r.Terms[0]
      }
    }
  }
  printTable("aesMixColAdd", aesMixColAdd)
  fmt.Println()
  printTable("aesMixColMul", aesMixColMul)
}
