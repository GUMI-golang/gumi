package gumi

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

func Request(path string) (data io.Reader, err error) {
	u, err := url.Parse(path)
	if err != nil {
		return data, err
	}
	//
	switch u.Scheme {
	case "file":
		if runtime.GOOS == "windows" {
			u.Path = strings.TrimPrefix(u.Path, "/")
		}
		f, err := os.Open(u.Path)
		if err != nil {
			return data, err
		}
		defer f.Close()
		btd, err := ioutil.ReadAll(f)
		if err != nil {
			return data, err
		}
		return bytes.NewReader(btd), nil
	case "http":
		fallthrough
	case "https":
		res, err := http.Get(path)
		if err != nil {
			return data, err
		}
		btd, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return data, err
		}
		return bytes.NewReader(btd), err
	case "gumi":
		u.Path = strings.TrimPrefix(u.Path, "/")
		dat, err := Asset(u.Path)
		if err != nil {
			return data, err
		}
		return bytes.NewReader(dat), nil
	default:
		return nil, errors.WithMessage(ErrorRequestSchema, u.Scheme+" is unknown")
	}
}
