package ast

import "sad-language/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement // root of AST
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatment struct {
	Token token.Token // LET token
	Name  *Identifier // identifier of the binding
	Value Expression
}

func (ls *LetStatment) statementNode()       {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
