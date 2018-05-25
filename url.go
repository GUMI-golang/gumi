package gumi

import (
	"bytes"
	"net/url"
	"strings"
	"github.com/GUMI-golang/gumi/gime"
	"mime"
	"path/filepath"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/giame"
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

	//gime.support uniform://
	gime.Support(func(u url.URL, hintmime *string) (gime.Value) {
		u.Path = strings.TrimPrefix(u.Path, "/")
		c, err := gcore.UnmarshalColor(u.Path)
		if err != nil {
			return err
		}
		return giame.NewUniformFiller(c)
	}, "uniform")


}
