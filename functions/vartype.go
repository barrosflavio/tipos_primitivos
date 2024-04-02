package functions

import (
	"fmt"
	"reflect"
)

func VarType(input interface{}){
	fmt.Println(reflect.TypeOf(input))
}