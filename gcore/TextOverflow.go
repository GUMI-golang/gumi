package gcore

// TODO
//
// If first is
// 		- 0x00, None
//			No arg
//			No work for this, just do
// 		- 0x01, AsPossible
//			No arg
//			Draw as much as possible character
// 		- 0x02, Ellipsis
//			No arg
//			Overflow text part change to '...'
// 		- 0x03, String
//			string argment
//			Work like ellipse, but using custom arguments instead '...'
// 		- 0x04, Linefeed
//			No arg
//			Anything that exceeds the range is treated as a line break.
type TextOverflow []byte

var (
	TextOverflowNone       = TextOverflow{0x00}
	TextOverflowAsPossible = TextOverflow{0x01}
	TextOverflowEllipsis   = TextOverflow{0x02}
	TextOverflowString     = TextOverflow{0x03}
	TextOverflowLinefeed   = TextOverflow{0x04}
)
