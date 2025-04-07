package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	_ "unsafe"
)

type X_xsltCompMatch struct {
	Unused [8]uint8
}

type X_xsltNumberData struct {
	Level                *libxml2.Char
	Count                *libxml2.Char
	From                 *libxml2.Char
	Value                *libxml2.Char
	Format               *libxml2.Char
	HasFormat            c.Int
	DigitsPerGroup       c.Int
	GroupingCharacter    c.Int
	GroupingCharacterLen c.Int
	Doc                  libxml2.DocPtr
	Node                 libxml2.NodePtr
	CountPat             *X_xsltCompMatch
	FromPat              *X_xsltCompMatch
}
type NumberData X_xsltNumberData
type NumberDataPtr *NumberData

type X_xsltFormatNumberInfo struct {
	IntegerHash       c.Int
	IntegerDigits     c.Int
	FracDigits        c.Int
	FracHash          c.Int
	Group             c.Int
	Multiplier        c.Int
	AddDecimal        int8
	IsMultiplierSet   int8
	IsNegativePattern int8
}
type FormatNumberInfo X_xsltFormatNumberInfo
type FormatNumberInfoPtr *FormatNumberInfo
