// Code generated by 'go generate'; DO NOT EDIT.

package win32api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	modgdi32    = windows.NewLazySystemDLL("gdi32.dll")
	modcomdlg32 = windows.NewLazySystemDLL("comdlg32.dll")
	modOpengl32 = windows.NewLazySystemDLL("Opengl32.dll")
	modShcore   = windows.NewLazySystemDLL("Shcore.dll")
	modshell32  = windows.NewLazySystemDLL("shell32.dll")
	modShell32  = windows.NewLazySystemDLL("Shell32.dll")
	modOle32    = windows.NewLazySystemDLL("Ole32.dll")

	procSetLastError              = modkernel32.NewProc("SetLastError")
	procGetWindowLongPtrW         = moduser32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtrW         = moduser32.NewProc("SetWindowLongPtrW")
	procWindowFromPoint           = moduser32.NewProc("WindowFromPoint")
	procGetModuleHandleW          = modkernel32.NewProc("GetModuleHandleW")
	procScreenToClient            = moduser32.NewProc("ScreenToClient")
	procPtInRegion                = modgdi32.NewProc("PtInRegion")
	procEndDeferWindowPos         = moduser32.NewProc("EndDeferWindowPos")
	procCreateSolidBrush          = modgdi32.NewProc("CreateSolidBrush")
	procDeleteObject              = modgdi32.NewProc("DeleteObject")
	procCreateFontW               = modgdi32.NewProc("CreateFontW")
	procCreateRectRgn             = modgdi32.NewProc("CreateRectRgn")
	procDestroyWindow             = moduser32.NewProc("DestroyWindow")
	procPostMessageW              = moduser32.NewProc("PostMessageW")
	procIsRectEmpty               = moduser32.NewProc("IsRectEmpty")
	procAdjustWindowRectEx        = moduser32.NewProc("AdjustWindowRectEx")
	procSetParent                 = moduser32.NewProc("SetParent")
	procSetWindowPos              = moduser32.NewProc("SetWindowPos")
	procDialogBoxParamW           = moduser32.NewProc("DialogBoxParamW")
	procEndDialog                 = moduser32.NewProc("EndDialog")
	procSetFocus                  = moduser32.NewProc("SetFocus")
	procFindTextW                 = modcomdlg32.NewProc("FindTextW")
	procCommDlgExtendedError      = modcomdlg32.NewProc("CommDlgExtendedError")
	procRegisterWindowMessageW    = moduser32.NewProc("RegisterWindowMessageW")
	procSetWindowTextW            = moduser32.NewProc("SetWindowTextW")
	procIsWindow                  = moduser32.NewProc("IsWindow")
	procClientToScreen            = moduser32.NewProc("ClientToScreen")
	procChoosePixelFormat         = modgdi32.NewProc("ChoosePixelFormat")
	procSetPixelFormat            = modgdi32.NewProc("SetPixelFormat")
	procSwapBuffers               = modgdi32.NewProc("SwapBuffers")
	procwglCreateContext          = modOpengl32.NewProc("wglCreateContext")
	procwglMakeCurrent            = modOpengl32.NewProc("wglMakeCurrent")
	procGetDpiForWindow           = moduser32.NewProc("GetDpiForWindow")
	procCreateWindowExW           = moduser32.NewProc("CreateWindowExW")
	procLoadIconW                 = moduser32.NewProc("LoadIconW")
	procLoadCursorW               = moduser32.NewProc("LoadCursorW")
	procRegisterClassExW          = moduser32.NewProc("RegisterClassExW")
	procGetClientRect             = moduser32.NewProc("GetClientRect")
	procShowWindow                = moduser32.NewProc("ShowWindow")
	procUpdateWindow              = moduser32.NewProc("UpdateWindow")
	procDefWindowProcW            = moduser32.NewProc("DefWindowProcW")
	procEnableNonClientDpiScaling = moduser32.NewProc("EnableNonClientDpiScaling")
	procGetDC                     = moduser32.NewProc("GetDC")
	procGetDeviceCaps             = modgdi32.NewProc("GetDeviceCaps")
	procReleaseDC                 = moduser32.NewProc("ReleaseDC")
	procGetProcessDpiAwareness    = modShcore.NewProc("GetProcessDpiAwareness")
	procCallWindowProcW           = moduser32.NewProc("CallWindowProcW")
	procEnableWindow              = moduser32.NewProc("EnableWindow")
	procSendMessageW              = moduser32.NewProc("SendMessageW")
	procBeginPaint                = moduser32.NewProc("BeginPaint")
	procEndPaint                  = moduser32.NewProc("EndPaint")
	procIsWindowEnabled           = moduser32.NewProc("IsWindowEnabled")
	procIsWindowVisible           = moduser32.NewProc("IsWindowVisible")
	procSetMenu                   = moduser32.NewProc("SetMenu")
	procBeginDeferWindowPos       = moduser32.NewProc("BeginDeferWindowPos")
	procDeferWindowPos            = moduser32.NewProc("DeferWindowPos")
	procGetMessageExtraInfo       = moduser32.NewProc("GetMessageExtraInfo")
	procGetMessageTime            = moduser32.NewProc("GetMessageTime")
	procGetSystemMetrics          = moduser32.NewProc("GetSystemMetrics")
	procGetDoubleClickTime        = moduser32.NewProc("GetDoubleClickTime")
	procSetCapture                = moduser32.NewProc("SetCapture")
	procGetKeyState               = moduser32.NewProc("GetKeyState")
	procGetCapture                = moduser32.NewProc("GetCapture")
	procReleaseCapture            = moduser32.NewProc("ReleaseCapture")
	procQueryPerformanceCounter   = modkernel32.NewProc("QueryPerformanceCounter")
	procQueryPerformanceFrequency = modkernel32.NewProc("QueryPerformanceFrequency")
	procTrackMouseEvent           = moduser32.NewProc("TrackMouseEvent")
	procGetCursorPos              = moduser32.NewProc("GetCursorPos")
	procGetKeyboardLayout         = moduser32.NewProc("GetKeyboardLayout")
	procVkKeyScanExW              = moduser32.NewProc("VkKeyScanExW")
	procGetMenu                   = moduser32.NewProc("GetMenu")
	procGetSubMenu                = moduser32.NewProc("GetSubMenu")
	procRemoveMenu                = moduser32.NewProc("RemoveMenu")
	procGetCommandLineW           = modkernel32.NewProc("GetCommandLineW")
	procSHGetFolderPathW          = modshell32.NewProc("SHGetFolderPathW")
	procSHGetKnownFolderPath      = modShell32.NewProc("SHGetKnownFolderPath")
	procCoTaskMemFree             = modOle32.NewProc("CoTaskMemFree")
	procFindResourceW             = modkernel32.NewProc("FindResourceW")
	procLoadResource              = modkernel32.NewProc("LoadResource")
	procSizeofResource            = modkernel32.NewProc("SizeofResource")
	procLockResource              = modkernel32.NewProc("LockResource")
)

func SetLastError(dwErrCode DWORD) {
	syscall.Syscall(procSetLastError.Addr(), 1, uintptr(dwErrCode), 0, 0)
	return
}

func getWindowLongPtr(hWnd HWND, nIndex int) (longPtr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowLongPtrW.Addr(), 2, uintptr(hWnd), uintptr(nIndex), 0)
	longPtr = uintptr(r0)
	if longPtr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setWindowLongPtr(hWnd HWND, nIndex int, dwNewLong uintptr) (longPtr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procSetWindowLongPtrW.Addr(), 3, uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	longPtr = uintptr(r0)
	if longPtr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func windowFromPoint(point uintptr) (r HWND) {
	r0, _, _ := syscall.Syscall(procWindowFromPoint.Addr(), 1, uintptr(point), 0, 0)
	r = HWND(r0)
	return
}

func GetModuleHandle(lpModuleName *uint16) (r HMODULE, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, uintptr(unsafe.Pointer(lpModuleName)), 0, 0)
	r = HMODULE(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ScreenToClient(hWnd HWND, lpPoint *Point) (r bool) {
	r0, _, _ := syscall.Syscall(procScreenToClient.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)), 0)
	r = r0 != 0
	return
}

func PtInRegion(hrgn HRGN, x int, y int) (r bool) {
	r0, _, _ := syscall.Syscall(procPtInRegion.Addr(), 3, uintptr(hrgn), uintptr(x), uintptr(y))
	r = r0 != 0
	return
}

func EndDeferWindowPos(hWinPosInfo HDWP) (err error) {
	r1, _, e1 := syscall.Syscall(procEndDeferWindowPos.Addr(), 1, uintptr(hWinPosInfo), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateSolidBrush(color COLORREF) (r HBRUSH) {
	r0, _, _ := syscall.Syscall(procCreateSolidBrush.Addr(), 1, uintptr(color), 0, 0)
	r = HBRUSH(r0)
	return
}

func DeleteObject(ho HGDIOBJ) (r bool) {
	r0, _, _ := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(ho), 0, 0)
	r = r0 != 0
	return
}

func CreateFont(cHeight int, cWidth int, cEscapement int, cOrientation int, cWeight int, bItalic bool, bUnderline bool, bStrikeOut bool, iCharSet DWORD, iOutPrecision DWORD, iClipPrecision DWORD, iQuality DWORD, iPitchAndFamily DWORD, pszFaceName *uint16) (r HFONT) {
	var _p0 uint32
	if bItalic {
		_p0 = 1
	} else {
		_p0 = 0
	}
	var _p1 uint32
	if bUnderline {
		_p1 = 1
	} else {
		_p1 = 0
	}
	var _p2 uint32
	if bStrikeOut {
		_p2 = 1
	} else {
		_p2 = 0
	}
	r0, _, _ := syscall.Syscall15(procCreateFontW.Addr(), 14, uintptr(cHeight), uintptr(cWidth), uintptr(cEscapement), uintptr(cOrientation), uintptr(cWeight), uintptr(_p0), uintptr(_p1), uintptr(_p2), uintptr(iCharSet), uintptr(iOutPrecision), uintptr(iClipPrecision), uintptr(iQuality), uintptr(iPitchAndFamily), uintptr(unsafe.Pointer(pszFaceName)), 0)
	r = HFONT(r0)
	return
}

func CreateRectRgn(x1 int, y1 int, x2 int, y2 int) (r HRGN) {
	r0, _, _ := syscall.Syscall6(procCreateRectRgn.Addr(), 4, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), 0, 0)
	r = HRGN(r0)
	return
}

func DestroyWindow(hWnd HWND) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyWindow.Addr(), 1, uintptr(hWnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func PostMessage(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (err error) {
	r1, _, e1 := syscall.Syscall6(procPostMessageW.Addr(), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func IsRectEmpty(lprc *Rect) (r bool) {
	r0, _, _ := syscall.Syscall(procIsRectEmpty.Addr(), 1, uintptr(unsafe.Pointer(lprc)), 0, 0)
	r = r0 != 0
	return
}

func AdjustWindowRectEx(lpRect *Rect, dwStyle DWORD, bMenu bool, dwExStyle DWORD) (err error) {
	var _p0 uint32
	if bMenu {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procAdjustWindowRectEx.Addr(), 4, uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle), uintptr(_p0), uintptr(dwExStyle), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetParent(hWndChild HWND, hWndNewParent HWND) (r HWND, err error) {
	r0, _, e1 := syscall.Syscall(procSetParent.Addr(), 2, uintptr(hWndChild), uintptr(hWndNewParent), 0)
	r = HWND(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetWindowPos(hWnd HWND, hWndInsertAfter HWND, X int, Y int, cx int, cy int, uFlags UINT) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetWindowPos.Addr(), 7, uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(X), uintptr(Y), uintptr(cx), uintptr(cy), uintptr(uFlags), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func dialogBoxParamW(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogFunc DLGPROC, dwInitParam LPARAM) (r INT_PTR, err error) {
	r0, _, e1 := syscall.Syscall6(procDialogBoxParamW.Addr(), 5, uintptr(hInstance), uintptr(unsafe.Pointer(lpTemplateName)), uintptr(hWndParent), uintptr(lpDialogFunc), uintptr(dwInitParam), 0)
	r = INT_PTR(r0)
	if r == 0xffffffff {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func EndDialog(hDlg HWND, nResult INT_PTR) (err error) {
	r1, _, e1 := syscall.Syscall(procEndDialog.Addr(), 2, uintptr(hDlg), uintptr(nResult), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetFocus(hWnd HWND) (r HWND, err error) {
	r0, _, e1 := syscall.Syscall(procSetFocus.Addr(), 1, uintptr(hWnd), 0, 0)
	r = HWND(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindText(Arg1 *Findreplace) (r HWND) {
	r0, _, _ := syscall.Syscall(procFindTextW.Addr(), 1, uintptr(unsafe.Pointer(Arg1)), 0, 0)
	r = HWND(r0)
	return
}

func CommDlgExtendedError() (r DWORD) {
	r0, _, _ := syscall.Syscall(procCommDlgExtendedError.Addr(), 0, 0, 0, 0)
	r = DWORD(r0)
	return
}

func RegisterWindowMessage(lpString *uint16) (r UINT, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterWindowMessageW.Addr(), 1, uintptr(unsafe.Pointer(lpString)), 0, 0)
	r = UINT(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetWindowText(hWnd HWND, lpString *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetWindowTextW.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpString)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func IsWindow(hWnd HWND) (r bool) {
	r0, _, _ := syscall.Syscall(procIsWindow.Addr(), 1, uintptr(hWnd), 0, 0)
	r = r0 != 0
	return
}

func ClientToScreen(hWnd HWND, lpPoint *Point) (r bool) {
	r0, _, _ := syscall.Syscall(procClientToScreen.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)), 0)
	r = r0 != 0
	return
}

func ChoosePixelFormat(hdc HDC, ppfd *Pixelformatdescriptor) (r int, err error) {
	r0, _, e1 := syscall.Syscall(procChoosePixelFormat.Addr(), 2, uintptr(hdc), uintptr(unsafe.Pointer(ppfd)), 0)
	r = int(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetPixelFormat(hdc HDC, format int, ppfd *Pixelformatdescriptor) (err error) {
	r1, _, e1 := syscall.Syscall(procSetPixelFormat.Addr(), 3, uintptr(hdc), uintptr(format), uintptr(unsafe.Pointer(ppfd)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SwapBuffers(Arg1 HDC) (err error) {
	r1, _, e1 := syscall.Syscall(procSwapBuffers.Addr(), 1, uintptr(Arg1), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WglCreateContext(Arg1 HDC) (r HGLRC, err error) {
	r0, _, e1 := syscall.Syscall(procwglCreateContext.Addr(), 1, uintptr(Arg1), 0, 0)
	r = HGLRC(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WglMakeCurrent(arg1 HDC, arg2 HGLRC) (err error) {
	r1, _, e1 := syscall.Syscall(procwglMakeCurrent.Addr(), 2, uintptr(arg1), uintptr(arg2), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetDpiForWindow(hwnd HWND) (r UINT) {
	r0, _, _ := syscall.Syscall(procGetDpiForWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	r = UINT(r0)
	return
}

func CreateWindowEx(dwExStyle DWORD, lpClassName *uint16, lpWindowName *uint16, dwStyle DWORD, X int, Y int, nWidth int, nHeight int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) (r HWND, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(dwExStyle), uintptr(unsafe.Pointer(lpClassName)), uintptr(unsafe.Pointer(lpWindowName)), uintptr(dwStyle), uintptr(X), uintptr(Y), uintptr(nWidth), uintptr(nHeight), uintptr(hWndParent), uintptr(hMenu), uintptr(hInstance), uintptr(lpParam))
	r = HWND(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadIcon(hInstance HINSTANCE, lpIconName *uint16) (r HICON, err error) {
	r0, _, e1 := syscall.Syscall(procLoadIconW.Addr(), 2, uintptr(hInstance), uintptr(unsafe.Pointer(lpIconName)), 0)
	r = HICON(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadCursor(hInstance HINSTANCE, lpCursorName *uint16) (r HCURSOR, err error) {
	r0, _, e1 := syscall.Syscall(procLoadCursorW.Addr(), 2, uintptr(hInstance), uintptr(unsafe.Pointer(lpCursorName)), 0)
	r = HCURSOR(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func RegisterClassEx(Arg1 *Wndclassex) (r ATOM, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassExW.Addr(), 1, uintptr(unsafe.Pointer(Arg1)), 0, 0)
	r = ATOM(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetClientRect(hWnd HWND, lpRect *Rect) (err error) {
	r1, _, e1 := syscall.Syscall(procGetClientRect.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ShowWindow(hWnd HWND, nCmdShow int) (r bool) {
	r0, _, _ := syscall.Syscall(procShowWindow.Addr(), 2, uintptr(hWnd), uintptr(nCmdShow), 0)
	r = r0 != 0
	return
}

func UpdateWindow(hWnd HWND) (r bool) {
	r0, _, _ := syscall.Syscall(procUpdateWindow.Addr(), 1, uintptr(hWnd), 0, 0)
	r = r0 != 0
	return
}

func DefWindowProc(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (r LRESULT) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	r = LRESULT(r0)
	return
}

func EnableNonClientDpiScaling(hwnd HWND) (err error) {
	r1, _, e1 := syscall.Syscall(procEnableNonClientDpiScaling.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetDC(hWnd HWND) (r HDC) {
	r0, _, _ := syscall.Syscall(procGetDC.Addr(), 1, uintptr(hWnd), 0, 0)
	r = HDC(r0)
	return
}

func GetDeviceCaps(hdc HDC, index int) (r int) {
	r0, _, _ := syscall.Syscall(procGetDeviceCaps.Addr(), 2, uintptr(hdc), uintptr(index), 0)
	r = int(r0)
	return
}

func ReleaseDC(hWnd HWND, hDC HDC) (r int) {
	r0, _, _ := syscall.Syscall(procReleaseDC.Addr(), 2, uintptr(hWnd), uintptr(hDC), 0)
	r = int(r0)
	return
}

func GetProcessDpiAwareness(hprocess HANDLE, value *ProcessDpiAwareness) (r HRESULT) {
	r0, _, _ := syscall.Syscall(procGetProcessDpiAwareness.Addr(), 2, uintptr(hprocess), uintptr(unsafe.Pointer(value)), 0)
	r = HRESULT(r0)
	return
}

func CallWindowProc(lpPrevWndFunc WNDPROC, hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (r LRESULT) {
	r0, _, _ := syscall.Syscall6(procCallWindowProcW.Addr(), 5, uintptr(lpPrevWndFunc), uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0)
	r = LRESULT(r0)
	return
}

func EnableWindow(hWnd HWND, bEnable bool) (r bool) {
	var _p0 uint32
	if bEnable {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall(procEnableWindow.Addr(), 2, uintptr(hWnd), uintptr(_p0), 0)
	r = r0 != 0
	return
}

func SendMessage(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) (r LRESULT) {
	r0, _, _ := syscall.Syscall6(procSendMessageW.Addr(), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	r = LRESULT(r0)
	return
}

func BeginPaint(hWnd HWND, lpPaint *Paintstruct) (r HDC) {
	r0, _, _ := syscall.Syscall(procBeginPaint.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)), 0)
	r = HDC(r0)
	return
}

func EndPaint(hWnd HWND, lpPaint *Paintstruct) {
	syscall.Syscall(procEndPaint.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)), 0)
	return
}

func IsWindowEnabled(hWnd HWND) (r bool) {
	r0, _, _ := syscall.Syscall(procIsWindowEnabled.Addr(), 1, uintptr(hWnd), 0, 0)
	r = r0 != 0
	return
}

func IsWindowVisible(hWnd HWND) (r bool) {
	r0, _, _ := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hWnd), 0, 0)
	r = r0 != 0
	return
}

func SetMenu(hWnd HWND, hMenu HMENU) (err error) {
	r1, _, e1 := syscall.Syscall(procSetMenu.Addr(), 2, uintptr(hWnd), uintptr(hMenu), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func BeginDeferWindowPos(nNumWindows int) (r HDWP, err error) {
	r0, _, e1 := syscall.Syscall(procBeginDeferWindowPos.Addr(), 1, uintptr(nNumWindows), 0, 0)
	r = HDWP(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DeferWindowPos(hWinPosInfo HDWP, hWnd HWND, hWndInsertAfter HWND, x int, y int, cx int, cy int, uFlags UINT) (r HDWP, err error) {
	r0, _, e1 := syscall.Syscall9(procDeferWindowPos.Addr(), 8, uintptr(hWinPosInfo), uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags), 0)
	r = HDWP(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetMessageExtraInfo() (r LPARAM) {
	r0, _, _ := syscall.Syscall(procGetMessageExtraInfo.Addr(), 0, 0, 0, 0)
	r = LPARAM(r0)
	return
}

func GetMessageTime() (r LONG) {
	r0, _, _ := syscall.Syscall(procGetMessageTime.Addr(), 0, 0, 0, 0)
	r = LONG(r0)
	return
}

func GetSystemMetrics(nIndex int) (r int) {
	r0, _, _ := syscall.Syscall(procGetSystemMetrics.Addr(), 1, uintptr(nIndex), 0, 0)
	r = int(r0)
	return
}

func GetDoubleClickTime() (r UINT) {
	r0, _, _ := syscall.Syscall(procGetDoubleClickTime.Addr(), 0, 0, 0, 0)
	r = UINT(r0)
	return
}

func SetCapture(hWnd HWND) (r HWND) {
	r0, _, _ := syscall.Syscall(procSetCapture.Addr(), 1, uintptr(hWnd), 0, 0)
	r = HWND(r0)
	return
}

func GetKeyState(nVirtKey int) (r SHORT) {
	r0, _, _ := syscall.Syscall(procGetKeyState.Addr(), 1, uintptr(nVirtKey), 0, 0)
	r = SHORT(r0)
	return
}

func GetCapture() (r HWND) {
	r0, _, _ := syscall.Syscall(procGetCapture.Addr(), 0, 0, 0, 0)
	r = HWND(r0)
	return
}

func ReleaseCapture() (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseCapture.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func QueryPerformanceCounter(lpPerformanceCount *LARGE_INTEGER) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(lpPerformanceCount)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func QueryPerformanceFrequency(lpFrequency *LARGE_INTEGER) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryPerformanceFrequency.Addr(), 1, uintptr(unsafe.Pointer(lpFrequency)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func TrackMouseEvent(lpEventTrack *Trackmouseevent) (err error) {
	r1, _, e1 := syscall.Syscall(procTrackMouseEvent.Addr(), 1, uintptr(unsafe.Pointer(lpEventTrack)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetCursorPos(lpPoint *Point) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCursorPos.Addr(), 1, uintptr(unsafe.Pointer(lpPoint)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetKeyboardLayout(idThread DWORD) (r HKL) {
	r0, _, _ := syscall.Syscall(procGetKeyboardLayout.Addr(), 1, uintptr(idThread), 0, 0)
	r = HKL(r0)
	return
}

func VkKeyScanEx(ch WCHAR, dwhkl HKL) (r SHORT) {
	r0, _, _ := syscall.Syscall(procVkKeyScanExW.Addr(), 2, uintptr(ch), uintptr(dwhkl), 0)
	r = SHORT(r0)
	return
}

func GetMenu(hWnd HWND) (r HMENU) {
	r0, _, _ := syscall.Syscall(procGetMenu.Addr(), 1, uintptr(hWnd), 0, 0)
	r = HMENU(r0)
	return
}

func GetSubMenu(hMenu HMENU, nPos int) (r HMENU) {
	r0, _, _ := syscall.Syscall(procGetSubMenu.Addr(), 2, uintptr(hMenu), uintptr(nPos), 0)
	r = HMENU(r0)
	return
}

func RemoveMenu(hMenu HMENU, uPosition UINT, uFlags UINT) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveMenu.Addr(), 3, uintptr(hMenu), uintptr(uPosition), uintptr(uFlags))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetCommandLine() (r *uint16) {
	r0, _, _ := syscall.Syscall(procGetCommandLineW.Addr(), 0, 0, 0, 0)
	r = (*uint16)(unsafe.Pointer(r0))
	return
}

func sHGetFolderPath(hwnd HWND, csidl int, hToken HANDLE, dwFlags DWORD, pszPath *uint16) (r HRESULT) {
	r0, _, _ := syscall.Syscall6(procSHGetFolderPathW.Addr(), 5, uintptr(hwnd), uintptr(csidl), uintptr(hToken), uintptr(dwFlags), uintptr(unsafe.Pointer(pszPath)), 0)
	r = HRESULT(r0)
	return
}

func sHGetKnownFolderPath(rfid REFKNOWNFOLDERID, dwFlags DWORD, hToken HANDLE, ppszPath *PWSTR) (r HRESULT) {
	r0, _, _ := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(rfid), uintptr(dwFlags), uintptr(hToken), uintptr(unsafe.Pointer(ppszPath)), 0, 0)
	r = HRESULT(r0)
	return
}

func CoTaskMemFree(pv LPVOID) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(pv), 0, 0)
	return
}

func FindResource(hModule HMODULE, lpName *uint16, lpType *uint16) (r HRSRC, err error) {
	r0, _, e1 := syscall.Syscall(procFindResourceW.Addr(), 3, uintptr(hModule), uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpType)))
	r = HRSRC(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) (r HGLOBAL, err error) {
	r0, _, e1 := syscall.Syscall(procLoadResource.Addr(), 2, uintptr(hModule), uintptr(hResInfo), 0)
	r = HGLOBAL(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) (r DWORD, err error) {
	r0, _, e1 := syscall.Syscall(procSizeofResource.Addr(), 2, uintptr(hModule), uintptr(hResInfo), 0)
	r = DWORD(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LockResource(hResData HGLOBAL) (r LPVOID, err error) {
	r0, _, e1 := syscall.Syscall(procLockResource.Addr(), 1, uintptr(hResData), 0, 0)
	r = LPVOID(r0)
	if r == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
