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
  pf := NewPrimeField(2)
  // irreducible polynomial
  p := NewPolynomial(pf)
  p.AddTerm(1, 8)
  p.AddTerm(1, 4)
  p.AddTerm(1, 3)
  p.AddTerm(1, 1)
  p.AddTerm(1, 0)
  // create GF(2^8) with irriducible polynomial p
  f := NewFiniteField(p)

  // allocate memory and compute results
  aesMixColAdd := make([][]byte, 0xFF+1)
  aesMixColMul := make([][]byte, 0xFF+1)
  for i := 0; i <= 0xFF; i++ {
    aesMixColAdd[i] = make([]byte, 0xFF+1)
    aesMixColMul[i] = make([]byte, 0xFF+1)
    a := NewPolynomial(pf)
    a.ReadByte(byte(i))
    for j := 0; j <= 0xFF; j++ {
      b := NewPolynomial(pf)
      b.ReadByte(byte(j))
      aesMixColAdd[i][j] = f.Add(a, b).WriteByte()
      aesMixColMul[i][j] = f.Mul(a, b).WriteByte()
    }
  }
  printTable("aesMixColAdd", aesMixColAdd)
  fmt.Println()
  printTable("aesMixColMul", aesMixColMul)
}
