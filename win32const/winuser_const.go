package win32const

import (
	"unsafe"

	"github.com/turutcrane/win32api"
)

// #include <windef.h>
// #include <winbase.h>
// #include <winuser.h>
// #include "shellscalingapi.h"
import "C"

const (
	sizeof_RECT          = C.sizeof_RECT
	sizeof_WNDCLASSEX    = C.sizeof_WNDCLASSEX
	sizeof_CREATESTRUCTW = C.sizeof_CREATESTRUCTW
	sizeof_PAINTSTRUCT   = C.sizeof_PAINTSTRUCT
)

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
	EsLeft        = C.ES_LEFT
	EsCenter      = C.ES_CENTER
	EsRight       = C.ES_RIGHT
	EsMultiline   = C.ES_MULTILINE
	EsUppercase   = C.ES_UPPERCASE
	EsLowercase   = C.ES_LOWERCASE
	EsPassword    = C.ES_PASSWORD
	EsAutovscroll = C.ES_AUTOVSCROLL
	EsAutohscroll = C.ES_AUTOHSCROLL
	EsNohidesel   = C.ES_NOHIDESEL
	EsOemconvert  = C.ES_OEMCONVERT
	EsReadonly    = C.ES_READONLY
	EsWantreturn  = C.ES_WANTRETURN
	EsNumber      = C.ES_NUMBER
)

const (
	EmGetsel              = C.EM_GETSEL
	EmSetsel              = C.EM_SETSEL
	EmGetrect             = C.EM_GETRECT
	EmSetrect             = C.EM_SETRECT
	EmSetrectnp           = C.EM_SETRECTNP
	EmScroll              = C.EM_SCROLL
	EmLinescroll          = C.EM_LINESCROLL
	EmScrollcaret         = C.EM_SCROLLCARET
	EmGetmodify           = C.EM_GETMODIFY
	EmSetmodify           = C.EM_SETMODIFY
	EmGetlinecount        = C.EM_GETLINECOUNT
	EmLineindex           = C.EM_LINEINDEX
	EmSethandle           = C.EM_SETHANDLE
	EmGethandle           = C.EM_GETHANDLE
	EmGetthumb            = C.EM_GETTHUMB
	EmLinelength          = C.EM_LINELENGTH
	EmReplacesel          = C.EM_REPLACESEL
	EmGetline             = C.EM_GETLINE
	EmLimittext           = C.EM_LIMITTEXT
	EmCanundo             = C.EM_CANUNDO
	EmUndo                = C.EM_UNDO
	EmFmtlines            = C.EM_FMTLINES
	EmLinefromchar        = C.EM_LINEFROMCHAR
	EmSettabstops         = C.EM_SETTABSTOPS
	EmSetpasswordchar     = C.EM_SETPASSWORDCHAR
	EmEmptyundobuffer     = C.EM_EMPTYUNDOBUFFER
	EmGetfirstvisibleline = C.EM_GETFIRSTVISIBLELINE
	EmSetreadonly         = C.EM_SETREADONLY
	EmSetwordbreakproc    = C.EM_SETWORDBREAKPROC
	EmGetwordbreakproc    = C.EM_GETWORDBREAKPROC
	EmGetpasswordchar     = C.EM_GETPASSWORDCHAR
	EmSetmargins          = C.EM_SETMARGINS
	EmGetmargins          = C.EM_GETMARGINS
	EmSetlimittext        = C.EM_SETLIMITTEXT
	EmGetlimittext        = C.EM_GETLIMITTEXT
	EmPosfromchar         = C.EM_POSFROMCHAR
	EmCharfrompos         = C.EM_CHARFROMPOS
	EmSetimestatus        = C.EM_SETIMESTATUS
	EmGetimestatus        = C.EM_GETIMESTATUS
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
	SizeRestored  = C.SIZE_RESTORED
	SizeMinimized = C.SIZE_MINIMIZED
	SizeMaximized = C.SIZE_MAXIMIZED
	SizeMaxshow   = C.SIZE_MAXSHOW
	SizeMaxhide   = C.SIZE_MAXHIDE
)

const (
	SwpNosize         = C.SWP_NOSIZE
	SwpNomove         = C.SWP_NOMOVE
	SwpNozorder       = C.SWP_NOZORDER
	SwpNoredraw       = C.SWP_NOREDRAW
	SwpNoactivate     = C.SWP_NOACTIVATE
	SwpFramechanged   = C.SWP_FRAMECHANGED
	SwpShowwindow     = C.SWP_SHOWWINDOW
	SwpHidewindow     = C.SWP_HIDEWINDOW
	SwpNocopybits     = C.SWP_NOCOPYBITS
	SwpNoownerzorder  = C.SWP_NOOWNERZORDER
	SwpNosendchanging = C.SWP_NOSENDCHANGING
	SwpDrawframe      = C.SWP_DRAWFRAME
	SwpNoreposition   = C.SWP_NOREPOSITION
	SwpDefererase     = C.SWP_DEFERERASE
	SwpAsyncwindowpos = C.SWP_ASYNCWINDOWPOS
)

const (
	VkLbutton           = C.VK_LBUTTON
	VkRbutton           = C.VK_RBUTTON
	VkCancel            = C.VK_CANCEL
	VkMbutton           = C.VK_MBUTTON
	VkXbutton1          = C.VK_XBUTTON1
	VkXbutton2          = C.VK_XBUTTON2
	VkBack              = C.VK_BACK
	VkTab               = C.VK_TAB
	VkClear             = C.VK_CLEAR
	VkReturn            = C.VK_RETURN
	VkShift             = C.VK_SHIFT
	VkControl           = C.VK_CONTROL
	VkMenu              = C.VK_MENU
	VkPause             = C.VK_PAUSE
	VkCapital           = C.VK_CAPITAL
	VkKana              = C.VK_KANA
	VkHangeul           = C.VK_HANGEUL
	VkHangul            = C.VK_HANGUL
	VkJunja             = C.VK_JUNJA
	VkFinal             = C.VK_FINAL
	VkHanja             = C.VK_HANJA
	VkKanji             = C.VK_KANJI
	VkEscape            = C.VK_ESCAPE
	VkConvert           = C.VK_CONVERT
	VkNonconvert        = C.VK_NONCONVERT
	VkAccept            = C.VK_ACCEPT
	VkModechange        = C.VK_MODECHANGE
	VkSpace             = C.VK_SPACE
	VkPrior             = C.VK_PRIOR
	VkNext              = C.VK_NEXT
	VkEnd               = C.VK_END
	VkHome              = C.VK_HOME
	VkLeft              = C.VK_LEFT
	VkUp                = C.VK_UP
	VkRight             = C.VK_RIGHT
	VkDown              = C.VK_DOWN
	VkSelect            = C.VK_SELECT
	VkPrint             = C.VK_PRINT
	VkExecute           = C.VK_EXECUTE
	VkSnapshot          = C.VK_SNAPSHOT
	VkInsert            = C.VK_INSERT
	VkDelete            = C.VK_DELETE
	VkHelp              = C.VK_HELP
	VkLwin              = C.VK_LWIN
	VkRwin              = C.VK_RWIN
	VkApps              = C.VK_APPS
	VkSleep             = C.VK_SLEEP
	VkNumpad0           = C.VK_NUMPAD0
	VkNumpad1           = C.VK_NUMPAD1
	VkNumpad2           = C.VK_NUMPAD2
	VkNumpad3           = C.VK_NUMPAD3
	VkNumpad4           = C.VK_NUMPAD4
	VkNumpad5           = C.VK_NUMPAD5
	VkNumpad6           = C.VK_NUMPAD6
	VkNumpad7           = C.VK_NUMPAD7
	VkNumpad8           = C.VK_NUMPAD8
	VkNumpad9           = C.VK_NUMPAD9
	VkMultiply          = C.VK_MULTIPLY
	VkAdd               = C.VK_ADD
	VkSeparator         = C.VK_SEPARATOR
	VkSubtract          = C.VK_SUBTRACT
	VkDecimal           = C.VK_DECIMAL
	VkDivide            = C.VK_DIVIDE
	VkF1                = C.VK_F1
	VkF2                = C.VK_F2
	VkF3                = C.VK_F3
	VkF4                = C.VK_F4
	VkF5                = C.VK_F5
	VkF6                = C.VK_F6
	VkF7                = C.VK_F7
	VkF8                = C.VK_F8
	VkF9                = C.VK_F9
	VkF10               = C.VK_F10
	VkF11               = C.VK_F11
	VkF12               = C.VK_F12
	VkF13               = C.VK_F13
	VkF14               = C.VK_F14
	VkF15               = C.VK_F15
	VkF16               = C.VK_F16
	VkF17               = C.VK_F17
	VkF18               = C.VK_F18
	VkF19               = C.VK_F19
	VkF20               = C.VK_F20
	VkF21               = C.VK_F21
	VkF22               = C.VK_F22
	VkF23               = C.VK_F23
	VkF24               = C.VK_F24
	VkNumlock           = C.VK_NUMLOCK
	VkScroll            = C.VK_SCROLL
	VkOemNecEqual       = C.VK_OEM_NEC_EQUAL
	VkOemFjJisho        = C.VK_OEM_FJ_JISHO
	VkOemFjMasshou      = C.VK_OEM_FJ_MASSHOU
	VkOemFjTouroku      = C.VK_OEM_FJ_TOUROKU
	VkOemFjLoya         = C.VK_OEM_FJ_LOYA
	VkOemFjRoya         = C.VK_OEM_FJ_ROYA
	VkLshift            = C.VK_LSHIFT
	VkRshift            = C.VK_RSHIFT
	VkLcontrol          = C.VK_LCONTROL
	VkRcontrol          = C.VK_RCONTROL
	VkLmenu             = C.VK_LMENU
	VkRmenu             = C.VK_RMENU
	VkBrowserBack       = C.VK_BROWSER_BACK
	VkBrowserForward    = C.VK_BROWSER_FORWARD
	VkBrowserRefresh    = C.VK_BROWSER_REFRESH
	VkBrowserStop       = C.VK_BROWSER_STOP
	VkBrowserSearch     = C.VK_BROWSER_SEARCH
	VkBrowserFavorites  = C.VK_BROWSER_FAVORITES
	VkBrowserHome       = C.VK_BROWSER_HOME
	VkVolumeMute        = C.VK_VOLUME_MUTE
	VkVolumeDown        = C.VK_VOLUME_DOWN
	VkVolumeUp          = C.VK_VOLUME_UP
	VkMediaNextTrack    = C.VK_MEDIA_NEXT_TRACK
	VkMediaPrevTrack    = C.VK_MEDIA_PREV_TRACK
	VkMediaStop         = C.VK_MEDIA_STOP
	VkMediaPlayPause    = C.VK_MEDIA_PLAY_PAUSE
	VkLaunchMail        = C.VK_LAUNCH_MAIL
	VkLaunchMediaSelect = C.VK_LAUNCH_MEDIA_SELECT
	VkLaunchApp1        = C.VK_LAUNCH_APP1
	VkLaunchApp2        = C.VK_LAUNCH_APP2
	VkOem1              = C.VK_OEM_1
	VkOemPlus           = C.VK_OEM_PLUS
	VkOemComma          = C.VK_OEM_COMMA
	VkOemMinus          = C.VK_OEM_MINUS
	VkOemPeriod         = C.VK_OEM_PERIOD
	VkOem2              = C.VK_OEM_2
	VkOem3              = C.VK_OEM_3
	VkOem4              = C.VK_OEM_4
	VkOem5              = C.VK_OEM_5
	VkOem6              = C.VK_OEM_6
	VkOem7              = C.VK_OEM_7
	VkOem8              = C.VK_OEM_8
	VkOemAx             = C.VK_OEM_AX
	VkOem102            = C.VK_OEM_102
	VkIcoHelp           = C.VK_ICO_HELP
	VkIco00             = C.VK_ICO_00
	VkProcesskey        = C.VK_PROCESSKEY
	VkIcoClear          = C.VK_ICO_CLEAR
	VkPacket            = C.VK_PACKET
	VkOemReset          = C.VK_OEM_RESET
	VkOemJump           = C.VK_OEM_JUMP
	VkOemPa1            = C.VK_OEM_PA1
	VkOemPa2            = C.VK_OEM_PA2
	VkOemPa3            = C.VK_OEM_PA3
	VkOemWsctrl         = C.VK_OEM_WSCTRL
	VkOemCusel          = C.VK_OEM_CUSEL
	VkOemAttn           = C.VK_OEM_ATTN
	VkOemFinish         = C.VK_OEM_FINISH
	VkOemCopy           = C.VK_OEM_COPY
	VkOemAuto           = C.VK_OEM_AUTO
	VkOemEnlw           = C.VK_OEM_ENLW
	VkOemBacktab        = C.VK_OEM_BACKTAB
	VkAttn              = C.VK_ATTN
	VkCrsel             = C.VK_CRSEL
	VkExsel             = C.VK_EXSEL
	VkEreof             = C.VK_EREOF
	VkPlay              = C.VK_PLAY
	VkZoom              = C.VK_ZOOM
	VkNoname            = C.VK_NONAME
	VkPa1               = C.VK_PA1
	VkOemClear          = C.VK_OEM_CLEAR
)

const (
	WaInactive    = C.WA_INACTIVE
	WaActive      = C.WA_ACTIVE
	WaClickactive = C.WA_CLICKACTIVE
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

const (
	ProcessDpiUnaware         win32api.ProcessDpiAwareness = C.PROCESS_DPI_UNAWARE
	ProcessSystemDpiAware     win32api.ProcessDpiAwareness = C.PROCESS_SYSTEM_DPI_AWARE
	ProcessPerMonitorDpiAware win32api.ProcessDpiAwareness = C.PROCESS_PER_MONITOR_DPI_AWARE
)

const (
	SOk = C.S_OK
)
