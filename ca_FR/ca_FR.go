package ca_FR

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ca_FR struct {
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

// New returns a new instance of translator for the 'ca_FR' locale
func New() locales.Translator {
	return &ca_FR{
		locale:                 "ca_FR",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x67, 0x65, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x72, 0xc3, 0xa7}, {0x61, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x69, 0x67}, {0x6a, 0x75, 0x6e, 0x79}, {0x6a, 0x75, 0x6c, 0x2e}, {0x61, 0x67, 0x2e}, {0x73, 0x65, 0x74, 0x2e}, {0x6f, 0x63, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x73, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x47, 0x4e}, {0x46, 0x42}, {0x4d, 0xc3, 0x87}, {0x41, 0x42}, {0x4d, 0x47}, {0x4a, 0x4e}, {0x4a, 0x4c}, {0x41, 0x47}, {0x53, 0x54}, {0x4f, 0x43}, {0x4e, 0x56}, {0x44, 0x53}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x64, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72}, {0x64, 0x65, 0x20, 0x66, 0x65, 0x62, 0x72, 0x65, 0x72}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x72, 0xc3, 0xa7}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x62, 0x72, 0x69, 0x6c}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x69, 0x67}, {0x64, 0x65, 0x20, 0x6a, 0x75, 0x6e, 0x79}, {0x64, 0x65, 0x20, 0x6a, 0x75, 0x6c, 0x69, 0x6f, 0x6c}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x67, 0x6f, 0x73, 0x74}, {0x64, 0x65, 0x20, 0x73, 0x65, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0xe2, 0x80, 0x99, 0x6f, 0x63, 0x74, 0x75, 0x62, 0x72, 0x65}, {0x64, 0x65, 0x20, 0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0x65, 0x20, 0x64, 0x65, 0x73, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x64, 0x67, 0x2e}, {0x64, 0x6c, 0x2e}, {0x64, 0x74, 0x2e}, {0x64, 0x63, 0x2e}, {0x64, 0x6a, 0x2e}, {0x64, 0x76, 0x2e}, {0x64, 0x73, 0x2e}},
		daysNarrow:             [][]uint8{{0x64, 0x67}, {0x64, 0x6c}, {0x64, 0x74}, {0x64, 0x63}, {0x64, 0x6a}, {0x64, 0x76}, {0x64, 0x73}},
		daysShort:              [][]uint8{{0x64, 0x67, 0x2e}, {0x64, 0x6c, 0x2e}, {0x64, 0x74, 0x2e}, {0x64, 0x63, 0x2e}, {0x64, 0x6a, 0x2e}, {0x64, 0x76, 0x2e}, {0x64, 0x73, 0x2e}},
		daysWide:               [][]uint8{{0x64, 0x69, 0x75, 0x6d, 0x65, 0x6e, 0x67, 0x65}, {0x64, 0x69, 0x6c, 0x6c, 0x75, 0x6e, 0x73}, {0x64, 0x69, 0x6d, 0x61, 0x72, 0x74, 0x73}, {0x64, 0x69, 0x6d, 0x65, 0x63, 0x72, 0x65, 0x73}, {0x64, 0x69, 0x6a, 0x6f, 0x75, 0x73}, {0x64, 0x69, 0x76, 0x65, 0x6e, 0x64, 0x72, 0x65, 0x73}, {0x64, 0x69, 0x73, 0x73, 0x61, 0x62, 0x74, 0x65}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasNarrow:             [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasWide:               [][]uint8{{0x61, 0x62, 0x61, 0x6e, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74}, {0x64, 0x65, 0x73, 0x70, 0x72, 0xc3, 0xa9, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74}},
		timezones:              map[string][]uint8{"CLT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x58, 0x69, 0x6c, 0x65}, "ACDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "COST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb2, 0x6d, 0x62, 0x69, 0x61}, "IST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x8d, 0x6e, 0x64, 0x69, 0x61}, "GFT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x61, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "WIT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x65, 0x73}, "EAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "UYST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "SRT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "ACWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "ADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0xc3, 0xa0, 0x6e, 0x74, 0x69, 0x63}, "ACWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "TMST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "BT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "HADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "WIB": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "WAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "PDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63}, "SAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x73, 0x75, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61}, "MYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0xc3, 0xa0, 0x69, 0x73, 0x69, 0x61}, "MDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6d, 0x75, 0x6e, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "AWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "GYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "CHADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ChST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "ART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "MST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6d, 0x75, 0x6e, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "AST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0xc3, 0xa0, 0x6e, 0x74, 0x69, 0x63}, "NZDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "HKT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "EST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "SGT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "JDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3}, "WEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "NZST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "GMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4d, 0x65, 0x72, 0x69, 0x64, 0x69, 0xc3, 0xa0, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "UYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "PST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63}, "COT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb2, 0x6d, 0x62, 0x69, 0x61}, "WITA": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "JST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3}, "MEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "MESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "ARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "OESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "LHDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "AEDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "CHAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ECT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x71, 0x75, 0x61, 0x64, 0x6f, 0x72}, "AKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "AKDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "BOT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0xc3, 0xad, 0x76, 0x69, 0x61}, "OEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "HNT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "LHST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "WART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "AEST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "VET": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0xc3, 0xa7, 0x75, 0x65, 0x6c, 0x61}, "HAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "ACST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CLST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x58, 0x69, 0x6c, 0x65}, "EDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "TMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "CAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}},
	}
}

// Locale returns the current translators string locale
func (ca *ca_FR) Locale() string {
	return ca.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ca_FR'
func (ca *ca_FR) PluralsCardinal() []locales.PluralRule {
	return ca.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ca_FR'
func (ca *ca_FR) PluralsOrdinal() []locales.PluralRule {
	return ca.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ca_FR'
func (ca *ca_FR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ca_FR'
func (ca *ca_FR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 3 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ca_FR'
func (ca *ca_FR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ca_FR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ca_FR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ca *ca_FR) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ca.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ca.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ca.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ca_FR'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ca.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ca.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ca.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ca.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ca.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ca_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca_FR) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := ca.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
