package gumi

const (
	STYLE_INVALID StyleData = iota
	//
	STYLE_Background StyleData = (iota << 16) | StyleData(TYPE_TEXTURE)
	//
	STYLE_LineWidth StyleData = (iota << 16) | StyleData(TYPE_INT)
	//
	STYLE_Font       StyleData = (iota << 16) | StyleData(TYPE_FONT)
	STYLE_TextColor  StyleData = (iota << 16) | StyleData(TYPE_COLOR)
	STYLE_TextHeight StyleData = (iota << 16) | StyleData(TYPE_INT)
	STYLE_TextAlign StyleData = (iota << 16) | StyleData(TYPE_ALIGN)

	//STYLE_TextOverflow StyleData = (iota << 16) | StyleData(TYPE_ALIGN)
)

var stylemap = map[string]StyleData{
	"background":  STYLE_Background,
	"line-width":  STYLE_LineWidth,

	"font":        STYLE_Font,
	"text-color":  STYLE_TextColor,
	"text-height": STYLE_TextHeight,
	"text-align" : STYLE_TextAlign,
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
