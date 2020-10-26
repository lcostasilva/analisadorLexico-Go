package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.column--
}

func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexVerifyBlock() string {
	var lit string
	keys := 0

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				panic("Problema na formatação do fechamento do bloco")
			}
		}

		l.pos.column++
		if r != '}' {
			if r == '{' {
				keys++
			}
			lit = lit + string(r)
		} else if keys > 0 {
			lit = lit + string(r)
			keys--
			if keys == 0 {
				return lit
			}
		} else {
			panic("Problema na formatação do inicio bloco")
		}
	}
}

func (l *Lexer) lexIfBlock() string {
	var lit string
	keys := 0

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				panic("IF com problema na formatação do fechamento do bloco")
			}
		}

		l.pos.column++
		if r != '}' {
			if r == '{' {
				keys++
			}
			lit = lit + string(r)
		} else if keys > 0 {
			lit = lit + string(r)
			keys--
			if keys == 0 {
				blockElse := l.lexVerifyElse()
				if blockElse == " else" {
					lit = lit + blockElse
				} else {
					return lit
				}
			}
		} else {
			panic("IF com problema na formatação do inicio bloco")
		}
	}
}

func (l *Lexer) lexVerifyElse() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				panic(err)
			}
		}

		l.pos.column++
		switch r {
		case '\n':
			l.resetPosition()
			return string(r)
		case '/':
			char, _, err := l.reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					panic(err)
				}
			}
			l.backup()
			if char != '/' && char != '*' {
				l.backup()
			}
			return ""
		default:
			if unicode.IsSpace(r) {
				lit = lit + string(r)
			} else if unicode.IsLetter(r) {
				l.backup()
				litIdent := l.lexIdent()
				if litIdent == "else" {
					lit = lit + litIdent
					return lit
				} else {
					panic("Erro apos a estrutura do bloco IF")
				}
			}
		}

	}
}

func (l *Lexer) lexVerifyArrayBlock() (Token, string) {
	var lit string
	numVetors := 0
	var token Token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				panic("Problema na formatação do bloco de array")
			}
		}

		l.pos.column++
		if r != '\n' {
			if r == '[' || r == ']' {
				numVetors++
				if numVetors == 2 {
					token = VETOR
				} else if numVetors > 2 && numVetors%2 == 0 {
					token = ARRAY
				}
			}
			if r == ',' {
				return token, lit
			}
			lit = lit + string(r)
		} else {
			return token, lit
		}
	}
}

func (l *Lexer) lexStringLit() string {
	var lit string = `"`
	for {
		char, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++
		if string(char) != `"` {
			lit = lit + string(char)
		} else {
			lit = lit + string(char)
			return lit
		}
	}
}
func (l *Lexer) lexCommet() (Position, Token, string) {
	var lit string = `/`
	char, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return l.pos, COMMENT, lit
		}
	}
	if char != '/' && char != '*' {
		l.backup()
		return l.pos, DIV, "/"
	}
	if char == '*' {
		lit = lit + string(char)
		for {
			char, _, err := l.reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					// at the end of the identifier
					panic("Programa sem fechamento de comentario em bloco")
				}
			}

			l.pos.column++
			if char != '/' {
				lit = lit + string(char)
			} else {
				r := []rune(lit)
				last := r[len(lit)-1]
				lit = lit + string(char)
				if last == '*' {
					return l.pos, COMMENT_BLOCK, lit
				}
			}
		}
	}
	lit = lit + string(char)
	for {
		char, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the identifier
				return l.pos, COMMENT, lit
			}
		}

		l.pos.column++
		if char != '\n' {
			lit = lit + string(char)
		} else {
			return l.pos, COMMENT, lit
		}
	}
}

func (t Token) String() string {
	return tokens[t]
}

func (r Result) String() string {

	return fmt.Sprintf("| lin: %d col: %d | => \ttoken: %s\t lit: %s |\n", r.pos.line, r.pos.column, r.token, r.lit)
}

func lexKeywords(lit string) Token {
	switch lit {
	case "break":
		return BREAK
	case "default":
		return DEFAULT
	case "func":
		return FUNC
	case "interface":
		return INTERFACE
	case "select":
		return SELECT
	case "case":
		return CASE
	case "defer":
		return DEFER
	case "go":
		return GO
	case "map":
		return MAP
	case "struct":
		return STRUCT
	case "chan":
		return CHAN
	case "else":
		return ELSE
	case "goto":
		return GOTO
	case "package":
		return PACKAGE
	case "switch":
		return SWITCH
	case "const":
		return CONST
	case "fallthrough":
		return FALLTHROUGH
	case "if":
		return IF
	case "range":
		return RANGE
	case "type":
		return TYPE
	case "continue":
		return CONTINUE
	case "for":
		return FOR
	case "import":
		return IMPORT
	case "return":
		return RETURN
	case "false":
		return FALSE
	case "true":
		return TRUE
	case "var":
		return VAR
	default:
		return IDENT
	}
}
