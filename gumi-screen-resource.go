package gumi

import (
	"github.com/GUMI-golang/gumi/gime"
	"net/url"
	"github.com/pkg/errors"
)

type Resource struct {
	*gime.URISupportMap
	Inners []MLGResourceStore
}

func NewResource() *Resource {
	temp := gime.URIMapWithDefault()
	res := &Resource{
		URISupportMap:temp,
		Inners:nil,
	}
	res.Support(func(u url.URL, hintmime *string) (gime.Value) {
		found := res.Resource(u.Path)
		if found == nil{
			return errors.Wrap(ErrorNotFound, u.String())
		}
		return found.Data()
	}, "resource")
	return res
}
func (s *Resource) Resource(path string) MLGResourceStore {
	for _, v := range s.Inners {
		if v.ID() == path {
			return v
		}
	}
	return nil
}