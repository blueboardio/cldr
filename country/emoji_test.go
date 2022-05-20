//go:build go1.10
// +build go1.10

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
		var runes []rune
		for {
			r, size := utf8.DecodeRuneInString(emoji)
			if !unicode.Is(unicode.Regional_Indicator, r) {
				t.Errorf("%s: unexpected rune %d", code, r)
			}
			runes = append(runes, r)
			emoji = emoji[size:] // next
			if emoji == "" {
				break
			}
		}
		r0, r1 := code.EmojiRunes()
		if r0 != runes[0] || r1 != runes[1] {
			t.Errorf("%s: invalid runes", code)
		}

		emoji = code.Emoji()
		var cc country.Emoji
		if err := cc.Set(code.Emoji()); err != nil {
			t.Errorf("%s: can't set from emoji: %v", code, err)
		} else if cc.Code != code {
			t.Errorf("%s: mismatch when restoring from emoji: %s", code, cc.Code)
		}
	}
}
