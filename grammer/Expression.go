package grammer

import (
	"bytes"
	"github.com/pkg/errors"
	"regexp"
)

var (
	ErrorGrammaticalDiscrepancy = errors.New("Grammer fail")
	ErrorNoReference            = errors.New("No matching reference")
)

type Expression interface {
	GetName() string
	GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error)
}
type (
	// common parent class for Implementation of 'Expression'
	nameBase struct {
		name string
	}

	// Reference for expression
	ExpressionRefer struct {
		nameBase
		id string
	}

	// Exact match up
	ExpressionPrefix struct {
		nameBase
		prefix string
	}
	// Regular expression matchup
	ExpressionRegexp struct {
		nameBase
		original string
		re       *regexp.Regexp
	}
	// Consecutive matching of 'Expression'
	ExpressionAnd struct {
		nameBase
		cond []Expression
	}
	// One 'Expression' match
	ExpressionOr struct {
		nameBase
		cond []Expression
	}
	// Multiple match
	ExpressionMultiple struct {
		nameBase
		e Expression
	}
	// One or no match
	ExpressionCanbe struct {
		nameBase
		e Expression
	}
)

func NewExpressionRefer(name string, to string) *ExpressionRefer {
	return &ExpressionRefer{
		nameBase: nameBase{name: name},
		id:       to,
	}
}
func NewExpressionPrefix(name string, prefix string) *ExpressionPrefix {
	return &ExpressionPrefix{
		nameBase: nameBase{name: name},
		prefix:prefix,
	}
}

func NewExpressionRegexp(name string, expr string) (*ExpressionRegexp, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &ExpressionRegexp{
		nameBase: nameBase{name: name},
		re:       r,
		original: expr,
	}, nil
}
func MustExpressionRegexp(name string, expr string) *ExpressionRegexp {
	e, err := NewExpressionRegexp(name, expr)
	if err != nil {
		panic(err)
	}
	return e
}
func NewExpressionAnd(name string, e ... Expression) *ExpressionAnd {
	return &ExpressionAnd{
		nameBase: nameBase{name: name},
		cond:e,
	}
}
func NewExpressionOr(name string, e ... Expression) *ExpressionOr {

	return &ExpressionOr{
		nameBase: nameBase{name: name},
		cond:e,
	}
}
func NewExpressionMultiple(name string, e Expression) *ExpressionMultiple {

	return &ExpressionMultiple{
		nameBase: nameBase{name: name},
		e:e,
	}
}

func NewExpressionCanbe(name string, e Expression) *ExpressionCanbe {
	return &ExpressionCanbe{
		nameBase: nameBase{name: name},
		e:e,
	}
}


func (s nameBase) GetName() string {
	return s.name
}
func (s *ExpressionRefer) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if elem, ok := grammer.gram[s.id]; ok {
		token.SetName(s.name)
		token.SetData(s.id)
		token.SetChildrun(token.Make(1)...)
		return elem.GrammerParsing(grammer, src, token.GetChildrun()[0])
	}
	return nil, errors.WithMessage(ErrorNoReference, "there is no name '"+s.id+"'")
}
func (s *ExpressionPrefix) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if bytes.HasPrefix(src, []byte(s.prefix)) {
		token.SetName(s.name)
		token.SetData(s.prefix)
		return bytes.TrimPrefix(src, []byte(s.prefix)), nil
	}
	return nil, errors.WithMessage(ErrorGrammaticalDiscrepancy, string(src))
}
func (s *ExpressionRegexp) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	res := s.re.Find(src)
	if len(res) == 0 {
		return nil, errors.WithMessage(ErrorGrammaticalDiscrepancy, string(src))
	}
	token.SetName(s.name)
	token.SetData(string(res))
	return src[len(res):], nil
}
func (s *ExpressionAnd) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	token.SetName(s.name)
	mk := token.Make(len(s.cond))
	token.SetChildrun(mk...)
	var err error
	for k, v := range s.cond {
		src, err = v.GrammerParsing(grammer, src, mk[k])
		if err != nil {
			return nil, err
		}
	}
	return src, nil
}
func (s *ExpressionOr) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	token.SetName(s.name)
	var err error
	var tk = token.Make(1)[0]
	var temp []byte
	for _, v := range s.cond {
		temp, err = v.GrammerParsing(grammer, src, tk)
		if err == nil {
			token.SetData(v.GetName())
			token.SetChildrun(tk)
			break
		}
	}
	if err != nil {
		return nil, errors.WithMessage(ErrorGrammaticalDiscrepancy, "No matching any condition")
	}
	return temp, nil
}
func (s *ExpressionMultiple) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	//s.
	panic("er")
}
func (s *ExpressionCanbe) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	token.SetName(s.name)
	var tk = token.Make(1)[0]
	var left []byte
	var err error
	if left, err = s.e.GrammerParsing(grammer, src, tk); err != nil {
		return src, nil
	}
	token.SetData(s.e.GetName())
	token.SetChildrun(tk)
	return left, nil
}
