package gime

import "github.com/pkg/errors"

var (
	RegistedSchema =  errors.New("Already registered schema")
	UndefinedSchema = errors.New("Schema not support")
	UnknownMime = errors.New("Unknown Mime type")
)
