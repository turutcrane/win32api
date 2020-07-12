package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	strcase "github.com/stoewer/go-strcase"
)

func main() {
	prefix := flag.String("prefix", "", "Const prefix ex) WS_")
	postfix := flag.String("postfix", "", "Const prefix ex) _CAHRSET")
	// prologue := flag.Bool("prologue", false, "Outpput preprocesser prologue")
	flag.Parse()
	if *prefix == "" && *postfix == "" {
		log.Fatalln("Usage: ex) mkwinconst -prefix WS_ -postfix _CHARSET")
	}

	// if *prologue {
	// 	fmt.Println("#define MAKEINTRESOURCE")
	// 	fmt.Println("#include <windows.h>")
	// }
	re := regexp.MustCompile("#define[\t ]+(" + *prefix + "[_[:alnum:]]+" + *postfix +")")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#define") {
			match := re.FindSubmatch([]byte(line))
			if match != nil {
				key := string(match[1])
				fmt.Printf("\t%s = C.%s\n", strcase.UpperCamelCase(key), key)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
