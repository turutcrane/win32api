package win32api

import (
	"syscall"
	"unsafe"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zwin32api_windows.go win32api.go

// #include <windows.h>
import "C"

type ATOM int16
type COLORREF DWORD
type DWORD uint32
type HBRUSH syscall.Handle
type HCURSOR syscall.Handle
type HDC syscall.Handle
type HICON syscall.Handle
type HINSTANCE syscall.Handle
type HMENU syscall.Handle
type HMODULE syscall.Handle
type HWND syscall.Handle
type HRESULT LONG
type LONG int32
type LPVOID uintptr
type LRESULT uintptr
type WPARAM uintptr // UINT_PTR
type LPARAM uintptr // LONG_PTR
type Rect struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

const sizeof_RECT = C.sizeof_RECT

type UINT uint32
type WNDPROC uintptr

//type LPCWSTR *uint16
//type LONG_PTR uintptr

type Wndclassex struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

const sizeof_WNDCLASSEX = C.sizeof_WNDCLASSEX

type Createstruct struct {
	CreateParams LPVOID
	Instance     HINSTANCE
	Menu         HMENU
	Parent       HWND
	Cy           int32
	Cx           int32
	X            int32
	Y            int32
	Style        LONG
	Name         *uint16
	Class        *uint16
	ExStyle      DWORD
}

const sizeof_CREATESTRUCTW = C.sizeof_CREATESTRUCTW

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

//sys GetModuleHandle(m *uint16) (handle HMODULE, err error) = GetModuleHandleW
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

//sys GetDpiForWindow(hwnd HWND) (dpi UINT) = user32.GetDpiForWindow
//sys CreateWindowEx(dwExStyle DWORD, lpClassName *uint16, lpWindowName *uint16, dwStyle DWORD, x int, y int, nWidth int, nHeight int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) (hwnd HWND, err error) = user32.CreateWindowExW
//sys LoadIcon(hInstance HINSTANCE, lpIconName *uint16) (icon HICON, err error) = user32.LoadIconW
//sys LoadCursor(hInstance HINSTANCE, lpCursorName *uint16) (cursor HCURSOR, err error) = user32.LoadCursorW
//sys RegisterClassEx(Arg1 *Wndclassex) (atm ATOM) = user32.RegisterClassExW
//sys GetClientRect(hWnd HWND, lpRect *Rect) (b bool, err error) [failretval==false] = user32.GetClientRect
//sys ShowWindow(hWnd HWND, nCmdShow int) (b bool)= user32.ShowWindow
//sys UpdateWindow(hWnd HWND) (b bool) = user32.UpdateWindow
//sys DefWindowProc(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (result LRESULT) = user32.DefWindowProcW
//sys EnableNonClientDpiScaling(hwnd HWND) (b bool, err error) [failretval==false] = user32.EnableNonClientDpiScaling
//sys GetDC(hwnd HWND) (hdc HDC) = user32.GetDC
//sys GetDeviceCaps(hdc HDC, index int) (i int) = user32.GetDeviceCaps
//sys ReleaseDC(hWnd HWND, hDC HDC) (b bool) = user32.ReleaseDC
//sys GetProcessDpiAwareness(hprocess syscall.Handle , value *ProcessDpiAwareness) (result HRESULT) = shcore.GetProcessDpiAwareness
//sys CallWindowProc(lpPrevWndFunc WNDPROC, hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (result LRESULT) = user32.CallWindowProcW

//sys CreateSolidBrush(color COLORREF) (hBrush HBRUSH) = gdi32.CreateSolidBrush
