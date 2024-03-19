package reflection

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionName returns the name of the function
func GetFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return fullName[strings.LastIndex(fullName, ".")+1:]
}
