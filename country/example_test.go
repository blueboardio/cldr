package country_test

import (
	"fmt"

	"github.com/blueboardio/cldr/country"
)

func Example() {
	FR := country.Code("FR")
	fmt.Println(FR.Emoji(), country.Countries[FR].Name)

	// Output:
	// 🇫🇷 France
}
