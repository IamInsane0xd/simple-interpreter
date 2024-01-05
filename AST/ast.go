package AST

import (
	"strconv"

	"simpleInterpreter/Token"
)

type Type int64

const (
	NtError Type = iota - 1
	NtProgram
	NtInteger
	NtBinOp
	NtUnOp
	NtIdentifier
	NtVarDecl
	NtVarAssign
)

var (
	ErrorProgram = NewProgram([]Node{})
	ErrorNode    = NewNode(NtError, Token.ErrorToken, 0, nil, nil, nil)
)

type Program struct {
	Type  Type
	Nodes []Node
}

type Node struct {
	Type  Type
	Token Token.Token
	Value int64
	Left  *Node
	Op    *Token.Token
	Right *Node
}

func NewProgram(nodes []Node) Program {
	return Program{
		Type:  NtProgram,
		Nodes: nodes,
	}
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

func NewIdentifierNode(token Token.Token) Node {
	return NewNode(NtIdentifier, token, 0, nil, nil, nil)
}

func NewVarDeclNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarDecl, token, 0, &left, nil, &right)
}

func NewVarAssignNode(token Token.Token, left Node, right Node) Node {
	return NewNode(NtVarAssign, token, 0, &left, nil, &right)
}

func NTypeToString(nType Type) string {
	switch nType {
	case NtError:
		return "ERROR"

	case NtProgram:
		return "PROGRAM"

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
		return ""
	}
}
