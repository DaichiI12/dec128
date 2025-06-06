package dec128

import "github.com/worldlycuisin/dec128/state"

// RoundDown (or Floor) rounds the decimal to the specified precision using Round Down method (https://en.wikipedia.org/wiki/Rounding#Rounding_down).
//
// Examples:
//
//	RoundDown(1.236, 2) = 1.23
//	RoundDown(1.235, 2) = 1.23
//	RoundDown(1.234, 2) = 1.23
//	RoundDown(-1.234, 2) = -1.24
//	RoundDown(-1.235, 2) = -1.24
//	RoundDown(-1.236, 2) = -1.24
func (self Dec128) RoundDown(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	q, r, s := self.coef.QuoRem64(Pow10Uint64[self.exp-prec])
	if s >= state.Error {
		return Dec128{state: s}
	}

	if self.state == state.Neg && r > 0 {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// RoundUp (or Ceil) rounds the decimal to the specified precision using Round Up method (https://en.wikipedia.org/wiki/Rounding#Rounding_up).
//
// Examples:
//
//	RoundUp(1.236, 2) = 1.24
//	RoundUp(1.235, 2) = 1.24
//	RoundUp(1.234, 2) = 1.24
//	RoundUp(-1.234, 2) = -1.23
//	RoundUp(-1.235, 2) = -1.23
//	RoundUp(-1.236, 2) = -1.23
func (self Dec128) RoundUp(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	q, r, s := self.coef.QuoRem64(Pow10Uint64[self.exp-prec])
	if s >= state.Error {
		return Dec128{state: s}
	}

	if self.state != state.Neg && r > 0 {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// RoundTowardZero rounds the decimal to the specified prec using Toward Zero method (https://en.wikipedia.org/wiki/Rounding#Rounding_toward_zero).
//
// Examples:
//
//	RoundTowardZero(1.236, 2) = 1.23
//	RoundTowardZero(1.235, 2) = 1.23
//	RoundTowardZero(1.234, 2) = 1.23
//	RoundTowardZero(-1.234, 2) = -1.23
//	RoundTowardZero(-1.235, 2) = -1.23
//	RoundTowardZero(-1.236, 2) = -1.23
func (self Dec128) RoundTowardZero(prec uint8) Dec128 {
	return self.Trunc(prec)
}

// RoundAwayFromZero rounds the decimal to the specified prec using Away From Zero method (https://en.wikipedia.org/wiki/Rounding#Rounding_away_from_zero).
//
// Examples:
//
//	RoundAwayFromZero(1.236, 2) = 1.24
//	RoundAwayFromZero(1.235, 2) = 1.24
//	RoundAwayFromZero(1.234, 2) = 1.24
//	RoundAwayFromZero(-1.234, 2) = -1.24
//	RoundAwayFromZero(-1.235, 2) = -1.24
//	RoundAwayFromZero(-1.236, 2) = -1.24
func (self Dec128) RoundAwayFromZero(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	q, r, s := self.coef.QuoRem64(Pow10Uint64[self.exp-prec])
	if s >= state.Error {
		return Dec128{state: s}
	}

	if r > 0 {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// RoundHalfTowardZero rounds the decimal to the specified prec using Half Toward Zero method (https://en.wikipedia.org/wiki/Rounding#Rounding_half_toward_zero).
//
// Examples:
//
//	RoundHalfTowardZero(1.236, 2) = 1.24
//	RoundHalfTowardZero(1.235, 2) = 1.23
//	RoundHalfTowardZero(1.234, 2) = 1.23
//	RoundHalfTowardZero(-1.234, 2) = -1.23
//	RoundHalfTowardZero(-1.235, 2) = -1.23
//	RoundHalfTowardZero(-1.236, 2) = -1.24
func (self Dec128) RoundHalfTowardZero(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	factor := Pow10Uint64[self.exp-prec]
	half := factor / 2

	q, r, s := self.coef.QuoRem64(factor)
	if s >= state.Error {
		return Dec128{state: s}
	}

	if half < r {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// RoundHalfAwayFromZero rounds the decimal to the specified prec using Half Away from Zero method (https://en.wikipedia.org/wiki/Rounding#Rounding_half_away_from_zero).
//
// Examples:
//
//	RoundHalfAwayFromZero(1.236, 2) = 1.24
//	RoundHalfAwayFromZero(1.235, 2) = 1.24
//	RoundHalfAwayFromZero(1.234, 2) = 1.23
//	RoundHalfAwayFromZero(-1.234, 2) = -1.23
//	RoundHalfAwayFromZero(-1.235, 2) = -1.24
//	RoundHalfAwayFromZero(-1.236, 2) = -1.24
func (self Dec128) RoundHalfAwayFromZero(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	factor := Pow10Uint64[self.exp-prec]
	half := factor / 2

	q, r, s := self.coef.QuoRem64(factor)
	if s >= state.Error {
		return Dec128{state: s}
	}

	if half <= r {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// RoundBank uses half up to even (banker's rounding) to round the decimal to the specified precision.
//
// Examples:
//
//	RoundBank(2.121, 2) = 2.12 ; rounded down
//	RoundBank(2.125, 2) = 2.12 ; rounded down, rounding digit is an even number
//	RoundBank(2.135, 2) = 2.14 ; rounded up, rounding digit is an odd number
//	RoundBank(2.1351, 2) = 2.14; rounded up
//	RoundBank(2.127, 2) = 2.13 ; rounded up
func (self Dec128) RoundBank(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	factor := Pow10Uint64[self.exp-prec]
	half := factor / 2

	q, r, s := self.coef.QuoRem64(factor)
	if s >= state.Error {
		return Dec128{state: s}
	}

	if half < r || (half == r && q.Lo%2 == 1) {
		q, s = q.Add64(1)
		if s >= state.Error {
			return Dec128{state: s}
		}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}

// Trunc returns 'self' after truncating the decimal to the specified precision.
//
// Examples:
//
//	Trunc(1.12345, 4) = 1.1234
//	Trunc(1.12335, 4) = 1.1233
func (self Dec128) Trunc(prec uint8) Dec128 {
	if self.state >= state.Error || prec >= self.exp {
		return self
	}

	q, _, s := self.coef.QuoRem64(Pow10Uint64[self.exp-prec])
	if s >= state.Error {
		return Dec128{state: s}
	}

	return Dec128{coef: q, exp: prec, state: self.state}
}
