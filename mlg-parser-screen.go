package gumi

import (
	"encoding/xml"
	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/pkg/errors"
	"strconv"
	"github.com/GUMI-golang/gumi/gime"
	"strings"
	"regexp"
	"log"
)

// xpath - screen
var (
	xpath_screen = goxpath.MustParse("/gumi/screen/*[1]")
	xpath_replace_resource = goxpath.MustParse("./processing-instruction()")
)
var re_XMLProcInst = regexp.MustCompile(`^@(?P<URI>.*)$`)
func (s *MLGBuilder) parseScreen() error {
	s.screen = NewScreen()
	s.screen.Meta = s.meta
	s.screen.Meta.scr = s.screen
	s.screen.Resource = s.resource

	//
	if temp, err := xpath_screen.ExecNode(s.nd); err == nil {
		err := screenLoad(s.screen, nil, s.space, temp[0].(tree.Elem))
		if err != nil {
			return errors.Wrap(ErrorParsingFail, err.Error())
		}
	} else {
		return errors.Wrap(ErrorParsingFail, err.Error())
	}
	return nil
}
func screenLoad(s *Screen, p *Pipe, space *Space, e tree.Elem) error {
	if v, ok := e.GetToken().(xml.StartElement); ok {
		tag := space.Find(v.Name.Space, v.Name.Local)
		if tag == nil {
			return errors.Wrap(ErrorParsingFail, xmlnamestring(v.Name)+" is not support")
		}
		curr := s.Pipeline.New(p, tag.New())
		// direct Style Setup
		for _, v := range e.GetAttrs() {
			tk := v.GetToken().(xml.Attr)
			if tk.Name.Space == "style"{
				// direct style modify
				if sd := FromStringStyleData(tk.Name.Local);sd != STYLE_INVALID{
					if temp := re_XMLProcInst.FindStringSubmatch(strings.TrimSpace(tk.Value)); len(temp) > 1{
						u := gime.URI(temp[1])
						if u != nil{
							v := s.Resource.Request(*u)
							if err, ok := v.(error); ok && err != nil {
								// TODO
								panic(err)
							}
							ok := curr.SetStyle(sd, v)
							if !ok{
								panic("Not ok")
							}
						}else {
							// TODO
							log.Println(temp[1], "can't parse as gime")
						}
					}else {
						v, err := sd.Type().Unmarshal(tk.Value)
						if err != nil {
							panic(err)
						}
						ok := curr.SetStyle(sd, v)
						if !ok{
							panic("Not ok")
						}
					}
				}else {
					// TODO
					panic(tk.Name.Local)
				}

			}
		}
		if vm, ok := curr.Elem.(ValueManager); ok {
			for _, v := range e.GetAttrs() {
				tk := v.GetToken().(xml.Attr)
				if tk.Name.Space != "style"{
					if temp := re_XMLProcInst.FindStringSubmatch(strings.TrimSpace(tk.Value)); len(temp) > 1{
						u := gime.URI(temp[1])
						if u != nil{
							v := s.Resource.Request(*u)
							if err, ok := v.(error); ok && err != nil {
								// TODO
								panic(err)
							}
							vm.SetValue(xmlnamestring(tk.Name), v)
						}else {
							// TODO
							log.Println(temp[1], "can't parse as gime")
						}
					}else {
						vm.SetValue(xmlnamestring(tk.Name), tk.Value)
					}
				}
			}
		}
		if dm, ok := curr.Elem.(DefaultManager); ok {
			res, err := xpath_replace_resource.ExecNode(e)
			if err == nil {
				if len(res) > 0{
					u := gime.URI(xpath_name.MustExec(res[0]).String() + res[0].ResValue())
					if u != nil{
						v := s.Resource.Request(*u)
						if err, ok := v.(error); ok && err != nil {
							// TODO
							panic(err)
						}
						dm.SetDefault(v)
					}else {
						// TODO
						log.Println(xpath_name.MustExec(res[0]).String() + res[0].ResValue(), "can't parse as gime")
					}

				}else {
					dm.SetDefault(e.ResValue())
				}
			}
		}
		//
		for _, n := range e.GetChildren() {

			if nche, ok := n.(tree.Elem); ok {
				err := screenLoad(s, curr, space, nche)
				if err != nil {
					return err
				}
			}

		}
		if len(curr.Childrun) > curr.Elem.MaxChildrun() {
			return errors.Wrap(ErrorParsingFail,
				xmlnamestring(v.Name)+" got too many childrun, can hold "+
					strconv.Itoa(curr.Elem.MaxChildrun())+", got "+
					strconv.Itoa(len(curr.Childrun)))
		}
	}
	return nil
}
