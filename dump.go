package dbg

import (
	pp "github.com/k0kubun/pp/v3"
)

// Dump wraps pp.Println for now.
func Dump(a ...interface{}) {
	pp.Println(a...)
}

// Fatal wraps pp.Fatal for now.
func Fatal(a ...interface{}) {
	pp.Fatal(a...)
}
