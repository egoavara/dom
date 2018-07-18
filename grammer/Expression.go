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
	GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error)
}
type HaveInnerExpressions interface {
	InnerExpressions() []Expression
}
type (
	// Reference for expression
	ExpressionRefer struct {
		id    string
		refer Expression
	}
	// Exact match up
	ExpressionPrefix struct {
		prefix string
	}
	// Regular expression matchup
	ExpressionRegexp struct {
		original string
		re       *regexp.Regexp
	}
	// Consecutive matching of 'Expression'
	ExpressionAnd struct {
		cond []Expression
	}
	// One 'Expression' match
	ExpressionOr struct {
		cond []Expression
	}
	// Except expression
	ExpressionExcept struct {
		ori, e Expression
	}
	// Multiple match
	ExpressionMultiple struct {
		e Expression
	}
	// One or no match
	ExpressionPossible struct {
		e Expression
	}
)



func NewExpressionRefer(to string) *ExpressionRefer {
	return &ExpressionRefer{
		id: to,
	}
}
func NewExpressionPrefix(prefix string) *ExpressionPrefix {
	return &ExpressionPrefix{

		prefix: prefix,
	}
}
func NewExpressionRegexp(expr string) (*ExpressionRegexp, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &ExpressionRegexp{

		re:       r,
		original: expr,
	}, nil
}
func MustExpressionRegexp(expr string) *ExpressionRegexp {
	e, err := NewExpressionRegexp(expr)
	if err != nil {
		panic(err)
	}
	return e
}
func NewExpressionAnd(e ...Expression) *ExpressionAnd {
	return &ExpressionAnd{

		cond: e,
	}
}
func NewExpressionOr(e ...Expression) *ExpressionOr {

	return &ExpressionOr{

		cond: e,
	}
}
func NewExpressionExcept(ori, e Expression) *ExpressionExcept {
	return &ExpressionExcept{
		ori:ori,
		e: e,
	}
}
func NewExpressionMultiple(e Expression) *ExpressionMultiple {
	return &ExpressionMultiple{

		e: e,
	}
}
func NewExpressionPossible(e Expression) *ExpressionPossible {
	return &ExpressionPossible{

		e: e,
	}
}

func (s *ExpressionRefer) PreGrammerBuild(g *Grammer) error {
	if elem, ok := g.gram[s.id]; ok {
		s.refer = elem
		return nil
	}
	return errors.WithMessage(ErrorNoReference, "there is no name '"+s.id+"'")
}

//
func (s *ExpressionRefer) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	//token.SetData(s.id)
	token.SetChildrun(token.Make(1)...)
	return s.refer.GrammerParsing(grammer, src, token.GetChildrun()[0])
}
func (s *ExpressionPrefix) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if bytes.HasPrefix(src, []byte(s.prefix)) {
		if rtk, ok := token.(ReferToken); ok{
			rtk.Reference(s)
		}
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
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	token.SetData(string(res))
	return src[len(res):], nil
}
func (s *ExpressionAnd) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	mk := token.Make(len(s.cond))
	token.SetChildrun(mk...)
	var err error
	for k, v := range s.cond {
		src, err = v.GrammerParsing(grammer, src, mk[k])
		if err != nil {
			return nil, err
		}
	}
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	return src, nil
}
func (s *ExpressionOr) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	var err error
	var tk = token.Make(1)[0]
	var temp []byte
	for _, v := range s.cond {
		temp, err = v.GrammerParsing(grammer, src, tk)
		if err == nil {
			token.SetChildrun(tk)
			break
		}
	}
	if err != nil {
		return nil, errors.WithMessage(ErrorGrammaticalDiscrepancy, "No matching any condition")
	}
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	return temp, nil
}
func (s *ExpressionExcept) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	// TODO
	panic("implement me")
}
func (s *ExpressionMultiple) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	// TODO
	panic("implement me")
}
func (s *ExpressionPossible) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {

	var tk = token.Make(1)[0]
	var left []byte
	var err error
	if left, err = s.e.GrammerParsing(grammer, src, tk); err != nil {
		return src, nil
	}
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	token.SetChildrun(tk)
	return left, nil
}



func (s *ExpressionAnd) InnerExpressions() []Expression {
	return s.cond
}
func (s *ExpressionOr) InnerExpressions() []Expression {
	return s.cond
}
func (s *ExpressionExcept) InnerExpressions() []Expression {
	return []Expression{s.ori, s.e}
}
func (s *ExpressionMultiple) InnerExpressions() []Expression {
	return []Expression{s.e}
}
func (s *ExpressionPossible) InnerExpressions() []Expression {
	return []Expression{s.e}
}

// common parent class for Implementation of 'Expression'
type nameBase struct {
	name string
}

func (s nameBase) GetName() string {
	return s.name
}
