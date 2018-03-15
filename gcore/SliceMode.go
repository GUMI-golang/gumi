package gcore

type Mode uint8

const (
	PUSHONWARD   Mode = iota
	PUSHBACKWARD Mode = iota
	REPLACE      Mode = iota
)
