package gumi

import (
	"bytes"
	"net/url"
	"strings"
	"github.com/GUMI-golang/gumi/gime"
	"mime"
	"path/filepath"
)

func init()  {
	//gime.support gumi://
	gime.Support(func(u url.URL, hintmime *string) (gime.Value) {
		u.Path = strings.TrimPrefix(u.Path, "/")
		dat, err := Asset(u.Path)
		if err != nil {
			return err
		}
		if hintmime == nil{
			temp := mime.TypeByExtension(filepath.Ext(u.Path))
			hintmime = &temp
		}
		return gime.Parse(*hintmime, bytes.NewReader(dat))
	}, "gumi")
}
