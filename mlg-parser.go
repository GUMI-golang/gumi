package gumi

import (
	"encoding/xml"
	"fmt"
	"strings"
	"io"
)

type Parser struct {
	// decoder
	dec *xml.Decoder
	// setting data
	PrintComment bool
	Verbose      bool
	// share data
	hieracy []mlgXMLTag
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		dec: xml.NewDecoder(r),
	}
}

func (s *Parser)clear()  {
	s.hieracy = nil
}
func (s *Parser)paddingPrintf(format string, args ...interface{})  {
	var padding = ""
	padcount := len(s.hieracy) - 1
	if padcount > 0{
		padding = strings.Repeat("    ", padcount)
	}
	res := fmt.Sprintf(format, args...)
	spl := strings.Split(res, "\n")
	for _, s := range spl{
		if len(s) == 0{
			continue
		}
		fmt.Println(padding + s)
	}
}
func (s *Parser) Step() error {
	tk, err := s.dec.Token()
	if err != nil{
		return err
	}
	switch t := tk.(type) {
	case xml.ProcInst:
		fmt.Println("Instruction :", t.Target, string(t.Inst))
	case xml.Directive:
		fmt.Println("Directive   :",string(t))
	case xml.Comment:
		if s.PrintComment || s.Verbose{
			s.paddingPrintf("# %s\n", string(t))
		}
	case xml.CharData:
	case xml.StartElement:
		mlgtagname := mlgXMLTag{
			namespace: t.Name.Space,
			name:      t.Name.Local,
		}
		s.hieracy = append(s.hieracy, mlgtagname)

		//
		if s.Verbose{
			s.paddingPrintf("%s\n", mlgtagname)
		}
	case xml.EndElement:
		s.hieracy = s.hieracy[:len(s.hieracy)-1]
	}
	return nil
}
func (s *Parser ) Parse() error {
	return nil
}
