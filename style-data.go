package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"image/color"
	"strconv"
)

type StyleData uint32
type StyleDataName uint16
type StyleDataType uint16

func makeStyleValue(name StyleDataName, valueType StyleDataType) StyleData {
	return StyleData((uint32(name) << 16) | uint32(valueType))
}
func (s StyleData) Type() StyleDataType {
	return StyleDataType(s)
}

//
func (s StyleDataType) Marshal(value interface{}) (string, error) {
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
	}
	panic(CriticalUnknownBehavior)
}
func (s StyleDataType) Unmarshal(value string) (interface{}, error) {
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
	}
	panic(CriticalUnknownBehavior)
}
func (s StyleDataType) Valid(value interface{}) bool {
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
	}
	panic(CriticalUnknownBehavior)
}
