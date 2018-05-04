package gumi

import (
	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"encoding/xml"
	"strings"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/gime"
	"fmt"
)

// xpath - resource
var (
	xpath_resource = goxpath.MustParse("/gumi/resource/")
)

// xpath-resource-values
var (
	xpath_name   = goxpath.MustParse("name()")
	xpath_id   = goxpath.MustParse("./@id")
	xpath_mime = goxpath.MustParse("./@mime")
	xpath_memory = goxpath.MustParse("./@memory")
)

func xmlnamestring(name xml.Name) string {
	if name.Space == "" {
		return name.Local
	}
	return name.Space + ":" + name.Local
}
func resourceAttr(elem tree.Elem) (id, mime string, memory bool) {
	id = xpath_id.MustExec(elem).String()
	mime = xpath_mime.MustExec(elem).String()
	memory, err := xpath_memory.ExecBool(elem)
	if err != nil {
		return
	}
	return
}


func (s *MLGBuilder) parseResource() error {
	s.resource = NewResource()
	resourceLoad(&s.resource.Inners, gcore.MustValue(xpath_resource.ExecNode(s.nd)).(tree.NodeSet)[0].(tree.Elem), "")
	return nil
}

func resourceLoad(res *[]MLGResourceStore, elem tree.Elem, path string) {
	name := xpath_name.MustExec(elem).String()
	switch name {
	case "uri":
		id, mime, memory := resourceAttr(elem)
		var symbol MLGResourceStore
		if memory{
			u := gime.URI(elem.ResValue())
			if u != nil {
				v := gime.Request(*u)
				if err := v.(error);err != nil {
					panic(err)
				}
				symbol = &MLGResourceMemory{
					id:   strings.TrimPrefix(path, "resource") + id,
					mime: mime,
					data: v,
				}
			}else{
				fmt.Println("Can't parse ", elem.ResValue())
			}
		}else{
			symbol = &MLGResourceLocation{
				id:   strings.TrimPrefix(path, "resource") + id,
				mime: mime,
				uri: elem.ResValue(),
			}
		}
		if symbol != nil{
			*res = append(*res, symbol)
		}
		return
	case "embed":
		id, mime, _ := resourceAttr(elem)
		symbol := &MLGResourceMemory{
			id:   strings.TrimPrefix(path, "resource") + id,
			mime: mime,
			data: elem.ResValue(),
		}
		*res = append(*res, symbol)
		return
	default:
		path += name + "/"
	}
	for _, v := range elem.GetChildren() {
		if v2, ok := v.(tree.Elem); ok {
			resourceLoad(res, v2, path)
		}
	}
}