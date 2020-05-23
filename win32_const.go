package win32api

import (
	"unsafe"
)

// #include <windows.h>
// #include "shellscalingapi.h"
import "C"

// index of GetDevoceCaps
const (
	Logpixelsx = C.LOGPIXELSX
	Logpixelsy = C.LOGPIXELSY
)

const (
	BsPushbutton      = C.BS_PUSHBUTTON
	BsDefpushbutton   = C.BS_DEFPUSHBUTTON
	BsCheckbox        = C.BS_CHECKBOX
	BsAutocheckbox    = C.BS_AUTOCHECKBOX
	BsRadiobutton     = C.BS_RADIOBUTTON
	Bs3state          = C.BS_3STATE
	BsAuto3state      = C.BS_AUTO3STATE
	BsGroupbox        = C.BS_GROUPBOX
	BsUserbutton      = C.BS_USERBUTTON
	BsAutoradiobutton = C.BS_AUTORADIOBUTTON
	BsPushbox         = C.BS_PUSHBOX
	BsOwnerdraw       = C.BS_OWNERDRAW
	BsTypemask        = C.BS_TYPEMASK
	BsLefttext        = C.BS_LEFTTEXT
	BsText            = C.BS_TEXT
	BsIcon            = C.BS_ICON
	BsBitmap          = C.BS_BITMAP
	BsLeft            = C.BS_LEFT
	BsRight           = C.BS_RIGHT
	BsCenter          = C.BS_CENTER
	BsTop             = C.BS_TOP
	BsBottom          = C.BS_BOTTOM
	BsVcenter         = C.BS_VCENTER
	BsPushlike        = C.BS_PUSHLIKE
	BsMultiline       = C.BS_MULTILINE
	BsNotify          = C.BS_NOTIFY
	BsFlat            = C.BS_FLAT
	BsRightbutton     = C.BS_RIGHTBUTTON
)

const (
	CsVredraw         = C.CS_VREDRAW
	CsHredraw         = C.CS_HREDRAW
	CsDblclks         = C.CS_DBLCLKS
	CsOwndc           = C.CS_OWNDC
	CsClassdc         = C.CS_CLASSDC
	CsParentdc        = C.CS_PARENTDC
	CsNoclose         = C.CS_NOCLOSE
	CsSavebits        = C.CS_SAVEBITS
	CsBytealignclient = C.CS_BYTEALIGNCLIENT
	CsBytealignwindow = C.CS_BYTEALIGNWINDOW
	CsGlobalclass     = C.CS_GLOBALCLASS
	CsIme             = C.CS_IME
	CsDropshadow      = C.CS_DROPSHADOW
)

const (
	CwUsedefault = C.CW_USEDEFAULT
)

const (
	GwlpWndproc    = C.GWLP_WNDPROC
	GwlpHinstance  = C.GWLP_HINSTANCE
	GwlpHwndparent = C.GWLP_HWNDPARENT
	GwlpUserdata   = C.GWLP_USERDATA
	GwlpId         = C.GWLP_ID
)

var (
	IdcArrow       = (*uint16)(unsafe.Pointer(C.IDC_ARROW))
	IdcIbeam       = (*uint16)(unsafe.Pointer(C.IDC_IBEAM))
	IdcWait        = (*uint16)(unsafe.Pointer(C.IDC_WAIT))
	IdcCross       = (*uint16)(unsafe.Pointer(C.IDC_CROSS))
	IdcUparrow     = (*uint16)(unsafe.Pointer(C.IDC_UPARROW))
	IdcSize        = (*uint16)(unsafe.Pointer(C.IDC_SIZE))
	IdcIcon        = (*uint16)(unsafe.Pointer(C.IDC_ICON))
	IdcSizenwse    = (*uint16)(unsafe.Pointer(C.IDC_SIZENWSE))
	IdcSizenesw    = (*uint16)(unsafe.Pointer(C.IDC_SIZENESW))
	IdcSizewe      = (*uint16)(unsafe.Pointer(C.IDC_SIZEWE))
	IdcSizens      = (*uint16)(unsafe.Pointer(C.IDC_SIZENS))
	IdcSizeall     = (*uint16)(unsafe.Pointer(C.IDC_SIZEALL))
	IdcNo          = (*uint16)(unsafe.Pointer(C.IDC_NO))
	IdcHand        = (*uint16)(unsafe.Pointer(C.IDC_HAND))
	IdcAppstarting = (*uint16)(unsafe.Pointer(C.IDC_APPSTARTING))
	IdcHelp        = (*uint16)(unsafe.Pointer(C.IDC_HELP))
)

const (
	WmNull                   = C.WM_NULL
	WmCreate                 = C.WM_CREATE
	WmDestroy                = C.WM_DESTROY
	WmMove                   = C.WM_MOVE
	WmSize                   = C.WM_SIZE
	WmActivate               = C.WM_ACTIVATE
	WmSetfocus               = C.WM_SETFOCUS
	WmKillfocus              = C.WM_KILLFOCUS
	WmEnable                 = C.WM_ENABLE
	WmSetredraw              = C.WM_SETREDRAW
	WmSettext                = C.WM_SETTEXT
	WmGettext                = C.WM_GETTEXT
	WmGettextlength          = C.WM_GETTEXTLENGTH
	WmPaint                  = C.WM_PAINT
	WmClose                  = C.WM_CLOSE
	WmQueryendsession        = C.WM_QUERYENDSESSION
	WmQueryopen              = C.WM_QUERYOPEN
	WmEndsession             = C.WM_ENDSESSION
	WmQuit                   = C.WM_QUIT
	WmErasebkgnd             = C.WM_ERASEBKGND
	WmSyscolorchange         = C.WM_SYSCOLORCHANGE
	WmShowwindow             = C.WM_SHOWWINDOW
	WmWininichange           = C.WM_WININICHANGE
	WmSettingchange          = C.WM_SETTINGCHANGE
	WmDevmodechange          = C.WM_DEVMODECHANGE
	WmActivateapp            = C.WM_ACTIVATEAPP
	WmFontchange             = C.WM_FONTCHANGE
	WmTimechange             = C.WM_TIMECHANGE
	WmCancelmode             = C.WM_CANCELMODE
	WmSetcursor              = C.WM_SETCURSOR
	WmMouseactivate          = C.WM_MOUSEACTIVATE
	WmChildactivate          = C.WM_CHILDACTIVATE
	WmQueuesync              = C.WM_QUEUESYNC
	WmGetminmaxinfo          = C.WM_GETMINMAXINFO
	WmPainticon              = C.WM_PAINTICON
	WmIconerasebkgnd         = C.WM_ICONERASEBKGND
	WmNextdlgctl             = C.WM_NEXTDLGCTL
	WmSpoolerstatus          = C.WM_SPOOLERSTATUS
	WmDrawitem               = C.WM_DRAWITEM
	WmMeasureitem            = C.WM_MEASUREITEM
	WmDeleteitem             = C.WM_DELETEITEM
	WmVkeytoitem             = C.WM_VKEYTOITEM
	WmChartoitem             = C.WM_CHARTOITEM
	WmSetfont                = C.WM_SETFONT
	WmGetfont                = C.WM_GETFONT
	WmSethotkey              = C.WM_SETHOTKEY
	WmGethotkey              = C.WM_GETHOTKEY
	WmQuerydragicon          = C.WM_QUERYDRAGICON
	WmCompareitem            = C.WM_COMPAREITEM
	WmGetobject              = C.WM_GETOBJECT
	WmCompacting             = C.WM_COMPACTING
	WmCommnotify             = C.WM_COMMNOTIFY
	WmWindowposchanging      = C.WM_WINDOWPOSCHANGING
	WmWindowposchanged       = C.WM_WINDOWPOSCHANGED
	WmPower                  = C.WM_POWER
	WmCopydata               = C.WM_COPYDATA
	WmCanceljournal          = C.WM_CANCELJOURNAL
	WmNotify                 = C.WM_NOTIFY
	WmInputlangchangerequest = C.WM_INPUTLANGCHANGEREQUEST
	WmInputlangchange        = C.WM_INPUTLANGCHANGE
	WmTcard                  = C.WM_TCARD
	WmHelp                   = C.WM_HELP
	WmUserchanged            = C.WM_USERCHANGED
	WmNotifyformat           = C.WM_NOTIFYFORMAT
	WmContextmenu            = C.WM_CONTEXTMENU
	WmStylechanging          = C.WM_STYLECHANGING
	WmStylechanged           = C.WM_STYLECHANGED
	WmDisplaychange          = C.WM_DISPLAYCHANGE
	WmGeticon                = C.WM_GETICON
	WmSeticon                = C.WM_SETICON
	WmNccreate               = C.WM_NCCREATE
	WmNcdestroy              = C.WM_NCDESTROY
	WmNccalcsize             = C.WM_NCCALCSIZE
	WmNchittest              = C.WM_NCHITTEST
	WmNcpaint                = C.WM_NCPAINT
	WmNcactivate             = C.WM_NCACTIVATE
	WmGetdlgcode             = C.WM_GETDLGCODE
	WmSyncpaint              = C.WM_SYNCPAINT
	WmNcmousemove            = C.WM_NCMOUSEMOVE
	WmNclbuttondown          = C.WM_NCLBUTTONDOWN
	WmNclbuttonup            = C.WM_NCLBUTTONUP
	WmNclbuttondblclk        = C.WM_NCLBUTTONDBLCLK
	WmNcrbuttondown          = C.WM_NCRBUTTONDOWN
	WmNcrbuttonup            = C.WM_NCRBUTTONUP
	WmNcrbuttondblclk        = C.WM_NCRBUTTONDBLCLK
	WmNcmbuttondown          = C.WM_NCMBUTTONDOWN
	WmNcmbuttonup            = C.WM_NCMBUTTONUP
	WmNcmbuttondblclk        = C.WM_NCMBUTTONDBLCLK
	WmNcxbuttondown          = C.WM_NCXBUTTONDOWN
	WmNcxbuttonup            = C.WM_NCXBUTTONUP
	WmNcxbuttondblclk        = C.WM_NCXBUTTONDBLCLK
	WmInputDeviceChange      = C.WM_INPUT_DEVICE_CHANGE
	WmInput                  = C.WM_INPUT
	WmKeyfirst               = C.WM_KEYFIRST
	WmKeydown                = C.WM_KEYDOWN
	WmKeyup                  = C.WM_KEYUP
	WmChar                   = C.WM_CHAR
	WmDeadchar               = C.WM_DEADCHAR
	WmSyskeydown             = C.WM_SYSKEYDOWN
	WmSyskeyup               = C.WM_SYSKEYUP
	WmSyschar                = C.WM_SYSCHAR
	WmSysdeadchar            = C.WM_SYSDEADCHAR
	WmUnichar                = C.WM_UNICHAR
	WmKeylast                = C.WM_KEYLAST
	WmImeStartcomposition    = C.WM_IME_STARTCOMPOSITION
	WmImeEndcomposition      = C.WM_IME_ENDCOMPOSITION
	WmImeComposition         = C.WM_IME_COMPOSITION
	WmImeKeylast             = C.WM_IME_KEYLAST
	WmInitdialog             = C.WM_INITDIALOG
	WmCommand                = C.WM_COMMAND
	WmSyscommand             = C.WM_SYSCOMMAND
	WmTimer                  = C.WM_TIMER
	WmHscroll                = C.WM_HSCROLL
	WmVscroll                = C.WM_VSCROLL
	WmInitmenu               = C.WM_INITMENU
	WmInitmenupopup          = C.WM_INITMENUPOPUP
	WmMenuselect             = C.WM_MENUSELECT
	WmGesture                = C.WM_GESTURE
	WmGesturenotify          = C.WM_GESTURENOTIFY
	WmMenuchar               = C.WM_MENUCHAR
	WmEnteridle              = C.WM_ENTERIDLE
	WmMenurbuttonup          = C.WM_MENURBUTTONUP
	WmMenudrag               = C.WM_MENUDRAG
	WmMenugetobject          = C.WM_MENUGETOBJECT
	WmUninitmenupopup        = C.WM_UNINITMENUPOPUP
	WmMenucommand            = C.WM_MENUCOMMAND
	WmChangeuistate          = C.WM_CHANGEUISTATE
	WmUpdateuistate          = C.WM_UPDATEUISTATE
	WmQueryuistate           = C.WM_QUERYUISTATE
	WmCtlcolormsgbox         = C.WM_CTLCOLORMSGBOX
	WmCtlcoloredit           = C.WM_CTLCOLOREDIT
	WmCtlcolorlistbox        = C.WM_CTLCOLORLISTBOX
	WmCtlcolorbtn            = C.WM_CTLCOLORBTN
	WmCtlcolordlg            = C.WM_CTLCOLORDLG
	WmCtlcolorscrollbar      = C.WM_CTLCOLORSCROLLBAR
	WmCtlcolorstatic         = C.WM_CTLCOLORSTATIC
	WmMousefirst             = C.WM_MOUSEFIRST
	WmMousemove              = C.WM_MOUSEMOVE
	WmLbuttondown            = C.WM_LBUTTONDOWN
	WmLbuttonup              = C.WM_LBUTTONUP
	WmLbuttondblclk          = C.WM_LBUTTONDBLCLK
	WmRbuttondown            = C.WM_RBUTTONDOWN
	WmRbuttonup              = C.WM_RBUTTONUP
	WmRbuttondblclk          = C.WM_RBUTTONDBLCLK
	WmMbuttondown            = C.WM_MBUTTONDOWN
	WmMbuttonup              = C.WM_MBUTTONUP
	WmMbuttondblclk          = C.WM_MBUTTONDBLCLK
	WmMousewheel             = C.WM_MOUSEWHEEL
	WmXbuttondown            = C.WM_XBUTTONDOWN
	WmXbuttonup              = C.WM_XBUTTONUP
	WmXbuttondblclk          = C.WM_XBUTTONDBLCLK
	WmMousehwheel            = C.WM_MOUSEHWHEEL
	WmMouselast              = C.WM_MOUSELAST
	WmParentnotify           = C.WM_PARENTNOTIFY
	WmEntermenuloop          = C.WM_ENTERMENULOOP
	WmExitmenuloop           = C.WM_EXITMENULOOP
	WmNextmenu               = C.WM_NEXTMENU
	WmSizing                 = C.WM_SIZING
	WmCapturechanged         = C.WM_CAPTURECHANGED
	WmMoving                 = C.WM_MOVING
	WmPowerbroadcast         = C.WM_POWERBROADCAST
	WmDevicechange           = C.WM_DEVICECHANGE
	WmMdicreate              = C.WM_MDICREATE
	WmMdidestroy             = C.WM_MDIDESTROY
	WmMdiactivate            = C.WM_MDIACTIVATE
	WmMdirestore             = C.WM_MDIRESTORE
	WmMdinext                = C.WM_MDINEXT
	WmMdimaximize            = C.WM_MDIMAXIMIZE
	WmMditile                = C.WM_MDITILE
	WmMdicascade             = C.WM_MDICASCADE
	WmMdiiconarrange         = C.WM_MDIICONARRANGE
	WmMdigetactive           = C.WM_MDIGETACTIVE
	WmMdisetmenu             = C.WM_MDISETMENU
	WmEntersizemove          = C.WM_ENTERSIZEMOVE
	WmExitsizemove           = C.WM_EXITSIZEMOVE
	WmDropfiles              = C.WM_DROPFILES
	WmMdirefreshmenu         = C.WM_MDIREFRESHMENU
	// WmPointerdevicechange = C.WM_POINTERDEVICECHANGE // #if WINVER >= 0x0602
	// WmPointerdeviceinrange = C.WM_POINTERDEVICEINRANGE
	// WmPointerdeviceoutofrange = C.WM_POINTERDEVICEOUTOFRANGE
	WmTouch = C.WM_TOUCH
	// WmNcpointerupdate = C.WM_NCPOINTERUPDATE  	// #if WINVER >= 0x0602
	// WmNcpointerdown = C.WM_NCPOINTERDOWN
	// WmNcpointerup = C.WM_NCPOINTERUP
	// WmPointerupdate = C.WM_POINTERUPDATE
	// WmPointerdown = C.WM_POINTERDOWN
	// WmPointerup = C.WM_POINTERUP
	// WmPointerenter = C.WM_POINTERENTER
	// WmPointerleave = C.WM_POINTERLEAVE
	// WmPointeractivate = C.WM_POINTERACTIVATE
	// WmPointercapturechanged = C.WM_POINTERCAPTURECHANGED
	// WmTouchhittesting = C.WM_TOUCHHITTESTING
	// WmPointerwheel = C.WM_POINTERWHEEL
	// WmPointerhwheel = C.WM_POINTERHWHEEL
	WmImeSetcontext                  = C.WM_IME_SETCONTEXT
	WmImeNotify                      = C.WM_IME_NOTIFY
	WmImeControl                     = C.WM_IME_CONTROL
	WmImeCompositionfull             = C.WM_IME_COMPOSITIONFULL
	WmImeSelect                      = C.WM_IME_SELECT
	WmImeChar                        = C.WM_IME_CHAR
	WmImeRequest                     = C.WM_IME_REQUEST
	WmImeKeydown                     = C.WM_IME_KEYDOWN
	WmImeKeyup                       = C.WM_IME_KEYUP
	WmMousehover                     = C.WM_MOUSEHOVER
	WmMouseleave                     = C.WM_MOUSELEAVE
	WmNcmousehover                   = C.WM_NCMOUSEHOVER
	WmNcmouseleave                   = C.WM_NCMOUSELEAVE
	WmWtssessionChange               = C.WM_WTSSESSION_CHANGE
	WmTabletFirst                    = C.WM_TABLET_FIRST
	WmTabletLast                     = C.WM_TABLET_LAST
	WmDpichanged                     = C.WM_DPICHANGED
	WmCut                            = C.WM_CUT
	WmCopy                           = C.WM_COPY
	WmPaste                          = C.WM_PASTE
	WmClear                          = C.WM_CLEAR
	WmUndo                           = C.WM_UNDO
	WmRenderformat                   = C.WM_RENDERFORMAT
	WmRenderallformats               = C.WM_RENDERALLFORMATS
	WmDestroyclipboard               = C.WM_DESTROYCLIPBOARD
	WmDrawclipboard                  = C.WM_DRAWCLIPBOARD
	WmPaintclipboard                 = C.WM_PAINTCLIPBOARD
	WmVscrollclipboard               = C.WM_VSCROLLCLIPBOARD
	WmSizeclipboard                  = C.WM_SIZECLIPBOARD
	WmAskcbformatname                = C.WM_ASKCBFORMATNAME
	WmChangecbchain                  = C.WM_CHANGECBCHAIN
	WmHscrollclipboard               = C.WM_HSCROLLCLIPBOARD
	WmQuerynewpalette                = C.WM_QUERYNEWPALETTE
	WmPaletteischanging              = C.WM_PALETTEISCHANGING
	WmPalettechanged                 = C.WM_PALETTECHANGED
	WmHotkey                         = C.WM_HOTKEY
	WmPrint                          = C.WM_PRINT
	WmPrintclient                    = C.WM_PRINTCLIENT
	WmAppcommand                     = C.WM_APPCOMMAND
	WmThemechanged                   = C.WM_THEMECHANGED
	WmClipboardupdate                = C.WM_CLIPBOARDUPDATE
	WmDwmcompositionchanged          = C.WM_DWMCOMPOSITIONCHANGED
	WmDwmncrenderingchanged          = C.WM_DWMNCRENDERINGCHANGED
	WmDwmcolorizationcolorchanged    = C.WM_DWMCOLORIZATIONCOLORCHANGED
	WmDwmwindowmaximizedchange       = C.WM_DWMWINDOWMAXIMIZEDCHANGE
	WmDwmsendiconicthumbnail         = C.WM_DWMSENDICONICTHUMBNAIL
	WmDwmsendiconiclivepreviewbitmap = C.WM_DWMSENDICONICLIVEPREVIEWBITMAP
	WmGettitlebarinfoex              = C.WM_GETTITLEBARINFOEX
	WmHandheldfirst                  = C.WM_HANDHELDFIRST
	WmHandheldlast                   = C.WM_HANDHELDLAST
	WmAfxfirst                       = C.WM_AFXFIRST
	WmAfxlast                        = C.WM_AFXLAST
	WmPenwinfirst                    = C.WM_PENWINFIRST
	WmPenwinlast                     = C.WM_PENWINLAST
	WmApp                            = C.WM_APP
	WmUser                           = C.WM_USER
)

const (
	SwHide            = C.SW_HIDE
	SwShownormal      = C.SW_SHOWNORMAL
	SwNormal          = C.SW_NORMAL
	SwShowminimized   = C.SW_SHOWMINIMIZED
	SwShowmaximized   = C.SW_SHOWMAXIMIZED
	SwMaximize        = C.SW_MAXIMIZE
	SwShownoactivate  = C.SW_SHOWNOACTIVATE
	SwShow            = C.SW_SHOW
	SwMinimize        = C.SW_MINIMIZE
	SwShowminnoactive = C.SW_SHOWMINNOACTIVE
	SwShowna          = C.SW_SHOWNA
	SwRestore         = C.SW_RESTORE
	SwShowdefault     = C.SW_SHOWDEFAULT
	SwForceminimize   = C.SW_FORCEMINIMIZE
	SwMax             = C.SW_MAX
	SwParentclosing   = C.SW_PARENTCLOSING
	SwOtherzoom       = C.SW_OTHERZOOM
	SwParentopening   = C.SW_PARENTOPENING
	SwOtherunzoom     = C.SW_OTHERUNZOOM
	SwScrollchildren  = C.SW_SCROLLCHILDREN
	SwInvalidate      = C.SW_INVALIDATE
	SwErase           = C.SW_ERASE
	SwSmoothscroll    = C.SW_SMOOTHSCROLL
)

const (
	WsOverlapped         = C.WS_OVERLAPPED
	WsPopup              = C.WS_POPUP
	WsChild              = C.WS_CHILD
	WsMinimize           = C.WS_MINIMIZE
	WsVisible            = C.WS_VISIBLE
	WsDisabled           = C.WS_DISABLED
	WsClipsiblings       = C.WS_CLIPSIBLINGS
	WsClipchildren       = C.WS_CLIPCHILDREN
	WsMaximize           = C.WS_MAXIMIZE
	WsCaption            = C.WS_CAPTION
	WsBorder             = C.WS_BORDER
	WsDlgframe           = C.WS_DLGFRAME
	WsVscroll            = C.WS_VSCROLL
	WsHscroll            = C.WS_HSCROLL
	WsSysmenu            = C.WS_SYSMENU
	WsThickframe         = C.WS_THICKFRAME
	WsGroup              = C.WS_GROUP
	WsTabstop            = C.WS_TABSTOP
	WsMinimizebox        = C.WS_MINIMIZEBOX
	WsMaximizebox        = C.WS_MAXIMIZEBOX
	WsTiled              = C.WS_TILED
	WsIconic             = C.WS_ICONIC
	WsSizebox            = C.WS_SIZEBOX
	WsTiledwindow        = C.WS_TILEDWINDOW
	WsOverlappedwindow   = C.WS_OVERLAPPEDWINDOW
	WsPopupwindow        = C.WS_POPUPWINDOW
	WsChildwindow        = C.WS_CHILDWINDOW
	WsExDlgmodalframe    = C.WS_EX_DLGMODALFRAME
	WsExNoparentnotify   = C.WS_EX_NOPARENTNOTIFY
	WsExTopmost          = C.WS_EX_TOPMOST
	WsExAcceptfiles      = C.WS_EX_ACCEPTFILES
	WsExTransparent      = C.WS_EX_TRANSPARENT
	WsExMdichild         = C.WS_EX_MDICHILD
	WsExToolwindow       = C.WS_EX_TOOLWINDOW
	WsExWindowedge       = C.WS_EX_WINDOWEDGE
	WsExClientedge       = C.WS_EX_CLIENTEDGE
	WsExContexthelp      = C.WS_EX_CONTEXTHELP
	WsExRight            = C.WS_EX_RIGHT
	WsExLeft             = C.WS_EX_LEFT
	WsExRtlreading       = C.WS_EX_RTLREADING
	WsExLtrreading       = C.WS_EX_LTRREADING
	WsExLeftscrollbar    = C.WS_EX_LEFTSCROLLBAR
	WsExRightscrollbar   = C.WS_EX_RIGHTSCROLLBAR
	WsExControlparent    = C.WS_EX_CONTROLPARENT
	WsExStaticedge       = C.WS_EX_STATICEDGE
	WsExAppwindow        = C.WS_EX_APPWINDOW
	WsExOverlappedwindow = C.WS_EX_OVERLAPPEDWINDOW
	WsExPalettewindow    = C.WS_EX_PALETTEWINDOW
	WsExLayered          = C.WS_EX_LAYERED
	WsExNoinheritlayout  = C.WS_EX_NOINHERITLAYOUT
	// WsExNoredirectionbitmap = C.WS_EX_NOREDIRECTIONBITMAP #if WINVER >= 0x0602
	WsExLayoutrtl   = C.WS_EX_LAYOUTRTL
	WsExComposited  = C.WS_EX_COMPOSITED
	WsExNoactivate  = C.WS_EX_NOACTIVATE
	WsActivecaption = C.WS_ACTIVECAPTION
)

type ProcessDpiAwareness int

const (
	ProcessDpiUnaware         ProcessDpiAwareness = C.PROCESS_DPI_UNAWARE
	ProcessSystemDpiAware     ProcessDpiAwareness = C.PROCESS_SYSTEM_DPI_AWARE
	ProcessPerMonitorDpiAware ProcessDpiAwareness = C.PROCESS_PER_MONITOR_DPI_AWARE
)

const (
	SOk = C.S_OK
)
