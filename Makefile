#

.PHONY: vet
vet: win32api.go zwin32api_windows.go win32_const.go
	go vet .

zwin32api_windows.go: win32api.go
	go generate win32api.go

