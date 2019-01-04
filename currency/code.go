/*
Copyright (c) 2018 BlueBoard SAS.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package currency

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

var ErrInvalidCurrencyCode = errors.New("invalid currency code")

// Code represents an ISO 4217 currency code: 3 ASCII letters, uppercase.
// See https://www.iso.org/fr/iso-4217-currency-codes.html.
type Code string

// IsValid returns false if the code is invalid.
func (cc Code) IsValid() bool {
	/*
		return len(cc) == 3 &&
			cc[0] >= 'A' && cc[0] <= 'Z' &&
			cc[1] >= 'A' && cc[1] <= 'Z' &&
			cc[2] >= 'A' && cc[2] <= 'Z'
	*/
	_, ok := ActiveCurrencies[cc]
	return ok
}

// String implements fmt.Stringer.
func (cc Code) String() string {
	return string(cc)
}

// Set implements flag.Value.
func (cc *Code) Set(src string) error {
	if !Code(src).IsValid() {
		return ErrInvalidCurrencyCode
	}
	*cc = Code(src)
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (cc Code) MarshalText() ([]byte, error) {
	return []byte(cc), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (cc *Code) UnmarshalText(src []byte) error {
	return cc.Set(string(src))
}

// Value implements database/sql/driver.Valuer.
func (cc Code) Value() (driver.Value, error) {
	if len(cc) == 0 {
		return nil, nil
	}
	return string(cc), nil
}

// Scan implements database/sql.Scanner.
func (cc *Code) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		*cc = ""
	case []byte:
		*cc = Code(src)
	case string:
		*cc = Code(src)
	default:
		return fmt.Errorf("unexpected value of type %T for %T", src, *cc)
	}
	return nil
}

type FractionInfo struct {
	Digits       uint
	Rounding     uint
	CashDigits   uint
	CashRounding uint
}

type Info struct {
	Code Code

	Fraction FractionInfo

	Countries []string
}
