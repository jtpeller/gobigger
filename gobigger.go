// ============================================================================
// = gobigger.go
// = 	Description		math/big wrappers for readability
// = 	Date			2025.02.05
// ============================================================================

/* Make sure you change the name! */
package rename_package_upon_import

import (
	"math/big"
	"math/rand"
)

// ========================================================
// CONSTANTS
// ========================================================

// how big the mantissa is going to be. For super extreme
//	precision, you could use a far larger number but this
//	will use much more space / memory.
const DEFAULT_FLOAT_PREC = 256

// ========================================================
// SHORTHAND TYPES
// ========================================================

type bint = big.Int
type bfloat = big.Float
type brat = big.Rat

// ========================================================
// WRAPPERS SECTION
// ========================================================

// ==============================================
// BIG INT WRAPPERS
// ==============================================

func zero() *bint { return big.NewInt(0) }

func inew(i int64) *bint { return big.NewInt(i) }

func abs(a *bint) *bint { return zero().Abs(a) }

func add(a, b *bint) *bint { 	return zero().Add(a, b) }

func and(a, b *bint) *bint { return zero().And(a, b) }

func nand(a, b *bint) *bint { return zero().AndNot(a, b) }

func binomial(a *bint, n, k int64) *bint { return a.Binomial(n, k) }

func cmp(a, b *bint) int { return a.Cmp(b) }

func cmpAbs(a, b *bint) int { return a.CmpAbs(b) }

func div(a, b *bint) *bint { return zero().Div(a, b) }

func divmod(a, b, m *bint) (*bint, *bint) { return zero().DivMod(a, b, m) }

func exp(a, b, m *bint) *bint { return zero().Exp(a, b, m) }

func fillBytes(a *bint, buf []byte) []byte { return a.FillBytes(buf) }

func gcd_w(x, y, a, b *bint) *bint { return zero().GCD(x, y, a, b) }

func lsh(a *bint, n uint) *bint { return zero().Lsh(a, n) }

func mod(a, b *bint) *bint { return zero().Mod(a, b) }

func modInverse(a, b *bint) *bint { return zero().ModInverse(a, b) }

func modSqrt(x, p *bint) *bint { return zero().ModSqrt(x, p) }

func mul(a, b *bint) *bint { return zero().Mul(a, b) }

func mulRange(a, b int64) *bint { return zero().MulRange(a, b) }

func neg(a *bint) *bint { return zero().Neg(a) }

func not(a *bint) *bint { return zero().Not(a) }

func or(a, b *bint) *bint { return zero().Or(a, b) }

func probablyPrime(a *bint, n int) bool { return a.ProbablyPrime(n) }

func quo(a, b *bint) *bint { return zero().Quo(a, b) }

func quoRem(a, b, r *bint) (*bint, *bint) { return zero().QuoRem(a, b, r) }

// underlying Rand function requires math/rand
func rnd(rnd *rand.Rand, n *bint) *bint { return zero().Rand(rnd, n) }

func rem(a, b *bint) *bint { return zero().Rem(a, b) }

func rsh(a *bint, n uint) *bint { return zero().Rsh(a, n) }

func sqrt(a *bint) *bint { return zero().Sqrt(a) }

func sub(a, b *bint) *bint { return zero().Sub(a, b) }

func xor(a, b *bint) *bint { return zero().Xor(a, b) }

// ==============================================
// BIG FLOAT WRAPPERS
// ==============================================

func fzero() *bfloat { return big.NewFloat(0) }

func fnew(a float64) *bfloat { return big.NewFloat(a) }

func parseFloat(s string, base int, prec uint, mod big.RoundingMode) (f *bfloat, b int, err error) { return big.ParseFloat(s, base, prec, mod) }

func fabs(a *bfloat) *bfloat { return fzero().Abs(a) }

func facc(a *bfloat) big.Accuracy { return a.Acc() }

func fadd(a, b *bfloat) *bfloat { return fzero().Add(a, b) }

func fappend(a *bfloat, buf []byte, fmt byte, prec int) []byte {return a.Append(buf, fmt, prec) }

func fcmp(a, b *bfloat) int { return a.Cmp(b) }

func fcopy(a, b *bfloat) *bfloat { return a.Copy(b) }

func fdiv(a, b *bfloat) *bfloat { return fzero().Quo(a, b) }

func ftofloat32(a *bfloat) (float32, big.Accuracy) { return a.Float32() }

func ftofloat64(a *bfloat) (float64, big.Accuracy) { return a.Float64() }

func fmul(a, b *bfloat) *bfloat { return fzero().Mul(a, b) }

func fneg(a *bfloat) *bfloat { return fzero().Neg(a) }

func fquo(a, b *bfloat) *bfloat { return fdiv(a, b) }

func fsqrt(a *bfloat) *bfloat { return fzero().Sqrt(a) }

func fsub(a, b *bfloat) *bfloat { return fzero().Sub(a, b) }

func trunc(a *bfloat) (*bint, big.Accuracy) { return a.Int(nil) }

// ==============================================
// BIG RAT WRAPPERS
// ==============================================

func rzero() *brat { return big.NewRat(0, 1)}

func rnew(a, b int64) *brat { return big.NewRat(a, b) }

func rabs(a *brat) *brat { return rzero().Abs(a) }

func radd(x, y *brat) *brat { return rzero().Add(x, y) }

func rcmp(a, b *brat) int { return a.Cmp(b) }

func rdiv(a, b *brat) *brat { return rzero().Quo(a, b) }

func rinv(a *brat) *brat { return rzero().Inv(a) }

func rmul(a, b *brat) *brat { return rzero().Mul(a, b) }

func rneg(a *brat) *brat { return rzero().Neg(a) }

func rsub(a, b *brat) *brat { return rzero().Sub(a, b) }

// =================================================================
// EXTENDED FUNCTIONALITY!
// =================================================================

// ==============================================
// BIG INT NEW FUNCTIONALITY
// ==============================================

// adds all big ints provided!
func addall(nums ...*bint) *bint {
	sum := zero()
	for _, n := range nums {
		sum = add(sum, n)
	}
	return sum
}

// subtracts all big ints provided!
func suball(start *bint, nums ...*bint) *bint {
	sum := addall(nums...)
	return sub(start, sum)
}

// multiplies all big ints provided!
func mulall(nums ...*bint) *bint {
	prod := inew(1)
	for _, n := range nums {
		prod = mul(prod, n)
	}
	return prod
}

// divides all big ints provided!
func divall(start *bint, nums ...*bint) *bint {
	quo := mulall(nums...)
	return div(start, quo)
}

// returns a big int slice of given length. Prepopulated with 0s
func bigSlice(len int64) []*bint {
	slice := make([]*bint, 0)
	for i := int64(0); i < len; i++ {
		slice = append(slice, big.NewInt(0))
	}
	return slice
}

// computes the factorial of the big.Int provided!
func fact(a *bint) *bint {
	prod := inew(1)
	for i := inew(2); lteq(i, a); inc(i) {
		prod = mul(prod, i)
	}
	return prod
}

// Inc(a) returns a+1
func inc(a *big.Int) *big.Int {
	return a.Add(a, inew(1))
}

// Dec(a) returns a-1
func dec(a *big.Int) *big.Int {
	return a.Sub(a, inew(1))
}

// Pow(a, e) returns a^e
func pow(a *big.Int, e *big.Int) *big.Int {
	return zero().Exp(a, e, big.NewInt(0))
}

// gcd() -- simplified wrapper for the one in math.Big
func gcd(a, b *big.Int) *big.Int {
	return zero().GCD(nil, nil, a, b)
}

/**
 * nCr(n, k) -- Compute the binomial of two big ints! 
 * 	AKA Binomial(n, k) or C(n, r)
 *  Will panic if: n < 0 || k < 0 || n < k
 */
func nCr(n, k *bint) *bint { 
	// error checking
	if lt(n, zero()) || lt(k, zero()) {
		panic("Neither n nor k can be negative!")
	} else if lt(n, k) {
		panic("n cannot be less than k")
	}

	// C(n,r) = n!/((n-r)!r!). Instead, do the shortcut. Modify k based on n
	// k > n/2
	if gt(k, div(n, inew(2))) {
		k = sub(n, k)		// k = n - k
	}
	c := inew(1)
	for i := inew(1); lteq(i, k); inc(i) {
		// c = (n - k + i) * c / i
		c = div(mul(add(sub(n, k), i), c), i)
	}
	return c
}

// computes k-permutations of n (P(n, k) aka nPr(n, r))
func nPr(n, k *bint) *bint {
	// check for problematic cases
	if gt(k, n) {
		return zero()
	}

	// return n!/(n-k)!
	return div(fact(n), fact(sub(n, k)))
}

// ### conversions

// converts a big.Int to big.Float
func itof(a *bint) *bfloat {
	r := new(big.Float)
	r.SetInt(a)
	r.SetPrec(DEFAULT_FLOAT_PREC)
	return r
}

// converts a big.Int to big.Rat
func itor(a *bint) *brat {
	r := new(big.Rat)
	r.SetFrac(a, inew(1))
	return r 
}

// ### comparisons

// returns true/false of a < b
func lt(a, b *bint) bool { return a.Cmp(b) == -1 }

// returns true/false of a == b
func equals(a, b *bint) bool { return a.Cmp(b) == 0 }

// returns true/false of a > b
func gt(a, b *bint) bool { return a.Cmp(b) == 1 }

// returns true/false of a <= b
func lteq(a, b *bint) bool { return lt(a, b) || equals(a, b) }

// returns true/false of a >= b
func gteq(a, b *bint) bool { return gt(a, b) || equals(a, b) }

// ==============================================
// BIG FLOAT NEW FUNCTIONALITY
// ==============================================

// adds all big floats provided!
func addall_f(nums ...*bfloat) *bfloat {
	sum := fzero()
	for _, n := range nums {
		sum = fadd(sum, n)
	}
	return sum
}

// subtracts all big floats provided!
func suball_f(start *bfloat, nums ...*bfloat) *bfloat {
	sum := addall_f(nums...)
	return fsub(start, sum)
}

// multiplies all big floats provided!
func mulall_f(nums ...*bfloat) *bfloat {
	sum := fnew(1)
	for _, n := range nums {
		sum = fmul(sum, n)
	}
	return sum
}

// divides all big floats provided!
func divall_f(start *bfloat, nums ...*bfloat) *bfloat {
	denom := mulall_f(nums...)
	return fdiv(start, denom)
}

// computes a^e, where a is a float!
func fpow(a *bfloat, e int64) *bfloat { 
	if e == 0 {
		return fnew(1)
	}
	r := fzero().Copy(a)
	for i := int64(0); i < e-1; i++ {
		r = fmul(r, a)
	}
	if e < 0 {		// negative exponents are reciprocals of positives
		return fdiv(fnew(1), r)
	}
	return r
}

// ### conversions

// truncates big float. Result is a big int.
func floor(a *bfloat) *bint {
	f, _ := a.Int(nil)
	return f
}

// converts big float to int via truncation.
func ftoi(a *bfloat) *bint { return floor(a) }

// rounds big float. Result is a big int.
func round(a *bfloat) *bint {
	a.Add(a, fnew(0.5))
	r, _ := a.Int(nil)
	return r
}

// ### comparisons

// returns true/false of a < b
func lt_f(a, b *bfloat) bool { return a.Cmp(b) == -1 }

// returns true/false of a == b
func equals_f(a, b *bfloat) bool { return a.Cmp(b) == 0 }

// returns true/false of a > b
func gt_f(a, b *bfloat) bool { return a.Cmp(b) == 1 }

// returns true/false of a <= b
func lteq_f(a, b *bfloat) bool { return lt_f(a, b) || equals_f(a, b) }

// returns true/false of a >= b
func gteq_f(a, b *bfloat) bool { return gt_f(a, b) || equals_f(a, b) }


// ==============================================
// BIG RAT NEW FUNCTIONALITY
// ==============================================

// adds all big rats provided!
func addall_r(nums ...*brat) *brat {
	sum := rzero()
	for _, n := range nums {
		sum = radd(sum, n)
	}
	return sum
}

// subtracts all big rats provided!
func suball_r(start *brat, nums ...*brat) *brat {
	sum := addall_r(nums...)
	return rsub(start, sum)
}

// multiplies all big rats provided!
func mulall_r(nums ...*brat) *brat {
	sum := rnew(1, 1)
	for _, n := range nums {
		sum = rmul(sum, n)
	}
	return sum
}

// divides all big rats provided!
func divall_r(start *brat, nums ...*brat) *brat {
	denom := mulall_r(nums...)
	return rdiv(start, denom)
}

// ### conversions

// converts big rat to float
func rtof(r *brat) *bfloat {
	return fdiv(itof(r.Num()), itof(r.Denom()))
}

// converts big rat to int
func rtoi(r *brat) *bint {
	return floor(rtof(r))
}

// ### comparisons

// returns true/false of a < b
func lt_r(a, b *brat) bool { return a.Cmp(b) == -1 }

// returns true/false of a == b
func equals_r(a, b *brat) bool { return a.Cmp(b) == 0 }

// returns true/false of a > b
func gt_r(a, b *brat) bool { return a.Cmp(b) == 1 }

// returns true/false of a <= b
func lteq_r(a, b *brat) bool { return lt_r(a, b) || equals_r(a, b) }

// returns true/false of a >= b
func gteq_r(a, b *brat) bool { return gt_r(a, b) || equals_r(a, b) }