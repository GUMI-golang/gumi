package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"image/color"
	"strconv"
	"github.com/GUMI-golang/giame"
	"encoding/json"
)

type StyleDatatype uint16
const (
	TYPE_          StyleDatatype = iota
	TYPE_INT       StyleDatatype = iota
	TYPE_FLOAT     StyleDatatype = iota
	TYPE_STRING    StyleDatatype = iota
	TYPE_SIZE      StyleDatatype = iota
	TYPE_FIXEDSIZE StyleDatatype = iota
	TYPE_BLANK     StyleDatatype = iota
	TYPE_ALIGN     StyleDatatype = iota
	TYPE_AXIS      StyleDatatype = iota
	TYPE_COLOR     StyleDatatype = iota
	TYPE_FONT      StyleDatatype = iota
	TYPE_TEXTURE   StyleDatatype = iota
)

func (s StyleDatatype) Marshal(value interface{}) (string, error) {
	switch s {
	case TYPE_INT:
		switch value.(type) {
		case int:
		case uint:
		case int8:
		case int16:
		case int32:
		case int64:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		default:
			return "", ErrorInvalidValue
		}
		return fmt.Sprint(value), nil
	case TYPE_FLOAT:
		switch value.(type) {
		case float32:
		case float64:
		default:
			return "", ErrorInvalidValue
		}
		return fmt.Sprint(value), nil
	case TYPE_STRING:
		if v, ok := value.(string);ok{
			return v, nil
		}
		return "", ErrorInvalidValue
	case TYPE_SIZE:
		if v, ok := value.(gcore.Size);ok{
			return gcore.MarshalSize(v), nil
		}
		return "", ErrorInvalidValue
	case TYPE_FIXEDSIZE:
		if v, ok := value.(gcore.FixedSize);ok{
			return gcore.MarshalFixedSize(v), nil
		}
		return "", ErrorInvalidValue
	case TYPE_BLANK:
		if v, ok := value.(gcore.Blank);ok{
			return gcore.MarshalBlank(v), nil
		}
		return "", ErrorInvalidValue
	case TYPE_ALIGN:
		if v, ok := value.(gcore.Align);ok{
			return gcore.MarshalAlign(v), nil
		}
		return "", ErrorInvalidValue
	case TYPE_AXIS:
		// TODO
	case TYPE_COLOR:
		if v, ok := value.(color.Color);ok{
			return gcore.MarshalColor(v), nil
		}
		return "", ErrorInvalidValue
	case TYPE_FONT:
		// TODO
	case TYPE_TEXTURE:
		if f, ok := value.(giame.Filler); ok{
			serialed := f.Serial()
			bts, err := json.Marshal(serialed)
			return string(bts), err
		}
		return "", ErrorInvalidValue
	}
	panic(CriticalUnknownBehavior)
}
func (s StyleDatatype) Unmarshal(value string) (interface{}, error) {
	switch s {
	case TYPE_INT:
		return strconv.ParseInt(value, 10, 64)
	case TYPE_FLOAT:
		return strconv.ParseFloat(value, 64)
	case TYPE_STRING:
		return value, nil
	case TYPE_SIZE:
		return gcore.UnmarshalSize(value)
	case TYPE_FIXEDSIZE:
		return gcore.UnmarshalFixedSize(value)
	case TYPE_BLANK:
		return gcore.UnmarshalBlank(value)
	case TYPE_ALIGN:
		return gcore.UnmarshalAlign(value)
	case TYPE_AXIS:
		// TODO
	case TYPE_COLOR:
		return gcore.UnmarshalColor(value)
	case TYPE_FONT:
		// TODO
	case TYPE_TEXTURE:
		var raw *giame.RawFiller
		err := json.Unmarshal([]byte(value), raw)
		if err != nil {
			return nil, err
		}
		return raw.Restore(), nil
	}
	panic(CriticalUnknownBehavior)
}
func (s StyleDatatype) Valid(value interface{}) bool {
	switch s {
	case TYPE_INT:
		switch value.(type) {
		case int:
		case uint:
		case int8:
		case int16:
		case int32:
		case int64:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		default:
			return false
		}
		return true
	case TYPE_FLOAT:
		switch value.(type) {
		case float32:
		case float64:
		default:
			return false
		}
		return true
	case TYPE_STRING:
		if _, ok := value.(string);ok{
			return true
		}
		return false
	case TYPE_SIZE:
		if _, ok := value.(gcore.Size);ok{
			return true
		}
		return false
	case TYPE_FIXEDSIZE:
		if _, ok := value.(gcore.FixedSize);ok{
			return true
		}
		return false
	case TYPE_BLANK:
		if _, ok := value.(gcore.Blank);ok{
			return true
		}
		return false
	case TYPE_ALIGN:
		if _, ok := value.(gcore.Align);ok{
			return true
		}
		return false
	case TYPE_AXIS:
		// TODO
	case TYPE_COLOR:
		if _, ok := value.(color.Color);ok{
			return true
		}
		return false
	case TYPE_FONT:
		// TODO
	case TYPE_TEXTURE:
		if _, ok := value.(giame.Filler);ok{
			return true
		}
		return false
	}
	panic(CriticalUnknownBehavior)
}
