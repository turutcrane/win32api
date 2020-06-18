#

.PHONY: vet
vet: win32api.go zwin32api_windows.go w32api_windows.go
	go vet

zwin32api_windows.go: win32api.go w32api_windows.go
	go generate win32api.go

w32api_windows.go: w32api.txt
	go run github.com/turutcrane/win32api/cmd/mkapisys -o w32api_windows.go -pkg win32api w32api.txt

