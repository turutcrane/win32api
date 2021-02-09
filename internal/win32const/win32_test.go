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

	if unsafe.Sizeof(win32api.GUID{}) != sizeof_GUID {
		t.Error("T28: GUID: size dose not match!", unsafe.Sizeof(win32api.GUID{}), sizeof_GUID)
	}
}

func TestIdc(t *testing.T) {
	if IdcArrow != uintptr(unsafe.Pointer(IDC_ARROW)) {
		t.Error("T47:IDC_ARROW")
	}
	if IdcIbeam != uintptr(unsafe.Pointer(IDC_IBEAM)) {
		t.Error("T50:IDC_IBEAM")
	}
	if IdcWait != uintptr(unsafe.Pointer(IDC_WAIT)) {
		t.Error("T43:IDC_WAIT")
	}
	if IdcCross != uintptr(unsafe.Pointer(IDC_CROSS)) {
		t.Error("T43:IDC_CROSS")
	}
	if IdcUparrow != uintptr(unsafe.Pointer(IDC_UPARROW)) {
		t.Error("T43:IDC_UPARROW")
	}
	if IdcSize != uintptr(unsafe.Pointer(IDC_SIZE)) {
		t.Error("T43:IDC_SIZE")
	}
	if IdcIcon != uintptr(unsafe.Pointer(IDC_ICON)) {
		t.Error("T43:IDC_ICON")
	}
	if IdcSizenwse != uintptr(unsafe.Pointer(IDC_SIZENWSE)) {
		t.Error("T43:IDC_SIZENWSE")
	}
	if IdcSizenesw != uintptr(unsafe.Pointer(IDC_SIZENESW)) {
		t.Error("T43:IDC_SIZENESW")
	}
	if IdcSizewe != uintptr(unsafe.Pointer(IDC_SIZEWE)) {
		t.Error("T43:IDC_SIZEWE")
	}
	if IdcSizens != uintptr(unsafe.Pointer(IDC_SIZENS)) {
		t.Error("T43:IDC_SIZENS")
	}
	if IdcSizeall != uintptr(unsafe.Pointer(IDC_SIZEALL)) {
		t.Error("T43:IDC_SIZEALL")
	}
	if IdcNo != uintptr(unsafe.Pointer(IDC_NO)) {
		t.Error("T43:IDC_NO")
	}
	if IdcHand != uintptr(unsafe.Pointer(IDC_HAND)) {
		t.Error("T43:IDC_HAND")
	}
	if IdcAppstarting != uintptr(unsafe.Pointer(IDC_APPSTARTING)) {
		t.Error("T43:IDC_APPSTARTING")
	}
	if IdcHelp != uintptr(unsafe.Pointer(IDC_HELP)) {
		t.Error("T43:IDC_HELP")
	}

}
