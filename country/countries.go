// Code generated by "go run countries_gen.go cldr-common-42.0.zip"; DO NOT EDIT.

package country

import "github.com/blueboardio/cldr/v2/currency"

type Info struct {
	Code       Code
	Name       string
	Currencies []currency.Code
}

// Countries is the list of countries from Unicode CLDR.
//
// Source: cldr-common-42.0.zip
//
// The following codes are removed:
//
//	QO (duplicate of UM)
//	ZZ (unknown)
var Countries = map[Code]*Info{
	"AC": {Code: "AC", Name: "Ascension Island", Currencies: []currency.Code{"SHP"}},
	"AD": {Code: "AD", Name: "Andorra", Currencies: []currency.Code{"EUR"}},
	"AE": {Code: "AE", Name: "United Arab Emirates", Currencies: []currency.Code{"AED"}},
	"AF": {Code: "AF", Name: "Afghanistan", Currencies: []currency.Code{"AFN"}},
	"AG": {Code: "AG", Name: "Antigua & Barbuda", Currencies: []currency.Code{"XCD"}},
	"AI": {Code: "AI", Name: "Anguilla", Currencies: []currency.Code{"XCD"}},
	"AL": {Code: "AL", Name: "Albania", Currencies: []currency.Code{"ALL"}},
	"AM": {Code: "AM", Name: "Armenia", Currencies: []currency.Code{"AMD"}},
	"AO": {Code: "AO", Name: "Angola", Currencies: []currency.Code{"AOA"}},
	"AQ": {Code: "AQ", Name: "Antarctica"},
	"AR": {Code: "AR", Name: "Argentina", Currencies: []currency.Code{"ARS"}},
	"AS": {Code: "AS", Name: "American Samoa", Currencies: []currency.Code{"USD"}},
	"AT": {Code: "AT", Name: "Austria", Currencies: []currency.Code{"EUR"}},
	"AU": {Code: "AU", Name: "Australia", Currencies: []currency.Code{"AUD"}},
	"AW": {Code: "AW", Name: "Aruba", Currencies: []currency.Code{"AWG"}},
	"AX": {Code: "AX", Name: "Åland Islands", Currencies: []currency.Code{"EUR"}},
	"AZ": {Code: "AZ", Name: "Azerbaijan", Currencies: []currency.Code{"AZN"}},
	"BA": {Code: "BA", Name: "Bosnia & Herzegovina", Currencies: []currency.Code{"BAM"}},
	"BB": {Code: "BB", Name: "Barbados", Currencies: []currency.Code{"BBD"}},
	"BD": {Code: "BD", Name: "Bangladesh", Currencies: []currency.Code{"BDT"}},
	"BE": {Code: "BE", Name: "Belgium", Currencies: []currency.Code{"EUR"}},
	"BF": {Code: "BF", Name: "Burkina Faso", Currencies: []currency.Code{"XOF"}},
	"BG": {Code: "BG", Name: "Bulgaria", Currencies: []currency.Code{"BGN"}},
	"BH": {Code: "BH", Name: "Bahrain", Currencies: []currency.Code{"BHD"}},
	"BI": {Code: "BI", Name: "Burundi", Currencies: []currency.Code{"BIF"}},
	"BJ": {Code: "BJ", Name: "Benin", Currencies: []currency.Code{"XOF"}},
	"BL": {Code: "BL", Name: "St. Barthélemy", Currencies: []currency.Code{"EUR"}},
	"BM": {Code: "BM", Name: "Bermuda", Currencies: []currency.Code{"BMD"}},
	"BN": {Code: "BN", Name: "Brunei", Currencies: []currency.Code{"BND"}},
	"BO": {Code: "BO", Name: "Bolivia", Currencies: []currency.Code{"BOB"}},
	"BQ": {Code: "BQ", Name: "Caribbean Netherlands", Currencies: []currency.Code{"USD"}},
	"BR": {Code: "BR", Name: "Brazil", Currencies: []currency.Code{"BRL"}},
	"BS": {Code: "BS", Name: "Bahamas", Currencies: []currency.Code{"BSD"}},
	"BT": {Code: "BT", Name: "Bhutan", Currencies: []currency.Code{"BTN", "INR"}},
	"BV": {Code: "BV", Name: "Bouvet Island", Currencies: []currency.Code{"NOK"}},
	"BW": {Code: "BW", Name: "Botswana", Currencies: []currency.Code{"BWP"}},
	"BY": {Code: "BY", Name: "Belarus", Currencies: []currency.Code{"BYN"}},
	"BZ": {Code: "BZ", Name: "Belize", Currencies: []currency.Code{"BZD"}},
	"CA": {Code: "CA", Name: "Canada", Currencies: []currency.Code{"CAD"}},
	"CC": {Code: "CC", Name: "Cocos (Keeling) Islands", Currencies: []currency.Code{"AUD"}},
	"CD": {Code: "CD", Name: "Congo - Kinshasa", Currencies: []currency.Code{"CDF"}},
	"CF": {Code: "CF", Name: "Central African Republic", Currencies: []currency.Code{"XAF"}},
	"CG": {Code: "CG", Name: "Congo - Brazzaville", Currencies: []currency.Code{"XAF"}},
	"CH": {Code: "CH", Name: "Switzerland", Currencies: []currency.Code{"CHF"}},
	"CI": {Code: "CI", Name: "Côte d’Ivoire", Currencies: []currency.Code{"XOF"}},
	"CK": {Code: "CK", Name: "Cook Islands", Currencies: []currency.Code{"NZD"}},
	"CL": {Code: "CL", Name: "Chile", Currencies: []currency.Code{"CLP"}},
	"CM": {Code: "CM", Name: "Cameroon", Currencies: []currency.Code{"XAF"}},
	"CN": {Code: "CN", Name: "China", Currencies: []currency.Code{"CNY"}},
	"CO": {Code: "CO", Name: "Colombia", Currencies: []currency.Code{"COP"}},
	"CP": {Code: "CP", Name: "Clipperton Island"},
	"CR": {Code: "CR", Name: "Costa Rica", Currencies: []currency.Code{"CRC"}},
	"CU": {Code: "CU", Name: "Cuba", Currencies: []currency.Code{"CUP", "CUC"}},
	"CV": {Code: "CV", Name: "Cape Verde", Currencies: []currency.Code{"CVE"}},
	"CW": {Code: "CW", Name: "Curaçao", Currencies: []currency.Code{"ANG"}},
	"CX": {Code: "CX", Name: "Christmas Island", Currencies: []currency.Code{"AUD"}},
	"CY": {Code: "CY", Name: "Cyprus", Currencies: []currency.Code{"EUR"}},
	"CZ": {Code: "CZ", Name: "Czechia", Currencies: []currency.Code{"CZK"}},
	"DE": {Code: "DE", Name: "Germany", Currencies: []currency.Code{"EUR"}},
	"DG": {Code: "DG", Name: "Diego Garcia", Currencies: []currency.Code{"USD"}},
	"DJ": {Code: "DJ", Name: "Djibouti", Currencies: []currency.Code{"DJF"}},
	"DK": {Code: "DK", Name: "Denmark", Currencies: []currency.Code{"DKK"}},
	"DM": {Code: "DM", Name: "Dominica", Currencies: []currency.Code{"XCD"}},
	"DO": {Code: "DO", Name: "Dominican Republic", Currencies: []currency.Code{"DOP"}},
	"DZ": {Code: "DZ", Name: "Algeria", Currencies: []currency.Code{"DZD"}},
	"EA": {Code: "EA", Name: "Ceuta & Melilla", Currencies: []currency.Code{"EUR"}},
	"EC": {Code: "EC", Name: "Ecuador", Currencies: []currency.Code{"USD"}},
	"EE": {Code: "EE", Name: "Estonia", Currencies: []currency.Code{"EUR"}},
	"EG": {Code: "EG", Name: "Egypt", Currencies: []currency.Code{"EGP"}},
	"EH": {Code: "EH", Name: "Western Sahara", Currencies: []currency.Code{"MAD"}},
	"ER": {Code: "ER", Name: "Eritrea", Currencies: []currency.Code{"ERN"}},
	"ES": {Code: "ES", Name: "Spain", Currencies: []currency.Code{"EUR"}},
	"ET": {Code: "ET", Name: "Ethiopia", Currencies: []currency.Code{"ETB"}},
	"EU": {Code: "EU", Name: "European Union", Currencies: []currency.Code{"EUR"}},
	"EZ": {Code: "EZ", Name: "Eurozone"},
	"FI": {Code: "FI", Name: "Finland", Currencies: []currency.Code{"EUR"}},
	"FJ": {Code: "FJ", Name: "Fiji", Currencies: []currency.Code{"FJD"}},
	"FK": {Code: "FK", Name: "Falkland Islands", Currencies: []currency.Code{"FKP"}},
	"FM": {Code: "FM", Name: "Micronesia", Currencies: []currency.Code{"USD"}},
	"FO": {Code: "FO", Name: "Faroe Islands", Currencies: []currency.Code{"DKK"}},
	"FR": {Code: "FR", Name: "France", Currencies: []currency.Code{"EUR"}},
	"GA": {Code: "GA", Name: "Gabon", Currencies: []currency.Code{"XAF"}},
	"GB": {Code: "GB", Name: "United Kingdom", Currencies: []currency.Code{"GBP"}},
	"GD": {Code: "GD", Name: "Grenada", Currencies: []currency.Code{"XCD"}},
	"GE": {Code: "GE", Name: "Georgia", Currencies: []currency.Code{"GEL"}},
	"GF": {Code: "GF", Name: "French Guiana", Currencies: []currency.Code{"EUR"}},
	"GG": {Code: "GG", Name: "Guernsey", Currencies: []currency.Code{"GBP"}},
	"GH": {Code: "GH", Name: "Ghana", Currencies: []currency.Code{"GHS"}},
	"GI": {Code: "GI", Name: "Gibraltar", Currencies: []currency.Code{"GIP"}},
	"GL": {Code: "GL", Name: "Greenland", Currencies: []currency.Code{"DKK"}},
	"GM": {Code: "GM", Name: "Gambia", Currencies: []currency.Code{"GMD"}},
	"GN": {Code: "GN", Name: "Guinea", Currencies: []currency.Code{"GNF"}},
	"GP": {Code: "GP", Name: "Guadeloupe", Currencies: []currency.Code{"EUR"}},
	"GQ": {Code: "GQ", Name: "Equatorial Guinea", Currencies: []currency.Code{"XAF"}},
	"GR": {Code: "GR", Name: "Greece", Currencies: []currency.Code{"EUR"}},
	"GS": {Code: "GS", Name: "South Georgia & South Sandwich Islands", Currencies: []currency.Code{"GBP"}},
	"GT": {Code: "GT", Name: "Guatemala", Currencies: []currency.Code{"GTQ"}},
	"GU": {Code: "GU", Name: "Guam", Currencies: []currency.Code{"USD"}},
	"GW": {Code: "GW", Name: "Guinea-Bissau", Currencies: []currency.Code{"XOF"}},
	"GY": {Code: "GY", Name: "Guyana", Currencies: []currency.Code{"GYD"}},
	"HK": {Code: "HK", Name: "Hong Kong SAR China", Currencies: []currency.Code{"HKD"}},
	"HM": {Code: "HM", Name: "Heard & McDonald Islands", Currencies: []currency.Code{"AUD"}},
	"HN": {Code: "HN", Name: "Honduras", Currencies: []currency.Code{"HNL"}},
	"HR": {Code: "HR", Name: "Croatia", Currencies: []currency.Code{"EUR"}},
	"HT": {Code: "HT", Name: "Haiti", Currencies: []currency.Code{"HTG", "USD"}},
	"HU": {Code: "HU", Name: "Hungary", Currencies: []currency.Code{"HUF"}},
	"IC": {Code: "IC", Name: "Canary Islands", Currencies: []currency.Code{"EUR"}},
	"ID": {Code: "ID", Name: "Indonesia", Currencies: []currency.Code{"IDR"}},
	"IE": {Code: "IE", Name: "Ireland", Currencies: []currency.Code{"EUR"}},
	"IL": {Code: "IL", Name: "Israel", Currencies: []currency.Code{"ILS"}},
	"IM": {Code: "IM", Name: "Isle of Man", Currencies: []currency.Code{"GBP"}},
	"IN": {Code: "IN", Name: "India", Currencies: []currency.Code{"INR"}},
	"IO": {Code: "IO", Name: "British Indian Ocean Territory", Currencies: []currency.Code{"USD"}},
	"IQ": {Code: "IQ", Name: "Iraq", Currencies: []currency.Code{"IQD"}},
	"IR": {Code: "IR", Name: "Iran", Currencies: []currency.Code{"IRR"}},
	"IS": {Code: "IS", Name: "Iceland", Currencies: []currency.Code{"ISK"}},
	"IT": {Code: "IT", Name: "Italy", Currencies: []currency.Code{"EUR"}},
	"JE": {Code: "JE", Name: "Jersey", Currencies: []currency.Code{"GBP"}},
	"JM": {Code: "JM", Name: "Jamaica", Currencies: []currency.Code{"JMD"}},
	"JO": {Code: "JO", Name: "Jordan", Currencies: []currency.Code{"JOD"}},
	"JP": {Code: "JP", Name: "Japan", Currencies: []currency.Code{"JPY"}},
	"KE": {Code: "KE", Name: "Kenya", Currencies: []currency.Code{"KES"}},
	"KG": {Code: "KG", Name: "Kyrgyzstan", Currencies: []currency.Code{"KGS"}},
	"KH": {Code: "KH", Name: "Cambodia", Currencies: []currency.Code{"KHR"}},
	"KI": {Code: "KI", Name: "Kiribati", Currencies: []currency.Code{"AUD"}},
	"KM": {Code: "KM", Name: "Comoros", Currencies: []currency.Code{"KMF"}},
	"KN": {Code: "KN", Name: "St. Kitts & Nevis", Currencies: []currency.Code{"XCD"}},
	"KP": {Code: "KP", Name: "North Korea", Currencies: []currency.Code{"KPW"}},
	"KR": {Code: "KR", Name: "South Korea", Currencies: []currency.Code{"KRW"}},
	"KW": {Code: "KW", Name: "Kuwait", Currencies: []currency.Code{"KWD"}},
	"KY": {Code: "KY", Name: "Cayman Islands", Currencies: []currency.Code{"KYD"}},
	"KZ": {Code: "KZ", Name: "Kazakhstan", Currencies: []currency.Code{"KZT"}},
	"LA": {Code: "LA", Name: "Laos", Currencies: []currency.Code{"LAK"}},
	"LB": {Code: "LB", Name: "Lebanon", Currencies: []currency.Code{"LBP"}},
	"LC": {Code: "LC", Name: "St. Lucia", Currencies: []currency.Code{"XCD"}},
	"LI": {Code: "LI", Name: "Liechtenstein", Currencies: []currency.Code{"CHF"}},
	"LK": {Code: "LK", Name: "Sri Lanka", Currencies: []currency.Code{"LKR"}},
	"LR": {Code: "LR", Name: "Liberia", Currencies: []currency.Code{"LRD"}},
	"LS": {Code: "LS", Name: "Lesotho", Currencies: []currency.Code{"ZAR", "LSL"}},
	"LT": {Code: "LT", Name: "Lithuania", Currencies: []currency.Code{"EUR"}},
	"LU": {Code: "LU", Name: "Luxembourg", Currencies: []currency.Code{"EUR"}},
	"LV": {Code: "LV", Name: "Latvia", Currencies: []currency.Code{"EUR"}},
	"LY": {Code: "LY", Name: "Libya", Currencies: []currency.Code{"LYD"}},
	"MA": {Code: "MA", Name: "Morocco", Currencies: []currency.Code{"MAD"}},
	"MC": {Code: "MC", Name: "Monaco", Currencies: []currency.Code{"EUR"}},
	"MD": {Code: "MD", Name: "Moldova", Currencies: []currency.Code{"MDL"}},
	"ME": {Code: "ME", Name: "Montenegro", Currencies: []currency.Code{"EUR"}},
	"MF": {Code: "MF", Name: "St. Martin", Currencies: []currency.Code{"EUR"}},
	"MG": {Code: "MG", Name: "Madagascar", Currencies: []currency.Code{"MGA"}},
	"MH": {Code: "MH", Name: "Marshall Islands", Currencies: []currency.Code{"USD"}},
	"MK": {Code: "MK", Name: "North Macedonia", Currencies: []currency.Code{"MKD"}},
	"ML": {Code: "ML", Name: "Mali", Currencies: []currency.Code{"XOF"}},
	"MM": {Code: "MM", Name: "Myanmar (Burma)", Currencies: []currency.Code{"MMK"}},
	"MN": {Code: "MN", Name: "Mongolia", Currencies: []currency.Code{"MNT"}},
	"MO": {Code: "MO", Name: "Macao SAR China", Currencies: []currency.Code{"MOP"}},
	"MP": {Code: "MP", Name: "Northern Mariana Islands", Currencies: []currency.Code{"USD"}},
	"MQ": {Code: "MQ", Name: "Martinique", Currencies: []currency.Code{"EUR"}},
	"MR": {Code: "MR", Name: "Mauritania", Currencies: []currency.Code{"MRU"}},
	"MS": {Code: "MS", Name: "Montserrat", Currencies: []currency.Code{"XCD"}},
	"MT": {Code: "MT", Name: "Malta", Currencies: []currency.Code{"EUR"}},
	"MU": {Code: "MU", Name: "Mauritius", Currencies: []currency.Code{"MUR"}},
	"MV": {Code: "MV", Name: "Maldives", Currencies: []currency.Code{"MVR"}},
	"MW": {Code: "MW", Name: "Malawi", Currencies: []currency.Code{"MWK"}},
	"MX": {Code: "MX", Name: "Mexico", Currencies: []currency.Code{"MXN"}},
	"MY": {Code: "MY", Name: "Malaysia", Currencies: []currency.Code{"MYR"}},
	"MZ": {Code: "MZ", Name: "Mozambique", Currencies: []currency.Code{"MZN"}},
	"NA": {Code: "NA", Name: "Namibia", Currencies: []currency.Code{"NAD", "ZAR"}},
	"NC": {Code: "NC", Name: "New Caledonia", Currencies: []currency.Code{"XPF"}},
	"NE": {Code: "NE", Name: "Niger", Currencies: []currency.Code{"XOF"}},
	"NF": {Code: "NF", Name: "Norfolk Island", Currencies: []currency.Code{"AUD"}},
	"NG": {Code: "NG", Name: "Nigeria", Currencies: []currency.Code{"NGN"}},
	"NI": {Code: "NI", Name: "Nicaragua", Currencies: []currency.Code{"NIO"}},
	"NL": {Code: "NL", Name: "Netherlands", Currencies: []currency.Code{"EUR"}},
	"NO": {Code: "NO", Name: "Norway", Currencies: []currency.Code{"NOK"}},
	"NP": {Code: "NP", Name: "Nepal", Currencies: []currency.Code{"NPR"}},
	"NR": {Code: "NR", Name: "Nauru", Currencies: []currency.Code{"AUD"}},
	"NU": {Code: "NU", Name: "Niue", Currencies: []currency.Code{"NZD"}},
	"NZ": {Code: "NZ", Name: "New Zealand", Currencies: []currency.Code{"NZD"}},
	"OM": {Code: "OM", Name: "Oman", Currencies: []currency.Code{"OMR"}},
	"PA": {Code: "PA", Name: "Panama", Currencies: []currency.Code{"PAB", "USD"}},
	"PE": {Code: "PE", Name: "Peru", Currencies: []currency.Code{"PEN"}},
	"PF": {Code: "PF", Name: "French Polynesia", Currencies: []currency.Code{"XPF"}},
	"PG": {Code: "PG", Name: "Papua New Guinea", Currencies: []currency.Code{"PGK"}},
	"PH": {Code: "PH", Name: "Philippines", Currencies: []currency.Code{"PHP"}},
	"PK": {Code: "PK", Name: "Pakistan", Currencies: []currency.Code{"PKR"}},
	"PL": {Code: "PL", Name: "Poland", Currencies: []currency.Code{"PLN"}},
	"PM": {Code: "PM", Name: "St. Pierre & Miquelon", Currencies: []currency.Code{"EUR"}},
	"PN": {Code: "PN", Name: "Pitcairn Islands", Currencies: []currency.Code{"NZD"}},
	"PR": {Code: "PR", Name: "Puerto Rico", Currencies: []currency.Code{"USD"}},
	"PS": {Code: "PS", Name: "Palestinian Territories", Currencies: []currency.Code{"ILS", "JOD"}},
	"PT": {Code: "PT", Name: "Portugal", Currencies: []currency.Code{"EUR"}},
	"PW": {Code: "PW", Name: "Palau", Currencies: []currency.Code{"USD"}},
	"PY": {Code: "PY", Name: "Paraguay", Currencies: []currency.Code{"PYG"}},
	"QA": {Code: "QA", Name: "Qatar", Currencies: []currency.Code{"QAR"}},
	"RE": {Code: "RE", Name: "Réunion", Currencies: []currency.Code{"EUR"}},
	"RO": {Code: "RO", Name: "Romania", Currencies: []currency.Code{"RON"}},
	"RS": {Code: "RS", Name: "Serbia", Currencies: []currency.Code{"RSD"}},
	"RU": {Code: "RU", Name: "Russia", Currencies: []currency.Code{"RUB"}},
	"RW": {Code: "RW", Name: "Rwanda", Currencies: []currency.Code{"RWF"}},
	"SA": {Code: "SA", Name: "Saudi Arabia", Currencies: []currency.Code{"SAR"}},
	"SB": {Code: "SB", Name: "Solomon Islands", Currencies: []currency.Code{"SBD"}},
	"SC": {Code: "SC", Name: "Seychelles", Currencies: []currency.Code{"SCR"}},
	"SD": {Code: "SD", Name: "Sudan", Currencies: []currency.Code{"SDG"}},
	"SE": {Code: "SE", Name: "Sweden", Currencies: []currency.Code{"SEK"}},
	"SG": {Code: "SG", Name: "Singapore", Currencies: []currency.Code{"SGD"}},
	"SH": {Code: "SH", Name: "St. Helena", Currencies: []currency.Code{"SHP"}},
	"SI": {Code: "SI", Name: "Slovenia", Currencies: []currency.Code{"EUR"}},
	"SJ": {Code: "SJ", Name: "Svalbard & Jan Mayen", Currencies: []currency.Code{"NOK"}},
	"SK": {Code: "SK", Name: "Slovakia", Currencies: []currency.Code{"EUR"}},
	"SL": {Code: "SL", Name: "Sierra Leone", Currencies: []currency.Code{"SLE"}},
	"SM": {Code: "SM", Name: "San Marino", Currencies: []currency.Code{"EUR"}},
	"SN": {Code: "SN", Name: "Senegal", Currencies: []currency.Code{"XOF"}},
	"SO": {Code: "SO", Name: "Somalia", Currencies: []currency.Code{"SOS"}},
	"SR": {Code: "SR", Name: "Suriname", Currencies: []currency.Code{"SRD"}},
	"SS": {Code: "SS", Name: "South Sudan", Currencies: []currency.Code{"SSP"}},
	"ST": {Code: "ST", Name: "São Tomé & Príncipe", Currencies: []currency.Code{"STN"}},
	"SV": {Code: "SV", Name: "El Salvador", Currencies: []currency.Code{"USD"}},
	"SX": {Code: "SX", Name: "Sint Maarten", Currencies: []currency.Code{"ANG"}},
	"SY": {Code: "SY", Name: "Syria", Currencies: []currency.Code{"SYP"}},
	"SZ": {Code: "SZ", Name: "Eswatini", Currencies: []currency.Code{"SZL"}},
	"TA": {Code: "TA", Name: "Tristan da Cunha", Currencies: []currency.Code{"GBP"}},
	"TC": {Code: "TC", Name: "Turks & Caicos Islands", Currencies: []currency.Code{"USD"}},
	"TD": {Code: "TD", Name: "Chad", Currencies: []currency.Code{"XAF"}},
	"TF": {Code: "TF", Name: "French Southern Territories", Currencies: []currency.Code{"EUR"}},
	"TG": {Code: "TG", Name: "Togo", Currencies: []currency.Code{"XOF"}},
	"TH": {Code: "TH", Name: "Thailand", Currencies: []currency.Code{"THB"}},
	"TJ": {Code: "TJ", Name: "Tajikistan", Currencies: []currency.Code{"TJS"}},
	"TK": {Code: "TK", Name: "Tokelau", Currencies: []currency.Code{"NZD"}},
	"TL": {Code: "TL", Name: "Timor-Leste", Currencies: []currency.Code{"USD"}},
	"TM": {Code: "TM", Name: "Turkmenistan", Currencies: []currency.Code{"TMT"}},
	"TN": {Code: "TN", Name: "Tunisia", Currencies: []currency.Code{"TND"}},
	"TO": {Code: "TO", Name: "Tonga", Currencies: []currency.Code{"TOP"}},
	"TR": {Code: "TR", Name: "Turkey", Currencies: []currency.Code{"TRY"}},
	"TT": {Code: "TT", Name: "Trinidad & Tobago", Currencies: []currency.Code{"TTD"}},
	"TV": {Code: "TV", Name: "Tuvalu", Currencies: []currency.Code{"AUD"}},
	"TW": {Code: "TW", Name: "Taiwan", Currencies: []currency.Code{"TWD"}},
	"TZ": {Code: "TZ", Name: "Tanzania", Currencies: []currency.Code{"TZS"}},
	"UA": {Code: "UA", Name: "Ukraine", Currencies: []currency.Code{"UAH"}},
	"UG": {Code: "UG", Name: "Uganda", Currencies: []currency.Code{"UGX"}},
	"UM": {Code: "UM", Name: "U.S. Outlying Islands", Currencies: []currency.Code{"USD"}},
	"UN": {Code: "UN", Name: "United Nations"},
	"US": {Code: "US", Name: "United States", Currencies: []currency.Code{"USD"}},
	"UY": {Code: "UY", Name: "Uruguay", Currencies: []currency.Code{"UYU"}},
	"UZ": {Code: "UZ", Name: "Uzbekistan", Currencies: []currency.Code{"UZS"}},
	"VA": {Code: "VA", Name: "Vatican City", Currencies: []currency.Code{"EUR"}},
	"VC": {Code: "VC", Name: "St. Vincent & Grenadines", Currencies: []currency.Code{"XCD"}},
	"VE": {Code: "VE", Name: "Venezuela", Currencies: []currency.Code{"VES"}},
	"VG": {Code: "VG", Name: "British Virgin Islands", Currencies: []currency.Code{"USD"}},
	"VI": {Code: "VI", Name: "U.S. Virgin Islands", Currencies: []currency.Code{"USD"}},
	"VN": {Code: "VN", Name: "Vietnam", Currencies: []currency.Code{"VND"}},
	"VU": {Code: "VU", Name: "Vanuatu", Currencies: []currency.Code{"VUV"}},
	"WF": {Code: "WF", Name: "Wallis & Futuna", Currencies: []currency.Code{"XPF"}},
	"WS": {Code: "WS", Name: "Samoa", Currencies: []currency.Code{"WST"}},
	"XK": {Code: "XK", Name: "Kosovo", Currencies: []currency.Code{"EUR"}},
	"YE": {Code: "YE", Name: "Yemen", Currencies: []currency.Code{"YER"}},
	"YT": {Code: "YT", Name: "Mayotte", Currencies: []currency.Code{"EUR"}},
	"ZA": {Code: "ZA", Name: "South Africa", Currencies: []currency.Code{"ZAR"}},
	"ZM": {Code: "ZM", Name: "Zambia", Currencies: []currency.Code{"ZMW"}},
	"ZW": {Code: "ZW", Name: "Zimbabwe", Currencies: []currency.Code{"USD"}},
}
