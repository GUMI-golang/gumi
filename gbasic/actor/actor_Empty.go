package actor

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
)

// Actor::Empty
//
// AEmpty exists only for the GUMI Tree as an element that does nothing
type AEmpty struct {
}


// GUMIFunction / GUMISize 			-> Define
func (AEmpty) GUMISize() gcore.Size {
	return gcore.Size{
		gcore.AUTOLENGTH,
		gcore.AUTOLENGTH,
	}
}


// fmt.Stringer / String					-> Define
func (s AEmpty) String() string {
	return fmt.Sprintf("%s", "AEmpty")
}

// Constructor
func AEmpty0() *AEmpty {
	return &AEmpty{}
}
