package gumi

func (s GUMIKey) String() string {
	switch s {
	case KEY_UNKNOWN:
		return "KEY : UNKNOWN"
	// Alphabet
	case KEY_A:
		return "KEY : A"
	case KEY_B:
		return "KEY : B"
	case KEY_C:
		return "KEY : C"
	case KEY_D:
		return "KEY : D"
	case KEY_E:
		return "KEY : E"
	case KEY_F:
		return "KEY : F"
	case KEY_G:
		return "KEY : G"
	case KEY_H:
		return "KEY : H"
	case KEY_I:
		return "KEY : I"
	case KEY_J:
		return "KEY : J"
	case KEY_K:
		return "KEY : K"
	case KEY_L:
		return "KEY : L"

	case KEY_M:
		return "KEY : M"
	case KEY_N:
		return "KEY : N"
	case KEY_O:
		return "KEY : O"
	case KEY_P:
		return "KEY : P"
	case KEY_Q:
		return "KEY : Q"
	case KEY_R:
		return "KEY : R"
	case KEY_S:
		return "KEY : S"
	case KEY_T:
		return "KEY : T"
	case KEY_U:
		return "KEY : U"
	case KEY_V:
		return "KEY : V"
	case KEY_W:
		return "KEY : W"
	case KEY_X:
		return "KEY : X"
	case KEY_Y:
		return "KEY : Y"
	case KEY_Z:
		return "KEY : Z"
	// Number
	case KEY_0:
		return "KEY : 0"
	case KEY_1:
		return "KEY : 1"
	case KEY_2:
		return "KEY : 2"
	case KEY_3:
		return "KEY : 3"
	case KEY_4:
		return "KEY : 4"
	case KEY_5:
		return "KEY : 5"
	case KEY_6:
		return "KEY : 6"
	case KEY_7:
		return "KEY : 7"
	case KEY_8:
		return "KEY : 8"
	case KEY_9:
		return "KEY : 9"
	// Special
	case KEY_BACKQUOTE: // `
		return "KEY : BACKQUOTE"
	case KEY_APOSTROPHE: // '
		return "KEY : APOSTROPHE"
	case KEY_LEFTBRACKET: // [
		return "KEY : LEFTBRACKET"
	case KEY_RIGHTBRACKET: // ]
		return "KEY : RIGHTBRACKET"
	case KEY_SEMICOLON: // ;
		return "KEY : SEMICOLON"
	case KEY_MINUS: // -
		return "KEY : MINUS"
	case KEY_EQUAL: // =
		return "KEY : EQUAL"
	case KEY_BACKSLASH: // \
		return "KEY : BACKSLASH"
	case KEY_SLASH: // /
		return "KEY : SLASH"
	case KEY_COMMA: // ,
		return "KEY : COMMA"
	case KEY_PERIOD: // .
		return "KEY : PERIOD"
	// Function
	case KEY_F1:
		return "KEY : F1"
	case KEY_F2:
		return "KEY : F2"
	case KEY_F3:
		return "KEY : F3"
	case KEY_F4:
		return "KEY : F4"
	case KEY_F5:
		return "KEY : F5"
	case KEY_F6:
		return "KEY : F6"
	case KEY_F7:
		return "KEY : F7"
	case KEY_F8:
		return "KEY : F8"
	case KEY_F9:
		return "KEY : F9"
	case KEY_F10:
		return "KEY : F10"
	case KEY_F11:
		return "KEY : F11"
	case KEY_F12:
		return "KEY : F12"
	// Control
	case KEY_ESCAPE:
		return "KEY : ESCAPE"
	case KEY_BACKSPACE:
		return "KEY : BACKSPACE"
	case KEY_DELETE:
		return "KEY : DELETE"
	case KEY_SPACE:
		return "KEY : SPACE"
	case KEY_ENTER:
		return "KEY : ENTER"
	case KEY_SHIFT: // mac : Shift
		return "KEY : SHIFT"
	case KEY_CONTROL: // mac : Command, Control
		return "KEY : CONTROL"
	case KEY_ALT: // mac : Option
		return "KEY : ALT"
	case KEY_TAB:
		return "KEY : TAB"
	// Arrow
	case KEY_UP:
		return "KEY : UP"
	case KEY_DOWN:
		return "KEY : DOWN"
	case KEY_LEFT:
		return "KEY : LEFT"
	case KEY_RIGHT:
		return "KEY : RIGHT"
	// Mouse
	case KEY_MOUSE1:
		return "KEY : MOUSE1"
	case KEY_MOUSE2:
		return "KEY : MOUSE2"
	case KEY_MOUSE3:
		return "KEY : MOUSE3"
	}
	return "UNDEFINE"
}
