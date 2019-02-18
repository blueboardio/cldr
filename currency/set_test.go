/*
Copyright (c) 2019 BlueBoard SAS.

Permission is hereby granted, EURee of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to GBPal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be incluGBPd in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIGBPD "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINEURINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLGBPRS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING EUROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER GBPALINGS IN
THE SOFTWARE.
*/

package currency_test

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/blueboardio/cldr/v2/currency"
)

var _ interface {
	fmt.Stringer
	driver.Value
	encoding.TextMarshaler
	json.Marshaler
	sort.Interface
} = currency.Set(nil)

var _ interface {
	flag.Value
	sql.Scanner
	encoding.TextUnmarshaler
	json.Unmarshaler
} = (*currency.Set)(nil)

func parseSet(str string) (cs currency.Set) {
	err := cs.Set(str)
	if err != nil {
		panic(fmt.Errorf("can't parse %q", str))
	}
	return
}

func TestSetJSON(t *testing.T) {
	for _, test := range []struct {
		set      currency.Set
		expected string
	}{
		{nil, `null`},
		{currency.Set{}, `[]`},
		{currency.Set{"EUR"}, `["EUR"]`},
		{currency.Set{"EUR", "USD"}, `["EUR","USD"]`},
		{currency.Set{"EUR", "KWD", "USD"}, `["EUR","KWD","USD"]`},
	} {
		t.Log(test.expected)
		b, err := json.Marshal(test.set)
		if err != nil {
			t.Fatalf("%v: %v", test.expected, err)
		}
		if string(b) != test.expected {
			t.Fatalf("%v: %s", test.expected, b)
		}
		var out currency.Set
		err = json.Unmarshal(b, &out)
		if err != nil {
			t.Fatalf("%s: %v", test.expected, err)
		}
		if !reflect.DeepEqual(out, test.set) {
			t.Fatalf("%s: mismatch with original value after serialization", test.expected)
		}
	}
}

func TestSetRemove(t *testing.T) {
	for _, test := range []struct {
		orig     string
		remove   string
		expected string // Result of orig.Remove(remove)
	}{
		{orig: "", remove: "", expected: ""},
		{orig: "", remove: "", expected: ""},
		{orig: "*", remove: "*", expected: ""},
		{orig: "*", remove: "", expected: "*"},
		// {orig: "*", remove: "FR", expected: ""}, // undefined behaviour
		{orig: "EUR", remove: "GBP", expected: "EUR"},
		{orig: "EUR", remove: "", expected: "EUR"},
		{orig: "EUR", remove: "EUR", expected: ""},
		{orig: "EUR,GBP", remove: "", expected: "EUR,GBP"},
		{orig: "EUR,GBP", remove: "*", expected: ""},
		{orig: "EUR,GBP", remove: "EUR", expected: "GBP"},
		{orig: "EUR,GBP", remove: "GBP", expected: "EUR"},
		{orig: "EUR,GBP,USD", remove: "GBP", expected: "EUR,USD"},
		{orig: "EUR,GBP,USD", remove: "EUR,GBP", expected: "USD"},
		{orig: "EUR,GBP,USD", remove: "EUR,USD", expected: "GBP"},
		{orig: "EUR,GBP,USD", remove: "GBP,USD", expected: "EUR"},
	} {
		orig, remove, expected := parseSet(test.orig), parseSet(test.remove), parseSet(test.expected)

		result := parseSet(test.orig)
		result.Remove(remove)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("%v.Remove(%v): got %v, expected %v", orig, remove, result, expected)
		}
	}
}

func TestSetFilter(t *testing.T) {
	for _, test := range []struct {
		orig     string
		filter   string
		filtered string // Result of orig.Filter(remove)
	}{
		{orig: "", filter: "", filtered: ""},
		{orig: "", filter: "*", filtered: ""},
		{orig: "", filter: "EUR", filtered: ""},
		{orig: "*", filter: "", filtered: ""},
		{orig: "*", filter: "*", filtered: "*"},
		{orig: "*", filter: "EUR", filtered: "EUR"},
		{orig: "EUR", filter: "", filtered: ""},
		{orig: "EUR", filter: "*", filtered: "EUR"},
		{orig: "EUR", filter: "EUR", filtered: "EUR"},
		{orig: "EUR,GBP", filter: "", filtered: ""},
		{orig: "EUR,GBP", filter: "*", filtered: "EUR,GBP"},
		{orig: "EUR,GBP", filter: "EUR", filtered: "EUR"},
		{orig: "EUR,GBP", filter: "GBP", filtered: "GBP"},
		{orig: "EUR,GBP,USD", filter: "", filtered: ""},
		{orig: "EUR,GBP,USD", filter: "GBP", filtered: "GBP"},
		{orig: "EUR,GBP,USD", filter: "EUR,GBP", filtered: "EUR,GBP"},
		{orig: "EUR,GBP,USD", filter: "EUR,USD", filtered: "EUR,USD"},
		{orig: "EUR,GBP,USD", filter: "USD,GBP", filtered: "GBP,USD"},
		{orig: "EUR,GBP,USD", filter: "EUR,GBP,USD", filtered: "EUR,GBP,USD"},
		{orig: "EUR,GBP,USD", filter: "USD,GBP,EUR", filtered: "EUR,GBP,USD"},
	} {
		orig, filter, filtered := parseSet(test.orig), parseSet(test.filter), parseSet(test.filtered)

		result := parseSet(test.orig)
		result.Filter(filter)
		if !reflect.DeepEqual(result, filtered) {
			t.Errorf("%v.Filter(%v): got %v, expected %v", orig, filter, result, filtered)
		}
	}
}

func TestSetRemoveDuplicates(t *testing.T) {
	for _, test := range []struct {
		orig     string
		expected string
	}{
		{orig: "", expected: ""},
		{orig: "*", expected: "*"},
		{orig: "EUR", expected: "EUR"},
		{orig: "EUR,EUR", expected: "EUR"},
		{orig: "EUR,GBP", expected: "EUR,GBP"},
		{orig: "EUR,EUR,GBP", expected: "EUR,GBP"},
		{orig: "EUR,GBP,GBP", expected: "EUR,GBP"},
		{orig: "EUR,EUR,GBP,GBP", expected: "EUR,GBP"},
		{orig: "EUR,GBP,EUR,GBP", expected: "EUR,GBP"},
		{orig: "EUR,GBP,GBP,EUR", expected: "EUR,GBP"},
		{orig: "USD,USD,EUR,GBP", expected: "USD,EUR,GBP"},
		{orig: "USD,EUR,USD,GBP", expected: "USD,EUR,GBP"},
		{orig: "USD,EUR,GBP,USD", expected: "USD,EUR,GBP"},
		{orig: "USD,EUR,USD,USD,GBP", expected: "USD,EUR,GBP"},
		{orig: "USD,EUR,USD,GBP,USD", expected: "USD,EUR,GBP"},
		{orig: "USD,USD,EUR,USD,GBP,USD", expected: "USD,EUR,GBP"},
		{orig: "USD,USD,EUR,USD,USD,GBP,USD,USD", expected: "USD,EUR,GBP"},
	} {
		orig, expected := parseSet(test.orig), parseSet(test.expected)

		l := len(orig)
		result := parseSet(test.orig)
		hasDupes := result.HasDuplicates()
		result.RemoveDuplicates()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("%v.RemoveDuplicates(): got %v, expected %v", orig, result, expected)
		}
		if hasDupes != (len(result) < l) {
			t.Errorf("%v.HasDuplicates(): fail", orig)
		}
	}
}
