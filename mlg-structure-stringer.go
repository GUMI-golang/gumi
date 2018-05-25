package gumi

import (
	"fmt"
)

func (s *Document) String() string {
	return fmt.Sprintf("Document(version:%s)", s.Version.String())
}


func (s *Meta) String() string {
	return fmt.Sprintf("Meta(title:%s, description:%s, locale:%s, frameRate:%s, size:%s)", s.Title,s.Description, s.Locale,  s.Framerate, s.size.String())
}
