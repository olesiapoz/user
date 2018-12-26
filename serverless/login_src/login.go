package main

import (
	"login"
)

// Main forwading to Hello
func Login(args map[string]interface{}) map[string]interface{} {

	return login.Login(args["__ow_headers"].(map[string]interface{}))

}
