package ast

import "github.com/lwlwilliam/monkey/token"

// AST 由各个相互关联的 Node 组成
type Node interface {
	// 返回节点对应 token 的 literal 值
	// 用于 debug 和测试
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

// AST 的根节点
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let 语句
// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// return 语句
// return <expression>;
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {

}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
