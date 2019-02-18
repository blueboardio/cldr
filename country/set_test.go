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

package country_test

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

	"github.com/blueboardio/cldr/v2/country"
	"github.com/blueboardio/cldr/v2/currency"
)

var _ interface {
	fmt.Stringer
	driver.Value
	encoding.TextMarshaler
	json.Marshaler
	sort.Interface
} = country.Set(nil)

var _ interface {
	flag.Value
	sql.Scanner
	encoding.TextUnmarshaler
	json.Unmarshaler
} = (*country.Set)(nil)

func parseSet(str string) (cs country.Set) {
	err := cs.Set(str)
	if err != nil {
		panic(fmt.Errorf("can't parse %q", str))
	}
	return
}

func TestSetJSON(t *testing.T) {
	for _, test := range []struct {
		set      country.Set
		expected string
	}{
		{nil, `null`},
		{country.Set{}, `[]`},
		{country.Set{"FR"}, `["FR"]`},
		{country.Set{"FR", "US"}, `["FR","US"]`},
		{country.Set{"DE", "FR", "IT"}, `["DE","FR","IT"]`},
	} {
		b, err := json.Marshal(test.set)
		if err != nil {
			t.Fatalf("%v: %v", test.expected, err)
		}
		if string(b) != test.expected {
			t.Fatalf("%v: %s", test.expected, b)
		}

		var out country.Set
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
		{orig: "FR", remove: "DE", expected: "FR"},
		{orig: "FR", remove: "", expected: "FR"},
		{orig: "FR", remove: "FR", expected: ""},
		{orig: "FR,DE", remove: "", expected: "FR,DE"},
		{orig: "FR,DE", remove: "*", expected: ""},
		{orig: "FR,DE", remove: "FR", expected: "DE"},
		{orig: "FR,DE", remove: "DE", expected: "FR"},
		{orig: "FR,DE,GB", remove: "DE", expected: "FR,GB"},
		{orig: "FR,DE,GB", remove: "FR,DE", expected: "GB"},
		{orig: "FR,DE,GB", remove: "FR,GB", expected: "DE"},
		{orig: "FR,DE,GB", remove: "GB,DE", expected: "FR"},
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
		{orig: "", filter: "FR", filtered: ""},
		{orig: "*", filter: "", filtered: ""},
		{orig: "*", filter: "*", filtered: "*"},
		{orig: "*", filter: "FR", filtered: "FR"},
		{orig: "FR", filter: "", filtered: ""},
		{orig: "FR", filter: "*", filtered: "FR"},
		{orig: "FR", filter: "FR", filtered: "FR"},
		{orig: "FR,DE", filter: "", filtered: ""},
		{orig: "FR,DE", filter: "*", filtered: "FR,DE"},
		{orig: "FR,DE", filter: "FR", filtered: "FR"},
		{orig: "FR,DE", filter: "DE", filtered: "DE"},
		{orig: "FR,DE,GB", filter: "", filtered: ""},
		{orig: "FR,DE,GB", filter: "*", filtered: "FR,DE,GB"},
		{orig: "FR,DE,GB", filter: "DE", filtered: "DE"},
		{orig: "FR,DE,GB", filter: "FR,DE", filtered: "FR,DE"},
		{orig: "FR,DE,GB", filter: "FR,GB", filtered: "FR,GB"},
		{orig: "FR,DE,GB", filter: "GB,DE", filtered: "DE,GB"},
		{orig: "FR,DE,GB", filter: "", filtered: ""},
		{orig: "FR,DE,GB", filter: "*", filtered: "FR,DE,GB"},
		{orig: "FR,DE,GB", filter: "FR,DE,GB", filtered: "FR,DE,GB"},
		{orig: "FR,DE,GB", filter: "GB,DE,FR", filtered: "FR,DE,GB"},
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
		{orig: "FR", expected: "FR"},
		{orig: "FR,FR", expected: "FR"},
		{orig: "FR,DE", expected: "FR,DE"},
		{orig: "FR,FR,DE", expected: "FR,DE"},
		{orig: "FR,DE,DE", expected: "FR,DE"},
		{orig: "FR,FR,DE,DE", expected: "FR,DE"},
		{orig: "FR,DE,FR,DE", expected: "FR,DE"},
		{orig: "FR,DE,DE,FR", expected: "FR,DE"},
		{orig: "GB,GB,FR,DE", expected: "GB,FR,DE"},
		{orig: "GB,FR,GB,DE", expected: "GB,FR,DE"},
		{orig: "GB,FR,DE,GB", expected: "GB,FR,DE"},
		{orig: "GB,FR,GB,GB,DE", expected: "GB,FR,DE"},
		{orig: "GB,FR,GB,DE,GB", expected: "GB,FR,DE"},
		{orig: "GB,GB,FR,GB,DE,GB", expected: "GB,FR,DE"},
		{orig: "GB,GB,FR,GB,GB,DE,GB,GB", expected: "GB,FR,DE"},
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

func currenciesString(currencies []currency.Code) string {
	if currencies == nil {
		return "*"
	}
	if len(currencies) == 0 {
		return ""
	}
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i] < currencies[j]
	})
	s := make([]byte, 4*len(currencies)-1)
	copy(s, currencies[0])
	for i := 1; i < len(currencies); i++ {
		s[4*i-1] = ','
		cu := currencies[i]
		copy(s[4*i:4*i+3], cu)
		//s[4*i] = cu[0]
		//s[4*i+1] = cu[1]
		//s[4*i+2] = cu[2]
	}
	return string(s)
}

func TestSetCurrencies(t *testing.T) {
	if parseSet("*").Currencies() != nil {
		t.Errorf("nil expected for *")
	}
	for _, test := range []struct {
		countries  string
		currencies string
	}{
		{countries: "*", currencies: "*"},
		{countries: "FR", currencies: "EUR"},
		{countries: "FR,DE", currencies: "EUR"},
		{countries: "FR,DE,GB", currencies: "EUR,GBP"},
		{countries: "FR,DE,GB,GS", currencies: "EUR,GBP"},
		{countries: "FR,DE,GB,BT", currencies: "BTN,EUR,GBP,INR"},
	} {
		co := parseSet(test.countries)
		cu := currenciesString(co.Currencies())
		if cu != test.currencies {
			t.Errorf("%q: got %q, expected %q", test.countries, cu, test.currencies)
		}
	}
}
