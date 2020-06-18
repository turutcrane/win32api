#

.PHONY: vet
vet: win32.y.go mkapisys.go
	go vet

win32.y.go: win32.y
	# go generate *.go
	goyacc.exe -o win32.y.go win32.y
	# goyacc.exe -o win32.y.go -l win32.y
