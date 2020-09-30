package win32api

import (
	"syscall"
	"unsafe"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zwin32api_windows.go win32api.go w32api_windows.go

type ATOM int16
type BOOL int32
type BYTE byte
type COLORREF DWORD
type DWORD uint32
type HANDLE syscall.Handle
type HBRUSH syscall.Handle
type HCURSOR syscall.Handle
type HDC syscall.Handle
type HDWP syscall.Handle
type HFONT syscall.Handle
type HGDIOBJ syscall.Handle
type HGLRC syscall.Handle
type HICON syscall.Handle
type HINSTANCE syscall.Handle
type HKL syscall.Handle
type HMENU syscall.Handle
type HMODULE syscall.Handle
type HWND syscall.Handle
type HRESULT LONG
type HRGN syscall.Handle
type LONG int32
type LPVOID uintptr
type LRESULT uintptr
type WPARAM uintptr // UINT_PTR
type LPARAM uintptr // LONG_PTR
type WORD int16
type SHORT int16
type LARGE_INTEGER int64
type WCHAR uint16

type ProcessDpiAwareness int

type UINT uint32
type INT_PTR uintptr
type WNDPROC uintptr
type DLGPROC uintptr
type LPFRHOOKPROC uintptr

// type WndProc func(hWnd win32api.HWND, message win32api.UINT, wParam win32api.WPARAM, lParam win32api.LPARAM) win32api.LRESULT
type WndProc func(hWnd HWND, message UINT, wParam WPARAM, lParam LPARAM) LRESULT

//type LPCWSTR *uint16
//type LONG_PTR uintptr

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func ToPCreatestruct(lparam LPARAM) *Createstruct {
	return (*Createstruct)(unsafe.Pointer(lparam))
}

func WndProcToWNDPROC(wp WndProc) WNDPROC {
	return WNDPROC(syscall.NewCallback(wp))
}

func Makepoints(lParam LPARAM) (points Points) {
	points.X = SHORT(lParam & 0xff)
	points.Y = SHORT((lParam >> 16) & 0xff)

	return points
}

//sys SetLastError(dwErrCode DWORD) = SetLastError

// GetWindowLongPtr does not return syscall.EINVAL
//sys getWindowLongPtr(hWnd HWND, nIndex int) (longPtr uintptr, err error) = user32.GetWindowLongPtrW
func GetWindowLongPtr(hWnd HWND, nIndex int) (longPtr uintptr, err error) {
	longPtr, err = getWindowLongPtr(hWnd, nIndex)
	if err == syscall.EINVAL {
		err = nil
	}
	return
}

// SetWindowLongPtr does not return syscall.EINTVAL
//sys setWindowLongPtr(hWnd HWND, nIndex int, dwNewLong uintptr) (longPtr uintptr, err error) = user32.SetWindowLongPtrW
func SetWindowLongPtr(hWnd HWND, nIndex int, dwNewLong uintptr) (longPtr uintptr, err error) {
	SetLastError(0)
	longPtr, err = setWindowLongPtr(hWnd, nIndex, dwNewLong)
	if err == syscall.EINVAL {
		err = nil
	}
	return
}

func BoolToBOOL(b bool) BOOL {
	if b {
		return 1
	}
	return 0
}

func BOOLToBool(b BOOL) bool {
	if b != 0 {
		return true
	}
	return false
}

func LOWORD(w WPARAM) uintptr {
	return uintptr(w & 0xffff)
}

func HIWORD(w WPARAM) uintptr {
	return uintptr((w >> 16) & 0xffff)
}

func GET_X_LPARAM(lParam LPARAM) int {
	low := lParam & 0xffff
	var x int
	if low & 0x8000 != 0 {
		x = int(low) - 0x10000
	} else {
		x = int(low)
	}
	return x
}

func GET_Y_LPARAM(lParam LPARAM) int {
	high := (lParam >> 16) & 0xffff
	var y int
	if high & 0x8000 != 0 {
		y = int(high) - 0x10000
	} else {
		y = int(high)
	}
	return y
}

func GET_WHEEL_DELTA_WPARAM(wParam WPARAM) int {
	high := (wParam >> 16) & 0xffff
	var delta int
	if high & 0x8000 != 0 {
		delta = int(high) - 0x10000
	} else {
		delta = int(high)
	}
	return delta
}

func LParamToPRect(lParam LPARAM) *Rect {
	return (*Rect)(unsafe.Pointer(lParam))
}

func DialogBoxParam(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) (r INT_PTR, err error) {
	r, err = dialogBoxParamW(hInstance, lpTemplateName, hWndParent, lpDialogFunc, dwInitParam)

	// If the function fails because the hWndParent parameter is invalid, the return value is zero.
	if err != nil && r == 0 {
		err = syscall.EINVAL
	}
	return r, err
}

//sys windowFromPoint(point uintptr) (r HWND) = user32.WindowFromPoint
func WindowFromPoint(point Point) (r HWND) {
	var xy uintptr
	x := uintptr(point.X)
	y := uintptr(point.Y)
	mask32 := (^uintptr(0) >> (64 - 32))
	xy = (x & mask32) | (y << 32)

	return windowFromPoint(xy)
}
