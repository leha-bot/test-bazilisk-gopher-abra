package main

import "fmt"

/*
#cgo CFLAGS: -Ilib/testcpplib
#cgo CXXFLAGS: -Ilib/testcpplib
#cgo LDFLAGS: -L./bin/ -ltestcpplib
#include <plusik.h>
*/
import "C"

func main() {
	fmt.Printf("calculate %d+%d=%d\n", 1, 2, C.plusik(1,2))
}

