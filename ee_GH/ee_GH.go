package ee_GH

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ee_GH struct {
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

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale:                 "ee_GH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x64, 0x7a, 0x76}, {0x64, 0x7a, 0x64}, {0x74, 0x65, 0x64}, {0x61, 0x66, 0xc9, 0x94}, {0x64, 0x61, 0x6d}, {0x6d, 0x61, 0x73}, {0x73, 0x69, 0x61}, {0x64, 0x65, 0x61}, {0x61, 0x6e, 0x79}, {0x6b, 0x65, 0x6c}, {0x61, 0x64, 0x65}, {0x64, 0x7a, 0x6d}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x64}, {0x64}, {0x74}, {0x61}, {0x64}, {0x6d}, {0x73}, {0x64}, {0x61}, {0x6b}, {0x61}, {0x64}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x64, 0x7a, 0x6f, 0x76, 0x65}, {0x64, 0x7a, 0x6f, 0x64, 0x7a, 0x65}, {0x74, 0x65, 0x64, 0x6f, 0x78, 0x65}, {0x61, 0x66, 0xc9, 0x94, 0x66, 0xc4, 0xa9, 0x65}, {0x64, 0x61, 0x6d, 0x61}, {0x6d, 0x61, 0x73, 0x61}, {0x73, 0x69, 0x61, 0x6d, 0x6c, 0xc9, 0x94, 0x6d}, {0x64, 0x65, 0x61, 0x73, 0x69, 0x61, 0x6d, 0x69, 0x6d, 0x65}, {0x61, 0x6e, 0x79, 0xc9, 0x94, 0x6e, 0x79, 0xc9, 0x94}, {0x6b, 0x65, 0x6c, 0x65}, {0x61, 0x64, 0x65, 0xc9, 0x9b, 0x6d, 0x65, 0x6b, 0x70, 0xc9, 0x94, 0x78, 0x65}, {0x64, 0x7a, 0x6f, 0x6d, 0x65}},
		daysAbbreviated:        [][]uint8{{0x6b, 0xc9, 0x94, 0x73}, {0x64, 0x7a, 0x6f}, {0x62, 0x6c, 0x61}, {0x6b, 0x75, 0xc9, 0x96}, {0x79, 0x61, 0x77}, {0x66, 0x69, 0xc9, 0x96}, {0x6d, 0x65, 0x6d}},
		daysNarrow:             [][]uint8{{0x6b}, {0x64}, {0x62}, {0x6b}, {0x79}, {0x66}, {0x6d}},
		daysShort:              [][]uint8{{0x6b, 0xc9, 0x94, 0x73}, {0x64, 0x7a, 0x6f}, {0x62, 0x6c, 0x61}, {0x6b, 0x75, 0xc9, 0x96}, {0x79, 0x61, 0x77}, {0x66, 0x69, 0xc9, 0x96}, {0x6d, 0x65, 0x6d}},
		daysWide:               [][]uint8{{0x6b, 0xc9, 0x94, 0x73, 0x69, 0xc9, 0x96, 0x61}, {0x64, 0x7a, 0x6f, 0xc9, 0x96, 0x61}, {0x62, 0x6c, 0x61, 0xc9, 0x96, 0x61}, {0x6b, 0x75, 0xc9, 0x96, 0x61}, {0x79, 0x61, 0x77, 0x6f, 0xc9, 0x96, 0x61}, {0x66, 0x69, 0xc9, 0x96, 0x61}, {0x6d, 0x65, 0x6d, 0x6c, 0x65, 0xc9, 0x96, 0x61}},
		periodsAbbreviated:     [][]uint8{{0xc5, 0x8b, 0x64, 0x69}, {0xc9, 0xa3, 0x65, 0x74, 0x72, 0xc9, 0x94}},
		periodsNarrow:          [][]uint8{{0xc5, 0x8b}, {0xc9, 0xa3}},
		periodsWide:            [][]uint8{{0xc5, 0x8b, 0x64, 0x69}, {0xc9, 0xa3, 0x65, 0x74, 0x72, 0xc9, 0x94}},
		erasAbbreviated:        [][]uint8{{0x68, 0x59}, {0x59, 0xc5, 0x8b}},
		erasNarrow:             [][]uint8{{0x68, 0x59}, {0x59, 0xc5, 0x8b}},
		erasWide:               [][]uint8{{0x48, 0x61, 0x66, 0x69, 0x20, 0x59, 0x65, 0x73, 0x75, 0x20, 0x56, 0x61, 0x20, 0x44, 0x6f, 0x20, 0xc5, 0x8b, 0x67, 0xc9, 0x94}, {0x59, 0x65, 0x73, 0x75, 0x20, 0xc5, 0x8a, 0xc9, 0x94, 0x6c, 0x69}},
		timezones:              map[string][]uint8{"AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "COT": {0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CLST": {0x54, 0x73, 0x69, 0x6c, 0x65, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "WART": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "MESZ": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "HAT": {0x4e, 0x69, 0x75, 0x66, 0x61, 0x75, 0x6e, 0xc9, 0x96, 0x6c, 0x61, 0x6e, 0xc9, 0x96, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "PDT": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "PST": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "TMST": {0x54, 0xc9, 0x9b, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "ACDT": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "COST": {0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CDT": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "SGT": {0x53, 0x47, 0x54}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "GFT": {0x46, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "∅∅∅": {0x42, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "ECT": {0x49, 0x6b, 0x75, 0x65, 0x64, 0xc9, 0x94, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "EST": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "AEST": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "AWST": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "BT": {0x42, 0x54}, "HKST": {0x48, 0xc9, 0x94, 0x6e, 0x67, 0x20, 0x4b, 0xc9, 0x94, 0x6e, 0x67, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "MST": {0x4d, 0x61, 0x6b, 0x61, 0x75, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "WIT": {0x57, 0x49, 0x54}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "HKT": {0x48, 0xc9, 0x94, 0x6e, 0x67, 0x20, 0x4b, 0xc9, 0x94, 0x6e, 0x67, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x20, 0x6d, 0x65}, "IST": {0x49, 0x53, 0x54}, "WARST": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "ChST": {0x43, 0x68, 0x53, 0x54}, "WAST": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CLT": {0x54, 0x73, 0x69, 0x6c, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "OESZ": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "ACWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0xc9, 0xa3, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0x66, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "HNT": {0x4e, 0x69, 0x75, 0x66, 0x61, 0x75, 0x6e, 0xc9, 0x96, 0x6c, 0x61, 0x6e, 0xc9, 0x96, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "WIB": {0x57, 0x49, 0x42}, "EAT": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "MYT": {0x4d, 0x59, 0x54}, "OEZ": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "ACWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0xc9, 0xa3, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0x66, 0x65, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "AWDT": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "MEZ": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "WAT": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "WITA": {0x57, 0x49, 0x54, 0x41}, "WEZ": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "WESZ": {0xc6, 0x94, 0x65, 0x74, 0x6f, 0xc9, 0x96, 0x6f, 0xc6, 0x92, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "JST": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "GYT": {0x47, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x20, 0x64, 0x7a, 0x6f, 0x6d, 0x65, 0xc5, 0x8b, 0xc9, 0x94, 0x6c, 0x69, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "EDT": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "AEDT": {0xc6, 0x94, 0x65, 0x64, 0x7a, 0x65, 0xc6, 0x92, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CAT": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "CST": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "MDT": {0x4d, 0x61, 0x6b, 0x61, 0x75, 0x20, 0xc5, 0x8b, 0x6b, 0x65, 0x6b, 0x65, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}, "TMT": {0x54, 0xc9, 0x9b, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "ACST": {0x54, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0xc9, 0x96, 0x6f, 0x61, 0x6e, 0x79, 0x69, 0x6d, 0x65}, "SAST": {0x41, 0x6e, 0x79, 0x69, 0x65, 0x68, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x67, 0x61, 0xc6, 0x92, 0x6f, 0xc6, 0x92, 0x6f, 0x6d, 0x65}},
	}
}

// Locale returns the current translators string locale
func (ee *ee_GH) Locale() string {
	return ee.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsCardinal() []locales.PluralRule {
	return ee.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsOrdinal() []locales.PluralRule {
	return ee.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ee_GH'
func (ee *ee_GH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ee *ee_GH) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(ee.decimal) + len(ee.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ee.decimal) - 1; j >= 0; j-- {
				b = append(b, ee.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ee.group) - 1; j >= 0; j-- {
					b = append(b, ee.group[j])
				}

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
		for j := len(ee.minus) - 1; j >= 0; j-- {
			b = append(b, ee.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(ee.decimal) + len(ee.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ee.decimal) - 1; j >= 0; j-- {
				b = append(b, ee.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ee.group) - 1; j >= 0; j-- {
					b = append(b, ee.group[j])
				}

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

		for j := len(ee.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ee.currencyNegativePrefix[j])
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
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ee.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ee.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ee_GH'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ee *ee_GH) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := ee.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
