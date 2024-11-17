package bptree

import "fmt"

var GLOBAL_DEBUG = true

func log(format string, a ...interface{}) (n int, err error) {
	if GLOBAL_DEBUG {
		return fmt.Printf(format,a...)
	}
	return 0,nil
}