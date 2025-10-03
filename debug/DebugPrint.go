package debug

import "fmt"

const (
	debug = false
)

func Println(a ...any) {
	if debug {
		fmt.Println(a...)
	}
}
