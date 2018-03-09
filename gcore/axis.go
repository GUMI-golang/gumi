package gcore

type Axis uint8

func (s Axis) String() string {
	switch s {
	default:
		return "Unknown"
	case AxisHorizontal:
		return "Horizontal"
	case AxisVertical:
		return "Vertical"
	case AxisBoth:
		return "Both"
	}
}
const (
	AxisVertical   Axis = 1 << iota
	AxisHorizontal Axis = 1 << iota
	AxisBoth            = AxisVertical | AxisHorizontal
)
