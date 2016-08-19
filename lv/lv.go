package lv

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type lv struct {
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

// New returns a new instance of translator for the 'lv' locale
func New() locales.Translator {
	return &lv{
		locale:                 "lv",
		pluralsCardinal:        []locales.PluralRule{1, 2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x52, 0x24}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46, 0x20}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x73}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44, 0x20}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x76, 0x2e}, {0x66, 0x65, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x72, 0x74, 0x73}, {0x61, 0x70, 0x72, 0x2e}, {0x6d, 0x61, 0x69, 0x6a, 0x73}, {0x6a, 0xc5, 0xab, 0x6e, 0x2e}, {0x6a, 0xc5, 0xab, 0x6c, 0x2e}, {0x61, 0x75, 0x67, 0x2e}, {0x73, 0x65, 0x70, 0x74, 0x2e}, {0x6f, 0x6b, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x76, 0xc4, 0x81, 0x72, 0x69, 0x73}, {0x66, 0x65, 0x62, 0x72, 0x75, 0xc4, 0x81, 0x72, 0x69, 0x73}, {0x6d, 0x61, 0x72, 0x74, 0x73}, {0x61, 0x70, 0x72, 0xc4, 0xab, 0x6c, 0x69, 0x73}, {0x6d, 0x61, 0x69, 0x6a, 0x73}, {0x6a, 0xc5, 0xab, 0x6e, 0x69, 0x6a, 0x73}, {0x6a, 0xc5, 0xab, 0x6c, 0x69, 0x6a, 0x73}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74, 0x73}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x73}, {0x6f, 0x6b, 0x74, 0x6f, 0x62, 0x72, 0x69, 0x73}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x73}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x73}},
		daysAbbreviated:        [][]uint8{{0x53, 0x76}, {0x50, 0x72}, {0x4f, 0x74}, {0x54, 0x72}, {0x43, 0x65}, {0x50, 0x6b}, {0x53, 0x65}},
		daysNarrow:             [][]uint8{{0x53}, {0x50}, {0x4f}, {0x54}, {0x43}, {0x50}, {0x53}},
		daysShort:              [][]uint8{{0x53, 0x76}, {0x50, 0x72}, {0x4f, 0x74}, {0x54, 0x72}, {0x43, 0x65}, {0x50, 0x6b}, {0x53, 0x65}},
		daysWide:               [][]uint8{{0x73, 0x76, 0xc4, 0x93, 0x74, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x70, 0x69, 0x72, 0x6d, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x6f, 0x74, 0x72, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x74, 0x72, 0x65, 0xc5, 0xa1, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x63, 0x65, 0x74, 0x75, 0x72, 0x74, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x70, 0x69, 0x65, 0x6b, 0x74, 0x64, 0x69, 0x65, 0x6e, 0x61}, {0x73, 0x65, 0x73, 0x74, 0x64, 0x69, 0x65, 0x6e, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x70, 0x72, 0x69, 0x65, 0x6b, 0xc5, 0xa1, 0x70, 0x2e}, {0x70, 0xc4, 0x93, 0x63, 0x70, 0x2e}},
		periodsNarrow:          [][]uint8{{0x70, 0x72, 0x69, 0x65, 0x6b, 0xc5, 0xa1, 0x70, 0x2e}, {0x70, 0xc4, 0x93, 0x63, 0x70, 0x2e}},
		periodsWide:            [][]uint8{{0x70, 0x72, 0x69, 0x65, 0x6b, 0xc5, 0xa1, 0x70, 0x75, 0x73, 0x64, 0x69, 0x65, 0x6e, 0xc4, 0x81}, {0x70, 0xc4, 0x93, 0x63, 0x70, 0x75, 0x73, 0x64, 0x69, 0x65, 0x6e, 0xc4, 0x81}},
		erasAbbreviated:        [][]uint8{{0x70, 0x2e, 0x6d, 0x2e, 0xc4, 0x93, 0x2e}, {0x6d, 0x2e, 0xc4, 0x93, 0x2e}},
		erasNarrow:             [][]uint8{{0x70, 0x2e, 0x6d, 0x2e, 0xc4, 0x93, 0x2e}, {0x6d, 0x2e, 0xc4, 0x93, 0x2e}},
		erasWide:               [][]uint8{{0x70, 0x69, 0x72, 0x6d, 0x73, 0x20, 0x6d, 0xc5, 0xab, 0x73, 0x75, 0x20, 0xc4, 0x93, 0x72, 0x61, 0x73}, {0x6d, 0xc5, 0xab, 0x73, 0x75, 0x20, 0xc4, 0x93, 0x72, 0xc4, 0x81}},
		timezones:              map[string][]uint8{"SAST": {0x44, 0x69, 0x65, 0x6e, 0x76, 0x69, 0x64, 0xc4, 0x81, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "MESZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x72, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x61, 0x20, 0x48, 0x61, 0x76, 0x61, 0x20, 0x73, 0x61, 0x6c, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AEST": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x61, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "BT": {0x42, 0x75, 0x74, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WARST": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc4, 0xab, 0x6e, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WEZ": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x72, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "EAT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0xc4, 0x81, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CAT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0xc4, 0x81, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "MDT": {0x4b, 0x61, 0x6c, 0x6e, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WAST": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0xc4, 0x81, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "JST": {0x4a, 0x61, 0x70, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "EDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ChST": {0xc4, 0x8c, 0x61, 0x6d, 0x6f, 0x72, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x61, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CHADT": {0xc4, 0x8c, 0x65, 0x74, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HAT": {0xc5, 0x85, 0xc5, 0xab, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "NZDT": {0x4a, 0x61, 0x75, 0x6e, 0x7a, 0xc4, 0x93, 0x6c, 0x61, 0x6e, 0x64, 0x65, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AKDT": {0x41, 0xc4, 0xbc, 0x61, 0x73, 0x6b, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x63, 0x75, 0xc4, 0x93, 0x6c, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WESZ": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WART": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc4, 0xab, 0x6e, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HNT": {0xc5, 0x85, 0xc5, 0xab, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "MEZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "OEZ": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HKST": {0x48, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0xc4, 0x8d, 0x75, 0x20, 0x47, 0x76, 0x69, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WIT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc4, 0x93, 0x7a, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AEDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x61, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "NZST": {0x4a, 0x61, 0x75, 0x6e, 0x7a, 0xc4, 0x93, 0x6c, 0x61, 0x6e, 0x64, 0x65, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "PST": {0x4b, 0x6c, 0x75, 0x73, 0xc4, 0x81, 0x20, 0x6f, 0x6b, 0x65, 0xc4, 0x81, 0x6e, 0x61, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HKT": {0x48, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AKST": {0x41, 0xc4, 0xbc, 0x61, 0x73, 0x6b, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WIB": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc4, 0x93, 0x7a, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "JDT": {0x4a, 0x61, 0x70, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ACDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "OESZ": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x69, 0x72, 0x6f, 0x70, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0xc5, 0xab, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HADT": {0x48, 0x61, 0x76, 0x61, 0x6a, 0x75, 0xe2, 0x80, 0x93, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "∅∅∅": {0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x61, 0x20, 0x48, 0x61, 0x76, 0x61, 0x20, 0x73, 0x61, 0x6c, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "GYT": {0x47, 0x61, 0x6a, 0xc4, 0x81, 0x6e, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ACWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x72, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CLT": {0xc4, 0x8c, 0xc4, 0xab, 0x6c, 0x65, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "HAST": {0x48, 0x61, 0x76, 0x61, 0x6a, 0x75, 0xe2, 0x80, 0x93, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ACWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x72, 0x69, 0x65, 0x74, 0x75, 0x6d, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "BOT": {0x42, 0x6f, 0x6c, 0xc4, 0xab, 0x76, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc4, 0xab, 0x6e, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc4, 0xab, 0x6e, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "PDT": {0x4b, 0x6c, 0x75, 0x73, 0xc4, 0x81, 0x20, 0x6f, 0x6b, 0x65, 0xc4, 0x81, 0x6e, 0x61, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CLST": {0xc4, 0x8c, 0xc4, 0xab, 0x6c, 0x65, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "MST": {0x4b, 0x61, 0x6c, 0x6e, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WITA": {0x43, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc4, 0x93, 0x7a, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "WAT": {0x52, 0x69, 0x65, 0x74, 0x75, 0x6d, 0xc4, 0x81, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x69, 0x7a, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "EST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x75, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "CHAST": {0xc4, 0x8c, 0x65, 0x74, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "ACST": {0x41, 0x75, 0x73, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc4, 0x81, 0x6c, 0x61, 0x69, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x7a, 0x69, 0x65, 0x6d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "GMT": {0x47, 0x72, 0x69, 0x6e, 0x69, 0xc4, 0x8d, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x61, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x61, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x73}},
	}
}

// Locale returns the current translators string locale
func (lv *lv) Locale() string {
	return lv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lv'
func (lv *lv) PluralsCardinal() []locales.PluralRule {
	return lv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lv'
func (lv *lv) PluralsOrdinal() []locales.PluralRule {
	return lv.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lv'
func (lv *lv) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)
	fMod100 := f % 100
	fMod10 := f % 10

	if (nMod10 == 0) || (nMod100 >= 11 && nMod100 <= 19) || (v == 2 && fMod100 >= 11 && fMod100 <= 19) {
		return locales.PluralRuleZero
	} else if (nMod10 == 1 && nMod100 != 11) || (v == 2 && fMod10 == 1 && fMod100 != 11) || (v != 2 && fMod10 == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lv'
func (lv *lv) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lv'
func (lv *lv) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := lv.CardinalPluralRule(num1, v1)
	end := lv.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleZero {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleZero {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleZero {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lv *lv) MonthAbbreviated(month time.Month) []byte {
	return lv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lv *lv) MonthsAbbreviated() [][]byte {
	return lv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lv *lv) MonthNarrow(month time.Month) []byte {
	return lv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lv *lv) MonthsNarrow() [][]byte {
	return lv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lv *lv) MonthWide(month time.Month) []byte {
	return lv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lv *lv) MonthsWide() [][]byte {
	return lv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lv *lv) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return lv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lv *lv) WeekdaysAbbreviated() [][]byte {
	return lv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lv *lv) WeekdayNarrow(weekday time.Weekday) []byte {
	return lv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lv *lv) WeekdaysNarrow() [][]byte {
	return lv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lv *lv) WeekdayShort(weekday time.Weekday) []byte {
	return lv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lv *lv) WeekdaysShort() [][]byte {
	return lv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lv *lv) WeekdayWide(weekday time.Weekday) []byte {
	return lv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lv *lv) WeekdaysWide() [][]byte {
	return lv.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lv' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(lv.decimal) + len(lv.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lv.group) - 1; j >= 0; j-- {
					b = append(b, lv.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lv' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lv *lv) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(lv.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lv.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lv.currencies[currency]
	l := len(s) + len(lv.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lv.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lv'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lv.currencies[currency]
	l := len(s) + len(lv.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, lv.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lv.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lv.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtDateShort(t time.Time) []byte {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x64, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lv.monthsAbbreviated[t.Month()]...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x64, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lv.monthsWide[t.Month()]...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, lv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x64, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, lv.monthsWide[t.Month()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'lv'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lv *lv) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
