//+build go1.10

package country_test

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/blueboardio/cldr/country"
)

func TestEmoji(t *testing.T) {
	for code, info := range country.Countries {
		emoji := country.Code(code).Emoji()
		t.Log(country.Code(code), emoji, info.Name)
		if len(emoji) != 8 {
			t.Errorf("%s: incorrect length %d", code, len(emoji))
		}
		for {
			r, size := utf8.DecodeRuneInString(emoji)
			if !unicode.Is(unicode.Regional_Indicator, r) {
				t.Errorf("%s: unexpected rune %d", code, r)
			}
			emoji = emoji[size:] // next
			if emoji == "" {
				break
			}
		}
	}
}
