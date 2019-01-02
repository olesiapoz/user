package main

import (
	"fmt"
	"register"
)

// Main forwading to Hello
func Register(args map[string]interface{}) map[string]interface{} {
	
	return register.Register(args)

}
