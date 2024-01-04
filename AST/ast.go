package AST

import (
	"strconv"

	"simpleInterpreter/Token"
)

type Type int64

const (
	NtError Type = iota
	NtInteger
	NtBinOp
	NtUnOp
	NtIdentifier
	NtVarDecl
	NtVarAssign
)

var (
	ErrorNode = NewNode(NtError, Token.ErrorToken, 0, nil, nil, nil)
)

type Node struct {
	Type  Type
	Token Token.Token
	Value int64
	Left  *Node
	Op    *Token.Token
	Right *Node
}

func NewNode(nType Type, token Token.Token, value int64, left *Node, op *Token.Token, right *Node) Node {
	return Node{
		Type:  nType,
		Token: token,
		Value: value,
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func NewIntegerNode(token Token.Token) (Node, error) {
	value, err := strconv.ParseInt(token.SValue, 10, 64)

	if err != nil {
		return ErrorNode, err
	}

	return NewNode(NtInteger, token, value, nil, nil, nil), nil
}

func NewBinOpNode(left Node, op Token.Token, right Node) Node {
	return NewNode(NtBinOp, op, 0, &left, &op, &right)
}

func NewUnOpNode(op Token.Token, right Node) Node {
	return NewNode(NtUnOp, op, 0, nil, &op, &right)
}

func NewIdentifierNode(token Token.Token, right Node) Node {
	return NewNode(NtIdentifier, token, 0, nil, nil, &right)
}

func NewVarDeclNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarDecl, token, 0, &left, nil, &right)
}

func NewVarAssignNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarAssign, token, 0, &left, nil, &right)
}
