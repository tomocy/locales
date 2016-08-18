package ii_CN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ii_CN struct {
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

// New returns a new instance of translator for the 'ii_CN' locale
func New() locales.Translator {
	return &ii_CN{
		locale:                 "ii_CN",
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
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}, {0x37}, {0x38}, {0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xea, 0x8b, 0x8d, 0xea, 0x86, 0xaa}, {0xea, 0x91, 0x8d, 0xea, 0x86, 0xaa}, {0xea, 0x8c, 0x95, 0xea, 0x86, 0xaa}, {0xea, 0x87, 0x96, 0xea, 0x86, 0xaa}, {0xea, 0x89, 0xac, 0xea, 0x86, 0xaa}, {0xea, 0x83, 0x98, 0xea, 0x86, 0xaa}, {0xea, 0x8f, 0x83, 0xea, 0x86, 0xaa}, {0xea, 0x89, 0x86, 0xea, 0x86, 0xaa}, {0xea, 0x88, 0xac, 0xea, 0x86, 0xaa}, {0xea, 0x8a, 0xb0, 0xea, 0x86, 0xaa}, {0xea, 0x8a, 0xb0, 0xea, 0x8a, 0xaa, 0xea, 0x86, 0xaa}, {0xea, 0x8a, 0xb0, 0xea, 0x91, 0x8b, 0xea, 0x86, 0xaa}},
		daysAbbreviated:        [][]uint8{{0xea, 0x91, 0xad, 0xea, 0x86, 0x8f}, {0xea, 0x86, 0x8f, 0xea, 0x8b, 0x8d}, {0xea, 0x86, 0x8f, 0xea, 0x91, 0x8d}, {0xea, 0x86, 0x8f, 0xea, 0x8c, 0x95}, {0xea, 0x86, 0x8f, 0xea, 0x87, 0x96}, {0xea, 0x86, 0x8f, 0xea, 0x89, 0xac}, {0xea, 0x86, 0x8f, 0xea, 0x83, 0x98}},
		daysNarrow:             [][]uint8{{0xea, 0x86, 0x8f}, {0xea, 0x8b, 0x8d}, {0xea, 0x91, 0x8d}, {0xea, 0x8c, 0x95}, {0xea, 0x87, 0x96}, {0xea, 0x89, 0xac}, {0xea, 0x83, 0x98}},
		daysWide:               [][]uint8{{0xea, 0x91, 0xad, 0xea, 0x86, 0x8f, 0xea, 0x91, 0x8d}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x8b, 0x8d}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x91, 0x8d}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x8c, 0x95}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x87, 0x96}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x89, 0xac}, {0xea, 0x86, 0x8f, 0xea, 0x8a, 0x82, 0xea, 0x83, 0x98}},
		periodsAbbreviated:     [][]uint8{{0xea, 0x8e, 0xb8, 0xea, 0x84, 0x91}, {0xea, 0x81, 0xaf, 0xea, 0x8b, 0x92}},
		periodsWide:            [][]uint8{{0xea, 0x8e, 0xb8, 0xea, 0x84, 0x91}, {0xea, 0x81, 0xaf, 0xea, 0x8b, 0x92}},
		erasAbbreviated:        [][]uint8{{0xea, 0x83, 0x85, 0xea, 0x8b, 0x8a, 0xea, 0x82, 0xbf}, {0xea, 0x83, 0x85, 0xea, 0x8b, 0x8a, 0xea, 0x8a, 0x82}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"MDT": {0x4d, 0x44, 0x54}, "VET": {0x56, 0x45, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "IST": {0x49, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "MST": {0x4d, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CDT": {0x43, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CST": {0x43, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "AEST": {0x41, 0x45, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "BT": {0x42, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "COST": {0x43, 0x4f, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "JST": {0x4a, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "AST": {0x41, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "PDT": {0x50, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "GFT": {0x47, 0x46, 0x54}, "EAT": {0x45, 0x41, 0x54}, "ADT": {0x41, 0x44, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "GMT": {0x47, 0x4d, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "BOT": {0x42, 0x4f, 0x54}, "UYT": {0x55, 0x59, 0x54}, "ART": {0x41, 0x52, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "LHDT": {0x4c, 0x48, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (ii *ii_CN) Locale() string {
	return ii.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ii_CN'
func (ii *ii_CN) PluralsCardinal() []locales.PluralRule {
	return ii.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ii_CN'
func (ii *ii_CN) PluralsOrdinal() []locales.PluralRule {
	return ii.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ii_CN'
func (ii *ii_CN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ii_CN'
func (ii *ii_CN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ii_CN'
func (ii *ii_CN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ii_CN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ii_CN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ii *ii_CN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(ii.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ii.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(ii.minus) - 1; j >= 0; j-- {
			b = append(b, ii.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ii.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ii_CN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(ii.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ii.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyNegativePrefix[j])
		}

		for j := len(ii.minus) - 1; j >= 0; j-- {
			b = append(b, ii.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, ii.currencyNegativeSuffix...)
	} else {

		b = append(b, ii.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ii_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ii *ii_CN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	if btz, ok := ii.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
