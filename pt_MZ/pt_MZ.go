package pt_MZ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pt_MZ struct {
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

// New returns a new instance of translator for the 'pt_MZ' locale
func New() locales.Translator {
	return &pt_MZ{
		locale:                 "pt_MZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x54, 0x6e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e}, {0x66, 0x65, 0x76}, {0x6d, 0x61, 0x72}, {0x61, 0x62, 0x72}, {0x6d, 0x61, 0x69}, {0x6a, 0x75, 0x6e}, {0x6a, 0x75, 0x6c}, {0x61, 0x67, 0x6f}, {0x73, 0x65, 0x74}, {0x6f, 0x75, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x7a}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x65, 0x69, 0x72, 0x6f}, {0x66, 0x65, 0x76, 0x65, 0x72, 0x65, 0x69, 0x72, 0x6f}, {0x6d, 0x61, 0x72, 0xc3, 0xa7, 0x6f}, {0x61, 0x62, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x69, 0x6f}, {0x6a, 0x75, 0x6e, 0x68, 0x6f}, {0x6a, 0x75, 0x6c, 0x68, 0x6f}, {0x61, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x73, 0x65, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x6f, 0x75, 0x74, 0x75, 0x62, 0x72, 0x6f}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x64, 0x65, 0x7a, 0x65, 0x6d, 0x62, 0x72, 0x6f}},
		daysAbbreviated:        [][]uint8{{0x64, 0x6f, 0x6d}, {0x73, 0x65, 0x67}, {0x74, 0x65, 0x72}, {0x71, 0x75, 0x61}, {0x71, 0x75, 0x69}, {0x73, 0x65, 0x78}, {0x73, 0xc3, 0xa1, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x53}, {0x54}, {0x51}, {0x51}, {0x53}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x6f, 0x6d}, {0x73, 0x65, 0x67}, {0x74, 0x65, 0x72}, {0x71, 0x75, 0x61}, {0x71, 0x75, 0x69}, {0x73, 0x65, 0x78}, {0x73, 0xc3, 0xa1, 0x62}},
		daysWide:               [][]uint8{{0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x6f}, {0x73, 0x65, 0x67, 0x75, 0x6e, 0x64, 0x61, 0x2d, 0x66, 0x65, 0x69, 0x72, 0x61}, {0x74, 0x65, 0x72, 0xc3, 0xa7, 0x61, 0x2d, 0x66, 0x65, 0x69, 0x72, 0x61}, {0x71, 0x75, 0x61, 0x72, 0x74, 0x61, 0x2d, 0x66, 0x65, 0x69, 0x72, 0x61}, {0x71, 0x75, 0x69, 0x6e, 0x74, 0x61, 0x2d, 0x66, 0x65, 0x69, 0x72, 0x61}, {0x73, 0x65, 0x78, 0x74, 0x61, 0x2d, 0x66, 0x65, 0x69, 0x72, 0x61}, {0x73, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x61, 0x6e, 0x74, 0x65, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x64, 0x65, 0x70, 0x6f, 0x69, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:              map[string][]uint8{"WAT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "COT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb4, 0x6d, 0x62, 0x69, 0x61}, "GYT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61}, "NZDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0xc3, 0xa2, 0x6e, 0x64, 0x69, 0x61}, "BOT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x42, 0x6f, 0x6c, 0xc3, 0xad, 0x76, 0x69, 0x61}, "SRT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "CLST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "ECT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x45, 0x71, 0x75, 0x61, 0x64, 0x6f, 0x72}, "WIB": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CAT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "OEZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AWDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HNT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x20, 0x4e, 0x6f, 0x76, 0x61}, "HAT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x20, 0x4e, 0x6f, 0x76, 0x61}, "TMT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa3, 0x6f}, "GMT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x4d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "BT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa3, 0x6f}, "LHST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HKT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "HAST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x48, 0x61, 0x76, 0x61, 0xc3, 0xad, 0x20, 0x65, 0x20, 0x49, 0x6c, 0x68, 0x61, 0x73, 0x20, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x61, 0x73}, "MST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x75}, "OESZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AWST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "JST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xa3, 0x6f}, "AKST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "ChST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "WAST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "UYT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "MDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x75}, "EAT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "NZST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0xc3, 0xa2, 0x6e, 0x64, 0x69, 0x61}, "MEZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ACST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ACDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ACWST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x2d, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CHADT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa2, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "UYST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "EDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WARST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WESZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ARST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "SAST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x6f, 0x20, 0x53, 0x75, 0x6c}, "MESZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "MYT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x4d, 0x61, 0x6c, 0xc3, 0xa1, 0x73, 0x69, 0x61}, "EST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HADT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x48, 0x61, 0x76, 0x61, 0xc3, 0xad, 0x20, 0x65, 0x20, 0x49, 0x6c, 0x68, 0x61, 0x73, 0x20, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x61, 0x73}, "TMST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa3, 0x6f}, "CST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "PST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "COST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb4, 0x6d, 0x62, 0x69, 0x61}, "AEDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "JDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xa3, 0x6f}, "∅∅∅": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x42, 0x72, 0x61, 0x73, 0xc3, 0xad, 0x6c, 0x69, 0x61}, "ART": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "CLT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "GFT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "IST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0xc3, 0x8d, 0x6e, 0x64, 0x69, 0x61}, "ACWDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x2d, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CHAST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "SGT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x61}, "AEST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WIT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AKDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "WEZ": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "PDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "WITA": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "LHDT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HKST": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "WART": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x50, 0x61, 0x64, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ADT": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x72, 0xc3, 0xa3, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa2, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "VET": {0x48, 0x6f, 0x72, 0xc3, 0xa1, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}},
	}
}

// Locale returns the current translators string locale
func (pt *pt_MZ) Locale() string {
	return pt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pt_MZ'
func (pt *pt_MZ) PluralsCardinal() []locales.PluralRule {
	return pt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pt_MZ'
func (pt *pt_MZ) PluralsOrdinal() []locales.PluralRule {
	return pt.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pt_MZ'
func (pt *pt_MZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pt_MZ'
func (pt *pt_MZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pt_MZ'
func (pt *pt_MZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pt.CardinalPluralRule(num1, v1)
	end := pt.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pt_MZ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(pt.decimal) + len(pt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pt.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pt_MZ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pt *pt_MZ) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(pt.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pt.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(pt.decimal) + len(pt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pt.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, pt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pt_MZ'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pt.currencies[currency]
	l := len(s) + len(pt.decimal) + len(pt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pt.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pt.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, pt.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, pt.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, pt.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, pt.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, pt.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, pt.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'pt_MZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (pt *pt_MZ) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pt.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := pt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
