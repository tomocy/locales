package zh_Hans

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type zh_Hans struct {
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

// New returns a new instance of translator for the 'zh_Hans' locale
func New() locales.Translator {
	return &zh_Hans{
		locale:                 "zh_Hans",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x31, 0xe6, 0x9c, 0x88}, {0x32, 0xe6, 0x9c, 0x88}, {0x33, 0xe6, 0x9c, 0x88}, {0x34, 0xe6, 0x9c, 0x88}, {0x35, 0xe6, 0x9c, 0x88}, {0x36, 0xe6, 0x9c, 0x88}, {0x37, 0xe6, 0x9c, 0x88}, {0x38, 0xe6, 0x9c, 0x88}, {0x39, 0xe6, 0x9c, 0x88}, {0x31, 0x30, 0xe6, 0x9c, 0x88}, {0x31, 0x31, 0xe6, 0x9c, 0x88}, {0x31, 0x32, 0xe6, 0x9c, 0x88}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}, {0x37}, {0x38}, {0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe4, 0xb8, 0x80, 0xe6, 0x9c, 0x88}, {0xe4, 0xba, 0x8c, 0xe6, 0x9c, 0x88}, {0xe4, 0xb8, 0x89, 0xe6, 0x9c, 0x88}, {0xe5, 0x9b, 0x9b, 0xe6, 0x9c, 0x88}, {0xe4, 0xba, 0x94, 0xe6, 0x9c, 0x88}, {0xe5, 0x85, 0xad, 0xe6, 0x9c, 0x88}, {0xe4, 0xb8, 0x83, 0xe6, 0x9c, 0x88}, {0xe5, 0x85, 0xab, 0xe6, 0x9c, 0x88}, {0xe4, 0xb9, 0x9d, 0xe6, 0x9c, 0x88}, {0xe5, 0x8d, 0x81, 0xe6, 0x9c, 0x88}, {0xe5, 0x8d, 0x81, 0xe4, 0xb8, 0x80, 0xe6, 0x9c, 0x88}, {0xe5, 0x8d, 0x81, 0xe4, 0xba, 0x8c, 0xe6, 0x9c, 0x88}},
		daysAbbreviated:        [][]uint8{{0xe5, 0x91, 0xa8, 0xe6, 0x97, 0xa5}, {0xe5, 0x91, 0xa8, 0xe4, 0xb8, 0x80}, {0xe5, 0x91, 0xa8, 0xe4, 0xba, 0x8c}, {0xe5, 0x91, 0xa8, 0xe4, 0xb8, 0x89}, {0xe5, 0x91, 0xa8, 0xe5, 0x9b, 0x9b}, {0xe5, 0x91, 0xa8, 0xe4, 0xba, 0x94}, {0xe5, 0x91, 0xa8, 0xe5, 0x85, 0xad}},
		daysNarrow:             [][]uint8{{0xe6, 0x97, 0xa5}, {0xe4, 0xb8, 0x80}, {0xe4, 0xba, 0x8c}, {0xe4, 0xb8, 0x89}, {0xe5, 0x9b, 0x9b}, {0xe4, 0xba, 0x94}, {0xe5, 0x85, 0xad}},
		daysShort:              [][]uint8{{0xe5, 0x91, 0xa8, 0xe6, 0x97, 0xa5}, {0xe5, 0x91, 0xa8, 0xe4, 0xb8, 0x80}, {0xe5, 0x91, 0xa8, 0xe4, 0xba, 0x8c}, {0xe5, 0x91, 0xa8, 0xe4, 0xb8, 0x89}, {0xe5, 0x91, 0xa8, 0xe5, 0x9b, 0x9b}, {0xe5, 0x91, 0xa8, 0xe4, 0xba, 0x94}, {0xe5, 0x91, 0xa8, 0xe5, 0x85, 0xad}},
		daysWide:               [][]uint8{{0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe6, 0x97, 0xa5}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe4, 0xb8, 0x80}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe4, 0xba, 0x8c}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe4, 0xb8, 0x89}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe5, 0x9b, 0x9b}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe4, 0xba, 0x94}, {0xe6, 0x98, 0x9f, 0xe6, 0x9c, 0x9f, 0xe5, 0x85, 0xad}},
		periodsAbbreviated:     [][]uint8{{0xe4, 0xb8, 0x8a, 0xe5, 0x8d, 0x88}, {0xe4, 0xb8, 0x8b, 0xe5, 0x8d, 0x88}},
		periodsNarrow:          [][]uint8{{0xe4, 0xb8, 0x8a, 0xe5, 0x8d, 0x88}, {0xe4, 0xb8, 0x8b, 0xe5, 0x8d, 0x88}},
		periodsWide:            [][]uint8{{0xe4, 0xb8, 0x8a, 0xe5, 0x8d, 0x88}, {0xe4, 0xb8, 0x8b, 0xe5, 0x8d, 0x88}},
		erasAbbreviated:        [][]uint8{{0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83, 0xe5, 0x89, 0x8d}, {0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83}},
		erasNarrow:             [][]uint8{{0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83, 0xe5, 0x89, 0x8d}, {0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83}},
		erasWide:               [][]uint8{{0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83, 0xe5, 0x89, 0x8d}, {0xe5, 0x85, 0xac, 0xe5, 0x85, 0x83}},
		timezones:              map[string][]uint8{"WIB": {0xe5, 0x8d, 0xb0, 0xe5, 0xba, 0xa6, 0xe5, 0xb0, 0xbc, 0xe8, 0xa5, 0xbf, 0xe4, 0xba, 0x9a, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "NZST": {0xe6, 0x96, 0xb0, 0xe8, 0xa5, 0xbf, 0xe5, 0x85, 0xb0, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CAT": {0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe9, 0x9d, 0x9e, 0xe6, 0xb4, 0xb2, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "COST": {0xe5, 0x93, 0xa5, 0xe4, 0xbc, 0xa6, 0xe6, 0xaf, 0x94, 0xe4, 0xba, 0x9a, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "SAST": {0xe5, 0x8d, 0x97, 0xe9, 0x83, 0xa8, 0xe9, 0x9d, 0x9e, 0xe6, 0xb4, 0xb2, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "OEZ": {0xe4, 0xb8, 0x9c, 0xe6, 0xac, 0xa7, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "EAT": {0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe9, 0x9d, 0x9e, 0xe6, 0xb4, 0xb2, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AKST": {0xe9, 0x98, 0xbf, 0xe6, 0x8b, 0x89, 0xe6, 0x96, 0xaf, 0xe5, 0x8a, 0xa0, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "NZDT": {0xe6, 0x96, 0xb0, 0xe8, 0xa5, 0xbf, 0xe5, 0x85, 0xb0, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "UYST": {0xe4, 0xb9, 0x8c, 0xe6, 0x8b, 0x89, 0xe5, 0x9c, 0xad, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ARST": {0xe9, 0x98, 0xbf, 0xe6, 0xa0, 0xb9, 0xe5, 0xbb, 0xb7, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ACWDT": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0xad, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WEZ": {0xe8, 0xa5, 0xbf, 0xe6, 0xac, 0xa7, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "MESZ": {0xe4, 0xb8, 0xad, 0xe6, 0xac, 0xa7, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "BOT": {0xe7, 0x8e, 0xbb, 0xe5, 0x88, 0xa9, 0xe7, 0xbb, 0xb4, 0xe4, 0xba, 0x9a, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "MST": {0xe6, 0xbe, 0xb3, 0xe9, 0x97, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "SGT": {0xe6, 0x96, 0xb0, 0xe5, 0x8a, 0xa0, 0xe5, 0x9d, 0xa1, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AWDT": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WAT": {0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe9, 0x9d, 0x9e, 0xe6, 0xb4, 0xb2, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WAST": {0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe9, 0x9d, 0x9e, 0xe6, 0xb4, 0xb2, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WITA": {0xe5, 0x8d, 0xb0, 0xe5, 0xba, 0xa6, 0xe5, 0xb0, 0xbc, 0xe8, 0xa5, 0xbf, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "SRT": {0xe8, 0x8b, 0x8f, 0xe9, 0x87, 0x8c, 0xe5, 0x8d, 0x97, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AWST": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "IST": {0xe5, 0x8d, 0xb0, 0xe5, 0xba, 0xa6, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HKT": {0xe9, 0xa6, 0x99, 0xe6, 0xb8, 0xaf, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ECT": {0xe5, 0x8e, 0x84, 0xe7, 0x93, 0x9c, 0xe5, 0xa4, 0x9a, 0xe5, 0xb0, 0x94, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HAT": {0xe7, 0xba, 0xbd, 0xe8, 0x8a, 0xac, 0xe5, 0x85, 0xb0, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "TMST": {0xe5, 0x9c, 0x9f, 0xe5, 0xba, 0x93, 0xe6, 0x9b, 0xbc, 0xe6, 0x96, 0xaf, 0xe5, 0x9d, 0xa6, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HADT": {0xe5, 0xa4, 0x8f, 0xe5, 0xa8, 0x81, 0xe5, 0xa4, 0xb7, 0x2d, 0xe9, 0x98, 0xbf, 0xe7, 0x95, 0x99, 0xe7, 0x94, 0xb3, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "OESZ": {0xe4, 0xb8, 0x9c, 0xe6, 0xac, 0xa7, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "PST": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe5, 0xa4, 0xaa, 0xe5, 0xb9, 0xb3, 0xe6, 0xb4, 0x8b, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "LHST": {0xe8, 0xb1, 0xaa, 0xe5, 0x8b, 0x8b, 0xe7, 0x88, 0xb5, 0xe5, 0xb2, 0x9b, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CHAST": {0xe6, 0x9f, 0xa5, 0xe5, 0x9d, 0xa6, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HNT": {0xe7, 0xba, 0xbd, 0xe8, 0x8a, 0xac, 0xe5, 0x85, 0xb0, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "VET": {0xe5, 0xa7, 0x94, 0xe5, 0x86, 0x85, 0xe7, 0x91, 0x9e, 0xe6, 0x8b, 0x89, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HAST": {0xe5, 0xa4, 0x8f, 0xe5, 0xa8, 0x81, 0xe5, 0xa4, 0xb7, 0x2d, 0xe9, 0x98, 0xbf, 0xe7, 0x95, 0x99, 0xe7, 0x94, 0xb3, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CHADT": {0xe6, 0x9f, 0xa5, 0xe5, 0x9d, 0xa6, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AEST": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "GFT": {0xe6, 0xb3, 0x95, 0xe5, 0xb1, 0x9e, 0xe5, 0x9c, 0xad, 0xe4, 0xba, 0x9a, 0xe9, 0x82, 0xa3, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "BT": {0xe4, 0xb8, 0x8d, 0xe4, 0xb8, 0xb9, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ACDT": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "PDT": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe5, 0xa4, 0xaa, 0xe5, 0xb9, 0xb3, 0xe6, 0xb4, 0x8b, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "LHDT": {0xe8, 0xb1, 0xaa, 0xe5, 0x8b, 0x8b, 0xe7, 0x88, 0xb5, 0xe5, 0xb2, 0x9b, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "EST": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AST": {0xe5, 0xa4, 0xa7, 0xe8, 0xa5, 0xbf, 0xe6, 0xb4, 0x8b, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CDT": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WART": {0xe9, 0x98, 0xbf, 0xe6, 0xa0, 0xb9, 0xe5, 0xbb, 0xb7, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AEDT": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "∅∅∅": {0xe5, 0xb7, 0xb4, 0xe8, 0xa5, 0xbf, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "MDT": {0xe6, 0xbe, 0xb3, 0xe9, 0x97, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "COT": {0xe5, 0x93, 0xa5, 0xe4, 0xbc, 0xa6, 0xe6, 0xaf, 0x94, 0xe4, 0xba, 0x9a, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ACWST": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0xad, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "MYT": {0xe9, 0xa9, 0xac, 0xe6, 0x9d, 0xa5, 0xe8, 0xa5, 0xbf, 0xe4, 0xba, 0x9a, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "TMT": {0xe5, 0x9c, 0x9f, 0xe5, 0xba, 0x93, 0xe6, 0x9b, 0xbc, 0xe6, 0x96, 0xaf, 0xe5, 0x9d, 0xa6, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ChST": {0xe6, 0x9f, 0xa5, 0xe8, 0x8e, 0xab, 0xe7, 0xbd, 0x97, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CLST": {0xe6, 0x99, 0xba, 0xe5, 0x88, 0xa9, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WARST": {0xe9, 0x98, 0xbf, 0xe6, 0xa0, 0xb9, 0xe5, 0xbb, 0xb7, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "GYT": {0xe5, 0x9c, 0xad, 0xe4, 0xba, 0x9a, 0xe9, 0x82, 0xa3, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "JST": {0xe6, 0x97, 0xa5, 0xe6, 0x9c, 0xac, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "JDT": {0xe6, 0x97, 0xa5, 0xe6, 0x9c, 0xac, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "MEZ": {0xe4, 0xb8, 0xad, 0xe6, 0xac, 0xa7, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "UYT": {0xe4, 0xb9, 0x8c, 0xe6, 0x8b, 0x89, 0xe5, 0x9c, 0xad, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CLT": {0xe6, 0x99, 0xba, 0xe5, 0x88, 0xa9, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "AKDT": {0xe9, 0x98, 0xbf, 0xe6, 0x8b, 0x89, 0xe6, 0x96, 0xaf, 0xe5, 0x8a, 0xa0, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ADT": {0xe5, 0xa4, 0xa7, 0xe8, 0xa5, 0xbf, 0xe6, 0xb4, 0x8b, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WESZ": {0xe8, 0xa5, 0xbf, 0xe6, 0xac, 0xa7, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ACST": {0xe6, 0xbe, 0xb3, 0xe5, 0xa4, 0xa7, 0xe5, 0x88, 0xa9, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "ART": {0xe9, 0x98, 0xbf, 0xe6, 0xa0, 0xb9, 0xe5, 0xbb, 0xb7, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "HKST": {0xe9, 0xa6, 0x99, 0xe6, 0xb8, 0xaf, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "GMT": {0xe6, 0xa0, 0xbc, 0xe6, 0x9e, 0x97, 0xe5, 0xb0, 0xbc, 0xe6, 0xb2, 0xbb, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "CST": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0xa0, 0x87, 0xe5, 0x87, 0x86, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "EDT": {0xe5, 0x8c, 0x97, 0xe7, 0xbe, 0x8e, 0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe4, 0xbb, 0xa4, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}, "WIT": {0xe5, 0x8d, 0xb0, 0xe5, 0xba, 0xa6, 0xe5, 0xb0, 0xbc, 0xe8, 0xa5, 0xbf, 0xe4, 0xba, 0x9a, 0xe4, 0xb8, 0x9c, 0xe9, 0x83, 0xa8, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4}},
	}
}

// Locale returns the current translators string locale
func (zh *zh_Hans) Locale() string {
	return zh.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'zh_Hans'
func (zh *zh_Hans) PluralsCardinal() []locales.PluralRule {
	return zh.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'zh_Hans'
func (zh *zh_Hans) PluralsOrdinal() []locales.PluralRule {
	return zh.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'zh_Hans'
func (zh *zh_Hans) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'zh_Hans'
func (zh *zh_Hans) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'zh_Hans'
func (zh *zh_Hans) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'zh_Hans' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(zh.decimal) + len(zh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zh.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zh.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'zh_Hans' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (zh *zh_Hans) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(zh.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zh.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zh.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, zh.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zh.currencies[currency]
	l := len(s) + len(zh.decimal) + len(zh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zh.group[0])
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
		b = append(b, zh.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, zh.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'zh_Hans'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zh.currencies[currency]
	l := len(s) + len(zh.decimal) + len(zh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zh.group[0])
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

		for j := len(zh.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, zh.currencyNegativePrefix[j])
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
			b = append(b, zh.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, zh.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)
	b = append(b, zh.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, zh.periodsAbbreviated[0]...)
	} else {
		b = append(b, zh.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zh.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, zh.periodsAbbreviated[0]...)
	} else {
		b = append(b, zh.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zh.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zh.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zh.periodsAbbreviated[0]...)
	} else {
		b = append(b, zh.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zh.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zh.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'zh_Hans'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zh *zh_Hans) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()
	if btz, ok := zh.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zh.periodsAbbreviated[0]...)
	} else {
		b = append(b, zh.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zh.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zh.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}
