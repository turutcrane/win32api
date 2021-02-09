#

.PHONY: vet
vet: w32api_windows.go zwin32api_windows.go makeconst
	go vet -unsafeptr=false

zwin32api_windows.go: win32api.go w32api_windows.go
	go generate win32api.go

w32api_windows.go: w32api.txt
	$(MAKE) -C ./cmd/mkapisys vet
	go run ./cmd/mkapisys -o w32api_windows.go -pkg win32api w32api.txt
	#go run github.com/turutcrane/win32api/cmd/mkapisys -o w32api_windows.go -pkg win32api w32api.txt

.PHONY: makecosnt
makeconst:
	$(MAKE) -C ./internal/win32const
	go generate win32api_const.go
