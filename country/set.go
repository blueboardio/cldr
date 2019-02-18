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

package country

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/blueboardio/cldr/v2/currency"
)

// Set represents a set of country codes.
// A set can be used to match a country code against the set. The zero value (nil) matches any country.
//
// The JSON representation (MarshalJSON) is an array of strings or null.
// The text representation (MarshalText) is codes separated by comma.
// The string representation (.String) is the text representation but with "*" if the set is empty.
type Set []Code

// Any returns a Set where .MatchesAny() == true
func Any() Set {
	return nil
}

// Copy returns an independent copy.
func (cs Set) Copy() Set {
	if cs == nil {
		return nil
	}
	return append(Set(nil), cs...)
}

var ErrInvalidSet = errors.New("invalid countries set")

// Len implements sort.Interface.
func (cs Set) Len() int {
	return len(cs)
}

// Less implements sort.Interface.
func (cs Set) Less(i, j int) bool {
	return cs[i] < cs[j]
}

// Swap implements sort.Interface.
func (cs Set) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// String implements fmt.Stringer.
func (cs Set) String() string {
	if cs == nil {
		return "*"
	}
	b, _ := cs.MarshalText()
	return string(b)
}

// Set implements flag.Value.
func (cs *Set) Set(src string) error {
	if src == "*" {
		*cs = nil
		return nil
	}
	return cs.UnmarshalText([]byte(src))
}

// MarshalText implements encoding.TextMarshaler.
func (cs Set) MarshalText() ([]byte, error) {
	var b []byte
	for _, c := range cs {
		b = append(b, c...)
		b = append(b, ',')
	}
	if len(b) > 0 {
		b = b[:len(b)-1] // Remove the last comma
	}
	return b, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (cs *Set) UnmarshalText(b []byte) error {
	switch len(b) {
	case 0:
		*cs = Set{}
		return nil
	case 1:
		if b[0] == '*' {
			*cs = nil
			return nil
		}
		return ErrInvalidSet
	case 2:
		c := Code(b)
		if !c.IsValid() {
			return ErrInvalidSet
		}
		*cs = Set{c}
	default:
		if (len(b)-2)%3 != 0 {
			return ErrInvalidSet
		}
		set := make(Set, 1+(len(b)-2)/3)
		set[0] = Code(b[0:2])
		if !set[0].IsValid() {
			return ErrInvalidSet
		}
		for i := 1; i < len(set); i++ {
			if b[3*i-1] != ',' {
				return ErrInvalidSet
			}
			set[i] = Code(b[3*i : 3*i+2])
			if !set[i].IsValid() {
				return ErrInvalidSet
			}
		}
		*cs = set
	}
	return nil
}

// MarshalJSON implements encoding/json.Marshaler.
func (cs Set) MarshalJSON() ([]byte, error) {
	if cs == nil {
		return []byte("null"), nil
	}
	if len(cs) == 0 {
		return []byte("[]"), nil
	}
	b := make([]byte, len(cs)*(1+2+1+1)-1+2)
	b[0] = '['
	p := 1
	for _, c := range cs {
		b[p] = '"'
		b[p+1] = c[0]
		b[p+2] = c[1]
		b[p+3] = '"'
		b[p+4] = ','
		p += 5
	}
	b[len(b)-1] = ']' // Overwrite the last comma
	return b, nil
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
func (cs *Set) UnmarshalJSON(b []byte) error {
	var raw []Code // anonymous
	err := json.Unmarshal(b, &raw)
	if err != nil {
		if err == ErrInvalidCountryCode {
			var raw2 []string
			_ = json.Unmarshal(b, &raw2)
			for _, c := range raw2 {
				if !Code(c).IsValid() {
					return fmt.Errorf("%q: invalid country code", c)
				}
			}
		}
		return err
	}
	*cs = raw
	return nil
}

// Value() implements database/sql/driver.Valuer.
func (cs Set) Value() (driver.Value, error) {
	if cs == nil {
		return nil, nil
	}
	return cs.MarshalJSON()
}

func (cs *Set) Scan(src interface{}) error {
	var b []byte
	switch src := src.(type) {
	case nil:
		*cs = nil
		return nil
	case []byte:
		if len(src) == 0 {
			*cs = Set{}
			return nil
		}
		b = src
	case string:
		if len(src) == 0 {
			*cs = Set{}
			return nil
		}
		b = []byte(src)
	default:
		return fmt.Errorf("invalid type %T for type %T", src, *cs)
	}
	if len(b) < 2 {
		return fmt.Errorf("invalid value %q for type %T", b, *cs)
	}
	// Provision for a move to a simpler representation
	if b[0] != '[' {
		return cs.UnmarshalText(b)
	}
	return cs.UnmarshalJSON(b)
}

// Contains returns true if cs contains c.
func (cs Set) Contains(c Code) bool {
	for _, cc := range cs {
		if cc == c {
			return true
		}
	}
	return false
}

// MatchesAny returns true if the set is empty.
func (cs Set) MatchesAny() bool {
	return cs == nil
}

// Matches returns true if the set is empty (matches anything) or contains c.
//
// Same as Contains except for the empty set case which return true for any country.
func (cs Set) Matches(c Code) bool {
	if cs.MatchesAny() {
		return true
	}
	return cs.Contains(c)
}

// Add appends c if not already contained in the set.
func (cs *Set) Add(c Code) {
	if cs.Matches(c) {
		return
	}
	*cs = append(*cs, c)
}

// Remove removes from cs the intersection of the two sets.
// If exclude is empty nothing is removed.
// Order is preserved.
func (cs *Set) Remove(exclude Set) {
	if exclude.MatchesAny() {
		if cs.MatchesAny() {
			*cs = Set{}
		} else {
			*cs = (*cs)[:0]
		}
		return
	}
	if len(exclude) == 0 {
		return
	}
	s := *cs
	removed := 0
	for i, cc := range s {
		if exclude.Contains(cc) {
			removed++
		} else {
			if removed > 0 { // shift
				s[i-removed] = s[i]
			}
		}
	}
	if removed > 0 {
		*cs = s[:len(s)-removed]
	}
}

// Filter removes countries that are not Matched by filter.
//
// The filter is assumed to not contain duplicates.
// An nil filter is a no-op (matches anything).
// Order is preserved.
func (cs *Set) Filter(filter Set) {
	if filter.MatchesAny() {
		return
	}
	if len(filter) == 0 {
		if *cs == nil {
			*cs = Set{}
		} else {
			*cs = (*cs)[:0]
		}
		return
	}
	if cs.MatchesAny() {
		*cs = filter.Copy()
		return
	}
	s := *cs
	removed := 0
Next:
	for i, cc := range s {
		for j, f := range filter {
			if cc == f {
				if removed > 0 { // shift
					s[i-removed] = s[i]
				}
				if j == 0 {
					filter = filter[1:]
				}
				continue Next
			}
		}
		removed++
	}
	if removed > 0 {
		*cs = s[:len(s)-removed]
	}
}

// HasDuplicates checks if cs contains duplicates.
func (cs Set) HasDuplicates() bool {
	l := len(cs)
	if l <= 1 {
		return false
	}
	for i := 1; i < l; i++ {
		if cs[:i].Contains(cs[i]) {
			return true
		}
	}
	return false
}

// RemoveDuplicates remove duplicates.
// Order is preserved. In case of a duplicate the first one is kept.
func (cs *Set) RemoveDuplicates() {
	l := len(*cs)
	if l <= 1 {
		return
	}
	removed := 0
	for i := 1; i < l; i++ {
		if (*cs)[:i-removed].Contains((*cs)[i]) {
			removed++
		} else {
			if removed > 0 {
				(*cs)[i-removed] = (*cs)[i]
			}
		}
	}
	if removed > 0 {
		*cs = (*cs)[:len(*cs)-removed]
	}
}

// Currencies returns the set of currencies of the countries set.
func (cs Set) Currencies() currency.Set {
	if cs.MatchesAny() {
		return currency.Any() // any currencies
	}
	cur := make(map[currency.Code]struct{}, len(cs))
	for _, c := range cs {
		for _, cu := range Countries[c].Currencies {
			cur[cu] = struct{}{}
		}
	}
	currencies := make(currency.Set, 0, len(cur))
	for cu := range cur {
		currencies = append(currencies, cu)
	}
	return currencies
}
