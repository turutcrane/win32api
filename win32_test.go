package win32api

import (
	"testing"
	"unsafe"
)

// 構造体は、サイズを C.sizeof_ と unsafe.Sizeof を比べる。
func TestStructSize(t *testing.T) {
	if unsafe.Sizeof(Wndclassex{}) != sizeof_WNDCLASSEX {
		t.Error("T10: size dose not match!", unsafe.Sizeof(Wndclassex{}), sizeof_WNDCLASSEX)
	}
	if unsafe.Sizeof(Createstruct{}) != sizeof_CREATESTRUCTW {
		t.Error("T13: Createstruct: size dose not match!", unsafe.Sizeof(Createstruct{}), sizeof_CREATESTRUCTW)
	}
	if unsafe.Sizeof(Rect{}) != sizeof_RECT {
		t.Error("T13: Createstruct: size dose not match!", unsafe.Sizeof(Createstruct{}), sizeof_RECT)
	}
}
