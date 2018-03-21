package mlgumi

import (
	"github.com/GUMI-golang/gumi"
)

type (
	XMLGUMI struct {
		Meta *XMLMeta
		Element *XMLElement
	}
	XMLMeta struct {
		Title string `xml:"title"`
		Description string `xml:"description"`
		Width int `xml:"width"`
		Height int `xml:"height"`
	}
	XMLElement struct {
		Root *XMLGUMIElem
	}
	XMLGUMIElem struct {
		Parent *XMLGUMIElem
		Childrun *XMLGUMIElem
		//
		Name string
		Args map[string]interface{}
		Build gumi.GUMI
	}
)

