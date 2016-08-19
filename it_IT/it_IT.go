package it_IT

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type it_IT struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
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

// New returns a new instance of translator for the 'it_IT' locale
func New() locales.Translator {
	return &it_IT{
		locale:                 "it_IT",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x67, 0x65, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x67}, {0x67, 0x69, 0x75}, {0x6c, 0x75, 0x67}, {0x61, 0x67, 0x6f}, {0x73, 0x65, 0x74}, {0x6f, 0x74, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x69, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x47}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x47}, {0x4c}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x67, 0x65, 0x6e, 0x6e, 0x61, 0x69, 0x6f}, {0x66, 0x65, 0x62, 0x62, 0x72, 0x61, 0x69, 0x6f}, {0x6d, 0x61, 0x72, 0x7a, 0x6f}, {0x61, 0x70, 0x72, 0x69, 0x6c, 0x65}, {0x6d, 0x61, 0x67, 0x67, 0x69, 0x6f}, {0x67, 0x69, 0x75, 0x67, 0x6e, 0x6f}, {0x6c, 0x75, 0x67, 0x6c, 0x69, 0x6f}, {0x61, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x73, 0x65, 0x74, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x6f, 0x74, 0x74, 0x6f, 0x62, 0x72, 0x65}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0x69, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e}, {0x6d, 0x61, 0x72}, {0x6d, 0x65, 0x72}, {0x67, 0x69, 0x6f}, {0x76, 0x65, 0x6e}, {0x73, 0x61, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x4d}, {0x47}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e}, {0x6d, 0x61, 0x72}, {0x6d, 0x65, 0x72}, {0x67, 0x69, 0x6f}, {0x76, 0x65, 0x6e}, {0x73, 0x61, 0x62}},
		daysWide:               [][]uint8{{0x64, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x63, 0x61}, {0x6c, 0x75, 0x6e, 0x65, 0x64, 0xc3, 0xac}, {0x6d, 0x61, 0x72, 0x74, 0x65, 0x64, 0xc3, 0xac}, {0x6d, 0x65, 0x72, 0x63, 0x6f, 0x6c, 0x65, 0x64, 0xc3, 0xac}, {0x67, 0x69, 0x6f, 0x76, 0x65, 0x64, 0xc3, 0xac}, {0x76, 0x65, 0x6e, 0x65, 0x72, 0x64, 0xc3, 0xac}, {0x73, 0x61, 0x62, 0x61, 0x74, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x6d, 0x2e}, {0x70, 0x2e}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasWide:               [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		timezones:              map[string][]uint8{"MST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x4d, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x6e, 0x65, 0x20, 0x52, 0x6f, 0x63, 0x63, 0x69, 0x6f, 0x73, 0x65, 0x20, 0x55, 0x53, 0x41}, "AST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "AWDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "ART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "HKT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "HKST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x69, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "AKDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "UYST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "WESZ": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "NZDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x4e, 0x75, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "JST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x47, 0x69, 0x61, 0x70, 0x70, 0x6f, 0x6e, 0x65}, "OESZ": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "CST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65, 0x20, 0x55, 0x53, 0x41}, "WART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "ADT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "COST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "PST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x6f, 0x20, 0x55, 0x53, 0x41}, "GYT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "BOT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "UYT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "VET": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "ACDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "ACWDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "MEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "NZST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x4e, 0x75, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "IST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x69, 0x61}, "ECT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "CHADT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "∅∅∅": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x65, 0x20, 0x41, 0x7a, 0x7a, 0x6f, 0x72, 0x72, 0x65}, "WAST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "AEDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "HADT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x65, 0x20, 0x49, 0x73, 0x6f, 0x6c, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x6e, 0x65}, "JDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x47, 0x69, 0x61, 0x70, 0x70, 0x6f, 0x6e, 0x65}, "HAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x65, 0x20, 0x49, 0x73, 0x6f, 0x6c, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x6e, 0x65}, "MYT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x4d, 0x61, 0x6c, 0x65, 0x73, 0x69, 0x61}, "CHAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "MESZ": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "AEST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "EAT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "ChST": {0x4f, 0x72, 0x61, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "CDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65, 0x20, 0x55, 0x53, 0x41}, "WEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "CAT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "WARST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "GMT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x69, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "GFT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x65}, "WIT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "WIB": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "EDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65, 0x20, 0x55, 0x53, 0x41}, "ACST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "SRT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "CLST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x69, 0x6c, 0x65}, "LHST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WAT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "EST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65, 0x20, 0x55, 0x53, 0x41}, "ACWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "COT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "SGT": {0x4f, 0x72, 0x61, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65}, "ARST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "TMST": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "OEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "SAST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x65}, "BT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "MDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x4d, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x6e, 0x65, 0x20, 0x52, 0x6f, 0x63, 0x63, 0x69, 0x6f, 0x73, 0x65, 0x20, 0x55, 0x53, 0x41}, "LHDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x69, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "AKST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "WITA": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "CLT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x69, 0x6c, 0x65}, "HNT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "PDT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x6f, 0x20, 0x55, 0x53, 0x41}, "TMT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "HAT": {0x4f, 0x72, 0x61, 0x20, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x69, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "AWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}},
	}
}

// Locale returns the current translators string locale
func (it *it_IT) Locale() string {
	return it.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'it_IT'
func (it *it_IT) PluralsCardinal() []locales.PluralRule {
	return it.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'it_IT'
func (it *it_IT) PluralsOrdinal() []locales.PluralRule {
	return it.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'it_IT'
func (it *it_IT) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'it_IT'
func (it *it_IT) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 11 || n == 8 || n == 80 || n == 800 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'it_IT'
func (it *it_IT) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := it.CardinalPluralRule(num1, v1)
	end := it.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (it *it_IT) MonthAbbreviated(month time.Month) []byte {
	return it.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (it *it_IT) MonthsAbbreviated() [][]byte {
	return it.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (it *it_IT) MonthNarrow(month time.Month) []byte {
	return it.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (it *it_IT) MonthsNarrow() [][]byte {
	return it.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (it *it_IT) MonthWide(month time.Month) []byte {
	return it.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (it *it_IT) MonthsWide() [][]byte {
	return it.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (it *it_IT) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return it.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (it *it_IT) WeekdaysAbbreviated() [][]byte {
	return it.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (it *it_IT) WeekdayNarrow(weekday time.Weekday) []byte {
	return it.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (it *it_IT) WeekdaysNarrow() [][]byte {
	return it.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (it *it_IT) WeekdayShort(weekday time.Weekday) []byte {
	return it.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (it *it_IT) WeekdaysShort() [][]byte {
	return it.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (it *it_IT) WeekdayWide(weekday time.Weekday) []byte {
	return it.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (it *it_IT) WeekdaysWide() [][]byte {
	return it.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'it_IT' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(it.decimal) + len(it.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, it.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, it.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, it.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'it_IT' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (it *it_IT) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(it.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, it.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, it.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, it.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := it.currencies[currency]
	l := len(s) + len(it.decimal) + len(it.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, it.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, it.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, it.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, it.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, it.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'it_IT'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := it.currencies[currency]
	l := len(s) + len(it.decimal) + len(it.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, it.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, it.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, it.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, it.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, it.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, it.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, it.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, it.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, it.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, it.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, it.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, it.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, it.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, it.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, it.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'it_IT'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (it *it_IT) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, it.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, it.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := it.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
