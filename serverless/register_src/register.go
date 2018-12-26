package main

import (
	"fmt"
	"register"
)

// Main forwading to Hello
func Register(args map[string]interface{}) map[string]interface{} {
	fmt.Println(args)
	return register.Register(args)

}
