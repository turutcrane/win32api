#

.PHONY: all test
all:  win32const_string.go

win32const_string.go: winuser_const.go
	go generate 

test:
	go test
