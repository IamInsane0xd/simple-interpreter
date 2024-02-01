package AST

import (
	"strconv"

	"simpleInterpreter/Token"
)

type Type int64

const (
	NtError Type = iota - 1
	NtProgram
	NtFunction
	NtInteger
	NtBinOp
	NtUnOp
	NtIdentifier
	NtVarDecl
	NtVarAssign
)

var (
	ErrorNode = NewNode(NtError, Token.ErrorToken, 0, nil, nil, nil, nil)
)

type Node struct {
	Type     Type
	Token    Token.Token
	Value    int64
	Left     *Node
	Op       *Token.Token
	Right    *Node
	Children *[]Node
}

func NewNode(nType Type, token Token.Token, value int64, left *Node, op *Token.Token, right *Node,
	children *[]Node) Node {
	return Node{
		Type:     nType,
		Token:    token,
		Value:    value,
		Left:     left,
		Op:       op,
		Right:    right,
		Children: children,
	}
}

func NewProgramNode(token Token.Token, children []Node) Node {
	return NewNode(NtProgram, token, 0, nil, nil, nil, &children)
}

func NewFunctionNode(token Token.Token, left Node, children []Node) Node {
	return NewNode(NtFunction, token, 0, &left, nil, nil, &children)
}

func NewIntegerNode(token Token.Token) (Node, error) {
	value, err := strconv.ParseInt(token.SValue, 10, 64)

	if err != nil {
		return ErrorNode, err
	}

	return NewNode(NtInteger, token, value, nil, nil, nil, nil), nil
}

func NewBinOpNode(left Node, op Token.Token, right Node) Node {
	return NewNode(NtBinOp, op, 0, &left, &op, &right, nil)
}

func NewUnOpNode(op Token.Token, right Node) Node {
	return NewNode(NtUnOp, op, 0, nil, &op, &right, nil)
}

func NewIdentifierNode(token Token.Token) Node {
	return NewNode(NtIdentifier, token, 0, nil, nil, nil, nil)
}

func NewVarDeclNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarDecl, token, 0, &left, nil, &right, nil)
}

func NewVarAssignNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarAssign, token, 0, &left, nil, &right, nil)
}

func NTypeToString(nType Type) string {
	switch nType {
	case NtError:
		return "ERROR"

	case NtProgram:
		return "PROGRAM"

	case NtFunction:
		return "FUNCTION"

	case NtInteger:
		return "INTEGER"

	case NtBinOp:
		return "BIN_OP"

	case NtUnOp:
		return "UN_OP"

	case NtIdentifier:
		return "IDENTIFIER"

	case NtVarDecl:
		return "VAR_DECL"

	case NtVarAssign:
		return "VAR_ASSIGN"

	default:
		return "NO_REPR"
	}
}
