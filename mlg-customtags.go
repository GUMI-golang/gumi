package gumi

var mlgTags = make(map[mlgXMLTag]MLGTag)

type mlgXMLTag struct {
	namespace string
	name string
}

func (s mlgXMLTag) String() string {
	if len(s.namespace) > 0{
		return s.namespace + ":"+s.name
	}
	if len(s.name) > 0{
		return s.name
	}
	return "< nil >"
}

func AddMLGNamespace(ns MLGNamespace) error {
	for _, tag := range ns.Tags{
		key := mlgXMLTag{
			namespace:ns.Namespace,
			name:tag.Name,
		}
		if _, ok := mlgTags[key]; ok {
			return ErrorNameConflict
		}
		mlgTags[key] = tag
	}
	return nil
}

type MLGNamespace struct {
	Namespace string
	Tags []MLGTag
}
type MLGTag struct {
	Name string
	Args []MLGArg
}

type MLGArg struct {
	Name  string
	SetFn func(value string) error
	GetFn func(value string) error
}

// TODO: Callback, setup Javascript
//type MLGCallback struct {
//	Name string
//	OnFn func(value string) error
//	ReferFn func(value string) error
//}
