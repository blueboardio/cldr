//+build go1.10

package country_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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

func emoji1(cc string) string {
	buf := [...]byte{240, 159, 135, 0, 240, 159, 135, 0}
	buf[7] = cc[1] + (166 - 'A')
	buf[3] = cc[0] + (166 - 'A')
	return string(buf[:])
}

func emoji2(cc string) string {
	return string([]byte{
		240, 159, 135, cc[0] + (166 - 'A'),
		240, 159, 135, cc[1] + (166 - 'A'),
	})
}

var sss string

func BenchmarkEmoji(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.Run("emoji1", func(b *testing.B) {
		cc := string([]byte{'A' + byte(rand.Intn(26)), 'A' + byte(rand.Intn(26))})
		sss = emoji1(cc)
	})
	b.Run("emoji2", func(b *testing.B) {
		cc := string([]byte{'A' + byte(rand.Intn(26)), 'A' + byte(rand.Intn(26))})
		sss = emoji2(cc)
	})
}
