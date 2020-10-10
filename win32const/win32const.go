package win32const
//go:generate go run golang.org/x/tools/cmd/stringer -output win32const_string.go -type MessageId 

import (
	"github.com/turutcrane/win32api"
)
type MessageId win32api.UINT