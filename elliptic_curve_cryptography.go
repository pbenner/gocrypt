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

type ECC struct {
  Curve EllipticCurve
  G     ECPoint
  N, H  *big.Int
}

func NewECC(p, a, b, x_, y_, n_, h_ *big.Int) ECC {
  curve := NewEllipticCurve(a, b, p)
  x := big.NewInt(0)
  x.Set(x_)
  y := big.NewInt(0)
  y.Set(y_)
  n := big.NewInt(0)
  n.Set(n_)
  h := big.NewInt(0)
  h.Set(h_)
  g := NewECPoint(x, y)
  return ECC{curve, g, n, h}
}

func (ecc ECC) Base() ECPoint {
  return ecc.G.Clone()
}

func (ecc ECC) Eval(a ECPoint, b *big.Int) ECPoint {
  if a.IsZero() {
    return ecc.Curve.MulInt(ecc.G, b)
  } else {
    return ecc.Curve.MulInt(a, b)
  }
}

/* -------------------------------------------------------------------------- */

var Secp521r1 ECC

func init() {
  p := big.NewInt(0)
  a := big.NewInt(0)
  b := big.NewInt(0)
  x := big.NewInt(0)
  y := big.NewInt(0)
  n := big.NewInt(0)
  h := big.NewInt(0)

  // secp521r1
  p.SetString("0x01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 0)
  a.SetString("0x01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFC", 0)
  b.SetString("0x0051953EB9618E1C9A1F929A21A0B68540EEA2DA725B99B315F3B8B489918EF109E156193951EC7E937B1652C0BD3BB1BF073573DF883D2C34F1EF451FD46B503F00", 0)
  x.SetString("0x00C6858E06B70404E9CD9E3ECB662395B4429C648139053FB521F828AF606B4D3DBAA14B5E77EFE75928FE1DC127A2FFA8DE3348B3C1856A429BF97E7E31C2E5BD66", 0)
  y.SetString("0x011839296A789A3BC0045C8A5FB42C7D1BD998F54449579B446817AFBD17273E662C97EE72995EF42640C550B9013FAD0761353C7086A272C24088BE94769FD16650", 0)
  n.SetString("0x01FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFA51868783BF2F966B7FCC0148F709A5D03BB5C9B8899C47AEBB6FB71E91386409", 0)
  h.SetString("1", 0)
  Secp521r1 = NewECC(p, a, b, x, y, n, h)

}
