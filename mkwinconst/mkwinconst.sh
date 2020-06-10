#!/bin/bash

INC=${2:-winuser}
cat /mingw64/x86_64-w64-mingw32/include/$INC.h | go run mkwinconst.go -prefix $1
# cat /mingw64/x86_64-w64-mingw32/include/$INC.h | go run mkwinconst.go -postfix $1
