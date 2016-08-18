package sah_RU

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sah_RU struct {
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

// New returns a new instance of translator for the 'sah_RU' locale
func New() locales.Translator {
	return &sah_RU{
		locale:                 "sah_RU",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0x4b},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0x4b},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xd0, 0xa2, 0xd0, 0xbe, 0xd1, 0x85, 0xd1, 0x81}, {0xd0, 0x9e, 0xd0, 0xbb, 0xd1, 0x83, 0xd0, 0xbd}, {0xd0, 0x9a, 0xd0, 0xbb, 0xd0, 0xbd, 0x5f, 0xd1, 0x82, 0xd1, 0x82, 0xd1, 0x80}, {0xd0, 0x9c, 0xd1, 0x83, 0xd1, 0x81, 0x5f, 0xd1, 0x83, 0xd1, 0x81, 0xd1, 0x82}, {0xd0, 0xab, 0xd0, 0xb0, 0xd0, 0xbc, 0x5f, 0xd0, 0xb9, 0xd0, 0xbd}, {0xd0, 0x91, 0xd1, 0x8d, 0xd1, 0x81, 0x5f, 0xd0, 0xb9, 0xd0, 0xbd}, {0xd0, 0x9e, 0xd1, 0x82, 0x5f, 0xd0, 0xb9, 0xd0, 0xbd}, {0xd0, 0x90, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb4, 0xd1, 0x8c, 0x5f, 0xd0, 0xb9, 0xd0, 0xbd}, {0xd0, 0x91, 0xd0, 0xbb, 0xd2, 0x95, 0xd0, 0xbd, 0x5f, 0xd0, 0xb9, 0xd0, 0xbd}, {0xd0, 0x90, 0xd0, 0xbb, 0xd1, 0x82}, {0xd0, 0xa1, 0xd1, 0x8d, 0xd1, 0x82}, {0xd0, 0x90, 0xd1, 0x85, 0xd1, 0x81}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xd0, 0xa2}, {0xd0, 0x9e}, {0xd0, 0x9a}, {0xd0, 0x9c}, {0xd0, 0xab}, {0xd0, 0x91}, {0xd0, 0x9e}, {0xd0, 0x90}, {0xd0, 0x91}, {0xd0, 0x90}, {0xd0, 0xa1}, {0xd0, 0x90}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xd0, 0xa2, 0xd0, 0xbe, 0xd1, 0x85, 0xd1, 0x81, 0xd1, 0x83, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd1, 0x83}, {0xd0, 0x9e, 0xd0, 0xbb, 0xd1, 0x83, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd1, 0x83}, {0xd0, 0x9a, 0xd1, 0x83, 0xd0, 0xbb, 0xd1, 0x83, 0xd0, 0xbd, 0x20, 0xd1, 0x82, 0xd1, 0x83, 0xd1, 0x82, 0xd0, 0xb0, 0xd1, 0x80}, {0xd0, 0x9c, 0xd1, 0x83, 0xd1, 0x83, 0xd1, 0x81, 0x20, 0xd1, 0x83, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb0, 0xd1, 0x80}, {0xd0, 0xab, 0xd0, 0xb0, 0xd0, 0xbc, 0x20, 0xd1, 0x8b, 0xd0, 0xb9, 0xd1, 0x8b, 0xd0, 0xbd}, {0xd0, 0x91, 0xd1, 0x8d, 0xd1, 0x81, 0x20, 0xd1, 0x8b, 0xd0, 0xb9, 0xd1, 0x8b, 0xd0, 0xbd}, {0xd0, 0x9e, 0xd1, 0x82, 0x20, 0xd1, 0x8b, 0xd0, 0xb9, 0xd1, 0x8b, 0xd0, 0xbd}, {0xd0, 0x90, 0xd1, 0x82, 0xd1, 0x8b, 0xd1, 0x80, 0xd0, 0xb4, 0xd1, 0x8c, 0xd1, 0x8b, 0xd1, 0x85, 0x20, 0xd1, 0x8b, 0xd0, 0xb9, 0xd1, 0x8b, 0xd0, 0xbd}, {0xd0, 0x91, 0xd0, 0xb0, 0xd0, 0xbb, 0xd0, 0xb0, 0xd2, 0x95, 0xd0, 0xb0, 0xd0, 0xbd, 0x20, 0xd1, 0x8b, 0xd0, 0xb9, 0xd1, 0x8b, 0xd0, 0xbd}, {0xd0, 0x90, 0xd0, 0xbb, 0xd1, 0x82, 0xd1, 0x8b, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd1, 0x8b}, {0xd0, 0xa1, 0xd1, 0x8d, 0xd1, 0x82, 0xd0, 0xb8, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd0, 0xb8}, {0xd0, 0x90, 0xd1, 0x85, 0xd1, 0x81, 0xd1, 0x8b, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd1, 0x8b}},
		daysAbbreviated:        [][]uint8{{0xd0, 0x91, 0xd1, 0x81}, {0xd0, 0x91, 0xd0, 0xbd}, {0xd0, 0x9e, 0xd0, 0xbf}, {0xd0, 0xa1, 0xd1, 0x8d}, {0xd0, 0xa7, 0xd0, 0xbf}, {0xd0, 0x91, 0xd1, 0x8d}, {0xd0, 0xa1, 0xd0, 0xb1}},
		daysNarrow:             [][]uint8{{0xd0, 0x91}, {0xd0, 0x91}, {0xd0, 0x9e}, {0xd0, 0xa1}, {0xd0, 0xa7}, {0xd0, 0x91}, {0xd0, 0xa1}},
		daysWide:               [][]uint8{{0xd0, 0x91, 0xd0, 0xb0, 0xd1, 0x81, 0xd0, 0xba, 0xd1, 0x8b, 0xd2, 0xbb, 0xd1, 0x8b, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd0, 0xb0}, {0xd0, 0x91, 0xd1, 0x8d, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0, 0xb8, 0xd1, 0x8d, 0xd0, 0xbb, 0xd0, 0xb8, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd0, 0xb8, 0xd0, 0xba}, {0xd0, 0x9e, 0xd0, 0xbf, 0xd1, 0x82, 0xd1, 0x83, 0xd0, 0xbe, 0xd1, 0x80, 0xd1, 0x83, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8c, 0xd1, 0x83, 0xd0, 0xba}, {0xd0, 0xa1, 0xd1, 0x8d, 0xd1, 0x80, 0xd1, 0x8d, 0xd0, 0xb4, 0xd1, 0x8d}, {0xd0, 0xa7, 0xd1, 0x8d, 0xd0, 0xbf, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x8d, 0xd1, 0x80}, {0xd0, 0x91, 0xd1, 0x8d, 0xd1, 0x8d, 0xd1, 0x82, 0xd0, 0xb8, 0xd2, 0xa5, 0xd1, 0x81, 0xd1, 0x8d}, {0xd0, 0xa1, 0xd1, 0x83, 0xd0, 0xb1, 0xd1, 0x83, 0xd0, 0xbe, 0xd1, 0x82, 0xd0, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0xd0, 0xad, 0xd0, 0x98}, {0xd0, 0xad, 0xd0, 0x9a}},
		periodsWide:            [][]uint8{{0xd0, 0xad, 0xd0, 0x98}, {0xd0, 0xad, 0xd0, 0x9a}},
		erasAbbreviated:        [][]uint8{{0xd0, 0xb1, 0x2e, 0x20, 0xd1, 0x8d, 0x2e, 0x20, 0xd0, 0xb8, 0x2e}, {0xd0, 0xb1, 0x2e, 0x20, 0xd1, 0x8d}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"JST": {0x4a, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "AST": {0x41, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "ACST": {0x41, 0x43, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "GMT": {0x47, 0x4d, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "IST": {0x49, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "PDT": {0x50, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "CDT": {0x43, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "SRT": {0x53, 0x52, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "ART": {0x41, 0x52, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "WAT": {0x57, 0x41, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "BT": {0x42, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "GFT": {0x47, 0x46, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "CAT": {0x43, 0x41, 0x54}, "CST": {0x43, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "ADT": {0x41, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "EAT": {0x45, 0x41, 0x54}, "MST": {0x4d, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "GYT": {0x47, 0x59, 0x54}, "WIT": {0x57, 0x49, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "LHST": {0x4c, 0x48, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (sah *sah_RU) Locale() string {
	return sah.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sah_RU'
func (sah *sah_RU) PluralsCardinal() []locales.PluralRule {
	return sah.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sah_RU'
func (sah *sah_RU) PluralsOrdinal() []locales.PluralRule {
	return sah.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sah_RU'
func (sah *sah_RU) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sah_RU'
func (sah *sah_RU) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sah_RU'
func (sah *sah_RU) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sah_RU' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sah_RU' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sah *sah_RU) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sah.currencies[currency]
	l := len(s) + len(sah.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sah.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(sah.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sah.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(sah.minus) - 1; j >= 0; j-- {
			b = append(b, sah.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sah.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sah_RU'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sah.currencies[currency]
	l := len(s) + len(sah.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sah.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sah.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sah.currencyNegativePrefix[j])
		}

		for j := len(sah.minus) - 1; j >= 0; j-- {
			b = append(b, sah.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(sah.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sah.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, sah.currencyNegativeSuffix...)
	} else {

		b = append(b, sah.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sah.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sah.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd1, 0x81, 0xd1, 0x8b, 0xd0, 0xbb, 0x27, 0x20}...)
	b = append(b, sah.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0xd0, 0xba, 0xd2, 0xaf, 0xd0, 0xbd, 0xd1, 0x8d, 0x27, 0x2c, 0x20}...)
	b = append(b, sah.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sah.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sah.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sah.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sah.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'sah_RU'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sah *sah_RU) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
