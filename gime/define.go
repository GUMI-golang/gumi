package gime

import (
	"github.com/pkg/errors"
	"net/url"
)

type URISupportMap struct {
	mapping map[string]func(u url.URL, hintmime *string) (Value)
}

func NewURISupportMap() *URISupportMap {
	return &URISupportMap{
		mapping: make(map[string]func(u url.URL, hintmime *string) (Value)),
	}
}
func (s *URISupportMap) Support(handleFn func(u url.URL, hintmime *string) (Value), schema ...string) error {
	for _, temp := range schema {
		if _, ok := s.mapping[temp]; ok {
			return RegistedSchema
		}
		s.mapping[temp] = handleFn
	}
	return nil
}
func (s *URISupportMap) Request(u url.URL, hint ... string) (Value) {
	//
	if fn, ok := s.mapping[u.Scheme]; ok {
		if len(hint) > 0{
			return fn(u, &hint[0])
		}else {
			return fn(u, nil)
		}
	}
	return errors.WithMessage(UndefinedSchema, u.Scheme+" is unknown")

}
func (s *URISupportMap) Copy() *URISupportMap {
	temp := make(map[string]func(u url.URL, hintmime *string) (Value))
	for k, v := range s.mapping{
		temp[k] = v
	}
	return &URISupportMap{
		mapping: temp,
	}
}
