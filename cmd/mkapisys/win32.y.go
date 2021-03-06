// Code generated by goyacc -o win32.y.go win32.y. DO NOT EDIT.

//line win32.y:2

package main

import __yyfmt__ "fmt"

//line win32.y:3

import (
	"log"
	"strconv"
	"text/scanner"
)

type Token struct {
	token   int
	literal string
	pos     scanner.Position // Position
}

type Funcdef struct {
	typespec  *Typespec
	funcname  Token
	params    []*Param
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
	QtErrorOnly
)

type FuncQual struct {
	qt  qualtype
	val string
}

type Typespec struct {
	name      Token
	isArray   bool
	arraySize int
	isPointer bool
}

type Param struct {
	typespec *Typespec
	pname    Token
}

type Typedef struct {
	name     Token
	members  []*Param
	defnames []*TypedefName
}

type TypedefName struct {
	name      Token
	isPointer bool
}

var keywords = map[string]int{
	"struct":     STRUCT,
	"typedef":    TYPEDEF,
	"ns":         NS,
	"failretval": FAILRETVAL,
	"funcname":   FUNCNAME,
	"noerr":      NOERR,
	"const":      CONST,
	"noreturn":   NORETURN,
	"erroronly":  ERRORONLY,
}

//line win32.y:79
type yySymType struct {
	yys   int
	token Token
	def   interface{}
	defs  []interface{}

	aElement interface{}

	params    []*Param
	funcquals []*FuncQual
	typespecs []*Typespec
	members   []*Param
	defnames  []*TypedefName
}

const IDENT = 57346
const STRUCT = 57347
const NUMBER = 57348
const NS = 57349
const FAILRETVAL = 57350
const FUNCNAME = 57351
const NOERR = 57352
const NORETURN = 57353
const TYPEDEF = 57354
const CONST = 57355
const ERRORONLY = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"STRUCT",
	"NUMBER",
	"NS",
	"FAILRETVAL",
	"FUNCNAME",
	"NOERR",
	"NORETURN",
	"TYPEDEF",
	"CONST",
	"ERRORONLY",
	"'('",
	"')'",
	"';'",
	"'@'",
	"'-'",
	"','",
	"'['",
	"']'",
	"'*'",
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line win32.y:304

type LexerWrapper struct {
	// s		  *Scanner
	scanner.Scanner
	// recentLit  string
	// recentPos  scanner.Position
	defs []interface{}
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

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 1,
	-2, 10,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 10,
	-1, 35,
	4, 7,
	-2, 25,
	-1, 36,
	25, 33,
	-2, 7,
	-1, 49,
	25, 33,
	-2, 7,
}

const yyPrivate = 57344

const yyLast = 66

var yyAct = [...]int{

	54, 41, 45, 42, 48, 56, 47, 36, 63, 52,
	43, 33, 61, 31, 44, 20, 21, 25, 22, 23,
	60, 6, 24, 13, 57, 46, 32, 13, 50, 49,
	35, 17, 5, 59, 37, 40, 18, 62, 47, 39,
	19, 15, 34, 30, 29, 28, 27, 1, 51, 53,
	14, 58, 16, 38, 7, 11, 10, 12, 9, 8,
	55, 4, 64, 3, 2, 26,
}
var yyPact = [...]int{

	9, -1000, 9, -1000, -1000, 18, 31, 5, -1000, -1000,
	-1000, -1000, -1000, 8, -1000, 42, 41, -1000, 40, -1000,
	39, 7, -1000, -1000, -1000, 38, 15, -1000, -1000, -17,
	-1000, -1000, 28, -1000, -1000, 18, 18, -1000, -6, -1000,
	2, -21, 12, 11, 18, -12, 34, -1000, 1, 18,
	-1000, -1000, 27, -1000, 3, -8, -1000, 33, -1000, -14,
	-1000, 1, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 65, 2, 47, 64, 63, 61, 60, 3, 59,
	58, 57, 56, 55, 54, 35, 53, 32, 0, 1,
	52,
}
var yyR1 = [...]int{

	0, 3, 3, 4, 4, 5, 15, 20, 20, 1,
	17, 17, 14, 14, 14, 14, 14, 9, 10, 10,
	10, 12, 13, 13, 11, 16, 16, 16, 8, 8,
	8, 2, 6, 19, 19, 18, 18, 7, 7,
}
var yyR2 = [...]int{

	0, 0, 2, 1, 1, 7, 2, 0, 1, 1,
	0, 2, 1, 1, 1, 1, 1, 3, 3, 4,
	3, 2, 2, 2, 3, 0, 1, 3, 2, 5,
	3, 1, 8, 0, 3, 1, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -3, -4, -5, -6, -17, 12, -14, -9, -10,
	-12, -13, -11, 18, -3, -15, -20, 13, 5, -17,
	7, 8, 10, 11, 14, 9, -1, 4, 4, 4,
	4, 6, 19, 4, 4, 15, 24, 6, -16, -8,
	-15, -19, -8, 16, 20, -2, 23, 4, 25, 17,
	17, -8, 21, -2, -18, -7, 4, 23, -19, 6,
	17, 20, 4, 22, -18,
}
var yyDef = [...]int{

	-2, -2, -2, 3, 4, 7, 0, 10, 12, 13,
	14, 15, 16, 0, 2, 0, 0, 8, 0, 11,
	0, 0, 21, 22, 23, 0, 0, 9, 6, 0,
	17, 18, 0, 20, 24, -2, -2, 19, 0, 26,
	0, 0, 0, 0, 7, 28, 0, 31, 0, -2,
	5, 27, 0, 30, 0, 35, 37, 0, 34, 0,
	32, 0, 38, 29, 36,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 23, 3, 20, 19, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 17,
	3, 3, 3, 3, 18, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 21, 3, 22, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 24, 3, 25,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line win32.y:109
		{
			yyVAL.defs = nil
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.defs = yyVAL.defs
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:116
		{
			yyVAL.defs = append([]interface{}{yyDollar[1].def}, yyDollar[2].defs...)
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.defs = yyVAL.defs
			}
		}
	case 5:
		yyDollar = yyS[yypt-7 : yypt+1]
//line win32.y:129
		{
			typespec := yyDollar[2].aElement.(*Typespec)
			yyVAL.def = &Funcdef{
				funcQuals: yyDollar[1].funcquals,
				typespec:  typespec,
				funcname:  yyDollar[3].token,
				params:    yyDollar[5].params,
			}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:141
		{
			yyVAL.aElement = &Typespec{
				name:      yyDollar[2].token,
				isArray:   false,
				arraySize: 0,
			}
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
//line win32.y:158
		{
			yyVAL.funcquals = []*FuncQual{}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:162
		{
			funcqual := yyDollar[1].aElement.(*FuncQual)
			yyVAL.funcquals = append([]*FuncQual{funcqual}, yyDollar[2].funcquals...)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:175
		{
			yyVAL.aElement = &FuncQual{QtNs, yyDollar[3].token.literal}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:181
		{
			yyVAL.aElement = &FuncQual{QtFailRetVal, yyDollar[3].token.literal}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line win32.y:185
		{
			yyVAL.aElement = &FuncQual{QtFailRetVal, "-" + yyDollar[4].token.literal}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:189
		{
			yyVAL.aElement = &FuncQual{QtFailRetVal, yyDollar[3].token.literal}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:194
		{
			yyVAL.aElement = &FuncQual{QtNoerr, ""}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:199
		{
			yyVAL.aElement = &FuncQual{QtNoreturn, ""}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:204
		{
			yyVAL.aElement = &FuncQual{QtErrorOnly, ""}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:209
		{
			yyVAL.aElement = &FuncQual{QtFuncname, yyDollar[3].token.literal}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
//line win32.y:214
		{
			yyVAL.params = []*Param{}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line win32.y:218
		{
			param := yyDollar[1].aElement.(*Param)
			yyVAL.params = []*Param{param}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:223
		{
			param := yyDollar[3].aElement.(*Param)
			yyVAL.params = append(yyDollar[1].params, param)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:230
		{
			typespec := yyDollar[1].aElement.(*Typespec)
			yyVAL.aElement = &Param{typespec, yyDollar[2].token}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
//line win32.y:235
		{
			i, err := strconv.ParseInt(yyDollar[4].token.literal, 10, 32)
			if err != nil {
				log.Panicln("T208:", err)
			}
			ts := yyDollar[1].aElement.(*Typespec)
			ts.isArray = true
			ts.arraySize = int(i)

			yyVAL.aElement = &Param{ts, yyDollar[2].token}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:247
		{
			ts := yyDollar[1].aElement.(*Typespec)
			ts.isPointer = true
			yyVAL.aElement = &Param{ts, yyDollar[3].token}
		}
	case 32:
		yyDollar = yyS[yypt-8 : yypt+1]
//line win32.y:257
		{
			yyVAL.def = &Typedef{
				name:     yyDollar[3].token,
				members:  yyDollar[5].members,
				defnames: yyDollar[7].defnames,
			}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
//line win32.y:267
		{
			yyVAL.members = []*Param{}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:271
		{
			member := yyDollar[1].aElement.(*Param)
			yyVAL.members = append([]*Param{member}, yyDollar[3].members...)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line win32.y:278
		{
			defname := yyDollar[1].aElement.(*TypedefName)
			yyVAL.defnames = []*TypedefName{defname}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line win32.y:283
		{
			defname := yyDollar[1].aElement.(*TypedefName)
			yyVAL.defnames = append([]*TypedefName{defname}, yyDollar[3].defnames...)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line win32.y:290
		{
			yyVAL.aElement = &TypedefName{
				name:      yyDollar[1].token,
				isPointer: false,
			}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line win32.y:297
		{
			yyVAL.aElement = &TypedefName{
				name:      yyDollar[2].token,
				isPointer: true,
			}
		}
	}
	goto yystack /* stack new state and value */
}
