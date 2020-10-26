package main

import (
	"io"
	"unicode"
	"unicode/utf8"
)

func (l *Lexer) Analyzer() (Position, Token, string) {
	string_lit, _ := utf8.DecodeRuneInString(`"`)
	for {
		char, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, "FIM"
			}
			panic(err)
		}

		l.pos.column++
		switch char {
		case '\n':
			l.resetPosition()
		case ';':
			return l.pos, SEMICOLON, ";"
		case ':':
			return l.pos, COLON, ":"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			char, _, _ := l.reader.ReadRune()
			l.backup()
			if char == '[' {
				tok, lit := l.lexVerifyArrayBlock()
				lit = "*" + lit
				return l.pos, tok, lit
			} else {
				return l.pos, MUL, "*"
			}
		case '%':
			return l.pos, REM, "%"
		case '&':
			return l.pos, AND, "&"
		case '|':
			return l.pos, OR, "|"
		case '^':
			return l.pos, XOR, "^"
		case '!':
			return l.pos, NOT, "!"
		case '/':
			return l.lexCommet()
		case '=':
			return l.pos, ASSIGN, "="
		case '(':
			return l.pos, LEFT_PAREN, "("
		case ')':
			return l.pos, RIGHT_PAREN, ")"
		case '[':
			l.backup()
			tok, lit := l.lexVerifyArrayBlock()
			return l.pos, tok, lit
		case '{':
			return l.pos, LEFT_BRACE, "{"
		case '}':
			return l.pos, RIGHT_BRACE, "}"
		case '>':
			return l.pos, LSS, ">"
		case '<':
			return l.pos, GTR, "<"
		case ',':
			return l.pos, COMMA, ","
		case '.':
			return l.pos, PERIOD, "."
		case string_lit:
			lit := l.lexStringLit()
			return l.pos, STRING, lit
		default:
			if unicode.IsSpace(char) {
				continue // nothing to do here, just move on
			} else if unicode.IsDigit(char) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(char) || char == '_' {
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				t := lexKeywords(lit)
				if t != IDENT {
					switch t {
					case FOR:
						lit = lit + l.lexVerifyBlock()
						break
					case SWITCH:
						lit = lit + l.lexVerifyBlock()
						break
					case IF:
						lit = lit + l.lexIfBlock()
						break
					}
				}
				return startPos, t, lit
			} else {
				return l.pos, UNKNOWN, string(char)
			}
		}
	}
}
