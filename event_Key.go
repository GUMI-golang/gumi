package gumi

type GUMIKey uint8

const (
	KEY_UNKNOWN GUMIKey = iota
	// Alphabet
	KEY_A GUMIKey = iota
	KEY_B GUMIKey = iota
	KEY_C GUMIKey = iota
	KEY_D GUMIKey = iota
	KEY_E GUMIKey = iota
	KEY_F GUMIKey = iota
	KEY_G GUMIKey = iota
	KEY_H GUMIKey = iota
	KEY_I GUMIKey = iota
	KEY_J GUMIKey = iota
	KEY_K GUMIKey = iota
	KEY_L GUMIKey = iota
	KEY_M GUMIKey = iota
	KEY_N GUMIKey = iota
	KEY_O GUMIKey = iota
	KEY_P GUMIKey = iota
	KEY_Q GUMIKey = iota
	KEY_R GUMIKey = iota
	KEY_S GUMIKey = iota
	KEY_T GUMIKey = iota
	KEY_U GUMIKey = iota
	KEY_V GUMIKey = iota
	KEY_W GUMIKey = iota
	KEY_X GUMIKey = iota
	KEY_Y GUMIKey = iota
	KEY_Z GUMIKey = iota
	// Number
	KEY_0 GUMIKey = iota
	KEY_1 GUMIKey = iota
	KEY_2 GUMIKey = iota
	KEY_3 GUMIKey = iota
	KEY_4 GUMIKey = iota
	KEY_5 GUMIKey = iota
	KEY_6 GUMIKey = iota
	KEY_7 GUMIKey = iota
	KEY_8 GUMIKey = iota
	KEY_9 GUMIKey = iota
	// Special
	KEY_BACKQUOTE    GUMIKey = iota // `
	KEY_APOSTROPHE   GUMIKey = iota // '
	KEY_LEFTBRACKET  GUMIKey = iota // [
	KEY_RIGHTBRACKET GUMIKey = iota // ]
	KEY_SEMICOLON    GUMIKey = iota // ;
	KEY_MINUS        GUMIKey = iota // -
	KEY_EQUAL        GUMIKey = iota // =
	KEY_BACKSLASH    GUMIKey = iota // \
	KEY_SLASH        GUMIKey = iota // /
	KEY_COMMA        GUMIKey = iota // ,
	KEY_PERIOD       GUMIKey = iota // .
	// Function
	KEY_F1  GUMIKey = iota
	KEY_F2  GUMIKey = iota
	KEY_F3  GUMIKey = iota
	KEY_F4  GUMIKey = iota
	KEY_F5  GUMIKey = iota
	KEY_F6  GUMIKey = iota
	KEY_F7  GUMIKey = iota
	KEY_F8  GUMIKey = iota
	KEY_F9  GUMIKey = iota
	KEY_F10 GUMIKey = iota
	KEY_F11 GUMIKey = iota
	KEY_F12 GUMIKey = iota
	// Control
	KEY_ESCAPE    GUMIKey = iota
	KEY_BACKSPACE GUMIKey = iota
	KEY_DELETE    GUMIKey = iota
	KEY_SPACE     GUMIKey = iota
	KEY_ENTER     GUMIKey = iota
	KEY_SHIFT     GUMIKey = iota // mac : Shift
	KEY_CONTROL   GUMIKey = iota // mac : Command, Control
	KEY_ALT       GUMIKey = iota // mac : Option
	KEY_TAB       GUMIKey = iota
	// Arrow
	KEY_UP    GUMIKey = iota
	KEY_DOWN  GUMIKey = iota
	KEY_LEFT  GUMIKey = iota
	KEY_RIGHT GUMIKey = iota
	// Mouse
	KEY_MOUSE1 GUMIKey = iota
	KEY_MOUSE2 GUMIKey = iota
	KEY_MOUSE3 GUMIKey = iota
)
