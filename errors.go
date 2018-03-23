package gumi

import "github.com/pkg/errors"

var (
	ErrorRequestSchema = errors.New("Schema error")
	ErrorNotNil = errors.New("Not allow nil to this")
	ErrorNameConflict = errors.New("Name conflict")
)