package gumi

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 1
	VERSION_PATCH = 0
)

func Version() (major, minor, patch int) {
	return VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH
}
