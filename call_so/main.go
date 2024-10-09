package main

//#cgo LDFLAGS: -L./lib -lshared
//#include "lib/lib.h"
import "C"
import (
	"fmt"
)

// #cgo LDFLAGS: -L./lib -lshared 动态链接
func main() {
	a, b := 3, 4
	result := C.add_numbers(C.int(a), C.int(b))
	fmt.Printf("Result of add_numbers(%d, %d): %v\n", a, b, result)
}
