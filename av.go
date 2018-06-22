package avgo

import "unsafe"

func GoBytes(cArray unsafe.Pointer, length int) []byte {
	return (*[1<<30]byte)(cArray)[:length]
}