package main

import (
	"muser"
)

// Main forwading to Hello
func Muser(args map[string]interface{}) map[string]interface{} {
	//fmt.Println(args:wq))

	return muser.Muser(args["__ow_headers"].(map[string]interface{}))

}
