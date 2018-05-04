package gumi

import "github.com/pkg/errors"

const (
	STYLE_INVALID    = iota
	STYLE_LineWidth  = (iota << 16) | StyleData(TYPE_INT)
	// TODO
	STYLE_Font       = (iota << 16) | StyleData(TYPE_FONT)
	STYLE_TextColor  = (iota << 16) | StyleData(TYPE_COLOR)
	STYLE_TextHeight = (iota << 16) | StyleData(TYPE_FLOAT)
)

var stylemap = map[string]StyleData{
	"line-width":  STYLE_LineWidth,
	"font":        STYLE_Font,
	"text-color":  STYLE_TextColor,
	"text-height": STYLE_TextHeight,
}

func FromStringStyleData(s string) StyleData {
	if v, ok := stylemap[s]; ok {
		return v
	}
	return STYLE_INVALID
}
func ListStringStyleData() []string {
	temp := []string{}
	for k := range stylemap {
		temp = append(temp, k)
	}
	return temp
}
func (s StyleData) String() string {
	for k, v := range stylemap {
		if s == v {
			return k
		}
	}
	panic(errors.New("Unknown Style"))
}

const (
	TYPE_          StyleDataType = iota
	TYPE_INT       StyleDataType = iota
	TYPE_FLOAT     StyleDataType = iota
	TYPE_STRING    StyleDataType = iota
	TYPE_SIZE      StyleDataType = iota
	TYPE_FIXEDSIZE StyleDataType = iota
	TYPE_BLANK     StyleDataType = iota
	TYPE_ALIGN     StyleDataType = iota
	TYPE_AXIS      StyleDataType = iota
	TYPE_COLOR     StyleDataType = iota
	TYPE_FONT      StyleDataType = iota
)
