package lkt

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type lkt struct {
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
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
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

// New returns a new instance of translator for the 'lkt' locale
func New() locales.Translator {
	return &lkt{
		locale:                 "lkt",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0x4b},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0x4b},
		monthsWide:             [][]uint8{[]uint8(nil), {0x57, 0x69, 0xc3, 0xb3, 0x74, 0x68, 0x65, 0xc8, 0x9f, 0x69, 0x6b, 0x61, 0x20, 0x57, 0xc3, 0xad}, {0x54, 0x68, 0x69, 0x79, 0xc3, 0xb3, 0xc8, 0x9f, 0x65, 0x79, 0x75, 0xc5, 0x8b, 0x6b, 0x61, 0x20, 0x57, 0xc3, 0xad}, {0x49, 0xc5, 0xa1, 0x74, 0xc3, 0xa1, 0x77, 0x69, 0xc4, 0x8d, 0x68, 0x61, 0x79, 0x61, 0x7a, 0x61, 0xc5, 0x8b, 0x20, 0x57, 0xc3, 0xad}, {0x50, 0xc8, 0x9f, 0x65, 0xc5, 0xbe, 0xc3, 0xad, 0x74, 0xc8, 0x9f, 0x6f, 0x20, 0x57, 0xc3, 0xad}, {0xc4, 0x8c, 0x68, 0x61, 0xc5, 0x8b, 0x77, 0xc3, 0xa1, 0x70, 0x65, 0x74, 0xc8, 0x9f, 0x6f, 0x20, 0x57, 0xc3, 0xad}, {0x57, 0xc3, 0xad, 0x70, 0x61, 0x7a, 0x75, 0x6b, 0xc8, 0x9f, 0x61, 0x2d, 0x77, 0x61, 0xc5, 0xa1, 0x74, 0xc3, 0xa9, 0x20, 0x57, 0xc3, 0xad}, {0xc4, 0x8c, 0x68, 0x61, 0xc5, 0x8b, 0x70, 0xc8, 0x9f, 0xc3, 0xa1, 0x73, 0x61, 0x70, 0x61, 0x20, 0x57, 0xc3, 0xad}, {0x57, 0x61, 0x73, 0xc3, 0xba, 0x74, 0xc8, 0x9f, 0x75, 0xc5, 0x8b, 0x20, 0x57, 0xc3, 0xad}, {0xc4, 0x8c, 0x68, 0x61, 0xc5, 0x8b, 0x77, 0xc3, 0xa1, 0x70, 0x65, 0xc7, 0xa7, 0x69, 0x20, 0x57, 0xc3, 0xad}, {0xc4, 0x8c, 0x68, 0x61, 0xc5, 0x8b, 0x77, 0xc3, 0xa1, 0x70, 0x65, 0x2d, 0x6b, 0x61, 0x73, 0x6e, 0xc3, 0xa1, 0x20, 0x57, 0xc3, 0xad}, {0x57, 0x61, 0x6e, 0xc3, 0xad, 0x79, 0x65, 0x74, 0x75, 0x20, 0x57, 0xc3, 0xad}, {0x54, 0xc8, 0x9f, 0x61, 0x68, 0xc3, 0xa9, 0x6b, 0x61, 0x70, 0xc5, 0xa1, 0x75, 0xc5, 0x8b, 0x20, 0x57, 0xc3, 0xad}},
		daysNarrow:             [][]uint8{{0x41}, {0x57}, {0x4e}, {0x59}, {0x54}, {0x5a}, {0x4f}},
		daysWide:               [][]uint8{{0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x77, 0x61, 0x6b, 0xc8, 0x9f, 0x61, 0xc5, 0x8b}, {0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x77, 0x61, 0xc5, 0x8b, 0xc5, 0xbe, 0x69}, {0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x6e, 0x75, 0xc5, 0x8b, 0x70, 0x61}, {0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x79, 0x61, 0x6d, 0x6e, 0x69}, {0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x74, 0x6f, 0x70, 0x61}, {0x41, 0xc5, 0x8b, 0x70, 0xc3, 0xa9, 0x74, 0x75, 0x7a, 0x61, 0x70, 0x74, 0x61, 0xc5, 0x8b}, {0x4f, 0x77, 0xc3, 0xa1, 0xc5, 0x8b, 0x67, 0x79, 0x75, 0xc5, 0xbe, 0x61, 0xc5, 0xbe, 0x61, 0x70, 0x69}},
		timezones:              map[string][]uint8{"MDT": {0x4d, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "IST": {0x49, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "ECT": {0x45, 0x43, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "WIT": {0x57, 0x49, 0x54}, "PDT": {0x50, 0x44, 0x54}, "JST": {0x4a, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "PST": {0x50, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CDT": {0x43, 0x44, 0x54}, "MST": {0x4d, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "EST": {0x45, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "ART": {0x41, 0x52, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "BT": {0x42, 0x54}, "CST": {0x43, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "GFT": {0x47, 0x46, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "AWST": {0x41, 0x57, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}},
	}
}

// Locale returns the current translators string locale
func (lkt *lkt) Locale() string {
	return lkt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lkt'
func (lkt *lkt) PluralsCardinal() []locales.PluralRule {
	return lkt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lkt'
func (lkt *lkt) PluralsOrdinal() []locales.PluralRule {
	return lkt.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lkt'
func (lkt *lkt) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lkt'
func (lkt *lkt) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lkt'
func (lkt *lkt) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lkt *lkt) MonthAbbreviated(month time.Month) []byte {
	return lkt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lkt *lkt) MonthsAbbreviated() [][]byte {
	return lkt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lkt *lkt) MonthNarrow(month time.Month) []byte {
	return lkt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lkt *lkt) MonthsNarrow() [][]byte {
	return lkt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lkt *lkt) MonthWide(month time.Month) []byte {
	return lkt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lkt *lkt) MonthsWide() [][]byte {
	return lkt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lkt *lkt) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return lkt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lkt *lkt) WeekdaysAbbreviated() [][]byte {
	return lkt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lkt *lkt) WeekdayNarrow(weekday time.Weekday) []byte {
	return lkt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lkt *lkt) WeekdaysNarrow() [][]byte {
	return lkt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lkt *lkt) WeekdayShort(weekday time.Weekday) []byte {
	return lkt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lkt *lkt) WeekdaysShort() [][]byte {
	return lkt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lkt *lkt) WeekdayWide(weekday time.Weekday) []byte {
	return lkt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lkt *lkt) WeekdaysWide() [][]byte {
	return lkt.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lkt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lkt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lkt *lkt) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lkt.currencies[currency]
	l := len(s) + len(lkt.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lkt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(lkt.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, lkt.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, lkt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lkt.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lkt'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lkt.currencies[currency]
	l := len(s) + len(lkt.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lkt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(lkt.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, lkt.currencyNegativePrefix[j])
		}

		b = append(b, lkt.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(lkt.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, lkt.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, lkt.currencyNegativeSuffix...)
	} else {

		b = append(b, lkt.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, lkt.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, lkt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, lkt.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, lkt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, lkt.periodsAbbreviated[0]...)
	} else {
		b = append(b, lkt.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, lkt.periodsAbbreviated[0]...)
	} else {
		b = append(b, lkt.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, lkt.periodsAbbreviated[0]...)
	} else {
		b = append(b, lkt.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'lkt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lkt *lkt) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lkt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, lkt.periodsAbbreviated[0]...)
	} else {
		b = append(b, lkt.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lkt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
