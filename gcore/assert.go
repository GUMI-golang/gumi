package gcore

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
func MustValue(val interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return val
}
