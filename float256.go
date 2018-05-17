// Package float256 implements a 256 bit precision floating point math library for high precision financial calculations
package float256

import (
	"math/big"
)

// Zero returns a new float with 256 bits of precision
func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

// New returns a new floating point number with 256 bits of precision containing the desired initial value
func New(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(256)
	return r
}

// FromString returns a float with 256 bit precision from a string or zero if invalid
// The reverse conversion is built into big.Float with .String() though it is unnecessary to do this to use with any of the fmt Print* functions
func FromString(s string) *big.Float {
	r, _ := Zero().SetString(s)
	r.SetPrec(256)
	return r
}

// Int returns the truncated value of the parameter as a *big.Int
func Int(a *big.Float) *big.Int {
	i, _ := a.Int(nil)
	return i
}

// FromInt returns a *big.Float from a *big.Int
func FromInt(i *big.Int) *big.Float {
	return Zero().SetInt(i)
}

// Int64 returns the value of the parameter as an int64 with the decimal truncated
func Int64(a *big.Float) int64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Int64()
	return i
}

// FromInt64 returns a *big.Float from an int64
func FromInt64(i int64) *big.Float {
	return Zero().SetInt64(i)
}

// Uint64 returns the value of the parameter as a uint64 with decimal truncated and sign positive (obviously)
func Uint64(a *big.Float) uint64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Uint64()
	return i
}

// FromUint64 returns a *big.Float from a uint64
func FromUint64(i uint64) *big.Float {
	return Zero().SetUint64(i)
}

// Equal returns true if x and y are equal
func Equal(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// Greater returns true if x > y
func Greater(x, y *big.Float) bool {
	return x.Cmp(y) == 1
}

// Lesser returns true if x < y
func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

// Exp returns a ** b
func Exp(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	for i:=uint64(0); i<e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

// Root returns the nth root of a
func Root(a *big.Float, n uint64) *big.Float {
	limit := Exp(New(2), 256)
	n1 := n-1
	n1f, rn := New(float64(n1)), Div(New(1.0), New(float64(n)))
	x, x0 := New(1.0), Zero()
	_ = x0
	for {
		potx, t2 := Div(New(1.0), x), a
		for b:=n1; b>0; b>>=1 {
			if b&1 == 1 {
				t2 = Mul(t2, potx)
			}
			potx = Mul(potx, potx)
		}
		x0, x = x, Mul(rn, Add(Mul(n1f, x), t2) )
		if Lesser(Mul(Abs(Sub(x, x0)), limit), x) { break } 
	}
	return x
}

// Abs returns a with sign set to positive
func Abs(a *big.Float) *big.Float {
	return Zero().Abs(a)
}

// Neg flips the sign of the parameter Float
func Neg(a *big.Float) *big.Float {
	return a.Neg(a)
}

// Sign returns 1 for positive, -1 for negative and 0 for 0s
func Sign(a *big.Float) int {
	return a.Sign()
}

// Sqrt returns the square root of the parameter
func Sqrt(a *big.Float) *big.Float {
	return Root(a, 2)
}

// Add returns a + b
func Add(a, b *big.Float) *big.Float {
	return Zero().Add(a, b)
}

// Sub returns a - b
func Sub(a, b *big.Float) *big.Float {
	return Zero().Sub(a, b)
}

// Mul returns a * b
func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

// Div returns a / b
func Div(a, b *big.Float) *big.Float {
	return Zero().Quo(a, b)
}

// Mod returns a % b
func Mod(a, b *big.Float) *big.Float {
	q := Div(a, b)
	i := Int(q)
	fi := Zero().SetInt(i)
	rem := Sub(a, Mul(b, fi))
	return rem
}

// TODO: write codec to convert to and from 256 bit fixed point with zero padding and no sign