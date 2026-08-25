package token

import (
	"strconv"
	"unicode"
)

type Token uint

// The list of tokens, mostly copied from Golang, along with some additional tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literal_end

	operator_beg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR

	RETURNS
	ADDRESS
	VALUE
	NEGATE
	AS

	keyword_end

	additional_beg
	// additional tokens, handled in an ad-hoc manner
	TILDE
	additional_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:    "முறி",
	CASE:     "நிலை",
	CHAN:     "வரிசை",
	CONST:    "நிலையான",
	CONTINUE: "தொடர்",

	//could just be left as இயல்பு
	DEFAULT: "இயல்புநிலை",

	DEFER:       "ஒத்திவை",
	ELSE:        "அல்லது",
	FALLTHROUGH: "அடுத்தநிலை",
	FOR:         "வரை",

	FUNC: "செயல்",

	GO:     "go",
	GOTO:   "செல்க",
	IF:     "எனில்",
	IMPORT: "இறக்கு",

	INTERFACE: "இடைமுகம்",
	MAP:       "அகராதி",
	PACKAGE:   "தொகுப்பு",
	RANGE:     "range",
	RETURN:    "திரும்பு",

	SELECT: "தேர்ந்தெடு",
	STRUCT: "வடிவம்",
	SWITCH: "மாற்று",
	TYPE:   "வகை",
	VAR:    "மாறக்கூடிய",

	RETURNS: "தரும்",
	ADDRESS: "முகவரி",
	VALUE:   "value",
	NEGATE:  "இல்லை",
	AS:      "என",

	TILDE: "~",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func IsKeyword(ident string) bool {
	_, isKeyword := keywords[ident]
	return isKeyword
}

func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENT
}

func IsIdentifier(name string) bool {
	// TODO: Have a check to make sure that Tamil combining marks are preceded by a
	// Tamil vowel or consonant. Maybe implement using a finite state machine.
	if name == "" || IsKeyword(name) {
		return false
	}
	for i, ch := range name {
		if !unicode.IsLetter(ch) && ch != '_' && (i == 0 || !unicode.IsDigit(ch) || !(unicode.Is(unicode.Tamil, ch) && unicode.IsMark(ch))) {
			return false
		}
	}
	return true
}
