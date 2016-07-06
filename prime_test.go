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
import "math/rand"

/* -------------------------------------------------------------------------- */

func TestFermatTest(t *testing.T) {

  r := rand.New(rand.NewSource(1))
  p := big.NewInt(1)

  if _, ok := p.SetString("20988936657440586486151264256610222593863921", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != true {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("-20988936657440586486151264256610222593863921", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != true {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("2", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != true {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("-2", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != true {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("3", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != true {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("0", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != false {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("1", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != false {
    t.Error("Fermat test failed")
  }

  if _, ok := p.SetString("-1", 10); !ok {
    panic("could not parse number")
  }
  if FermatTest(p, 100, r) != false {
    t.Error("Fermat test failed")
  }

}
