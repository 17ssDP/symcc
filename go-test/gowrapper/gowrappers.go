package gowrapper

//#cgo CFLAGS: -I/home/dp/Documents/fuzzing/symcc/runtime
//#cgo LDFLAGS: -L/home/dp/Documents/fuzzing/symcc/SymRuntime-prefix/src/SymRuntime-build -lSymRuntime -Wl,-rpath,/home/dp/Documents/fuzzing/symcc/SymRuntime-prefix/src/SymRuntime-build
//#include <LibcWrappers.h>
import (
	"C"
)

import (
	"os"
	"unsafe"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	return f, err
}

func Make(size int) []byte {
	buffer := make([]byte, size)
	return buffer
}

func Read(f *os.File, b []byte, size int) (n int, err error) {
	n, err = f.Read(b)
	return n, err
}

func OpenSymbolized(path string) (*os.File, error) {
	f, err := os.Open(path)
	C.setNullReturnExpression()
	C.setInputFileDescriptor(C.int(uintptr(f.Fd())))
	C.setInputOffset(C.int(0))
	return f, err
}

func MakeSymbolized(size int) []byte {
	buffer := make([]byte, size)
	var f func(int) []byte
	f = MakeSymbolized
	ptr := (*int)(unsafe.Pointer(&f))
	m := *((*int)(unsafe.Pointer(ptr)))
	maddr := *((*int)(unsafe.Pointer(uintptr(m))))
	C.tryMallocAlternative(C.int(size), C.longlong(int64(maddr)))
	C.setNullReturnExpression()
	return buffer
}

func ReadSymbolized(f *os.File, b []byte, size int) (n int, err error) {
	C.tryReadAlternative(C.CBytes(b), C.size_t(size))
	n, err = f.Read(b)
	C.setNullReturnExpression()
	C.readSymbolic(C.CBytes(b), C.int(n))
	return n, err
}
