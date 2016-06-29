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

/* -------------------------------------------------------------------------- */

type FiniteField struct {
  P   PrimeField
  N   int
  IP *Polynomial
}

/* -------------------------------------------------------------------------- */

func NewFiniteField(p, n int, ip *Polynomial) FiniteField {
  return FiniteField{NewPrimeField(p), n, ip}
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) modp(r *Polynomial) *Polynomial {
  for k, v := range r.Terms {
    r.Terms[k] = float64(f.P.Modp(int(v)))
  }
  r.Clean()
  return r
}

/* arithmetics without mod p (i.e. only coefficients are treated as elements
 * of the prime field)
 * -------------------------------------------------------------------------- */

func (r *Polynomial) netModCoeff(p PrimeField) *Polynomial {
  for k, v := range r.Terms {
    r.Terms[k] = float64(p.Modp(int(v)))
  }
  r.Clean()
  return r
}

func (r *Polynomial) netAddTerm(c float64, e int, p PrimeField) {
  r.AddTerm(c, e)
  r.netModCoeff(p)
}

func (r *Polynomial) netAdd(a, b *Polynomial, p PrimeField) {
  r.Add(a, b)
  r.netModCoeff(p)
}

func (r *Polynomial) netSub(a, b *Polynomial, p PrimeField) {
  r.Sub(a, b)
  r.netModCoeff(p)
}

func (r *Polynomial) netMul(a, b *Polynomial, p PrimeField) {
  r.Mul(a, b)
  r.netModCoeff(p)
}

func (r1 *Polynomial) netdiv(a, b, r2 *Polynomial, p PrimeField) {
  z := NewPolynomial()
  t := NewPolynomial()
  q := NewPolynomial()
  r := a.Clone()
  if b.Equals(z) {
    panic("Div(): division by zero")
  }
  c2, e2 := b.Lead()
  for !r.Equals(z) && r.Degree() >= b.Degree() {
    c1, e1 := r.Lead()
    t.Clear()
    t.netAddTerm(float64(p.Div(int(c1), int(c2))), e1 - e2, p)
    q.netAddTerm(float64(p.Div(int(c1), int(c2))), e1 - e2, p)
    t.netMul(t, b, p)
    r.netSub(r, t, p)
  }
  r1.Terms = q.Terms
  // save remainder if s is given
  if r2 != nil {
    r2.Terms = r.Terms
  }
}

func (r *Polynomial) netDiv(a, b *Polynomial, p PrimeField) {
  r.netdiv(a, b, nil, p)
}

func (r *Polynomial) netMod(a, b *Polynomial, p PrimeField) {
  r.netdiv(a, b, r, p)
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) Add(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.netAdd(a, b, f.P)
  return r
}

func (f FiniteField) Sub(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.netSub(a, b, f.P)
  return r
}

func (f FiniteField) Mul(a, b *Polynomial) *Polynomial {
  r := NewPolynomial()
  r.netMul(a, b, f.P)
  r.Mod(r, f.IP)
  return r
}

func (f FiniteField) Div(a, b *Polynomial) *Polynomial {
  _, _, t := f.EEA(f.IP, b)
  r := f.Mul(a, t)
  return r
}

/* -------------------------------------------------------------------------- */

func (f FiniteField) EEA(ri, rj *Polynomial) (*Polynomial, *Polynomial, *Polynomial) {

  z0 := NewPolynomial()
  si := NewPolynomial()
  si.AddTerm(1, 0)
  ti := NewPolynomial()
  ri  = ri.Clone()
  // j = i+1
  sj := NewPolynomial()
  tj := NewPolynomial()
  tj.AddTerm(1, 0)
  qj := NewPolynomial()
  rj  = rj.Clone()
  // k = j+1
  sk := NewPolynomial()
  tk := NewPolynomial()
  rk := NewPolynomial()

  for !rj.Equals(z0) {
    // r_i = r_i-2 mod r_i-1
    rk.netMod(ri, rj, f.P)
    // q_i-1 = (r_i-2 - r_i)/r_i-1
    qj.netSub(ri, rk, f.P)
    qj.netDiv(qj, rj, f.P)
    // s_i = s_i-2 - q_i-1*s_i-1  
    sk.netMul(qj, sj, f.P)
    sk.netSub(si, sk, f.P)
    // t_i = t_i-2 - q_i-1*t_i-1
    tk.netMul(qj, tj, f.P)
    tk.netSub(ti, tk, f.P)

    si, sj, sk = sj, sk, si
    ti, tj, tk = tj, tk, ti
    ri, rj, rk = rj, rk, ri
  }
  // gcd(r0, r1) = ri = s r_0 + t r_1
  return ri, si, ti
}
