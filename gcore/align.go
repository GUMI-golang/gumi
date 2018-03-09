package gcore

const (
	AlignTop        = 0x00
	AlignVertical   = 0x01
	AlignBottom     = 0x02
	AlignLeft       = 0x00
	AlignHorizontal = 0x10
	AlignRight      = 0x20
	//

	AlignCenter = AlignHorizontal | AlignVertical
)

type Align uint8

func ParseAlign(a Align) (v uint8, h uint8) {
	if a&AlignBottom == AlignBottom {
		v = AlignBottom
	} else if a&AlignVertical == AlignVertical {
		v = AlignVertical
	} else {
		v = AlignTop
	}
	if a&AlignRight == AlignRight {
		h = AlignRight
	} else if a&AlignHorizontal == AlignHorizontal {
		h = AlignHorizontal
	} else {
		h = AlignLeft
	}
	return

}
