package win32const

// #include <windef.h>
// #include <winbase.h>
// #include <shlobj.h>
import "C"
import (
	"github.com/turutcrane/win32api"
)

// shlobj.
const (
	CsidlFlagCreate             = C.CSIDL_FLAG_CREATE
	CsidlPersonal               = C.CSIDL_PERSONAL
	CsidlMypictures             = C.CSIDL_MYPICTURES
	CsidlAppdata                = C.CSIDL_APPDATA
	CsidlMymusic                = C.CSIDL_MYMUSIC
	CsidlMyvideo                = C.CSIDL_MYVIDEO
	CsidlDesktop                = C.CSIDL_DESKTOP
	CsidlInternet               = C.CSIDL_INTERNET
	CsidlPrograms               = C.CSIDL_PROGRAMS
	CsidlControls               = C.CSIDL_CONTROLS
	CsidlPrinters               = C.CSIDL_PRINTERS
	CsidlFavorites              = C.CSIDL_FAVORITES
	CsidlStartup                = C.CSIDL_STARTUP
	CsidlRecent                 = C.CSIDL_RECENT
	CsidlSendto                 = C.CSIDL_SENDTO
	CsidlBitbucket              = C.CSIDL_BITBUCKET
	CsidlStartmenu              = C.CSIDL_STARTMENU
	CsidlMydocuments            = C.CSIDL_MYDOCUMENTS
	CsidlDesktopdirectory       = C.CSIDL_DESKTOPDIRECTORY
	CsidlDrives                 = C.CSIDL_DRIVES
	CsidlNetwork                = C.CSIDL_NETWORK
	CsidlNethood                = C.CSIDL_NETHOOD
	CsidlFonts                  = C.CSIDL_FONTS
	CsidlTemplates              = C.CSIDL_TEMPLATES
	CsidlCommonStartmenu        = C.CSIDL_COMMON_STARTMENU
	CsidlCommonPrograms         = C.CSIDL_COMMON_PROGRAMS
	CsidlCommonStartup          = C.CSIDL_COMMON_STARTUP
	CsidlCommonDesktopdirectory = C.CSIDL_COMMON_DESKTOPDIRECTORY
	CsidlPrinthood              = C.CSIDL_PRINTHOOD
	CsidlLocalAppdata           = C.CSIDL_LOCAL_APPDATA
	CsidlAltstartup             = C.CSIDL_ALTSTARTUP
	CsidlCommonAltstartup       = C.CSIDL_COMMON_ALTSTARTUP
	CsidlCommonFavorites        = C.CSIDL_COMMON_FAVORITES
	CsidlInternetCache          = C.CSIDL_INTERNET_CACHE
	CsidlCookies                = C.CSIDL_COOKIES
	CsidlHistory                = C.CSIDL_HISTORY
	CsidlCommonAppdata          = C.CSIDL_COMMON_APPDATA
	CsidlWindows                = C.CSIDL_WINDOWS
	CsidlSystem                 = C.CSIDL_SYSTEM
	CsidlProgramFiles           = C.CSIDL_PROGRAM_FILES
	CsidlProfile                = C.CSIDL_PROFILE
	CsidlSystemx86              = C.CSIDL_SYSTEMX86
	CsidlProgramFilesx86        = C.CSIDL_PROGRAM_FILESX86
	CsidlProgramFilesCommon     = C.CSIDL_PROGRAM_FILES_COMMON
	CsidlProgramFilesCommonx86  = C.CSIDL_PROGRAM_FILES_COMMONX86
	CsidlCommonTemplates        = C.CSIDL_COMMON_TEMPLATES
	CsidlCommonDocuments        = C.CSIDL_COMMON_DOCUMENTS
	CsidlCommonAdmintools       = C.CSIDL_COMMON_ADMINTOOLS
	CsidlAdmintools             = C.CSIDL_ADMINTOOLS
	CsidlConnections            = C.CSIDL_CONNECTIONS
	CsidlCommonMusic            = C.CSIDL_COMMON_MUSIC
	CsidlCommonPictures         = C.CSIDL_COMMON_PICTURES
	CsidlCommonVideo            = C.CSIDL_COMMON_VIDEO
	CsidlResources              = C.CSIDL_RESOURCES
	CsidlResourcesLocalized     = C.CSIDL_RESOURCES_LOCALIZED
	CsidlCommonOemLinks         = C.CSIDL_COMMON_OEM_LINKS
	CsidlCdburnArea             = C.CSIDL_CDBURN_AREA
	CsidlComputersnearme        = C.CSIDL_COMPUTERSNEARME
	CsidlFlagDontVerify         = C.CSIDL_FLAG_DONT_VERIFY
	CsidlFlagDontUnexpand       = C.CSIDL_FLAG_DONT_UNEXPAND
	CsidlFlagNoAlias            = C.CSIDL_FLAG_NO_ALIAS
	CsidlFlagPerUserInit        = C.CSIDL_FLAG_PER_USER_INIT
	CsidlFlagMask               = C.CSIDL_FLAG_MASK
)

const (
	KfFlagDefault                   win32api.KnownFolderFlag = C.KF_FLAG_DEFAULT
	KfFlagNoAppcontainerRedirection win32api.KnownFolderFlag = C.KF_FLAG_NO_APPCONTAINER_REDIRECTION
	KfFlagCreate                    win32api.KnownFolderFlag = C.KF_FLAG_CREATE
	KfFlagDontVerify                win32api.KnownFolderFlag = C.KF_FLAG_DONT_VERIFY
	KfFlagDontUnexpand              win32api.KnownFolderFlag = C.KF_FLAG_DONT_UNEXPAND
	KfFlagNoAlias                   win32api.KnownFolderFlag = C.KF_FLAG_NO_ALIAS
	KfFlagInit                      win32api.KnownFolderFlag = C.KF_FLAG_INIT
	KfFlagDefaultPath               win32api.KnownFolderFlag = C.KF_FLAG_DEFAULT_PATH
	KfFlagNotParentRelative         win32api.KnownFolderFlag = C.KF_FLAG_NOT_PARENT_RELATIVE
	KfFlagSimpleIdlist              win32api.KnownFolderFlag = C.KF_FLAG_SIMPLE_IDLIST
	KfFlagAliasOnly                 win32api.KnownFolderFlag = C.KF_FLAG_ALIAS_ONLY
)
