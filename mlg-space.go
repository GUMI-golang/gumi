package gumi

import "github.com/pkg/errors"

type Space struct {
	m map[string]map[string]*SupportTag
}
type SpaceSupport func(space *Space) error

func NewSpace(supports ...SpaceSupport) (*Space, error) {
	temp := &Space{
		m: make(map[string]map[string]*SupportTag),
	}
	for _, v := range supports {
		if err := v(temp); err != nil {
			return nil, err
		}
	}
	return temp, nil
}
func (s *Space) Support(ns string, tag *SupportTag) error {
	if tag == nil {
		return ErrorNotNil
	}
	testcase := tag.New()
	if _, ok := testcase.(DefaultManager);ok{
		if testcase.MaxChildrun() != 0{
			return errors.Wrap(ErrorViolation, "Default Managered GUMI Must have 0 MaxChildrun")
		}
	}
	if sp, ok := s.m[ns]; ok {
		if _, ok := sp[tag.Name]; ok {
			return errors.Wrap(ErrorNameConflict, ns+":"+tag.Name+" is exist")
		}
		sp[tag.Name] = tag
	} else {
		s.m[ns] = map[string]*SupportTag{
			tag.Name: tag,
		}
	}
	return nil
}
func (s *Space) Find(ns, name string) *SupportTag {
	if ns == "" {
		for _, v := range s.m {
			if v2, ok := v[name]; ok {
				return v2
			}
		}
	} else {
		if v, ok := s.m[ns]; ok {
			return v[name]
		}
	}
	return nil
}

type (
	SupportTag struct {
		Name       string
		New        func() GUMI
	}
	TagArg struct {
		Name  string
		SetFn func(gumi GUMI, value string) error
		GetFn func(gumi GUMI) (string, error)
	}
	// TODO: Callback, setup Javascript
	//TagCallback struct {
	//	Name string
	//}
)
