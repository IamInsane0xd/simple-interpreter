package Parser

import (
	"errors"
	"fmt"

	"simpleInterpreter/AST"
	"simpleInterpreter/Lexer"
	"simpleInterpreter/Misc"
	"simpleInterpreter/Token"
)

type Parser struct {
	lexer        Lexer.Lexer
	currentToken Token.Token
}

func NewParser(lexer Lexer.Lexer) (Parser, error) {
	firstToken, err := lexer.GetNextToken()

	if err != nil {
		return Parser{Lexer.NewLexer(""), Token.NewToken(Token.TtError, "0")}, err
	}

	return Parser{lexer, firstToken}, nil
}

func (p *Parser) error(expected Token.Type) error {
	return errors.New(fmt.Sprintf("error: unexpected token `%s`, expected `%s`",
		Token.TTypeToString(p.currentToken.Type), Token.TTypeToString(expected)))
}

func (p *Parser) eat(tType Token.Type) error {
	if p.currentToken.Type != tType {
		return p.error(tType)
	}

	var err error
	p.currentToken, err = p.lexer.GetNextToken()

	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) expr() (AST.Node, error) {
	left, err := p.term()

	if err != nil {
		return AST.ErrorNode, err
	}

	for Misc.SliceContainsItem(Token.ExprTypes, p.currentToken.Type) {
		token := p.currentToken

		switch token.Type {
		case Token.TtPlus:
			err := p.eat(Token.TtPlus)

			if err != nil {
				return AST.ErrorNode, err
			}

			right, err := p.term()

			if err != nil {
				return AST.ErrorNode, err
			}

			return AST.NewBinOpNode(left, token, right), nil

		case Token.TtMinus:
			err := p.eat(Token.TtMinus)

			if err != nil {
				return AST.ErrorNode, err
			}

			right, err := p.term()

			if err != nil {
				return AST.ErrorNode, err
			}

			return AST.NewBinOpNode(left, token, right), nil

		default:
			break // * Do nothing
		}
	}

	return left, nil
}

func (p *Parser) term() (AST.Node, error) {
	left, err := p.factor()

	if err != nil {
		return AST.ErrorNode, err
	}

	for Misc.SliceContainsItem(Token.TermTypes, p.currentToken.Type) {
		token := p.currentToken

		switch token.Type {
		case Token.TtStar:
			err := p.eat(Token.TtStar)

			if err != nil {
				return AST.ErrorNode, err
			}

			right, err := p.factor()

			if err != nil {
				return AST.ErrorNode, err
			}

			return AST.NewBinOpNode(left, token, right), nil

		case Token.TtSlash:
			err := p.eat(Token.TtSlash)

			if err != nil {
				return AST.ErrorNode, err
			}

			right, err := p.factor()

			if err != nil {
				return AST.ErrorNode, err
			}

			return AST.NewBinOpNode(left, token, right), nil

		default:
			break // * Do nothing
		}
	}

	return left, nil
}

func (p *Parser) factor() (AST.Node, error) {
	switch p.currentToken.Type {
	case Token.TtInteger:
		token := p.currentToken
		err := p.eat(Token.TtInteger)

		if err != nil {
			return AST.ErrorNode, err
		}

		return AST.NewIntegerNode(token)

	case Token.TtLparen:
		err := p.eat(Token.TtLparen)

		if err != nil {
			return AST.ErrorNode, err
		}

		expr, err := p.expr()

		if err != nil {
			return AST.ErrorNode, err
		}

		err = p.eat(Token.TtRparen)

		if err != nil {
			return AST.ErrorNode, err
		}

		return expr, nil

	case Token.TtPlus:
		token := p.currentToken
		err := p.eat(Token.TtPlus)

		if err != nil {
			return AST.ErrorNode, err
		}

		factor, err := p.factor()

		if err != nil {
			return AST.ErrorNode, err
		}

		return AST.NewUnOpNode(token, factor), nil

	case Token.TtMinus:
		token := p.currentToken
		err := p.eat(Token.TtMinus)

		if err != nil {
			return AST.ErrorNode, err
		}

		factor, err := p.factor()

		if err != nil {
			return AST.ErrorNode, err
		}

		return AST.NewUnOpNode(token, factor), nil

	default:
		break // * Do nothing
	}

	return AST.ErrorNode, p.error(Token.TtInteger)
}

func (p *Parser) Parse() (AST.Node, error) {
	return p.expr()
}
