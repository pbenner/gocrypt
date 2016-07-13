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

import "crypto/elliptic"

/* -------------------------------------------------------------------------- */

func TestECC1(t *testing.T) {

  aPub := big.NewInt(0)
  aPrv := big.NewInt(0)

  // read two large integers
  aPubTmp, _ := Base64{}.ReadString(
      "MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBYdTpIsDoEHMCeVXZhf8LYfKET7kg" +
      "anL0zqZeVCZ4GTMZAAn60i4sNywpfpq00ujbzbFBcTMh9ZlPaXoi8hwtgQUBZjeE" +
      "6hjDsGSEo4c3cy+slSBJKrv7zRHBwHy5IL6ZYQmCOEg23+LXk5TRS7BSOgnU9nWE" +
      "n9qD5QtTWLUM8HJVcic=")
  aPub.SetBytes(aPubTmp)

  aPrvTmp, _ := Base64{}.ReadString(
      "MIHuAgEAMBAGByqGSM49AgEGBSuBBAAjBIHWMIHTAgEBBEIBQOUuE8ufDf+Q+FFx" +
      "xc3UQlHloubU4fXa9HEk//48aBGdGZj2uxIyoUiLO9PLTHu823kK9WfezMIpIkl/" +
      "7J7oAYKhgYkDgYYABAFh1OkiwOgQcwJ5VdmF/wth8oRPuSBqcvTOpl5UJngZMxkA" +
      "CfrSLiw3LCl+mrTS6NvNsUFxMyH1mU9peiLyHC2BBQFmN4TqGMOwZISjhzdzL6yV" +
      "IEkqu/vNEcHAfLkgvplhCYI4SDbf4teTlNFLsFI6CdT2dYSf2oPlC1NYtQzwclVy" +
      "Jw==")
  aPrv.SetBytes(aPrvTmp)

  // multiply base point with the two integers
  a := NullECPoint()
  a  = Secp521r1.Eval(a, aPub)
  a  = Secp521r1.Eval(a, aPrv)

  // repeat the same with the go crypt library
  d := elliptic.P521()
  Bx, By := d.ScalarBaseMult(aPub.Bytes())
  Bx, By  = d.ScalarMult(Bx, By, aPrv.Bytes())

  if a.GetX().Cmp(Bx) != 0 {
    t.Error("elliptic curve cryptography test failed")
  }
  if a.GetY().Cmp(By) != 0 {
    t.Error("elliptic curve cryptography test failed")
  }

}
