package Token

import (
	"fmt"
)

type Type int32

const (
	TtError Type = iota - 1
	TtEof
	TtInteger
	TtPlus
	TtMinus
	TtStar
	TtSlash
	TtLParen
	TtRParen
	TtLCurly
	TtRCurly
	TtEquals
	TtSemi
	TtColon
	TtKeyword
	TtIdentifier
)

var (
	ErrorToken = NewToken(TtError, "0")

	ExprTypes = []Type{TtPlus, TtMinus}
	TermTypes = []Type{TtStar, TtSlash}
)

type Token struct {
	Type   Type
	SValue string
}

func NewToken(tType Type, sValue string) Token {
	return Token{tType, sValue}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Token<%s, %s>", TTypeToString(t.Type), t.SValue)
}

func TTypeToString(tType Type) string {
	switch tType {
	case TtError:
		return "ERROR"

	case TtEof:
		return "EOF"

	case TtInteger:
		return "INTEGER"

	case TtPlus:
		return "PLUS"

	case TtMinus:
		return "MINUS"

	case TtStar:
		return "STAR"

	case TtSlash:
		return "SLASH"

	case TtLParen:
		return "LPAREN"

	case TtRParen:
		return "RPAREN"

	case TtLCurly:
		return "LCURLY"

	case TtRCurly:
		return "RCURLY"

	case TtEquals:
		return "EQUALS"

	case TtSemi:
		return "SEMI"

	case TtColon:
		return "COLON"

	case TtKeyword:
		return "KEYWORD"

	case TtIdentifier:
		return "IDENTIFIER"

	default:
		return "NO_REPR"
	}
}
