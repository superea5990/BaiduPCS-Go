package cachepool

import (
	"reflect"
	"unsafe"
)

//go:linkname mallocgc runtime.mallocgc
func mallocgc(size uintptr, typ uintptr, needzero bool) unsafe.Pointer

// RawByteSlice allocates a byte slice with the given size.
// Originally it was linked to runtime.rawbyteslice via go:linkname,
// but that internal symbol was removed in newer Go releases,
// so it falls back to the standard make() allocation.
func RawByteSlice(size int) (b []byte) {
	return make([]byte, size)
}

// RawMalloc allocates a new slice. The slice is not zeroed.
func RawMalloc(size int) unsafe.Pointer {
	return mallocgc(uintptr(size), 0, false)
}

// RawMallocByteSlice allocates a new byte slice. The slice is not zeroed.
func RawMallocByteSlice(size int) []byte {
	p := mallocgc(uintptr(size), 0, false)
	b := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(p),
		Len:  size,
		Cap:  size,
	}))
	return b
}
