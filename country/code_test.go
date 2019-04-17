/*
Copyright © 2018-2019 BlueBoard SAS.

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

package country_test

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"flag"
	"fmt"
	"testing"

	"github.com/blueboardio/cldr/v2/country"
)

var _ interface {
	fmt.Stringer
	driver.Valuer
} = country.Code("FR")

var _ interface {
	sql.Scanner
	driver.Valuer
	flag.Value
} = (*country.Code)(nil)

var reservedCodes []string

func init() {
	reservedCodes = []string{"AA", "ZZ"}
	c := [2]byte{}
	cs := c[:]
	// QM to QZ
	c[0] = 'Q'
	for b := byte('M'); b <= 'Z'; b++ {
		c[1] = b
		reservedCodes = append(reservedCodes, string(cs))
	}
	c[0] = 'X'
	for b := byte('A'); b <= 'Z'; b++ {
		if b == 'K' { // Kosovo
			continue
		}
		c[1] = b
		reservedCodes = append(reservedCodes, string(cs))
	}
}

func TestCode(t *testing.T) {
	for _, code := range reservedCodes {
		if country.Code(code).IsValid() {
			t.Errorf("%s is a reserved code. IsValid() should return false.", code)
		}
	}

	for code := range country.Countries {
		if !code.IsValid() {
			t.Errorf("%s is a known code. IsValid() should return true.", code)
		}
	}
}

func TestEmojiJSON(t *testing.T) {
	var s struct {
		Country country.Emoji `json:"country"`
	}
	s.Country.Code = "FR"

	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"country":"🇫🇷"}` {
		t.Fatalf("JSON marshaling failure: %s", b)
	}
	t.Logf("%s", b)
}
