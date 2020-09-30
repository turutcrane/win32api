package win32const

import (
	"testing"
	"unsafe"

	"github.com/turutcrane/win32api"
)

// 構造体は、サイズを C.sizeof_ と unsafe.Sizeof を比べる。
func TestStructSize(t *testing.T) {
	if unsafe.Sizeof(win32api.Wndclassex{}) != sizeof_WNDCLASSEX {
		t.Error("T13: size dose not match!", unsafe.Sizeof(win32api.Wndclassex{}), sizeof_WNDCLASSEX)
	}
	if unsafe.Sizeof(win32api.Createstruct{}) != sizeof_CREATESTRUCTW {
		t.Error("T16: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Createstruct{}), sizeof_CREATESTRUCTW)
	}
	if unsafe.Sizeof(win32api.Rect{}) != sizeof_RECT {
		t.Error("T19: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Createstruct{}), sizeof_RECT)
	}
	if unsafe.Sizeof(win32api.Paintstruct{}) != sizeof_PAINTSTRUCT {
		t.Error("T22: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Paintstruct{}), sizeof_PAINTSTRUCT)
	}
	if unsafe.Sizeof(win32api.Findreplace{}) != sizeof_FINDREPLACE {
		t.Error("T25: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Findreplace{}), sizeof_FINDREPLACE)
	}
	if unsafe.Sizeof(win32api.Pixelformatdescriptor{}) != sizeof_PIXELFORMATDESCRIPTOR {
		t.Error("T28: Createstruct: size dose not match!", unsafe.Sizeof(win32api.Pixelformatdescriptor{}), sizeof_PIXELFORMATDESCRIPTOR)
	}

	if unsafe.Sizeof(win32api.LARGE_INTEGER(0)) != sizeof_LONGLONG {
		t.Error("T31: LARGE_INTEGER: size dose not match!", unsafe.Sizeof(win32api.LARGE_INTEGER(0)), sizeof_LONGLONG)
	}

	if unsafe.Sizeof(win32api.Trackmouseevent{}) != sizeof_TRACKMOUSEEVENT {
		t.Error("T28: Trackmouseevent: size dose not match!", unsafe.Sizeof(win32api.Trackmouseevent{}), sizeof_TRACKMOUSEEVENT)

	}
}
