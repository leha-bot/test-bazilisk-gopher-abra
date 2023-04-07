package abra

/*
#cgo CFLAGS: -Ilib/testcpplib
#cgo CXXFLAGS: -Ilib/testcpplib
#cgo LDFLAGS: -L./bin/ -ltestcpplib
#include <plusik.h>
*/
import "C"

// Plusik simply sums two arguments.
func Plusik(a,b int) (C.int) {
	return C.plusik(1,2);
}

