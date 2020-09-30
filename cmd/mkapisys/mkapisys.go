package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"

	strcase "github.com/stoewer/go-strcase"
)

//go:generate goyacc.exe -o win32.y.go win32.y

var typeMap = map[string]string{
	"LPCWSTR": "*uint16",
	"LPWSTR":  "*uint16",
	"LPCTSTR": "*uint16",
	"BOOL":    "bool",
	// "UINT":    "uint32",
	
	"PROCESS_DPI_AWARENESS": "ProcessDpiAwareness",
}

func main() {
	out := flag.String("o", "", "output file")
	pkg := flag.String("pkg", "main", "package")
	var buf bytes.Buffer
	flag.Parse()

	buf.WriteString("// Code generated by mkapisys. DO NOT EDIT.\n")
	fmt.Fprintf(&buf, "package %s\n", *pkg)
	for _, arg := range flag.Args() {
		scanner := new(scanner.Scanner)
		// body, err := ioutil.ReadFile(arg)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// scanner.Init(string(body))
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln("T23", arg, err)
		}
		scanner.Init(f)
		for _, def := range Parse(scanner) {
			switch d := def.(type) {
			case *Funcdef:
				outFuncdef(d, &buf)
			case *Typedef:
				outTypedef(d, &buf)
			default:
				log.Fatalln("T48: Not Supportd", d)
			}
		}
	}
	if *out != "" {
		err := ioutil.WriteFile(*out, buf.Bytes(), 0644)
		if err != nil {
			log.Panicf("writing output: %s", err)
		}
	} else {
		os.Stdout.Write(buf.Bytes())
	}
}

func removeLastW(name string) string {
	return strings.TrimSuffix(name, "W")
}

func upFirst(name string) string {
	first := []rune(name)[0]
	if !unicode.IsUpper(first) {
		return string(unicode.ToUpper(first)) + name[1:]
	}
	return name
}

func convTypeParam(p Param) string {
	pname := p.pname.literal
	upperIndex := strings.IndexFunc(pname, unicode.IsUpper)
	if upperIndex > 0 {
		prefix := pname[:upperIndex]
		if prefix == "b" && convType(*p.typespec) == "DWORD" {
			return "bool"
		}
	}
	return convType(*p.typespec)
}

// type string
func convType(t Typespec) (ts string) {
	var ok bool
	if ts, ok = typeMap[t.name.literal]; !ok {
		ts = t.name.literal
	}
	if t.isPointer {
		ts = "*" + ts
	}
	if t.isArray {
		ts = "[" + strconv.FormatInt(int64(t.arraySize), 10) + "]" + ts
	}
	return ts
}

func outFuncdef(d *Funcdef, out *bytes.Buffer) {
	var ns, fv, fn string
	var noerr bool
	var noreturn bool
	var errorOnly bool
	for _, q := range d.funcQuals {
		switch q.qt {
		case QtNs:
			ns = q.val
		case QtFailRetVal:
			fv = q.val
		case QtNoerr:
			noerr = true
		case QtNoreturn:
			noreturn = true
		case QtErrorOnly:
			errorOnly = true
		case QtFuncname:
			fn = q.val
		default:
			log.Panicln("T41:", q)
		}
	}
	if noerr && errorOnly {
		log.Panicln("T132: can not both noerr and erroronly", d.funcQuals)
	}

	if fn == "" {
		fmt.Fprintf(out, "//sys %s(", upFirst(removeLastW(d.funcname.literal)))
	} else {
		fmt.Fprintf(out, "//sys %s(", fn)
	}

	for i, p := range d.params {
		if i > 0 {
			out.WriteString(", ")
		}
		fmt.Fprintf(out, "%s %s", p.pname.literal, convTypeParam(*p))
	}
	fmt.Fprintf(out, ")")
	if !noreturn {
		if errorOnly {
			out.WriteString(" (err error)")
		} else {
			fmt.Fprintf(out, " (r %s", convType(*d.typespec))
			if !noerr {
				out.WriteString(", err error")
			}
			out.WriteString(")")	
		}
	}
	if fv != "" {
		fmt.Fprintf(out, " [failretval==%s]", fv)
	}
	if ns != "" {
		fmt.Fprintf(out, " = %s.%s\n", ns, d.funcname.literal)
	} else {
		fmt.Fprintf(out, " = %s\n", d.funcname.literal)
	}
}

func outTypedef(td *Typedef, out *bytes.Buffer) {
	cStructName := td.defnames[0].name.literal
	goStructName := strcase.UpperCamelCase(removeLastW(cStructName))
	fmt.Fprintf(out, "type %s struct {\n", goStructName)
	for _, m := range td.members {
		name := strings.TrimLeftFunc(m.pname.literal, unicode.IsLower)
		var goname string
		if len(name) == 0 {
			goname = strings.Title(m.pname.literal)
		} else {
			goname = strings.Title(name)
		}
		fmt.Fprintf(out, "\t%s %s\n", goname, convTypeSt(*m.typespec))
	}
	fmt.Fprintf(out, "}\n\n")

	for _, name := range td.defnames {
		cName := name.name.literal
		if name.isPointer {
			typeMap[cName] = "*" + goStructName
		} else {
			typeMap[cName] = goStructName
		}
	}
}

// type string for struct member
func convTypeSt(ty Typespec) (tyStr string) {
	switch ty.name.literal {
	case "BOOL":	// prevent from changing member size
		tyStr = ty.name.literal
	case "int":
		tyStr = "int32"
	default:
		tyStr = convType(ty)
	}
	return tyStr
}
