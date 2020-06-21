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
type HBRUSH syscall.Handle
type HCURSOR syscall.Handle
type HDC syscall.Handle
type HDWP syscall.Handle
type HFONT uintptr
type HGDIOBJ uintptr
type HICON syscall.Handle
type HINSTANCE syscall.Handle
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

type ProcessDpiAwareness int

type UINT uint32
type WNDPROC uintptr

//type LPCWSTR *uint16
//type LONG_PTR uintptr

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func ToPCreatestruct(lparam LPARAM) *Createstruct {
	return (*Createstruct)(unsafe.Pointer(lparam))
}

func Makepoints(lParam LPARAM) (points Points) {
	points.X = SHORT(lParam & 0xff)
	points.Y = SHORT((lParam >> 16) & 0xff)

	return points
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
//sys CreateWindowEx (dwExStyle DWORD, lpClassName *uint16, lpWindowName *uint16, dwStyle DWORD, x int, y int, nWidth int, nHeight int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) (hwnd HWND, err error) = user32.CreateWindowExW
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
//sys EnableWindow(hWnd HWND, bEnable bool) (b bool) = user32.EnableWindow
//sys SendMessage(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (resust LRESULT)= user32.SendMessageW
//sys BeginPaint(hWnd HWND, lpPaint *Paintstruct) (hdc HDC) = user32.BeginPaint
//sys EndPaint(hWnd HWND, lpPaint *Paintstruct) = user32.EndPaint
//sys IsWindowEnabled(hWnd HWND) (b bool) = user32.IsWindowEnabled
//sys SetWindowPos(hWnd HWND, hWndInsertAfter HWND, X int, Y int, cx int, cy int, uFlags UINT) (b bool, err error) [failretval==false] = user32.SetWindowPos
//sys IsWindowVisible(hWnd HWND) (b bool) = user32.IsWindowVisible
//sys SetMenu(hWnd HWND, hMenu HMENU) (b bool, err error) [failretval==false] = user32.SetMenu
//sys BeginDeferWindowPos(nNumWindows int) (hdwp HDWP, err error) = user32.BeginDeferWindowPos
//sys DeferWindowPos(hWinPosInfo HDWP, hWnd HWND, hWndInsertAfter HWND, x int, y int, cx int, cy int, uFlags UINT) (hdwp HDWP, err error) = user32.DeferWindowPos

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

func LParamToPRect(lParam LPARAM) *Rect {
	return (*Rect)(unsafe.Pointer(lParam))
}
