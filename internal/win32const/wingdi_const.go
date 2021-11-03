package win32const

// #include <windef.h>
// #include <winbase.h>
// #include <wingdi.h>
import "C"

// wingdi.h
const (
	AnsiCharset        = C.ANSI_CHARSET
	DefaultCharset     = C.DEFAULT_CHARSET
	SymbolCharset      = C.SYMBOL_CHARSET
	ShiftjisCharset    = C.SHIFTJIS_CHARSET
	HangeulCharset     = C.HANGEUL_CHARSET
	HangulCharset      = C.HANGUL_CHARSET
	Gb2312Charset      = C.GB2312_CHARSET
	Chinesebig5Charset = C.CHINESEBIG5_CHARSET
	OemCharset         = C.OEM_CHARSET
	JohabCharset       = C.JOHAB_CHARSET
	HebrewCharset      = C.HEBREW_CHARSET
	ArabicCharset      = C.ARABIC_CHARSET
	GreekCharset       = C.GREEK_CHARSET
	TurkishCharset     = C.TURKISH_CHARSET
	VietnameseCharset  = C.VIETNAMESE_CHARSET
	ThaiCharset        = C.THAI_CHARSET
	EasteuropeCharset  = C.EASTEUROPE_CHARSET
	RussianCharset     = C.RUSSIAN_CHARSET
	MacCharset         = C.MAC_CHARSET
	BalticCharset      = C.BALTIC_CHARSET
)

const (
	DefaultQuality          = C.DEFAULT_QUALITY
	DraftQuality            = C.DRAFT_QUALITY
	ProofQuality            = C.PROOF_QUALITY
	NonantialiasedQuality   = C.NONANTIALIASED_QUALITY
	AntialiasedQuality      = C.ANTIALIASED_QUALITY
	CleartypeQuality        = C.CLEARTYPE_QUALITY
	CleartypeNaturalQuality = C.CLEARTYPE_NATURAL_QUALITY
)

const (
	DefaultPitch  = C.DEFAULT_PITCH
	FixedPitch    = C.FIXED_PITCH
	VariablePitch = C.VARIABLE_PITCH
)

const (
	ClipToPath          = C.CLIP_TO_PATH
	ClipDefaultPrecis   = C.CLIP_DEFAULT_PRECIS
	ClipCharacterPrecis = C.CLIP_CHARACTER_PRECIS
	ClipStrokePrecis    = C.CLIP_STROKE_PRECIS
	ClipMask            = C.CLIP_MASK
	ClipLhAngles        = C.CLIP_LH_ANGLES
	ClipTtAlways        = C.CLIP_TT_ALWAYS
	ClipDfaDisable      = C.CLIP_DFA_DISABLE
	ClipEmbedded        = C.CLIP_EMBEDDED
)

const (
	FfDontcare   = C.FF_DONTCARE
	FfRoman      = C.FF_ROMAN
	FfSwiss      = C.FF_SWISS
	FfModern     = C.FF_MODERN
	FfScript     = C.FF_SCRIPT
	FfDecorative = C.FF_DECORATIVE
)

const (
	FwDontcare   = C.FW_DONTCARE
	FwThin       = C.FW_THIN
	FwExtralight = C.FW_EXTRALIGHT
	FwLight      = C.FW_LIGHT
	FwNormal     = C.FW_NORMAL
	FwMedium     = C.FW_MEDIUM
	FwSemibold   = C.FW_SEMIBOLD
	FwBold       = C.FW_BOLD
	FwExtrabold  = C.FW_EXTRABOLD
	FwHeavy      = C.FW_HEAVY
	FwUltralight = C.FW_ULTRALIGHT
	FwRegular    = C.FW_REGULAR
	FwDemibold   = C.FW_DEMIBOLD
	FwUltrabold  = C.FW_ULTRABOLD
	FwBlack      = C.FW_BLACK
)

const (
	OutDefaultPrecis       = C.OUT_DEFAULT_PRECIS
	OutStringPrecis        = C.OUT_STRING_PRECIS
	OutCharacterPrecis     = C.OUT_CHARACTER_PRECIS
	OutStrokePrecis        = C.OUT_STROKE_PRECIS
	OutTtPrecis            = C.OUT_TT_PRECIS
	OutDevicePrecis        = C.OUT_DEVICE_PRECIS
	OutRasterPrecis        = C.OUT_RASTER_PRECIS
	OutTtOnlyPrecis        = C.OUT_TT_ONLY_PRECIS
	OutOutlinePrecis       = C.OUT_OUTLINE_PRECIS
	OutScreenOutlinePrecis = C.OUT_SCREEN_OUTLINE_PRECIS
	OutPsOnlyPrecis        = C.OUT_PS_ONLY_PRECIS
)

const (
	PfdTypeRgba             = C.PFD_TYPE_RGBA
	PfdTypeColorindex       = C.PFD_TYPE_COLORINDEX
	PfdMainPlane            = C.PFD_MAIN_PLANE
	PfdOverlayPlane         = C.PFD_OVERLAY_PLANE
	PfdUnderlayPlane        = C.PFD_UNDERLAY_PLANE
	PfdDoublebuffer         = C.PFD_DOUBLEBUFFER
	PfdStereo               = C.PFD_STEREO
	PfdDrawToWindow         = C.PFD_DRAW_TO_WINDOW
	PfdDrawToBitmap         = C.PFD_DRAW_TO_BITMAP
	PfdSupportGdi           = C.PFD_SUPPORT_GDI
	PfdSupportOpengl        = C.PFD_SUPPORT_OPENGL
	PfdGenericFormat        = C.PFD_GENERIC_FORMAT
	PfdNeedPalette          = C.PFD_NEED_PALETTE
	PfdNeedSystemPalette    = C.PFD_NEED_SYSTEM_PALETTE
	PfdSwapExchange         = C.PFD_SWAP_EXCHANGE
	PfdSwapCopy             = C.PFD_SWAP_COPY
	PfdSwapLayerBuffers     = C.PFD_SWAP_LAYER_BUFFERS
	PfdGenericAccelerated   = C.PFD_GENERIC_ACCELERATED
	PfdSupportDirectdraw    = C.PFD_SUPPORT_DIRECTDRAW
	PfdDirect3dAccelerated  = C.PFD_DIRECT3D_ACCELERATED
	PfdSupportComposition   = C.PFD_SUPPORT_COMPOSITION
	PfdDepthDontcare        = C.PFD_DEPTH_DONTCARE
	PfdDoublebufferDontcare = C.PFD_DOUBLEBUFFER_DONTCARE
	PfdStereoDontcare       = C.PFD_STEREO_DONTCARE
)

const (
	BiRgb       = C.BI_RGB
	BiRle8      = C.BI_RLE8
	BiRle4      = C.BI_RLE4
	BiBitfields = C.BI_BITFIELDS
	BiJpeg      = C.BI_JPEG
	BiPng       = C.BI_PNG
)

// index of GetDeviceCaps
const (
	Technology = C.TECHNOLOGY
	Horzsize   = C.HORZSIZE
	Vertsize   = C.VERTSIZE
	Horzres    = C.HORZRES
	Vertres    = C.VERTRES
	Bitspixel  = C.BITSPIXEL
	Planes     = C.PLANES
	Logpixelsx = C.LOGPIXELSX
	Logpixelsy = C.LOGPIXELSY
)

/*
#define NUMBRUSHES 16
#define NUMPENS 18
#define NUMMARKERS 20
#define NUMFONTS 22
#define NUMCOLORS 24
#define PDEVICESIZE 26
#define CURVECAPS 28
#define LINECAPS 30
#define POLYGONALCAPS 32
#define TEXTCAPS 34
#define CLIPCAPS 36
#define RASTERCAPS 38
#define ASPECTX 40
#define ASPECTY 42
#define ASPECTXY 44

#define LOGPIXELSX 88
#define LOGPIXELSY 90

#define SIZEPALETTE 104
#define NUMRESERVED 106
#define COLORRES 108

#define PHYSICALWIDTH 110
#define PHYSICALHEIGHT 111
#define PHYSICALOFFSETX 112
#define PHYSICALOFFSETY 113
#define SCALINGFACTORX 114
#define SCALINGFACTORY 115

#define VREFRESH 116
#define DESKTOPVERTRES 117
#define DESKTOPHORZRES 118
#define BLTALIGNMENT 119

#define SHADEBLENDCAPS 120
#define COLORMGMTCAPS 121

#ifndef NOGDICAPMASKS
#define DT_PLOTTER 0
#define DT_RASDISPLAY 1
#define DT_RASPRINTER 2
#define DT_RASCAMERA 3
#define DT_CHARSTREAM 4
#define DT_METAFILE 5
#define DT_DISPFILE 6

#define CC_NONE 0
#define CC_CIRCLES 1
#define CC_PIE 2
#define CC_CHORD 4
#define CC_ELLIPSES 8
#define CC_WIDE 16
#define CC_STYLED 32
#define CC_WIDESTYLED 64
#define CC_INTERIORS 128
#define CC_ROUNDRECT 256

#define LC_NONE 0
#define LC_POLYLINE 2
#define LC_MARKER 4
#define LC_POLYMARKER 8
#define LC_WIDE 16
#define LC_STYLED 32
#define LC_WIDESTYLED 64
#define LC_INTERIORS 128

#define PC_NONE 0
#define PC_POLYGON 1
#define PC_RECTANGLE 2
#define PC_WINDPOLYGON 4
#define PC_TRAPEZOID 4
#define PC_SCANLINE 8
#define PC_WIDE 16
#define PC_STYLED 32
#define PC_WIDESTYLED 64
#define PC_INTERIORS 128
#define PC_POLYPOLYGON 256
#define PC_PATHS 512

#define CP_NONE 0
#define CP_RECTANGLE 1
#define CP_REGION 2

#define TC_OP_CHARACTER 0x00000001
#define TC_OP_STROKE 0x00000002
#define TC_CP_STROKE 0x00000004
#define TC_CR_90 0x00000008
#define TC_CR_ANY 0x00000010
#define TC_SF_X_YINDEP 0x00000020
#define TC_SA_DOUBLE 0x00000040
#define TC_SA_INTEGER 0x00000080
#define TC_SA_CONTIN 0x00000100
#define TC_EA_DOUBLE 0x00000200
#define TC_IA_ABLE 0x00000400
#define TC_UA_ABLE 0x00000800
#define TC_SO_ABLE 0x00001000
#define TC_RA_ABLE 0x00002000
#define TC_VA_ABLE 0x00004000
#define TC_RESERVED 0x00008000
#define TC_SCROLLBLT 0x00010000
#endif

#define RC_NONE
#define RC_BITBLT 1
#define RC_BANDING 2
#define RC_SCALING 4
#define RC_BITMAP64 8
#define RC_GDI20_OUTPUT 0x0010
#define RC_GDI20_STATE 0x0020
#define RC_SAVEBITMAP 0x0040
#define RC_DI_BITMAP 0x0080
#define RC_PALETTE 0x0100
#define RC_DIBTODEV 0x0200
#define RC_BIGFONT 0x0400
#define RC_STRETCHBLT 0x0800
#define RC_FLOODFILL 0x1000
#define RC_STRETCHDIB 0x2000
#define RC_OP_DX_OUTPUT 0x4000
#define RC_DEVBITS 0x8000

#define SB_NONE 0x00000000
#define SB_CONST_ALPHA 0x00000001
#define SB_PIXEL_ALPHA 0x00000002
#define SB_PREMULT_ALPHA 0x00000004

#define SB_GRAD_RECT 0x00000010
#define SB_GRAD_TRI 0x00000020

#define CM_NONE 0x00000000
#define CM_DEVICE_ICM 0x00000001
#define CM_GAMMA_RAMP 0x00000002
#define CM_CMYK_COLOR 0x00000004
*/
