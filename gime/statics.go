package gime

import (
	"net/url"
	"github.com/GUMI-golang/gumi/gcore"
)

func MustURI(uri string) url.URL {
	return *gcore.MustValue(url.ParseRequestURI(uri)).(*url.URL)
}
func URI(uri string) *url.URL {
	d, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil
	}
	return d
}
func Request(u url.URL, hints ... string) (Value) {
	return defaultMapping.Request(u, hints...)
}
func Support(handleFn func(u url.URL, hintmime *string) (Value), schema ... string) (err error) {
	return defaultMapping.Support(handleFn, schema...)
}
func URIMapWithDefault() *URISupportMap {
	return defaultMapping.Copy()
}