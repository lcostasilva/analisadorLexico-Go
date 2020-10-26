package main

import "bufio"

const (
	MAX_LEXEME_LENGTH = 128
)

type Token int

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

type Result struct {
	pos   Position
	token Token
	lit   string
}

const (
	//Characters
	LETTER     = iota
	DIGIT      = iota
	WHITESPACE = iota

	//Tokens
	EOF                  = iota
	EXPR_END             = iota
	KEYWORKS             = iota
	HASH                 = iota
	DOUBLE_INTERROGATION = iota
	RIGHT_ASSIGN         = iota
	INTERROGATION        = iota
	TRUE                 = iota
	FALSE                = iota
	ADDRESS              = iota

	//Literais_start
	IDENT  = iota // main
	INT    = iota // 12345
	FLOAT  = iota // 123.45
	IMAG   = iota // 123.45i
	CHAR   = iota // 'a'
	STRING = iota // "abc"

	//Operadores_start
	ADD = iota // +
	SUB = iota // -
	MUL = iota // *
	DIV = iota // /
	REM = iota // %
	AND = iota // &
	OR  = iota // |
	XOR = iota // ^
	LSS = iota // <
	GTR = iota // >

	ASSIGN = iota // =
	NOT    = iota // !

	LEFT_PAREN  = iota // (
	VETOR       = iota // [
	ARRAY       = iota // [
	LEFT_BRACE  = iota // {
	RIGHT_PAREN = iota // )
	RIGHT_BRACE = iota // }

	COMMA     = iota // ,
	PERIOD    = iota // .
	SEMICOLON = iota // ;
	COLON     = iota // :

	//Keywords_start
	BREAK    = iota
	CASE     = iota
	CHAN     = iota
	CONST    = iota
	CONTINUE = iota

	DEFAULT     = iota
	DEFER       = iota
	ELSE        = iota
	FALLTHROUGH = iota
	FOR         = iota

	FUNC   = iota
	GO     = iota
	GOTO   = iota
	IF     = iota
	IMPORT = iota

	INTERFACE = iota
	MAP       = iota
	PACKAGE   = iota
	RANGE     = iota
	RETURN    = iota

	SELECT = iota
	STRUCT = iota
	SWITCH = iota
	TYPE   = iota
	VAR    = iota

	/*---------------------------*/
	COMMENT       = iota
	COMMENT_BLOCK = iota
	UNKNOWN       = iota
)

var tokens = []string{
	EOF:                  "EOF",
	CHAR:                 "CHAR",
	LETTER:               "LETTER",
	DIGIT:                "DIGIT",
	WHITESPACE:           "WHITESPACE",
	IDENT:                "IDENT",
	ASSIGN:               "ASSIGN",
	ADD:                  "ADD",
	SUB:                  "SUB",
	DIV:                  "DIV",
	LEFT_PAREN:           "LEFT_PAREN",
	RIGHT_PAREN:          "RIGHT_PAREN",
	EXPR_END:             "EXPR_END",
	VETOR:                "VETOR",
	ARRAY:                "ARRAY",
	LEFT_BRACE:           "LEFT_BRACE",
	RIGHT_BRACE:          "RIGHT_BRACE",
	GTR:                  "GTR",
	LSS:                  "LSS",
	KEYWORKS:             "KEYWORKS",
	CONTINUE:             "CONTINUE",
	BREAK:                "BREAK",
	OR:                   "OR",
	HASH:                 "HASH",
	NOT:                  "NOT",
	XOR:                  "XOR",
	MUL:                  "MUL",
	IF:                   "IF",
	GO:                   "GO",
	AND:                  "AND",
	FLOAT:                "FLOAT",
	STRING:               "STRING",
	RETURN:               "RETURN",
	TRUE:                 "TRUE",
	TYPE:                 "TYPE",
	IMPORT:               "IMPORT",
	FALLTHROUGH:          "FALLTHROUGH",
	CONST:                "CONST",
	SELECT:               "SELECT",
	IMAG:                 "IMAG",
	INT:                  "INT",
	PACKAGE:              "PACKAGE",
	SWITCH:               "SWITCH",
	DOUBLE_INTERROGATION: "DOUBLE_INTERROGATION",
	RIGHT_ASSIGN:         "RIGHT_ASSIGN",
	GOTO:                 "GOTO",
	INTERFACE:            "INTERFACE",
	FUNC:                 "FUNC",
	DEFER:                "DEFER",
	INTERROGATION:        "INTERROGATION",
	SEMICOLON:            "SEMICOLON",
	ELSE:                 "ELSE",
	CASE:                 "CASE",
	CHAN:                 "CHAN",
	DEFAULT:              "DEFAULT",
	RANGE:                "RANGE",
	FALSE:                "FALSE",
	ADDRESS:              "ADDRESS",
	COMMA:                "COMMA",
	PERIOD:               "PERIOD",
	COLON:                "COLON",
	FOR:                  "FOR",
	VAR:                  "VAR",
	MAP:                  "MAP",
	COMMENT:              "COMMENT",
	COMMENT_BLOCK:        "COMMENT_BLOCK",
	UNKNOWN:              "UNKNOWN",
	STRUCT:               "STRUCT",
}
