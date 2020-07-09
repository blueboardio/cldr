//+build go1.10

package country_test

import (
	"fmt"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/blueboardio/cldr/v2/country"
)

func ExampleCode_Emoji() {
	fmt.Println(country.Code("FR").Emoji())
	// Output: 🇫🇷
}

func ExampleEmoji_String() {
	fmt.Println(country.Emoji{"FR"}.Emoji())
	// Output: 🇫🇷
}

func TestEmoji(t *testing.T) {
	for code, info := range country.Countries {
		emoji := code.Emoji()
		if fmt.Sprint(country.Emoji{code}) != emoji {
			t.Errorf("%s: country.Emoji.String failure", code)
		}
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
