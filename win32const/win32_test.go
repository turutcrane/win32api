package win32const

import (
	"testing"
	"unsafe"

	"github.com/turutcrane/win32api"
)

// 構造体は、サイズを C.sizeof_ と unsafe.Sizeof を比べる。
func TestStructSize(t *testing.T) {
	if unsafe.Sizeof(win32api.Wndclassex{}) != sizeof_WNDCLASSEX {
		t.Error("T10: size dose not match!", unsafe.Sizeof(win32api.Wndclassex{}), sizeof_WNDCLASSEX)
	}
	if unsafe.Sizeof(win32api.Createstruct{}) != sizeof_CREATESTRUCTW {
		t.Error("T13: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Createstruct{}), sizeof_CREATESTRUCTW)
	}
	if unsafe.Sizeof(win32api.Rect{}) != sizeof_RECT {
		t.Error("T13: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Createstruct{}), sizeof_RECT)
	}
}
