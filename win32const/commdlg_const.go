package win32const

// #include <windef.h>
// #include <winbase.h>
// #include <winuser.h>
// #include <wingdi.h>
// #include <commdlg.h>
import "C"

// commdlg.h
const (
	FrDown                 = C.FR_DOWN
	FrWholeword            = C.FR_WHOLEWORD
	FrMatchcase            = C.FR_MATCHCASE
	FrFindnext             = C.FR_FINDNEXT
	FrReplace              = C.FR_REPLACE
	FrReplaceall           = C.FR_REPLACEALL
	FrDialogterm           = C.FR_DIALOGTERM
	FrShowhelp             = C.FR_SHOWHELP
	FrEnablehook           = C.FR_ENABLEHOOK
	FrEnabletemplate       = C.FR_ENABLETEMPLATE
	FrNoupdown             = C.FR_NOUPDOWN
	FrNomatchcase          = C.FR_NOMATCHCASE
	FrNowholeword          = C.FR_NOWHOLEWORD
	FrEnabletemplatehandle = C.FR_ENABLETEMPLATEHANDLE
	FrHideupdown           = C.FR_HIDEUPDOWN
	FrHidematchcase        = C.FR_HIDEMATCHCASE
	FrHidewholeword        = C.FR_HIDEWHOLEWORD
	FrRaw                  = C.FR_RAW
	FrMatchdiac            = C.FR_MATCHDIAC
	FrMatchkashida         = C.FR_MATCHKASHIDA
	FrMatchalefhamza       = C.FR_MATCHALEFHAMZA
)

const (
	Lbselchstring = C.LBSELCHSTRING
	Sharevistring = C.SHAREVISTRING
	Fileokstring  = C.FILEOKSTRING
	Colorokstring = C.COLOROKSTRING
	Setrgbstring  = C.SETRGBSTRING
	Helpmsgstring = C.HELPMSGSTRING
	Findmsgstring = C.FINDMSGSTRING
)
