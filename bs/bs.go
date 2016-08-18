package bs

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type bs struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsShort            [][]byte
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

// New returns a new instance of translator for the 'bs' locale
func New() locales.Translator {
	return &bs{
		locale:                 "bs",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x4b, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x6b, 0x6e}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x64, 0x69, 0x6e, 0x2e}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x6a}, {0x6a, 0x75, 0x6e}, {0x6a, 0x75, 0x6c}, {0x61, 0x75, 0x67}, {0x73, 0x65, 0x70}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x6a}, {0x66}, {0x6d}, {0x61}, {0x6d}, {0x6a}, {0x6a}, {0x61}, {0x73}, {0x6f}, {0x6e}, {0x64}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0x61, 0x72}, {0x66, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72}, {0x6d, 0x61, 0x72, 0x74}, {0x61, 0x70, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x6a}, {0x6a, 0x75, 0x6e, 0x69}, {0x6a, 0x75, 0x6c, 0x69}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x61, 0x72}, {0x6f, 0x6b, 0x74, 0x6f, 0x62, 0x61, 0x72}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x61, 0x72}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x61, 0x72}},
		daysAbbreviated:        [][]uint8{{0x6e, 0x65, 0x64}, {0x70, 0x6f, 0x6e}, {0x75, 0x74, 0x6f}, {0x73, 0x72, 0x69}, {0xc4, 0x8d, 0x65, 0x74}, {0x70, 0x65, 0x74}, {0x73, 0x75, 0x62}},
		daysNarrow:             [][]uint8{{0x4e}, {0x50}, {0x55}, {0x53}, {0xc4, 0x8c}, {0x50}, {0x53}},
		daysShort:              [][]uint8{{0x6e, 0x65, 0x64}, {0x70, 0x6f, 0x6e}, {0x75, 0x74, 0x6f}, {0x73, 0x72, 0x69}, {0xc4, 0x8d, 0x65, 0x74}, {0x70, 0x65, 0x74}, {0x73, 0x75, 0x62}},
		daysWide:               [][]uint8{{0x6e, 0x65, 0x64, 0x6a, 0x65, 0x6c, 0x6a, 0x61}, {0x70, 0x6f, 0x6e, 0x65, 0x64, 0x6a, 0x65, 0x6c, 0x6a, 0x61, 0x6b}, {0x75, 0x74, 0x6f, 0x72, 0x61, 0x6b}, {0x73, 0x72, 0x69, 0x6a, 0x65, 0x64, 0x61}, {0xc4, 0x8d, 0x65, 0x74, 0x76, 0x72, 0x74, 0x61, 0x6b}, {0x70, 0x65, 0x74, 0x61, 0x6b}, {0x73, 0x75, 0x62, 0x6f, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x70, 0x72, 0x69, 0x6a, 0x65, 0x70, 0x6f, 0x64, 0x6e, 0x65}, {0x70, 0x6f, 0x70, 0x6f, 0x64, 0x6e, 0x65}},
		periodsNarrow:          [][]uint8{{0x70, 0x72, 0x69, 0x6a, 0x65, 0x70, 0x6f, 0x64, 0x6e, 0x65}, {0x70, 0x6f, 0x70, 0x6f, 0x64, 0x6e, 0x65}},
		periodsWide:            [][]uint8{{0x70, 0x72, 0x69, 0x6a, 0x65, 0x20, 0x70, 0x6f, 0x64, 0x6e, 0x65}, {0x70, 0x6f, 0x70, 0x6f, 0x64, 0x6e, 0x65}},
		erasAbbreviated:        [][]uint8{{0x70, 0x2e, 0x20, 0x6e, 0x2e, 0x20, 0x65, 0x2e}, {0x6e, 0x2e, 0x20, 0x65, 0x2e}},
		erasNarrow:             [][]uint8{{0x70, 0x72, 0x2e, 0x6e, 0x2e, 0x65, 0x2e}, {0x41, 0x44}},
		erasWide:               [][]uint8{{0x50, 0x72, 0x69, 0x6a, 0x65, 0x20, 0x6e, 0x6f, 0x76, 0x65, 0x20, 0x65, 0x72, 0x65}, {0x4e, 0x6f, 0x76, 0x65, 0x20, 0x65, 0x72, 0x65}},
		timezones:              map[string][]uint8{"∅∅∅": {0x42, 0x72, 0x61, 0x7a, 0x69, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "COT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AKDT": {0x41, 0x6c, 0x6a, 0x61, 0x73, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "WAST": {0x5a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "ACST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WART": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "GFT": {0x56, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x75, 0x73, 0x6b, 0x65, 0x20, 0x47, 0x76, 0x61, 0x6a, 0x61, 0x6e, 0x65}, "ACWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "LHDT": {0x4c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x73, 0x74, 0x72, 0x76, 0x75, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x61, 0x75}, "WIT": {0x49, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "NZST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6e, 0x6f, 0x76, 0x6f, 0x7a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "OESZ": {0x49, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "BT": {0x42, 0x75, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "UYT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x75, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "CDT": {0x4c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x67, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "HKT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WESZ": {0x5a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "MESZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "GYT": {0x47, 0x76, 0x61, 0x6a, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WIB": {0x5a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "EAT": {0x49, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "ACWST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "TMT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x74, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "GMT": {0x56, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x70, 0x6f, 0x20, 0x47, 0x72, 0x69, 0x6e, 0x76, 0x69, 0xc4, 0x8d, 0x75}, "WAT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "ART": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "EST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "AEST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AWST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ACDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "PDT": {0x4c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x67, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "LHST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x6e, 0x61, 0x20, 0x4f, 0x73, 0x74, 0x72, 0x76, 0x75, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x61, 0x75}, "WARST": {0x5a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "AEDT": {0x49, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "MDT": {0x4d, 0x61, 0x6b, 0x61, 0x6f, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0x6a, 0x65, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "CST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "CHAST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0xc4, 0x8d, 0x61, 0x74, 0x61, 0x6d, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "JST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HAST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x68, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0xc4, 0x87, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MYT": {0x4d, 0x61, 0x6c, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AKST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x61, 0x6c, 0x6a, 0x61, 0x73, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x63, 0x75, 0x65, 0x6c, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "NZDT": {0x4e, 0x6f, 0x76, 0x6f, 0x7a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "SAST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6a, 0x75, 0xc5, 0xbe, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "SGT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CLST": {0xc4, 0x8c, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "EDT": {0x4c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x67, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "CLT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0xc4, 0x8d, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "IST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CHADT": {0xc4, 0x8c, 0x61, 0x74, 0x61, 0x6d, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "HNT": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6e, 0x6a, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HAT": {0x4e, 0x6a, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "HADT": {0x48, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0xc4, 0x87, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x65, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "MST": {0x4d, 0x61, 0x6b, 0x61, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65}, "WITA": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "OEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ChST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0xc4, 0x8d, 0x61, 0x6d, 0x6f, 0x72, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CAT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WEZ": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "ADT": {0x4c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x6b, 0x6f, 0x67, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "PST": {0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x28, 0x53, 0x41, 0x44, 0x29}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AWDT": {0x5a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x72, 0x61, 0xc4, 0x8d, 0x75, 0x6e, 0x61, 0x6e, 0x6a, 0x65, 0x20, 0x76, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x61}},
	}
}

// Locale returns the current translators string locale
func (bs *bs) Locale() string {
	return bs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bs'
func (bs *bs) PluralsCardinal() []locales.PluralRule {
	return bs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bs'
func (bs *bs) PluralsOrdinal() []locales.PluralRule {
	return bs.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14) || (fMod10 >= 2 && fMod10 <= 4 && fMod100 < 12 && fMod100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bs'
func (bs *bs) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bs.CardinalPluralRule(num1, v1)
	end := bs.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bs.decimal) + len(bs.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bs *bs) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bs.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bs.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(bs.decimal) + len(bs.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(bs.decimal) + len(bs.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bs.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtDateShort(t time.Time) []byte {

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

	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, bs.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'bs'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bs *bs) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := bs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
