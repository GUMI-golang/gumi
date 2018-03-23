package gumi

type (
	XMLGUMI struct {
		Meta *XMLMeta
		Screen *XMLScreen
	}
	XMLMeta struct {
		Title string `xml:"title"`
		Description string `xml:"description"`
		Width int `xml:"width"`
		Height int `xml:"height"`
	}
	XMLResources struct {
		Resources map[string]string `xml:"resources"`
		Upload bool `xml:"upload"`
	}
	XMLScreen struct {
		Root *XMLGUMIElem
	}
	XMLGUMIElem struct {
		Parent *XMLGUMIElem
		Childrun *XMLGUMIElem
		//
		Name string
		Args map[string]interface{}
		//Build gumi.GUMI
	}
)

