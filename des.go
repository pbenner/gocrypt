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

package lib

/* -------------------------------------------------------------------------- */

//import "fmt"

/* -------------------------------------------------------------------------- */

type DESCipher struct {
  FeistelNetwork
}

/* -------------------------------------------------------------------------- */

func fExpansion(input, output []byte) {
  table := [][]int{
    { 1, 47},
    { 2},
    { 3},
    { 4,  6},
    { 7,  5},
    { 8},
    { 9},
    {10, 12},
    {13, 11},
    {14},
    {15},
    {16, 18},
    {19, 17},
    {20},
    {21},
    {22, 24},
    {25, 23},
    {26},
    {27},
    {28, 30},
    {31, 29},
    {32},
    {33},
    {34, 36},
    {37, 35},
    {38},
    {39},
    {40, 42},
    {43, 41},
    {44},
    {45},
    {46,  0} }
  RemapBits(input, output, table)
}

func NewDESCipher() {
}
