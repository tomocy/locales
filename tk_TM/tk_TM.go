package tk_TM

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type tk_TM struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'tk_TM' locale
func New() locales.Translator {
	return &tk_TM{
		locale:                 "tk_TM",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xc3, 0xbd, 0x61, 0x6e}, {0x66, 0x65, 0x77}, {0x6d, 0x61, 0x72, 0x74}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0xc3, 0xbd}, {0x69, 0xc3, 0xbd, 0x75, 0x6e}, {0x69, 0xc3, 0xbd, 0x75, 0x6c}, {0x61, 0x77, 0x67}, {0x73, 0x65, 0x6e}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0xc3, 0xbd}, {0x64, 0x65, 0x6b}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xc3, 0x9d}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x49}, {0x49}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xc3, 0xbd, 0x61, 0x6e, 0x77, 0x61, 0x72}, {0x66, 0x65, 0x77, 0x72, 0x61, 0x6c}, {0x6d, 0x61, 0x72, 0x74}, {0x61, 0x70, 0x72, 0x65, 0x6c}, {0x6d, 0x61, 0xc3, 0xbd}, {0x69, 0xc3, 0xbd, 0x75, 0x6e}, {0x69, 0xc3, 0xbd, 0x75, 0x6c}, {0x61, 0x77, 0x67, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x6e, 0x74, 0xc3, 0xbd, 0x61, 0x62, 0x72}, {0x6f, 0x6b, 0x74, 0xc3, 0xbd, 0x61, 0x62, 0x72}, {0x6e, 0x6f, 0xc3, 0xbd, 0x61, 0x62, 0x72}, {0x64, 0x65, 0x6b, 0x61, 0x62, 0x72}},
		daysAbbreviated:        [][]uint8{{0xc3, 0xbd, 0x62}, {0x64, 0x62}, {0x73, 0x62}, {0xc3, 0xa7, 0x62}, {0x70, 0x62}, {0x61, 0x6e}, {0xc5, 0x9f, 0x62}},
		daysNarrow:             [][]uint8{{0xc3, 0x9d}, {0x44}, {0x53}, {0xc3, 0x87}, {0x50}, {0x41}, {0xc5, 0x9e}},
		daysWide:               [][]uint8{{0xc3, 0xbd, 0x65, 0x6b, 0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}, {0x64, 0x75, 0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}, {0x73, 0x69, 0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}, {0xc3, 0xa7, 0x61, 0x72, 0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}, {0x70, 0x65, 0x6e, 0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}, {0x61, 0x6e, 0x6e, 0x61}, {0xc5, 0x9f, 0x65, 0x6e, 0x62, 0x65}},
		timezones:              map[string][]uint8{"ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "NZST": {0x54, 0xc3, 0xa4, 0x7a, 0x65, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "OEZ": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "HNT": {0x4e, 0xc3, 0xbd, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "HKT": {0x47, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "AEDT": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "MESZ": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "BT": {0x42, 0x75, 0x74, 0x61, 0x6e}, "JDT": {0xc3, 0x9d, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "ACWST": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x67, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x74, 0x61, 0x72, 0x61, 0x70, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x77, 0x69, 0xc3, 0xbd, 0x61}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x61, 0x75, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "AKST": {0x41, 0x6c, 0xc3, 0xbd, 0x61, 0x73, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "VET": {0x57, 0x65, 0x6e, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x61}, "CLST": {0xc3, 0x87, 0x69, 0x6c, 0x69, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "CHAST": {0xc3, 0x87, 0x61, 0x74, 0x65, 0x6d, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "WIT": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0xc3, 0xbd, 0x61}, "∅∅∅": {0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "TMT": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "TMST": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "GYT": {0x47, 0x61, 0xc3, 0xbd, 0x61, 0x6e, 0x61}, "HKST": {0x47, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "ChST": {0xc3, 0x87, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "PST": {0xc3, 0x9d, 0x75, 0x77, 0x61, 0xc5, 0x9f, 0x20, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "EDT": {0x47, 0xc3, 0xbc, 0x6e, 0x6f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "SAST": {0x47, 0xc3, 0xbc, 0x6e, 0x6f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "WARST": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "MEZ": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "CLT": {0xc3, 0x87, 0x69, 0x6c, 0x69, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "CDT": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "AKDT": {0x41, 0x6c, 0xc3, 0xbd, 0x61, 0x73, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "WITA": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0xc3, 0xbd, 0x61}, "JST": {0xc3, 0x9d, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "CAT": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "WART": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "GMT": {0x47, 0x72, 0x69, 0x6e, 0x77, 0x69, 0xc3, 0xa7, 0x20, 0x62, 0x6f, 0xc3, 0xbd, 0x75, 0x6e, 0xc3, 0xa7, 0x61, 0x20, 0x6f, 0x72, 0x74, 0x61, 0x20, 0x77, 0x61, 0x67, 0x74}, "WIB": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0xc3, 0xbd, 0x61}, "EAT": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "PDT": {0xc3, 0x9d, 0x75, 0x77, 0x61, 0xc5, 0x9f, 0x20, 0x75, 0x6d, 0x6d, 0x61, 0x6e, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "HADT": {0x47, 0x61, 0x77, 0x61, 0xc3, 0xbd, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x77, 0x61, 0xc3, 0xbd, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x77, 0x61, 0xc3, 0xbd, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "OESZ": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "MDT": {0x44, 0x61, 0x67, 0x6c, 0x79, 0x6b, 0x20, 0xc3, 0xbd, 0x65, 0x72, 0x69, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74, 0x20, 0x28, 0x41, 0x42, 0xc5, 0x9e, 0x29}, "WAST": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "CHADT": {0xc3, 0x87, 0x61, 0x74, 0x65, 0x6d, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "ACDT": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "WAT": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "WESZ": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0xc3, 0xbd, 0x7a, 0x69, 0xc3, 0xbd, 0x61}, "CST": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "MST": {0x44, 0x61, 0x67, 0x6c, 0x79, 0x6b, 0x20, 0xc3, 0xbd, 0x65, 0x72, 0x69, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74, 0x20, 0x28, 0x41, 0x42, 0xc5, 0x9e, 0x29}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x73, 0x75, 0x7a, 0x20, 0x47, 0x77, 0x69, 0x61, 0x6e, 0x61}, "AWDT": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "WEZ": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0xc3, 0x9d, 0x65, 0x77, 0x72, 0x6f, 0x70, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "NZDT": {0x54, 0xc3, 0xa4, 0x7a, 0x65, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "IST": {0x48, 0x69, 0x6e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ACWDT": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x67, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x74, 0x61, 0x72, 0x61, 0x70, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x61, 0x75, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "HAT": {0x4e, 0xc3, 0xbd, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x2c, 0x20, 0x74, 0x6f, 0x6d, 0x75, 0x73, 0x6b, 0x79, 0x20, 0x77, 0x61, 0x67, 0x74}, "AEST": {0x47, 0xc3, 0xbc, 0x6e, 0x64, 0x6f, 0x67, 0x61, 0x72, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "AWST": {0x47, 0xc3, 0xbc, 0x6e, 0x62, 0x61, 0x74, 0x61, 0x72, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "ACST": {0x4d, 0x65, 0x72, 0x6b, 0x65, 0x7a, 0x69, 0x20, 0x41, 0x77, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0xc3, 0xbd, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "EST": {0x47, 0xc3, 0xbc, 0x6e, 0x6f, 0x72, 0x74, 0x61, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}, "ECT": {0x45, 0x6b, 0x77, 0x61, 0x64, 0x6f, 0x72}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "HAST": {0x47, 0x61, 0x77, 0x61, 0xc3, 0xbd, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x77, 0x61, 0x67, 0x74}},
	}
}

// Locale returns the current translators string locale
func (tk *tk_TM) Locale() string {
	return tk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'tk_TM'
func (tk *tk_TM) PluralsCardinal() []locales.PluralRule {
	return tk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'tk_TM'
func (tk *tk_TM) PluralsOrdinal() []locales.PluralRule {
	return tk.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'tk_TM'
func (tk *tk_TM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'tk_TM'
func (tk *tk_TM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'tk_TM'
func (tk *tk_TM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (tk *tk_TM) MonthAbbreviated(month time.Month) []byte {
	return tk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (tk *tk_TM) MonthsAbbreviated() [][]byte {
	return tk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (tk *tk_TM) MonthNarrow(month time.Month) []byte {
	return tk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (tk *tk_TM) MonthsNarrow() [][]byte {
	return tk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (tk *tk_TM) MonthWide(month time.Month) []byte {
	return tk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (tk *tk_TM) MonthsWide() [][]byte {
	return tk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (tk *tk_TM) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return tk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (tk *tk_TM) WeekdaysAbbreviated() [][]byte {
	return tk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (tk *tk_TM) WeekdayNarrow(weekday time.Weekday) []byte {
	return tk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (tk *tk_TM) WeekdaysNarrow() [][]byte {
	return tk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (tk *tk_TM) WeekdayShort(weekday time.Weekday) []byte {
	return tk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (tk *tk_TM) WeekdaysShort() [][]byte {
	return tk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (tk *tk_TM) WeekdayWide(weekday time.Weekday) []byte {
	return tk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (tk *tk_TM) WeekdaysWide() [][]byte {
	return tk.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'tk_TM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(tk.decimal) + len(tk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tk.group) - 1; j >= 0; j-- {
					b = append(b, tk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'tk_TM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (tk *tk_TM) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(tk.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, tk.percentSuffix...)

	b = append(b, tk.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tk.currencies[currency]
	l := len(s) + len(tk.decimal) + len(tk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tk.group) - 1; j >= 0; j-- {
					b = append(b, tk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, tk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'tk_TM'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tk.currencies[currency]
	l := len(s) + len(tk.decimal) + len(tk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(tk.group) - 1; j >= 0; j-- {
					b = append(b, tk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, tk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, tk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, tk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tk.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tk.daysWide[t.Weekday()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'tk_TM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (tk *tk_TM) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := tk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
