// gumi.go define GUMI interface
// GUMI is elements of GUI
// Every elements can render(graphic), affect(event), update(information) must implements this interface
package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
)

// GUMI is a collection of basic elements
type GUMI interface {
	GUMIFunction
	fmt.Stringer
}

// GUMI Root is special case of GUMI
// GUMI Root help to find Screen
// Mostly root locate root position on GUMI Tree
// But it is not necessary
type GUMIRoot interface {
	GUMI
	Screen() *Screen
}

type GUMIFunction interface {
	GUMISize() gcore.Size
}
