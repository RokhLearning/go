package main

import (
	"fmt"
)

type myStruct struct{}

func main() {
	generic("stirng")
	generic(123)
	generic(myStruct{})
	generic(1.32)

}
func generic(a any) {

	switch a.(type) {
	case int:
		i, result := a.(int)
		i++
		fmt.Printf("this is int %+v %+v\n", i, result)
	case string:
		fmt.Printf("this is string\n")
	case myStruct:
		fmt.Printf("this is myStruct\n")
	default:
		fmt.Printf("unknown\n")
	}
}
