package fr_LU

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fr_LU struct {
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

// New returns a new instance of translator for the 'fr_LU' locale
func New() locales.Translator {
	return &fr_LU{
		locale:                 "fr_LU",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0xd9, 0xab},
		group:                  []byte{0x2e},
		minus:                  []byte{0xe2, 0x80, 0x8f, 0xe2, 0x88, 0x92},
		percent:                []byte{0xd9, 0xaa},
		perMille:               []byte{0xd8, 0x89},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x76, 0x2e}, {0x66, 0xc3, 0xa9, 0x76, 0x72, 0x2e}, {0x6d, 0x61, 0x72, 0x73}, {0x61, 0x76, 0x72, 0x2e}, {0x6d, 0x61, 0x69}, {0x6a, 0x75, 0x69, 0x6e}, {0x6a, 0x75, 0x69, 0x6c, 0x2e}, {0x61, 0x6f, 0xc3, 0xbb, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x2e}, {0x6f, 0x63, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0xc3, 0xa9, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x76, 0x69, 0x65, 0x72}, {0x66, 0xc3, 0xa9, 0x76, 0x72, 0x69, 0x65, 0x72}, {0x6d, 0x61, 0x72, 0x73}, {0x61, 0x76, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x69}, {0x6a, 0x75, 0x69, 0x6e}, {0x6a, 0x75, 0x69, 0x6c, 0x6c, 0x65, 0x74}, {0x61, 0x6f, 0xc3, 0xbb, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x6f, 0x63, 0x74, 0x6f, 0x62, 0x72, 0x65}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0xc3, 0xa9, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x64, 0x69, 0x6d, 0x2e}, {0x6c, 0x75, 0x6e, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x6d, 0x65, 0x72, 0x2e}, {0x6a, 0x65, 0x75, 0x2e}, {0x76, 0x65, 0x6e, 0x2e}, {0x73, 0x61, 0x6d, 0x2e}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x4d}, {0x4a}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x69}, {0x6c, 0x75}, {0x6d, 0x61}, {0x6d, 0x65}, {0x6a, 0x65}, {0x76, 0x65}, {0x73, 0x61}},
		daysWide:               [][]uint8{{0x64, 0x69, 0x6d, 0x61, 0x6e, 0x63, 0x68, 0x65}, {0x6c, 0x75, 0x6e, 0x64, 0x69}, {0x6d, 0x61, 0x72, 0x64, 0x69}, {0x6d, 0x65, 0x72, 0x63, 0x72, 0x65, 0x64, 0x69}, {0x6a, 0x65, 0x75, 0x64, 0x69}, {0x76, 0x65, 0x6e, 0x64, 0x72, 0x65, 0x64, 0x69}, {0x73, 0x61, 0x6d, 0x65, 0x64, 0x69}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x61, 0x76, 0x2e, 0x20, 0x4a, 0x2e, 0x2d, 0x43, 0x2e}, {0x61, 0x70, 0x2e, 0x20, 0x4a, 0x2e, 0x2d, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{{0x61, 0x76, 0x2e, 0x20, 0x4a, 0x2e, 0x2d, 0x43, 0x2e}, {0x61, 0x70, 0x2e, 0x20, 0x4a, 0x2e, 0x2d, 0x43, 0x2e}},
		erasWide:               [][]uint8{{0x61, 0x76, 0x61, 0x6e, 0x74, 0x20, 0x4a, 0xc3, 0xa9, 0x73, 0x75, 0x73, 0x2d, 0x43, 0x68, 0x72, 0x69, 0x73, 0x74}, {0x61, 0x70, 0x72, 0xc3, 0xa8, 0x73, 0x20, 0x4a, 0xc3, 0xa9, 0x73, 0x75, 0x73, 0x2d, 0x43, 0x68, 0x72, 0x69, 0x73, 0x74}},
		timezones:              map[string][]uint8{"OEZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74}, "UYST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "WITA": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x75, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x65, 0x6e}, "SRT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x75, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "SAST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x71, 0x75, 0x65, 0x20, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x65}, "SGT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x75, 0x72}, "EST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x69, 0x6e}, "VET": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x75, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "GMT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6d, 0x6f, 0x79, 0x65, 0x6e, 0x6e, 0x65, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "UYT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "HADT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x20, 0x2d, 0x20, 0x41, 0x6c, 0xc3, 0xa9, 0x6f, 0x75, 0x74, 0x69, 0x65, 0x6e, 0x6e, 0x65, 0x73}, "PDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x71, 0x75, 0x65}, "ARST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65}, "MYT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x69, 0x73, 0x69, 0x65}, "ECT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x89, 0x71, 0x75, 0x61, 0x74, 0x65, 0x75, 0x72}, "MDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x6f}, "ACWDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "CHADT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x73, 0x20, 0xc3, 0xae, 0x6c, 0x65, 0x73, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AEST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "CDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65}, "PST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x71, 0x75, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x69, 0x6e}, "HKST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "AKST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "AST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x71, 0x75, 0x65}, "ADT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x71, 0x75, 0x65}, "WEZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74}, "NZST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x4e, 0x6f, 0x75, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x2d, 0x5a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x65}, "GFT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x65, 0x20, 0x66, 0x72, 0x61, 0x6e, 0xc3, 0xa7, 0x61, 0x69, 0x73, 0x65}, "JDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e}, "CST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x69, 0x6e}, "IST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x65}, "ACWST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "CHAST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x73, 0x20, 0xc3, 0xae, 0x6c, 0x65, 0x73, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "GYT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x75, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "JST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e}, "AWST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "NZDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x4e, 0x6f, 0x75, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x2d, 0x5a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x65}, "BOT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x65}, "WAT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x71, 0x75, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74}, "WAST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x71, 0x75, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74}, "MST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x6f}, "ART": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65}, "HKT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "HNT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x65, 0x2d, 0x4e, 0x65, 0x75, 0x76, 0x65}, "AEDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "TMST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0xc3, 0xa9, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ACST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "COT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x65}, "COST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x65}, "AWDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "∅∅∅": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x42, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61}, "ACDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65}, "CLT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x69}, "CLST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x75, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x69}, "TMT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x75, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0xc3, 0xa9, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "AKDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "WART": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e}, "WARST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e}, "WIB": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x65, 0x6e}, "ChST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x73, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "WESZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x75, 0x65, 0x73, 0x74}, "BT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x75, 0x20, 0x42, 0x68, 0x6f, 0x75, 0x74, 0x61, 0x6e}, "CAT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x71, 0x75, 0x65, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "OESZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74}, "MEZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "LHST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "EAT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x66, 0x72, 0x69, 0x71, 0x75, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74}, "HAST": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x20, 0x2d, 0x20, 0x41, 0x6c, 0xc3, 0xa9, 0x6f, 0x75, 0x74, 0x69, 0x65, 0x6e, 0x6e, 0x65, 0x73}, "LHDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WIT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x73, 0x69, 0x65, 0x6e}, "MESZ": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "EDT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74}, "HAT": {0x68, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0xa9, 0x74, 0xc3, 0xa9, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x65, 0x2d, 0x4e, 0x65, 0x75, 0x76, 0x65}},
	}
}

// Locale returns the current translators string locale
func (fr *fr_LU) Locale() string {
	return fr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fr_LU'
func (fr *fr_LU) PluralsCardinal() []locales.PluralRule {
	return fr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fr_LU'
func (fr *fr_LU) PluralsOrdinal() []locales.PluralRule {
	return fr.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fr_LU'
func (fr *fr_LU) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fr_LU'
func (fr *fr_LU) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fr_LU'
func (fr *fr_LU) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := fr.CardinalPluralRule(num1, v1)
	end := fr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fr_LU' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fr_LU' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fr *fr_LU) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fr.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fr.percentSuffix...)

	b = append(b, fr.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fr_LU'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(fr.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fr.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtDateShort(t time.Time) []byte {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, fr.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'fr_LU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fr *fr_LU) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := fr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
