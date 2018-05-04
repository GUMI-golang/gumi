package gcore

import (
	"time"
	"fmt"
)

type Framerate int

func (s Framerate) Duration() time.Duration {
	return time.Second / time.Duration(s)
}
func (s Framerate) String() string{
	return fmt.Sprintf("%d hz", s)
}
func ParseFramerate(f string) Framerate {
	var fr int
	_, err := fmt.Sscanf(f,"%d hz", &fr)
	if err != nil {
		fmt.Sscanf(f,"%d", &fr)
	}
	return Framerate(fr)
}