// Package float256 is a high precision fixed point library built from math/big's big.Float with a fixed maximum precision of 42 integer places and precision to 63 decimal places intended for use in a cryptocurrency ledger for a token with a maximum supply less than 42 bits of integer place token denomination of total supply. The maximum value the integer part can hold is 4,398,046,511,103.
// It has a serialisation codec that stores this as a 42.214 fixed precision value and converts it back to big.Float.
// During calculations it is possible to work with negative numbers, for convenience, however they cannot be imported using strings or integers or encoded and the functions return nil to signify the input is invalid.
package float256

import (
	"math/big"
)

// Zero returns a new float with 256 bits of precision. This can be a good way to generate a new variable with minimal typing for input conversion functions particularly.
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

// FromString accepts a properly formed string of integer or decimal value and returns a big.Float to the caller. Negative values and values greater than 4,398,046,511,103 are invalid and rejected by returning nil.
func FromString(s string) *big.Float {
	r, success := Zero().SetString(s)
	if r.Sign()<0 { return nil }
	if r.MantExp(nil)>42 { return nil }
	if success { return r }
	return nil
}

// Int returns the truncated value of the parameter as a *big.Int
func Int(a *big.Float) *big.Int {
	i, _ := a.Int(nil)
	return i
}

// FromInt returns a *big.Float from a *big.Int. Negative values and values greater than 4,398,046,511,103 are invalid and rejected by returning nil.
func FromInt(i *big.Int) *big.Float {
	r := Zero().SetInt(i)
	if r.Sign()<0 { return nil }
	if r.MantExp(nil)>42 { return nil }
	return r
}

// Int64 returns the value of the parameter as an int64 with the decimal truncated
func Int64(a *big.Float) int64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Int64()
	return i
}

// FromInt64 returns a *big.Float from an int64. Negative values and values greater than 4,398,046,511,103 are invalid and rejected by returning nil.
func FromInt64(i int64) *big.Float {
	r := Zero().SetInt64(i)
	if r.Sign()<0 { return nil }
	if r.MantExp(nil)>42 { return nil}
	return r
}

// Uint64 returns the value of the parameter as a uint64 with decimal truncated and sign positive (obviously)
func Uint64(a *big.Float) uint64 {
	r := Zero()
	r.Copy(a)
	i, _ := r.Uint64()
	return i
}

// FromUint64 returns a *big.Float from a uint64. Values greater than 4,398,046,511,103 are invalid and rejected by returning nil as they exceed the precision of the serialization format.
func FromUint64(i uint64) *big.Float {
	r := Zero().SetUint64(i)
	if r.Sign()<0 { return nil }
	if r.MantExp(nil)>42 { return nil }
	return r
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

// Root returns the nth root of a to 255 significant bits (if it is larger this function gets stuck in an infinite loop)
func Root(a *big.Float, n uint64) *big.Float {
	limit := Exp(New(2), 255)
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

// Neg flips the sign of the parameter Float. Negative values are rejected by the Encode function as they are invalid in a cryptocurrency ledger.
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

// Encode takes a big.Float, left shifts and truncates to 214 bits of decimal precision and maximum 42 bits of integer precision and returns a byte slice that can be stored on disk or sent over a network. Negative values and values greater than 4,398,046,511,103 are invalid and rejected by returning nil.
func Encode(a *big.Float) []byte {
	if a.Sign()<0 { return nil }
	mantissa := New(0)
	a.MantExp(mantissa)
	exp := a.MantExp(nil)
	if exp > 42 { return nil }
	bytes := Int(Mul(Exp(New(2), uint64(214+exp)), mantissa)).Bytes()
	return append(make([]byte, 32-len(bytes)), bytes...)
}

// Decode takes a value that has been created by Encode, loads it into a big.Float, changes the exponent to shift the point to 214 bits of decimal precision and 42 bits of integer precision (exponent)
func Decode(b []byte) *big.Float {
	decoded := big.NewInt(0).SetBytes(b)
	decodedF := Zero().SetInt(decoded)
	return Zero().SetMantExp(decodedF, -214)
}