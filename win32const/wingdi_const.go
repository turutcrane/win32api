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
