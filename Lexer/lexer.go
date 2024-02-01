package Lexer

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"simpleInterpreter/Misc"
	"simpleInterpreter/Token"
)

func Keywords() []string {
	return []string{"var", "func"}
}

type Lexer struct {
	text        string
	pos         int
	currentRune rune
}

func NewLexer(text string) Lexer {
	if len(text) == 0 {
		return Lexer{text, 0, 0}
	}

	return Lexer{text, 0, rune(text[0])}
}

func (l *Lexer) advance() {
	l.pos++

	if l.pos >= len(l.text) {
		l.currentRune = 0
		return
	}

	l.currentRune = rune(l.text[l.pos])
}

func (l *Lexer) Peek(offset int) rune {
	index := l.pos + offset

	if index >= len(l.text) {
		return 0
	}

	char := rune(l.text[index])

	for unicode.IsDigit(char) {
		index++

		if index >= len(l.text) {
			return 0
		}

		char = rune(l.text[index])
	}

	return char
}

func (l *Lexer) GetNextToken() (Token.Token, error) {
	for l.currentRune != 0 {
		if unicode.IsSpace(l.currentRune) {
			l.skipWhiteSpace()
			continue
		}

		if unicode.IsDigit(l.currentRune) {
			token, err := l.integer()
			return token, err
		}

		if unicode.IsLetter(l.currentRune) || l.currentRune == '_' {
			token := l.identifier()
			return token, nil
		}

		switch l.currentRune {
		case '+':
			l.advance()
			return Token.NewToken(Token.TtPlus, "+"), nil

		case '-':
			l.advance()
			return Token.NewToken(Token.TtMinus, "-"), nil

		case '*':
			l.advance()
			return Token.NewToken(Token.TtStar, "*"), nil

		case '/':
			l.advance()
			return Token.NewToken(Token.TtSlash, "/"), nil

		case '(':
			l.advance()
			return Token.NewToken(Token.TtLParen, "("), nil

		case ')':
			l.advance()
			return Token.NewToken(Token.TtRParen, ")"), nil

		case '{':
			l.advance()
			return Token.NewToken(Token.TtLCurly, "{"), nil

		case '}':
			l.advance()
			return Token.NewToken(Token.TtRCurly, "}"), nil

		case '=':
			l.advance()
			return Token.NewToken(Token.TtEquals, "="), nil

		case ';':
			l.advance()
			return Token.NewToken(Token.TtSemi, ";"), nil

		case ':':
			l.advance()
			return Token.NewToken(Token.TtColon, ":"), nil

		default:
			return Token.ErrorToken, errors.New(fmt.Sprintf("error: unrecognized character `%c`", l.currentRune))
		}
	}

	return Token.NewToken(Token.TtEof, "0"), nil
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.currentRune) {
		l.advance()
	}
}

func (l *Lexer) integer() (Token.Token, error) {
	var sValue string

	for unicode.IsDigit(l.currentRune) {
		sValue += string(l.currentRune)
		l.advance()
	}

	_, err := strconv.ParseInt(sValue, 10, 32)

	if err != nil {
		return Token.ErrorToken, err
	}

	return Token.NewToken(Token.TtInteger, sValue), nil
}

func (l *Lexer) identifier() Token.Token {
	var sValue string

	for unicode.IsLetter(l.currentRune) || unicode.IsDigit(l.currentRune) || l.currentRune == '_' {
		sValue += string(l.currentRune)
		l.advance()
	}

	tType := Token.TtIdentifier

	if Misc.SliceContainsItem(Keywords(), sValue) {
		tType = Token.TtKeyword
	}

	return Token.NewToken(tType, sValue)
}
