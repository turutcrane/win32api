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
	False = 0
	True  = 1
)

const (
	sizeof_RECT                  = C.sizeof_RECT
	sizeof_WNDCLASSEX            = C.sizeof_WNDCLASSEX
	sizeof_CREATESTRUCTW         = C.sizeof_CREATESTRUCTW
	sizeof_PAINTSTRUCT           = C.sizeof_PAINTSTRUCT
	sizeof_FINDREPLACE           = C.sizeof_FINDREPLACE
	sizeof_PIXELFORMATDESCRIPTOR = C.sizeof_PIXELFORMATDESCRIPTOR
	sizeof_LONGLONG              = C.sizeof_LONGLONG
	sizeof_TRACKMOUSEEVENT       = C.sizeof_TRACKMOUSEEVENT
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
	// GwlWndproc    = C.GWL_WNDPROC
	// GwlHinstance  = C.GWL_HINSTANCE
	// GwlHwndparent = C.GWL_HWNDPARENT
	GwlStyle   = C.GWL_STYLE
	GwlExstyle = C.GWL_EXSTYLE
	// GwlUserdata   = C.GWL_USERDATA
	GwlId = C.GWL_ID
)

const (
	GwlpWndproc    = C.GWLP_WNDPROC
	GwlpHinstance  = C.GWLP_HINSTANCE
	GwlpHwndparent = C.GWLP_HWNDPARENT
	GwlpUserdata   = C.GWLP_USERDATA
	GwlpId         = C.GWLP_ID
)

const (
	Hterror       = C.HTERROR
	Httransparent = C.HTTRANSPARENT
	Htnowhere     = C.HTNOWHERE
	Htclient      = C.HTCLIENT
	Htcaption     = C.HTCAPTION
	Htsysmenu     = C.HTSYSMENU
	Htgrowbox     = C.HTGROWBOX
	Htsize        = C.HTSIZE
	Htmenu        = C.HTMENU
	Hthscroll     = C.HTHSCROLL
	Htvscroll     = C.HTVSCROLL
	Htminbutton   = C.HTMINBUTTON
	Htmaxbutton   = C.HTMAXBUTTON
	Htleft        = C.HTLEFT
	Htright       = C.HTRIGHT
	Httop         = C.HTTOP
	Httopleft     = C.HTTOPLEFT
	Httopright    = C.HTTOPRIGHT
	Htbottom      = C.HTBOTTOM
	Htbottomleft  = C.HTBOTTOMLEFT
	Htbottomright = C.HTBOTTOMRIGHT
	Htborder      = C.HTBORDER
	Htreduce      = C.HTREDUCE
	Htzoom        = C.HTZOOM
	Htsizefirst   = C.HTSIZEFIRST
	Htsizelast    = C.HTSIZELAST
	Htobject      = C.HTOBJECT
	Htclose       = C.HTCLOSE
	Hthelp        = C.HTHELP
)

var (
	Idok       = C.IDOK
	Idcancel   = C.IDCANCEL
	Idabort    = C.IDABORT
	Idretry    = C.IDRETRY
	Idignore   = C.IDIGNORE
	Idyes      = C.IDYES
	Idno       = C.IDNO
	Idclose    = C.IDCLOSE
	Idhelp     = C.IDHELP
	Idtryagain = C.IDTRYAGAIN
	Idcontinue = C.IDCONTINUE
	Idtimeout  = C.IDTIMEOUT
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
	ObjidWindow            = C.OBJID_WINDOW
	ObjidSysmenu           = C.OBJID_SYSMENU
	ObjidTitlebar          = C.OBJID_TITLEBAR
	ObjidMenu              = C.OBJID_MENU
	ObjidClient            = C.OBJID_CLIENT
	ObjidVscroll           = C.OBJID_VSCROLL
	ObjidHscroll           = C.OBJID_HSCROLL
	ObjidSizegrip          = C.OBJID_SIZEGRIP
	ObjidCaret             = C.OBJID_CARET
	ObjidCursor            = C.OBJID_CURSOR
	ObjidAlert             = C.OBJID_ALERT
	ObjidSound             = C.OBJID_SOUND
	ObjidQueryclassnameidx = C.OBJID_QUERYCLASSNAMEIDX
	ObjidNativeom          = C.OBJID_NATIVEOM
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
	KfExtended = C.KF_EXTENDED
	KfDlgmode  = C.KF_DLGMODE
	KfMenumode = C.KF_MENUMODE
	KfAltdown  = C.KF_ALTDOWN
	KfRepeat   = C.KF_REPEAT
	KfUp       = C.KF_UP
)

const (
	WaInactive    = C.WA_INACTIVE
	WaActive      = C.WA_ACTIVE
	WaClickactive = C.WA_CLICKACTIVE
)

const (
	WmNull                   MessageId = C.WM_NULL
	WmCreate                 MessageId = C.WM_CREATE
	WmDestroy                MessageId = C.WM_DESTROY
	WmMove                   MessageId = C.WM_MOVE
	WmSize                   MessageId = C.WM_SIZE
	WmActivate               MessageId = C.WM_ACTIVATE
	WmSetfocus               MessageId = C.WM_SETFOCUS
	WmKillfocus              MessageId = C.WM_KILLFOCUS
	WmEnable                 MessageId = C.WM_ENABLE
	WmSetredraw              MessageId = C.WM_SETREDRAW
	WmSettext                MessageId = C.WM_SETTEXT
	WmGettext                MessageId = C.WM_GETTEXT
	WmGettextlength          MessageId = C.WM_GETTEXTLENGTH
	WmPaint                  MessageId = C.WM_PAINT
	WmClose                  MessageId = C.WM_CLOSE
	WmQueryendsession        MessageId = C.WM_QUERYENDSESSION
	WmQueryopen              MessageId = C.WM_QUERYOPEN
	WmEndsession             MessageId = C.WM_ENDSESSION
	WmQuit                   MessageId = C.WM_QUIT
	WmErasebkgnd             MessageId = C.WM_ERASEBKGND
	WmSyscolorchange         MessageId = C.WM_SYSCOLORCHANGE
	WmShowwindow             MessageId = C.WM_SHOWWINDOW
	WmWininichange           MessageId = C.WM_WININICHANGE
	WmSettingchange          MessageId = C.WM_SETTINGCHANGE
	WmDevmodechange          MessageId = C.WM_DEVMODECHANGE
	WmActivateapp            MessageId = C.WM_ACTIVATEAPP
	WmFontchange             MessageId = C.WM_FONTCHANGE
	WmTimechange             MessageId = C.WM_TIMECHANGE
	WmCancelmode             MessageId = C.WM_CANCELMODE
	WmSetcursor              MessageId = C.WM_SETCURSOR
	WmMouseactivate          MessageId = C.WM_MOUSEACTIVATE
	WmChildactivate          MessageId = C.WM_CHILDACTIVATE
	WmQueuesync              MessageId = C.WM_QUEUESYNC
	WmGetminmaxinfo          MessageId = C.WM_GETMINMAXINFO
	WmPainticon              MessageId = C.WM_PAINTICON
	WmIconerasebkgnd         MessageId = C.WM_ICONERASEBKGND
	WmNextdlgctl             MessageId = C.WM_NEXTDLGCTL
	WmSpoolerstatus          MessageId = C.WM_SPOOLERSTATUS
	WmDrawitem               MessageId = C.WM_DRAWITEM
	WmMeasureitem            MessageId = C.WM_MEASUREITEM
	WmDeleteitem             MessageId = C.WM_DELETEITEM
	WmVkeytoitem             MessageId = C.WM_VKEYTOITEM
	WmChartoitem             MessageId = C.WM_CHARTOITEM
	WmSetfont                MessageId = C.WM_SETFONT
	WmGetfont                MessageId = C.WM_GETFONT
	WmSethotkey              MessageId = C.WM_SETHOTKEY
	WmGethotkey              MessageId = C.WM_GETHOTKEY
	WmQuerydragicon          MessageId = C.WM_QUERYDRAGICON
	WmCompareitem            MessageId = C.WM_COMPAREITEM
	WmGetobject              MessageId = C.WM_GETOBJECT
	WmCompacting             MessageId = C.WM_COMPACTING
	WmCommnotify             MessageId = C.WM_COMMNOTIFY
	WmWindowposchanging      MessageId = C.WM_WINDOWPOSCHANGING
	WmWindowposchanged       MessageId = C.WM_WINDOWPOSCHANGED
	WmPower                  MessageId = C.WM_POWER
	WmCopydata               MessageId = C.WM_COPYDATA
	WmCanceljournal          MessageId = C.WM_CANCELJOURNAL
	WmNotify                 MessageId = C.WM_NOTIFY
	WmInputlangchangerequest MessageId = C.WM_INPUTLANGCHANGEREQUEST
	WmInputlangchange        MessageId = C.WM_INPUTLANGCHANGE
	WmTcard                  MessageId = C.WM_TCARD
	WmHelp                   MessageId = C.WM_HELP
	WmUserchanged            MessageId = C.WM_USERCHANGED
	WmNotifyformat           MessageId = C.WM_NOTIFYFORMAT
	WmContextmenu            MessageId = C.WM_CONTEXTMENU
	WmStylechanging          MessageId = C.WM_STYLECHANGING
	WmStylechanged           MessageId = C.WM_STYLECHANGED
	WmDisplaychange          MessageId = C.WM_DISPLAYCHANGE
	WmGeticon                MessageId = C.WM_GETICON
	WmSeticon                MessageId = C.WM_SETICON
	WmNccreate               MessageId = C.WM_NCCREATE
	WmNcdestroy              MessageId = C.WM_NCDESTROY
	WmNccalcsize             MessageId = C.WM_NCCALCSIZE
	WmNchittest              MessageId = C.WM_NCHITTEST
	WmNcpaint                MessageId = C.WM_NCPAINT
	WmNcactivate             MessageId = C.WM_NCACTIVATE
	WmGetdlgcode             MessageId = C.WM_GETDLGCODE
	WmSyncpaint              MessageId = C.WM_SYNCPAINT
	WmNcmousemove            MessageId = C.WM_NCMOUSEMOVE
	WmNclbuttondown          MessageId = C.WM_NCLBUTTONDOWN
	WmNclbuttonup            MessageId = C.WM_NCLBUTTONUP
	WmNclbuttondblclk        MessageId = C.WM_NCLBUTTONDBLCLK
	WmNcrbuttondown          MessageId = C.WM_NCRBUTTONDOWN
	WmNcrbuttonup            MessageId = C.WM_NCRBUTTONUP
	WmNcrbuttondblclk        MessageId = C.WM_NCRBUTTONDBLCLK
	WmNcmbuttondown          MessageId = C.WM_NCMBUTTONDOWN
	WmNcmbuttonup            MessageId = C.WM_NCMBUTTONUP
	WmNcmbuttondblclk        MessageId = C.WM_NCMBUTTONDBLCLK
	WmNcxbuttondown          MessageId = C.WM_NCXBUTTONDOWN
	WmNcxbuttonup            MessageId = C.WM_NCXBUTTONUP
	WmNcxbuttondblclk        MessageId = C.WM_NCXBUTTONDBLCLK
	WmInputDeviceChange      MessageId = C.WM_INPUT_DEVICE_CHANGE
	WmInput                  MessageId = C.WM_INPUT
	WmKeyfirst               MessageId = C.WM_KEYFIRST
	WmKeydown                MessageId = C.WM_KEYDOWN
	WmKeyup                  MessageId = C.WM_KEYUP
	WmChar                   MessageId = C.WM_CHAR
	WmDeadchar               MessageId = C.WM_DEADCHAR
	WmSyskeydown             MessageId = C.WM_SYSKEYDOWN
	WmSyskeyup               MessageId = C.WM_SYSKEYUP
	WmSyschar                MessageId = C.WM_SYSCHAR
	WmSysdeadchar            MessageId = C.WM_SYSDEADCHAR
	WmUnichar                MessageId = C.WM_UNICHAR
	WmKeylast                MessageId = C.WM_KEYLAST
	WmImeStartcomposition    MessageId = C.WM_IME_STARTCOMPOSITION
	WmImeEndcomposition      MessageId = C.WM_IME_ENDCOMPOSITION
	WmImeComposition         MessageId = C.WM_IME_COMPOSITION
	WmImeKeylast             MessageId = C.WM_IME_KEYLAST
	WmInitdialog             MessageId = C.WM_INITDIALOG
	WmCommand                MessageId = C.WM_COMMAND
	WmSyscommand             MessageId = C.WM_SYSCOMMAND
	WmTimer                  MessageId = C.WM_TIMER
	WmHscroll                MessageId = C.WM_HSCROLL
	WmVscroll                MessageId = C.WM_VSCROLL
	WmInitmenu               MessageId = C.WM_INITMENU
	WmInitmenupopup          MessageId = C.WM_INITMENUPOPUP
	WmMenuselect             MessageId = C.WM_MENUSELECT
	WmGesture                MessageId = C.WM_GESTURE
	WmGesturenotify          MessageId = C.WM_GESTURENOTIFY
	WmMenuchar               MessageId = C.WM_MENUCHAR
	WmEnteridle              MessageId = C.WM_ENTERIDLE
	WmMenurbuttonup          MessageId = C.WM_MENURBUTTONUP
	WmMenudrag               MessageId = C.WM_MENUDRAG
	WmMenugetobject          MessageId = C.WM_MENUGETOBJECT
	WmUninitmenupopup        MessageId = C.WM_UNINITMENUPOPUP
	WmMenucommand            MessageId = C.WM_MENUCOMMAND
	WmChangeuistate          MessageId = C.WM_CHANGEUISTATE
	WmUpdateuistate          MessageId = C.WM_UPDATEUISTATE
	WmQueryuistate           MessageId = C.WM_QUERYUISTATE
	WmCtlcolormsgbox         MessageId = C.WM_CTLCOLORMSGBOX
	WmCtlcoloredit           MessageId = C.WM_CTLCOLOREDIT
	WmCtlcolorlistbox        MessageId = C.WM_CTLCOLORLISTBOX
	WmCtlcolorbtn            MessageId = C.WM_CTLCOLORBTN
	WmCtlcolordlg            MessageId = C.WM_CTLCOLORDLG
	WmCtlcolorscrollbar      MessageId = C.WM_CTLCOLORSCROLLBAR
	WmCtlcolorstatic         MessageId = C.WM_CTLCOLORSTATIC
	WmMousefirst             MessageId = C.WM_MOUSEFIRST
	WmMousemove              MessageId = C.WM_MOUSEMOVE
	WmLbuttondown            MessageId = C.WM_LBUTTONDOWN
	WmLbuttonup              MessageId = C.WM_LBUTTONUP
	WmLbuttondblclk          MessageId = C.WM_LBUTTONDBLCLK
	WmRbuttondown            MessageId = C.WM_RBUTTONDOWN
	WmRbuttonup              MessageId = C.WM_RBUTTONUP
	WmRbuttondblclk          MessageId = C.WM_RBUTTONDBLCLK
	WmMbuttondown            MessageId = C.WM_MBUTTONDOWN
	WmMbuttonup              MessageId = C.WM_MBUTTONUP
	WmMbuttondblclk          MessageId = C.WM_MBUTTONDBLCLK
	WmMousewheel             MessageId = C.WM_MOUSEWHEEL
	WmXbuttondown            MessageId = C.WM_XBUTTONDOWN
	WmXbuttonup              MessageId = C.WM_XBUTTONUP
	WmXbuttondblclk          MessageId = C.WM_XBUTTONDBLCLK
	WmMousehwheel            MessageId = C.WM_MOUSEHWHEEL
	WmMouselast              MessageId = C.WM_MOUSELAST
	WmParentnotify           MessageId = C.WM_PARENTNOTIFY
	WmEntermenuloop          MessageId = C.WM_ENTERMENULOOP
	WmExitmenuloop           MessageId = C.WM_EXITMENULOOP
	WmNextmenu               MessageId = C.WM_NEXTMENU
	WmSizing                 MessageId = C.WM_SIZING
	WmCapturechanged         MessageId = C.WM_CAPTURECHANGED
	WmMoving                 MessageId = C.WM_MOVING
	WmPowerbroadcast         MessageId = C.WM_POWERBROADCAST
	WmDevicechange           MessageId = C.WM_DEVICECHANGE
	WmMdicreate              MessageId = C.WM_MDICREATE
	WmMdidestroy             MessageId = C.WM_MDIDESTROY
	WmMdiactivate            MessageId = C.WM_MDIACTIVATE
	WmMdirestore             MessageId = C.WM_MDIRESTORE
	WmMdinext                MessageId = C.WM_MDINEXT
	WmMdimaximize            MessageId = C.WM_MDIMAXIMIZE
	WmMditile                MessageId = C.WM_MDITILE
	WmMdicascade             MessageId = C.WM_MDICASCADE
	WmMdiiconarrange         MessageId = C.WM_MDIICONARRANGE
	WmMdigetactive           MessageId = C.WM_MDIGETACTIVE
	WmMdisetmenu             MessageId = C.WM_MDISETMENU
	WmEntersizemove          MessageId = C.WM_ENTERSIZEMOVE
	WmExitsizemove           MessageId = C.WM_EXITSIZEMOVE
	WmDropfiles              MessageId = C.WM_DROPFILES
	WmMdirefreshmenu         MessageId = C.WM_MDIREFRESHMENU
	// WmPointerdevicechange MessageId = C.WM_POINTERDEVICECHANGE // #if WINVER >= 0x0602
	// WmPointerdeviceinrange MessageId = C.WM_POINTERDEVICEINRANGE
	// WmPointerdeviceoutofrange MessageId = C.WM_POINTERDEVICEOUTOFRANGE
	WmTouch MessageId = C.WM_TOUCH
	// WmNcpointerupdate MessageId = C.WM_NCPOINTERUPDATE  	// #if WINVER >= 0x0602
	// WmNcpointerdown MessageId = C.WM_NCPOINTERDOWN
	// WmNcpointerup MessageId = C.WM_NCPOINTERUP
	// WmPointerupdate MessageId = C.WM_POINTERUPDATE
	// WmPointerdown MessageId = C.WM_POINTERDOWN
	// WmPointerup MessageId = C.WM_POINTERUP
	// WmPointerenter MessageId = C.WM_POINTERENTER
	// WmPointerleave MessageId = C.WM_POINTERLEAVE
	// WmPointeractivate MessageId = C.WM_POINTERACTIVATE
	// WmPointercapturechanged MessageId = C.WM_POINTERCAPTURECHANGED
	// WmTouchhittesting MessageId = C.WM_TOUCHHITTESTING
	// WmPointerwheel MessageId = C.WM_POINTERWHEEL
	// WmPointerhwheel MessageId = C.WM_POINTERHWHEEL
	WmImeSetcontext                  MessageId = C.WM_IME_SETCONTEXT
	WmImeNotify                      MessageId = C.WM_IME_NOTIFY
	WmImeControl                     MessageId = C.WM_IME_CONTROL
	WmImeCompositionfull             MessageId = C.WM_IME_COMPOSITIONFULL
	WmImeSelect                      MessageId = C.WM_IME_SELECT
	WmImeChar                        MessageId = C.WM_IME_CHAR
	WmImeRequest                     MessageId = C.WM_IME_REQUEST
	WmImeKeydown                     MessageId = C.WM_IME_KEYDOWN
	WmImeKeyup                       MessageId = C.WM_IME_KEYUP
	WmMousehover                     MessageId = C.WM_MOUSEHOVER
	WmMouseleave                     MessageId = C.WM_MOUSELEAVE
	WmNcmousehover                   MessageId = C.WM_NCMOUSEHOVER
	WmNcmouseleave                   MessageId = C.WM_NCMOUSELEAVE
	WmWtssessionChange               MessageId = C.WM_WTSSESSION_CHANGE
	WmTabletFirst                    MessageId = C.WM_TABLET_FIRST
	WmTabletLast                     MessageId = C.WM_TABLET_LAST
	WmDpichanged                     MessageId = C.WM_DPICHANGED
	WmCut                            MessageId = C.WM_CUT
	WmCopy                           MessageId = C.WM_COPY
	WmPaste                          MessageId = C.WM_PASTE
	WmClear                          MessageId = C.WM_CLEAR
	WmUndo                           MessageId = C.WM_UNDO
	WmRenderformat                   MessageId = C.WM_RENDERFORMAT
	WmRenderallformats               MessageId = C.WM_RENDERALLFORMATS
	WmDestroyclipboard               MessageId = C.WM_DESTROYCLIPBOARD
	WmDrawclipboard                  MessageId = C.WM_DRAWCLIPBOARD
	WmPaintclipboard                 MessageId = C.WM_PAINTCLIPBOARD
	WmVscrollclipboard               MessageId = C.WM_VSCROLLCLIPBOARD
	WmSizeclipboard                  MessageId = C.WM_SIZECLIPBOARD
	WmAskcbformatname                MessageId = C.WM_ASKCBFORMATNAME
	WmChangecbchain                  MessageId = C.WM_CHANGECBCHAIN
	WmHscrollclipboard               MessageId = C.WM_HSCROLLCLIPBOARD
	WmQuerynewpalette                MessageId = C.WM_QUERYNEWPALETTE
	WmPaletteischanging              MessageId = C.WM_PALETTEISCHANGING
	WmPalettechanged                 MessageId = C.WM_PALETTECHANGED
	WmHotkey                         MessageId = C.WM_HOTKEY
	WmPrint                          MessageId = C.WM_PRINT
	WmPrintclient                    MessageId = C.WM_PRINTCLIENT
	WmAppcommand                     MessageId = C.WM_APPCOMMAND
	WmThemechanged                   MessageId = C.WM_THEMECHANGED
	WmClipboardupdate                MessageId = C.WM_CLIPBOARDUPDATE
	WmDwmcompositionchanged          MessageId = C.WM_DWMCOMPOSITIONCHANGED
	WmDwmncrenderingchanged          MessageId = C.WM_DWMNCRENDERINGCHANGED
	WmDwmcolorizationcolorchanged    MessageId = C.WM_DWMCOLORIZATIONCOLORCHANGED
	WmDwmwindowmaximizedchange       MessageId = C.WM_DWMWINDOWMAXIMIZEDCHANGE
	WmDwmsendiconicthumbnail         MessageId = C.WM_DWMSENDICONICTHUMBNAIL
	WmDwmsendiconiclivepreviewbitmap MessageId = C.WM_DWMSENDICONICLIVEPREVIEWBITMAP
	WmGettitlebarinfoex              MessageId = C.WM_GETTITLEBARINFOEX
	WmHandheldfirst                  MessageId = C.WM_HANDHELDFIRST
	WmHandheldlast                   MessageId = C.WM_HANDHELDLAST
	WmAfxfirst                       MessageId = C.WM_AFXFIRST
	WmAfxlast                        MessageId = C.WM_AFXLAST
	WmPenwinfirst                    MessageId = C.WM_PENWINFIRST
	WmPenwinlast                     MessageId = C.WM_PENWINLAST
	WmApp                            MessageId = C.WM_APP
	WmUser                           MessageId = C.WM_USER
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

const (
	SmCxscreen                    = C.SM_CXSCREEN
	SmCyscreen                    = C.SM_CYSCREEN
	SmCxvscroll                   = C.SM_CXVSCROLL
	SmCyhscroll                   = C.SM_CYHSCROLL
	SmCycaption                   = C.SM_CYCAPTION
	SmCxborder                    = C.SM_CXBORDER
	SmCyborder                    = C.SM_CYBORDER
	SmCxdlgframe                  = C.SM_CXDLGFRAME
	SmCydlgframe                  = C.SM_CYDLGFRAME
	SmCyvthumb                    = C.SM_CYVTHUMB
	SmCxhthumb                    = C.SM_CXHTHUMB
	SmCxicon                      = C.SM_CXICON
	SmCyicon                      = C.SM_CYICON
	SmCxcursor                    = C.SM_CXCURSOR
	SmCycursor                    = C.SM_CYCURSOR
	SmCymenu                      = C.SM_CYMENU
	SmCxfullscreen                = C.SM_CXFULLSCREEN
	SmCyfullscreen                = C.SM_CYFULLSCREEN
	SmCykanjiwindow               = C.SM_CYKANJIWINDOW
	SmMousepresent                = C.SM_MOUSEPRESENT
	SmCyvscroll                   = C.SM_CYVSCROLL
	SmCxhscroll                   = C.SM_CXHSCROLL
	SmDebug                       = C.SM_DEBUG
	SmSwapbutton                  = C.SM_SWAPBUTTON
	SmReserved1                   = C.SM_RESERVED1
	SmReserved2                   = C.SM_RESERVED2
	SmReserved3                   = C.SM_RESERVED3
	SmReserved4                   = C.SM_RESERVED4
	SmCxmin                       = C.SM_CXMIN
	SmCymin                       = C.SM_CYMIN
	SmCxsize                      = C.SM_CXSIZE
	SmCysize                      = C.SM_CYSIZE
	SmCxframe                     = C.SM_CXFRAME
	SmCyframe                     = C.SM_CYFRAME
	SmCxmintrack                  = C.SM_CXMINTRACK
	SmCymintrack                  = C.SM_CYMINTRACK
	SmCxdoubleclk                 = C.SM_CXDOUBLECLK
	SmCydoubleclk                 = C.SM_CYDOUBLECLK
	SmCxiconspacing               = C.SM_CXICONSPACING
	SmCyiconspacing               = C.SM_CYICONSPACING
	SmMenudropalignment           = C.SM_MENUDROPALIGNMENT
	SmPenwindows                  = C.SM_PENWINDOWS
	SmDbcsenabled                 = C.SM_DBCSENABLED
	SmCmousebuttons               = C.SM_CMOUSEBUTTONS
	SmCxfixedframe                = C.SM_CXFIXEDFRAME
	SmCyfixedframe                = C.SM_CYFIXEDFRAME
	SmCxsizeframe                 = C.SM_CXSIZEFRAME
	SmCysizeframe                 = C.SM_CYSIZEFRAME
	SmSecure                      = C.SM_SECURE
	SmCxedge                      = C.SM_CXEDGE
	SmCyedge                      = C.SM_CYEDGE
	SmCxminspacing                = C.SM_CXMINSPACING
	SmCyminspacing                = C.SM_CYMINSPACING
	SmCxsmicon                    = C.SM_CXSMICON
	SmCysmicon                    = C.SM_CYSMICON
	SmCysmcaption                 = C.SM_CYSMCAPTION
	SmCxsmsize                    = C.SM_CXSMSIZE
	SmCysmsize                    = C.SM_CYSMSIZE
	SmCxmenusize                  = C.SM_CXMENUSIZE
	SmCymenusize                  = C.SM_CYMENUSIZE
	SmArrange                     = C.SM_ARRANGE
	SmCxminimized                 = C.SM_CXMINIMIZED
	SmCyminimized                 = C.SM_CYMINIMIZED
	SmCxmaxtrack                  = C.SM_CXMAXTRACK
	SmCymaxtrack                  = C.SM_CYMAXTRACK
	SmCxmaximized                 = C.SM_CXMAXIMIZED
	SmCymaximized                 = C.SM_CYMAXIMIZED
	SmNetwork                     = C.SM_NETWORK
	SmCleanboot                   = C.SM_CLEANBOOT
	SmCxdrag                      = C.SM_CXDRAG
	SmCydrag                      = C.SM_CYDRAG
	SmShowsounds                  = C.SM_SHOWSOUNDS
	SmCxmenucheck                 = C.SM_CXMENUCHECK
	SmCymenucheck                 = C.SM_CYMENUCHECK
	SmSlowmachine                 = C.SM_SLOWMACHINE
	SmMideastenabled              = C.SM_MIDEASTENABLED
	SmMousewheelpresent           = C.SM_MOUSEWHEELPRESENT
	SmXvirtualscreen              = C.SM_XVIRTUALSCREEN
	SmYvirtualscreen              = C.SM_YVIRTUALSCREEN
	SmCxvirtualscreen             = C.SM_CXVIRTUALSCREEN
	SmCyvirtualscreen             = C.SM_CYVIRTUALSCREEN
	SmCmonitors                   = C.SM_CMONITORS
	SmSamedisplayformat           = C.SM_SAMEDISPLAYFORMAT
	SmImmenabled                  = C.SM_IMMENABLED
	SmCxfocusborder               = C.SM_CXFOCUSBORDER
	SmCyfocusborder               = C.SM_CYFOCUSBORDER
	SmTabletpc                    = C.SM_TABLETPC
	SmMediacenter                 = C.SM_MEDIACENTER
	SmStarter                     = C.SM_STARTER
	SmServerr2                    = C.SM_SERVERR2
	SmMousehorizontalwheelpresent = C.SM_MOUSEHORIZONTALWHEELPRESENT
	SmCxpaddedborder              = C.SM_CXPADDEDBORDER
	SmDigitizer                   = C.SM_DIGITIZER
	SmMaximumtouches              = C.SM_MAXIMUMTOUCHES
	SmCmetrics                    = C.SM_CMETRICS
	SmRemotesession               = C.SM_REMOTESESSION
	SmShuttingdown                = C.SM_SHUTTINGDOWN
	SmRemotecontrol               = C.SM_REMOTECONTROL
	SmCaretblinkingenabled        = C.SM_CARETBLINKINGENABLED
	// SmConvertibleslatemode = C.SM_CONVERTIBLESLATEMODE  // WINVER >= 0x0602
	// SmSystemdocked = C.SM_SYSTEMDOCKED // WINVER >= 0x0602
)

const (
	MkLbutton  = C.MK_LBUTTON
	MkRbutton  = C.MK_RBUTTON
	MkShift    = C.MK_SHIFT
	MkControl  = C.MK_CONTROL
	MkMbutton  = C.MK_MBUTTON
	MkXbutton1 = C.MK_XBUTTON1
	MkXbutton2 = C.MK_XBUTTON2
)

const (
	TmeHover     = C.TME_HOVER
	TmeLeave     = C.TME_LEAVE
	TmeNonclient = C.TME_NONCLIENT
	TmeQuery     = C.TME_QUERY
	TmeCancel    = C.TME_CANCEL
)

const (
	MfInsert          = C.MF_INSERT
	MfChange          = C.MF_CHANGE
	MfAppend          = C.MF_APPEND
	MfDelete          = C.MF_DELETE
	MfRemove          = C.MF_REMOVE
	MfBycommand       = C.MF_BYCOMMAND
	MfByposition      = C.MF_BYPOSITION
	MfSeparator       = C.MF_SEPARATOR
	MfEnabled         = C.MF_ENABLED
	MfGrayed          = C.MF_GRAYED
	MfDisabled        = C.MF_DISABLED
	MfUnchecked       = C.MF_UNCHECKED
	MfChecked         = C.MF_CHECKED
	MfUsecheckbitmaps = C.MF_USECHECKBITMAPS
	MfString          = C.MF_STRING
	MfBitmap          = C.MF_BITMAP
	MfOwnerdraw       = C.MF_OWNERDRAW
	MfPopup           = C.MF_POPUP
	MfMenubarbreak    = C.MF_MENUBARBREAK
	MfMenubreak       = C.MF_MENUBREAK
	MfUnhilite        = C.MF_UNHILITE
	MfHilite          = C.MF_HILITE
	MfDefault         = C.MF_DEFAULT
	MfSysmenu         = C.MF_SYSMENU
	MfHelp            = C.MF_HELP
	MfRightjustify    = C.MF_RIGHTJUSTIFY
	MfMouseselect     = C.MF_MOUSESELECT
	MfEnd             = C.MF_END
)
