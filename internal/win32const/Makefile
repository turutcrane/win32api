#

.PHONY: all test
all: win32api_const.tmp

win32api_const.tmp: commdlg_const.go  shlobj_const.go win32const.go  wingdi_const.go  winuser_const.go
	go run github.com/turutcrane/win32api/cmd/mkgoconst > win32api_const.tmp
	cp win32api_const.tmp ../../win32api_const.go

test:
	go test
