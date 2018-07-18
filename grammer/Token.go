package grammer

import (
	"fmt"
	"io"
	"strings"
)

type Token interface {
	Make(i int) []Token
	//
	GetName() string
	SetName(name string)
	GetData() string
	SetData(data string)
	GetChildrun() []Token
	SetChildrun(childrun ...Token)
	//
	fmt.Stringer
	RecursivePrint(printer io.Writer)
	recursivePrint(printer io.Writer, depth int)
}

type NamedToken struct {
	Name     string
	Data     string
	Childrun []Token
}

func (s *NamedToken) Make(i int) []Token {
	var temo = make([]Token, i)
	for i := range temo{
		temo[i] = new(NamedToken)
	}
	return temo
}
func (s *NamedToken) GetName() string {
	return s.Name
}
func (s *NamedToken) SetName(name string) {
	s.Name = name
}
func (s *NamedToken) GetData() string {
	return s.Data
}
func (s *NamedToken) SetData(data string) {
	s.Data = data
}
func (s *NamedToken) GetChildrun() []Token {
	return s.Childrun
}
func (s *NamedToken) SetChildrun(childrun ...Token) {
	s.Childrun = childrun
}
func (s *NamedToken) String() string {
	if len(s.Data) > 0 {
		return fmt.Sprintf("%s(%s)", s.Name, s.Data)
	}
	return fmt.Sprintf("%s", s.Name)
}
func (s *NamedToken) RecursivePrint(printer io.Writer) {
	s.recursivePrint(printer, 0)
}
func (s *NamedToken) recursivePrint(printer io.Writer, depth int) {
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

type UnnamedToken struct {
	Data     string
	Childrun []Token
}

func (s *UnnamedToken) Make(i int) []Token {
	var temo = make([]Token, i)
	for i := range temo{
		temo[i] = new(UnnamedToken)
	}
	return temo
}
func (s *UnnamedToken) GetName() string {
	return ""
}
func (s *UnnamedToken) SetName(name string) {

}
func (s *UnnamedToken) GetData() string {
	return s.Data
}
func (s *UnnamedToken) SetData(data string) {
	s.Data = data
}
func (s *UnnamedToken) GetChildrun() []Token {
	return s.Childrun
}
func (s *UnnamedToken) SetChildrun(childrun ...Token) {
	s.Childrun = childrun
}
func (s *UnnamedToken) String() string {
	return fmt.Sprintf("(%s)", s.Data)
}
func (s *UnnamedToken) RecursivePrint(printer io.Writer) {
	s.recursivePrint(printer, 0)
}
func (s *UnnamedToken) recursivePrint(printer io.Writer, depth int) {
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
