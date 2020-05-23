#!/bin/bash

cat /mingw64/x86_64-w64-mingw32/include/winuser.h | go run mkwinconst.go -prefix $1
