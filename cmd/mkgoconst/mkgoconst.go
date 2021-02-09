package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"io"
	"os"

	// "go/parser"
	"go/token"
	"go/types"
	"log"
	"strings"
)

const win32constDir = "../../win32const"

var p1 *string

func init() {
	p1 = flag.String("path", ".", "parse dir")
	flag.Parse()
}
func main() {
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, `
package win32api
//go:generate go run golang.org/x/tools/cmd/stringer -output win32api_const_string.go -type MessageId 
const (
`)
	mainPkg := types.NewPackage("main", "main")
	// win32constPkg, err := importer.Default().Import("github.com/turutcrane/win32api/internal/win32const")
	w32cFset := token.NewFileSet()
	win32constPkg, err := importer.ForCompiler(w32cFset, "source", nil).Import("github.com/turutcrane/win32api/internal/win32const")
	// win32constPkg, err := importer.ForCompiler(w32cFset, "source", nil).Import(win32constDir)
	if err != nil {
		log.Panicln("T34:", err)
	}
	mainPkg.SetImports([]*types.Package{
		win32constPkg,
	})
	// fmt.Println(win32constPkg.Scope().Names())
	mainPkg.Scope().Insert(types.NewPkgName(token.NoPos, mainPkg, "win32const", win32constPkg))

	// astFst := token.NewFileSet()
	// pkgs, err := parser.ParseDir(astFst, win32constDir, nil, parser.ParseComments)
	// if err != nil {
	// 	log.Panicln("T49:", err)
	// }
	// for _, p := range pkgs {
	// 	for n, f := range p.Files {
	// 		if !strings.HasSuffix(n, "_string.go") {
	// 			for _, d := range f.Decls {
	// 				switch decl := d.(type) {
	// 				case *ast.GenDecl:
	// 					if decl.Tok == token.CONST {
	// 						for _, s := range decl.Specs {
	// 							outOneValue(s, w32cFset, mainPkg)
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	win32constScope := win32constPkg.Scope()
	fset := token.NewFileSet()
	for _, n := range win32constScope.Names() {
		if !strings.HasPrefix(n, "_") {
			o := win32constScope.Lookup(n)
			if o.Exported() {
				var v types.TypeAndValue
				if v, err = types.Eval(fset, mainPkg, token.NoPos, "win32const."+n); err != nil {
					log.Panicln("T51:", err, o)
				}
				switch t := o.Type().(type) {
				case *types.Basic:
					i := t.Info()
					if i&(types.IsInteger|types.IsString) != 0 {
						fmt.Fprintf(buf, "%s = %v // %T\n", n, v.Value, v.Value)
					}
				case *types.Named:
					ut := t.Obj().Type().Underlying().(*types.Basic)
					if ut.Info()&(types.IsInteger|types.IsString) != 0 {
						fmt.Fprintf(buf, "%s %s = %v\n", n, t.Obj().Name(), v.Value)
					}
				}
			}
			// fmt.Println(o.Type())
		}
	}

	fmt.Fprint(buf, `)
`)
	formatted := goFormat(buf.Bytes())
	os.Stdout.Write(formatted)
}

func outOneValue(decl ast.Spec, fset *token.FileSet, mainPkg *types.Package, w io.Writer) {
	d := decl.(*ast.ValueSpec)
	if d.Names[0].IsExported() {
		n := d.Names[0].Name
		if v, err := types.Eval(fset, mainPkg, token.NoPos, "win32const."+n); err == nil {
			switch t := v.Type.(type) {
			case *types.Basic:
				i := t.Info()
				if i&(types.IsInteger|types.IsString) != 0 {
					fmt.Fprintf(w, "%s = %v // %T\n", n, v.Value, v.Value)
				}
			case *types.Named:
				ut := t.Obj().Type().Underlying().(*types.Basic)
				if ut.Info()&(types.IsInteger|types.IsString) != 0 {
					fmt.Fprintf(w, "%s %s = %v\n", n, t.Obj().Name(), v.Value)
				}
			default:
				log.Panicln("T106:", n, t, v)
			}
		} else {
			log.Panicln("T97:", err)
		}
	}
}

func  goFormat(src []byte) []byte {
	formatted, err := format.Source(src)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return src
	}
	return formatted
}
