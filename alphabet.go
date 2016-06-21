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

type AsciiAlphabet struct {
  i, j int
}

/* field access
 * -------------------------------------------------------------------------- */

func (alphabet AsciiAlphabet) Geti() int {
  return alphabet.i
}

func (alphabet AsciiAlphabet) Getj() int {
  return alphabet.j
}

/* -------------------------------------------------------------------------- */

var StdAsciiAlphabet AsciiAlphabet = AsciiAlphabet{32, 126}
var RstAsciiAlphabet AsciiAlphabet = AsciiAlphabet{65, 122}
