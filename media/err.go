package media

import "github.com/pkg/errors"

var (
	CriticalRootOverride = errors.New("Root can't override")
)
