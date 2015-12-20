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

package main

/* -------------------------------------------------------------------------- */

//import   "fmt"
import   "bytes"
import   "bufio"
import   "os"
import   "code.google.com/p/getopt"

import . "github.com/pbenner/autodiff"
import . "github.com/pbenner/0100101101"

/* utility
 * -------------------------------------------------------------------------- */

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readFile(filename string) string {
  f, err := os.Open(filename)
  check(err)

  var text bytes.Buffer
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    text.WriteString(scanner.Text())
    text.WriteString(" ")
  }

  return text.String()
}

/* -------------------------------------------------------------------------- */

func runSampler(n int, trainingFile, textFile string, verbose bool) {

  // get a new random cipher
//  alphabet := StdAsciiAlphabet
  alphabet := RstAsciiAlphabet
  cipher1  := NewAsciiPermutationCipher(alphabet)

  // read text files
  trainingText := readFile(trainingFile)
     plainText := readFile(textFile)

  // estimate the transition matrix on the given text
  t := estimateTransitions(ProbabilityType, trainingText)

  cipherText := cipher1.Encrypt(NewMessage(plainText))

  sampler(n, alphabet, cipherText, t, verbose)
}

/* -------------------------------------------------------------------------- */

func main() {

  optIterations := getopt.IntLong   ("iterations", 'i',  1000, "number of iterations     [default: 1000]")
  optHelp       := getopt.BoolLong  ("help",       'h',        "print help")
  optVerbose    := getopt.BoolLong  ("verbose",    'v',        "print verbose output")

  getopt.SetParameters("<TrainingFile TextFile>")
  getopt.Parse()

  if *optHelp {
    getopt.Usage()
    os.Exit(0)
  }
  if len(getopt.Args()) != 2 {
    getopt.Usage()
    os.Exit(1)
  }
  trainingFile := getopt.Args()[0]
      textFile := getopt.Args()[1]

  runSampler(*optIterations, trainingFile, textFile, *optVerbose)

}
