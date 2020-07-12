%{

package main

import (
	"log"
	"strconv"
	"text/scanner"
)

type Token struct {
	token int
	literal string
	pos scanner.Position // Position
}

type Funcdef struct {
	typespec *Typespec
	funcname Token
	params []*Param
	funcQuals []*FuncQual
}

type qualtype int
const (
	QtUnknown qualtype = iota
	QtNs
	QtFailRetVal
	QtFuncname
	QtNoerr
	QtNoreturn
)

type FuncQual struct {
	qt qualtype
	val string
}

type Typespec struct {
	name Token
	isArray bool
	arraySize int
	isPointer bool
}

type Param struct {
	typespec *Typespec
	pname Token
}

type Typedef struct {
	name Token
	members []*Param
	defnames []*TypedefName
}

type TypedefName struct {
	name Token
	isPointer bool
}

var keywords = map[string]int{
	"struct":	STRUCT,
	"typedef":	TYPEDEF,
	"ns":		NS,
	"failretval":	FAILRETVAL,
	"funcname":	FUNCNAME,
	"noerr":	NOERR,
	"const":	CONST,
	"noreturn":     NORETURN,
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	token		Token
	def		interface{}
	defs		[]interface{}

	aElement	interface{}

	params		[]*Param
	funcquals	[]*FuncQual
	typespecs	[]*Typespec
	members		[]*Param
	defnames	[]*TypedefName
}

%type<token> funcname pname

%type<defs>	 defs
%type<def>	  def funcdef typedef
%type<aElement> defname param nsqual failretvalqual funcnamequal noerrqual noreturnqual funcqual typespec

%type<params>   params 
%type<funcquals>	funcquals
%type<defnames> defnames
%type<members>  members

%token<token> IDENT STRUCT NUMBER NS FAILRETVAL FUNCNAME NOERR NORETURN TYPEDEF CONST

%%

defs:
	{
		$$ = nil
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.defs = $$
		}
	}
	| def defs
	{
		$$ = append([]interface{}{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.defs = $$
		}
	}

def
	: funcdef
	| typedef

funcdef
	: funcquals typespec funcname '(' params ')' ';'
	  {
		  typespec :=  $2.(*Typespec)
		  $$ = &Funcdef{
			  funcQuals: $1,
			  typespec: typespec,
			  funcname: $3,
			  params: $5,
		  }
	  }

typespec
	: constqual IDENT
	{
		$$ = &Typespec{
			name: $2,
			isArray: false,
			arraySize: 0,
		}
	}

constqual
	:
	| CONST

funcname
	: IDENT

funcquals
	:
	{
		$$ = []*FuncQual{}
	}
	| funcqual funcquals
	{
		funcqual := $1.(*FuncQual)
		$$ = append([]*FuncQual{funcqual}, $2...)
	}

funcqual
	: nsqual
	| failretvalqual
	| noerrqual
	| noreturnqual
	| funcnamequal

nsqual: '@' NS IDENT
	{
		$$ = &FuncQual{QtNs, $3.literal}
	}

failretvalqual
	: '@' FAILRETVAL NUMBER
	{
		$$ = &FuncQual{QtFailRetVal, $3.literal}
	}
	| '@' FAILRETVAL '-' NUMBER
	{
		$$ = &FuncQual{QtFailRetVal, "-" + $4.literal}
	}
	| '@' FAILRETVAL IDENT
	{
		$$ = &FuncQual{QtFailRetVal, $3.literal}
	}

noerrqual: '@' NOERR
	{
		$$ = &FuncQual{QtNoerr, ""}
	}

noreturnqual: '@' NORETURN
	{
		$$ = &FuncQual{QtNoreturn, ""}
	}

funcnamequal: '@' FUNCNAME IDENT
	{
		$$ = &FuncQual{QtFuncname, $3.literal}
	}

params:
	{
		$$ = []*Param{}
	}
	| param
	{
		param := $1.(*Param)
		$$ = []*Param{param}
	}
	| params ',' param
	{
		param := $3.(*Param)
		$$ = append($1, param)
	}

param
	: typespec pname
	{
		typespec := $1.(*Typespec)
		$$ = &Param{typespec, $2}
	}
	| typespec pname '[' NUMBER ']'
	{
		i, err := strconv.ParseInt($4.literal, 10, 32)
		if err != nil {
			log.Panicln("T208:", err)
		}
		ts := $1.(*Typespec)
		ts.isArray = true
		ts.arraySize = int(i)

		$$ = &Param{ts, $2}
	}
	| typespec '*' pname
	{
		ts := $1.(*Typespec)
		ts.isPointer = true
		$$ = &Param{ts, $3}
	}

pname: IDENT

typedef
	: TYPEDEF STRUCT IDENT '{' members '}' defnames ';'
	{
		$$ = &Typedef {
			name: $3,
			members: $5,
			defnames: $7,
		}
	}

members
	:
	{
		$$ = []*Param{}
	}
	| param ';' members
	{
		member := $1.(*Param)
		$$ = append([]*Param{member}, $3...)
	}

defnames
	: defname
	{
		defname := $1.(*TypedefName)
		$$ = []*TypedefName{defname}
	}
	| defname ',' defnames
	{
		defname := $1.(*TypedefName)
		$$ = append([]*TypedefName{defname}, $3...)
	}
 
defname
	: IDENT
	{
		$$ = &TypedefName{
			name: $1,
			isPointer: false,
		}
	}
	| '*' IDENT
	{
		$$ = &TypedefName{
			name: $2,
			isPointer: true,
		}
	}

%%

type LexerWrapper struct {
	// s		  *Scanner
	scanner.Scanner
	// recentLit  string
	// recentPos  scanner.Position
	defs		  []interface{}
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	token := int(l.Scan())
	switch token {
	case scanner.Int:
		token = NUMBER
	case scanner.Ident:
		if keyword, ok := keywords[l.TokenText()]; ok {
			token = keyword
		} else {
			token = IDENT
		}
	}

	lval.token = Token{token: token, literal: l.TokenText(), pos: l.Position}
	return token
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.Position.Line, l.Position.Column, l.TokenText(), e)
		// l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *scanner.Scanner) []interface{} {
	l := LexerWrapper{*s, nil}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.defs
}
