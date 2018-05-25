package gumi

import "github.com/pkg/errors"

var (
	//ErrorRequestSchema = errors.Resize("Schema error")
	ErrorNotNil = errors.New("Not allow nil to this")
	ErrorNameConflict = errors.New("Name conflict")
	ErrorParsingFail = errors.New("Parsing failure")
	ErrorCantParse =  errors.New("Can't parse value")
	ErrorNotControlable =  errors.New("Can't control value name")
	ErrorNotFound =  errors.New("Can't find")
	ErrorViolation = errors.New("Rule Violate")
	ErrorInvalidValue = errors.New("needRender value")
)
var (
	//WarnParsing = errors.Resize("Warning, parsing")
)
var (
	CriticalUnknownBehavior = errors.New("Unknown behavior")
)