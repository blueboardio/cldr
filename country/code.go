/*
Copyright © 2018-2020 BlueBoard SAS.

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

package country

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

var ErrInvalidCountryCode = errors.New("invalid country code")

// Code represents an ISO 3166 country code: 2 ASCII letters, uppercase.
// See https://www.iso.org/fr/iso-3166-country-codes.html.
type Code string

// String implements fmt.Stringer.
func (cc Code) String() string {
	return string(cc)
}

// IsValid returns true if cc is a known country code.
// See Countries.
func (cc Code) IsValid() bool {
	if len(cc) != 2 {
		return false
	}
	_, found := Countries[cc]
	return found
}

// Set implements flag.Value.
func (cc *Code) Set(src string) error {
	if !Code(src).IsValid() {
		return ErrInvalidCountryCode
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

// Emoji converts "FR" to "🇫🇷".
func (cc Code) Emoji() string {
	buf := [...]byte{240, 159, 135, 0, 240, 159, 135, 0}
	buf[3] = cc[0] + (166 - 'A')
	buf[7] = cc[1] + (166 - 'A')
	return string(buf[:])
}

// EmojiRunes returns the 2 runes that represents the emoji for the country code.
func (cc Code) EmojiRunes() (rune, rune) {
	return 0x1f1a5 + rune(cc[0]), 0x1f1a5 + rune(cc[1])
}

// Emoji wraps a country Code to have an external representation as a flag emoji.
//     "FR" => "🇫🇷"
type Emoji struct {
	Code
}

// String implements fmt.Stringer.
func (cc Emoji) String() string {
	return cc.Code.Emoji()
}

// Runes returns the 2 runes that represents the emoji for the country code.
func (cc Emoji) Runes() (rune, rune) {
	return cc.Code.EmojiRunes()
}

// MarshalText implements encoding.TextMarshaler.
func (cc Emoji) MarshalText() ([]byte, error) {
	buf := [...]byte{240, 159, 135, 0, 240, 159, 135, 0}
	buf[3] = cc.Code[0] + (166 - 'A')
	buf[7] = cc.Code[1] + (166 - 'A')
	return buf[:], nil
}

// UnmarshalText forbids the use of Emoji as an encoding.TextUnmarshaler.
func (cc *Emoji) UnmarshalText(b []byte) error {
	if len(b) != 8 ||
		b[0] != 240 || b[4] != 240 ||
		b[1] != 159 || b[5] != 159 ||
		b[2] != 135 || b[6] != 135 {
		return errors.New("invalid country emoji")
	}
	return cc.Code.UnmarshalText([]byte{b[3] - 166 + 'A', b[7] - 166 + 'A'})
}

// Set forbids the use of Emoji as a flag.Value.
func (cc *Emoji) Set(s string) error {
	return cc.UnmarshalText([]byte(s))
}
