package gumi

import (
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/ChrisTrenkamp/goxpath"
)

// xpath - meta
var (
	xpath_meta_title       = goxpath.MustParse("/gumi/meta/title[last()]")
	xpath_meta_framerate   = goxpath.MustParse("/gumi/meta/framerate[last()]")
	xpath_meta_size        = goxpath.MustParse("/gumi/meta/size[last()]")
	xpath_meta_locale      = goxpath.MustParse("/gumi/meta/locale[last()]")
	xpath_meta_description = goxpath.MustParse("/gumi/meta/description[last()]")
)

func (s *MLGBuilder) parseMeta() error {
	var err error
	s.meta = new(Meta)
	s.meta.Title = xpath_meta_title.MustExec(s.nd).String()
	s.meta.Description = xpath_meta_description.MustExec(s.nd).String()
	s.meta.Locale = xpath_meta_locale.MustExec(s.nd).String()
	s.meta.size, err = gcore.UnmarshalFixedSize(xpath_meta_size.MustExec(s.nd).String())
	if err != nil {
		return err
	}
	s.meta.Framerate = gcore.ParseFramerate(xpath_meta_framerate.MustExec(s.nd).String())
	return nil
}