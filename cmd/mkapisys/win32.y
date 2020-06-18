%{

package main

import (
    "log"
    "text/scanner"
)

type Token struct {
    token int
    literal string
    pos scanner.Position // Position
}

type Funcdef struct {
    typespec Token
    funcname Token
    params []*Param
    funcQuals []*FuncQual
}

type qualtype int
const (
    QtUnknown qualtype = iota
    QtNs
    QtFailRetVal
    QtNoerr
)

type FuncQual struct {
    qt qualtype
    val string
}

type Param struct {
    typespec Token
    pname Token
}

var keywords = map[string]int{
	"struct":     STRUCT,
	"ns":         NS,
	"failretval": FAILRETVAL,
	"noerr":      NOERR,
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
    funcdef *Funcdef
    funcdefs []*Funcdef
    token Token
    param *Param
    params []*Param
    funcqual *FuncQual
    funcquals []*FuncQual
}

%type<funcdefs> funcdefs
%type<funcdef> funcdef
%type<token> typespec funcname pname
%type<params> params 
%type<param> param
%type<funcqual> nsqual failretvalqual noerrqual funcqual
%type<funcquals> funcquals

%token<token> IDENT STRUCT NUMBER NS FAILRETVAL NOERR

%%

funcdefs
    :
	{
		$$ = nil
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.funcdefs = $$
		}
	}
    | funcdef funcdefs
	{
		$$ = append([]*Funcdef{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.funcdefs = $$
		}
	}

funcdef
    : funcquals typespec funcname '(' params ')' ';'
      {
          $$ = &Funcdef{
              funcQuals: $1,
              typespec: $2,
              funcname: $3,
              params: $5,
          }
      }

typespec
    : IDENT

funcname
    : IDENT

funcquals
    :
    {
        $$ = []*FuncQual{}
    }
    | funcqual funcquals
    {
        $$ = append([]*FuncQual{$1}, $2...)
    }

funcqual
    : nsqual
    | failretvalqual
    | noerrqual

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

params:
      {
          $$ = []*Param{}
      }
    | param
      {
          $$ = []*Param{$1}
      }
    | params ',' param
      {
          $$ = append($1, $3)
      }

param: typespec pname
    {
        $$ = &Param{$1, $2}
    }

pname: IDENT

%%

type LexerWrapper struct {
	// s          *Scanner
    scanner.Scanner
	// recentLit  string
	// recentPos  scanner.Position
	funcdefs []*Funcdef
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

// func (l *LexerWrapper) Lex(lval *yySymType) int {
// 	tok, lit, pos := l.s.Scan()
// 	if tok == EOF {
// 		return 0
// 	}
// 	lval.token = Token{tok: tok, lit: lit, pos: pos}
// 	l.recentLit = lit
// 	l.recentPos = pos
// 	return tok
// }

// func (l *LexerWrapper) Error(e string) {
// 	log.Fatalf("Line %d, Column %d: %q %s",
// 		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
// }

func Parse(s *scanner.Scanner) []*Funcdef {
	l := LexerWrapper{*s, []*Funcdef{}}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.funcdefs
}
