package grammer

import (
	"fmt"
	"io"
	"strings"
)

type Token interface {
	Make(i int) []Token
	//
	Total() string
	//
	GetData() string
	SetData(data string)
	GetChildrun() []Token
	SetChildrun(childrun ...Token)
	//
	fmt.Stringer
	RecursivePrint(printer io.Writer)
	recursivePrint(printer io.Writer, depth int)
}
type ReferToken interface {
	Token
	Reference(expression Expression)
	GetExpression() Expression
}

func NewDefaultToken() Token {
	return new(DefaultToken)
}

type DefaultToken struct {
	Data     string
	Childrun []Token
}

func (s *DefaultToken) Make(i int) []Token {
	var temo = make([]Token, i)
	for i := range temo {
		temo[i] = NewDefaultToken()
	}
	return temo
}
func (s *DefaultToken) Total() string {
	res := s.Data
	for _, child := range s.Childrun {
		res += child.Total()
	}
	return res
}
func (s *DefaultToken) GetData() string {
	return s.Data
}
func (s *DefaultToken) SetData(data string) {
	s.Data = data
}
func (s *DefaultToken) GetChildrun() []Token {
	return s.Childrun
}
func (s *DefaultToken) SetChildrun(childrun ...Token) {
	s.Childrun = childrun
}
func (s *DefaultToken) String() string {
	return fmt.Sprintf("%s", s.Data)
}
func (s *DefaultToken) RecursivePrint(printer io.Writer) {
	s.recursivePrint(printer, 0)
}
func (s *DefaultToken) recursivePrint(printer io.Writer, depth int) {
	printer.Write([]byte(
		fmt.Sprintf(
			"%s%s",
			strings.Repeat("    ", depth),
			s.String(),
		) + "\n",
	))
	for _, c := range s.Childrun {
		c.recursivePrint(printer, depth+1)
	}
	return
}

func NewReferenceToken() Token {
	return new(ReferenceToken)
}

type ReferenceToken struct {
	Expression Expression
	Data       string
	Childrun   []Token
}

func (s *ReferenceToken) Reference(expression Expression) {
	s.Expression = expression
}
func (s *ReferenceToken) GetExpression() Expression {
	return s.Expression
}
func (s *ReferenceToken) Make(i int) []Token {
	var temo = make([]Token, i)
	for i := range temo {
		temo[i] = NewReferenceToken()
	}
	return temo
}
func (s *ReferenceToken) Total() string {
	res := s.Data
	for _, child := range s.Childrun {
		res += child.Total()
	}
	return res
}
func (s *ReferenceToken) GetData() string {
	return s.Data
}
func (s *ReferenceToken) SetData(data string) {
	s.Data = data
}
func (s *ReferenceToken) GetChildrun() []Token {
	return s.Childrun
}
func (s *ReferenceToken) SetChildrun(childrun ...Token) {
	s.Childrun = childrun
}
func (s *ReferenceToken) String() string {
	return fmt.Sprintf("%s", s.Data)
}
func (s *ReferenceToken) RecursivePrint(printer io.Writer) {
	s.recursivePrint(printer, 0)
}
func (s *ReferenceToken) recursivePrint(printer io.Writer, depth int) {
	printer.Write([]byte(
		fmt.Sprintf(
			"%s%s",
			strings.Repeat("    ", depth),
			s.String(),
		) + "\n",
	))
	for _, c := range s.Childrun {
		c.recursivePrint(printer, depth+1)
	}
	return
}
