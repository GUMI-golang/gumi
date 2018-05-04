package gime

import (
	"runtime"
	"strings"
	"os"
	"net/url"
	"mime"
	"path/filepath"
	"net/http"
)
var defaultMapping = NewURISupportMap()

func init() {
	defaultMapping.Support(func(u url.URL, hintmime *string) (Value) {
		if runtime.GOOS == "windows" {
			u.Path = strings.TrimPrefix(u.Path, "/")
		}
		f, err := os.Open(u.Path)
		if err != nil {
			return  err
		}
		defer f.Close()
		if hintmime == nil{
			temp := mime.TypeByExtension(filepath.Ext(u.Path))
			hintmime = &temp
		}
		return Parse(*hintmime, f)
	}, "file")
	defaultMapping.Support(func(u url.URL, hintmime *string) (Value) {
		res, err := http.Get(u.String())
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if hintmime == nil{
			temp := res.Header.Get("Content-Type")
			hintmime = &temp
		}
		return Parse(*hintmime, res.Body)
	}, "http", "https")
	defaultMapping.Support(func(u url.URL, hintmime *string) (Value) {
		return u.Host
	}, "echo")
}