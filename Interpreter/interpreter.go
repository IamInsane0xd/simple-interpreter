package Interpreter

import (
	"errors"
	"fmt"

	"simpleInterpreter/AST"
	"simpleInterpreter/Parser"
	"simpleInterpreter/Token"
)

var (
	symbolTable = make(map[string]int64)
)

type Interpreter struct {
	parser Parser.Parser
}

func NewInterpreter(parser Parser.Parser) Interpreter {
	return Interpreter{parser}
}

func visit(n *AST.Node) (int64, error) {
	switch n.Type {
	case AST.NtProgram:
		return visitProgram(n)

	case AST.NtInteger:
		return visitInteger(n), nil

	case AST.NtBinOp:
		return visitBinOp(n)

	case AST.NtUnOp:
		return visitUnOp(n)

	case AST.NtIdentifier:
		return visitIdentifierNode(n)

	case AST.NtVarDecl:
		return visitVarDeclNode(n)

	case AST.NtVarAssign:
		return visitVarAssignNode(n)

	default:
		break // * Do nothing
	}

	return 0, errors.New(fmt.Sprintf("error: unrecognized node type %s", AST.NTypeToString(n.Type)))
}

func visitProgram(n *AST.Node) (int64, error) {
	var value int64
	var err error

	for _, node := range *n.Children {
		value, err = visit(&node)

		if err != nil {
			return 0, err
		}
	}

	return value, nil
}

func visitInteger(n *AST.Node) int64 {
	return n.Value
}

func visitBinOp(n *AST.Node) (int64, error) {
	leftValue, err := visit(n.Left)

	if err != nil {
		return 0, err
	}

	rightValue, err := visit(n.Right)

	if err != nil {
		return 0, err
	}

	switch n.Op.Type {
	case Token.TtPlus:
		leftValue += rightValue
		break

	case Token.TtMinus:
		leftValue -= rightValue
		break

	case Token.TtStar:
		leftValue *= rightValue
		break

	case Token.TtSlash:
		if rightValue == 0 {
			return 0, errors.New("error: division by zero")
		}

		leftValue /= rightValue
		break

	default:
		break // * Do nothing
	}

	return leftValue, nil
}

func visitUnOp(n *AST.Node) (int64, error) {
	rightValue, err := visit(n.Right)

	if err != nil {
		return 0, err
	}

	switch n.Op.Type {
	case Token.TtPlus:
		break

	case Token.TtMinus:
		rightValue = -rightValue
		break

	default:
		break // * Do nothing
	}

	return rightValue, nil
}

func visitIdentifierNode(n *AST.Node) (int64, error) {
	value, exists := symbolTable[n.Token.SValue]

	if !exists {
		return 0, errors.New(fmt.Sprintf("error: variable %s does not exist", n.Token.SValue))
	}

	return value, nil
}

func visitVarDeclNode(n *AST.Node) (int64, error) {
	symbol := n.Left.Token.SValue
	_, exists := symbolTable[symbol]

	if exists {
		return 0, errors.New(fmt.Sprintf("error: variable %s already exists", symbol))
	}

	value, err := visit(n.Right)

	if err != nil {
		return 0, err
	}

	symbolTable[symbol] = value
	return value, nil
}

func visitVarAssignNode(n *AST.Node) (int64, error) {
	symbol := n.Left.Token.SValue
	_, exists := symbolTable[symbol]

	if !exists {
		return 0, errors.New(fmt.Sprintf("error: variable %s does not exist", symbol))
	}

	value, err := visit(n.Right)

	if err != nil {
		return 0, err
	}

	symbolTable[symbol] = value
	return value, nil
}

func (i *Interpreter) Interpret() (int64, error) {
	root, err := i.parser.Parse()

	if err != nil {
		return 0, err
	}

	result, err := visit(&root)

	if err != nil {
		return 0, err
	}

	return result, nil
}
