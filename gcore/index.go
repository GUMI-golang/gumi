package gcore

type Index int

func (s Index) Raw() int {
	return int(s)
}
func (s Index) Indexize(length int) int {
	// If there is no length, Index not exist either
	if length <= 0 {
		return IndexNotExist
	}
	temp := int(s)
	if 0 <= temp && temp < length {
		// normal index
		return temp
	} else if 1 <= -temp && -temp <= length {
		//
		return length + temp
	}
	return IndexNotExist
}
const (
	IndexNotExist = -1
)
const (
	First Index = 0
	Last  Index = -1
)
