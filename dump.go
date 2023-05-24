package dbg

import (
	"os"

	pp "github.com/k0kubun/pp/v3"
	"github.com/mattn/go-isatty"
)

func init() {
	pp.ColoringEnabled = isatty.IsTerminal(os.Stdout.Fd())
}

// Dump wraps pp.Println for now.
func Dump(a ...interface{}) {
	pp.Println(a...)
}

// Fatal wraps pp.Fatal for now.
func Fatal(a ...interface{}) {
	pp.Fatal(a...)
}
