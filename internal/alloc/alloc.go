package alloc

// #include <stdlib.h>
// #include <windef.h>
import "C"
import "unsafe"

const (
	MaxPath = C.MAX_PATH
	SizeofTCHAR = C.sizeof_TCHAR
	SizeofWCHAR = C.sizeof_WCHAR
	SizeofPWSTR = C.sizeof_PWSTR
)

func Malloc(length int) uintptr {
	return uintptr(C.malloc(C.size_t(length)))
}

func Free(ptr uintptr) {
	C.free(unsafe.Pointer(ptr))
}