package gumi

import (
	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/verGo"
	"github.com/pkg/errors"
)

// xpath - document
var (
	xpath_document_gumi_version = goxpath.MustParse(`/processing-instruction("gumi-version")`)
)

func (s *MLGBuilder) parseDocument() error {
	s.document = new(Document)
	if strver := gcore.MustValue(xpath_document_gumi_version.ExecNode(s.nd)).(tree.NodeSet); len(strver) != 0 {
		if ver := verGo.ParseVersion(strver.String()); ver != verGo.Invalid {
			s.document.Version = ver
		} else {
			return errors.Wrap(ErrorParsingFail, "version parsing fail")
		}
	}
	return nil
}
