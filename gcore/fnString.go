package gcore

import (
	"regexp"
	"strings"
	"unicode"
	"strconv"
	"github.com/pkg/errors"
)

var re_fn = regexp.MustCompile(`(?P<name>.*)\((?P<args>.*)\)`)


type FunctionArgument func(value string) (interface{}, error)
var (
	FnArgInt = FunctionArgument(func(value string) (interface{}, error){
		return strconv.ParseInt(value, 10, 64)
	})
	FnArgUint = FunctionArgument(func(value string) (interface{}, error){
		return strconv.ParseUint(value, 10, 64)
	})
	FnArgFloat = FunctionArgument(func(value string) (interface{}, error){
		return strconv.ParseFloat(value, 64)
	})
	FnArgString = FunctionArgument(func(value string) (interface{}, error){
		if !strings.HasPrefix(value, "'") || !strings.HasSuffix(value, "'"){
			return nil, errors.New("String must starts with ', ends with '")
		}
		return value[1:len(value) - 1], nil
	})
)
type FunctionCaller struct {
	Fns []Function
}
type Function struct {
	Name string
	Args []FunctionArgument
	Callback func(args ... interface{}) []interface{}
}
func (s FunctionCaller ) Call(script string) (result []interface{}, err error) {
	var (
		name string
		args []string
		cargs []interface{}
		target *Function
	)
	if temp := re_fn.FindStringSubmatch(script); len(temp) < 3{
		name = script
	}else {
		name = temp[1]
		args = strings.FieldsFunc(temp[2], func(r rune) bool {
			return unicode.IsSpace(r) || r == ','
		})
	}
	for _, fn := range s.Fns {
		if fn.Name == name{
			target = &fn
			break
		}
	}
	if target == nil{
		return nil, errors.New("No matching name")
	}
	if len(target.Args) != len(args){
		return nil, errors.New("No matching function")
	}
	cargs = make([]interface{}, len(args))
	for i, parser := range target.Args {
		cargs[i], err = parser(args[i])
		if err != nil {
			return nil, err
		}
	}

	return target.Callback(cargs...), nil
}