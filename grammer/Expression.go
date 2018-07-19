package grammer

import (
	"bytes"
	"github.com/pkg/errors"
	"regexp"
	"fmt"
)

var (
	ErrorGrammaticalDiscrepancy = errors.New("Grammer fail")
	ErrorNoReference            = errors.New("No matching reference")
	ErrorException            = errors.New("Exception")
)

type Expression interface {
	GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error)
	fmt.Stringer
}
type HaveInnerExpressions interface {
	InnerExpressions() []Expression
}
type (
	ExpressionBaseToken struct {
		e Expression
	}
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

func (s *ExpressionRefer) ID() string {
	return s.id
}

func NewExpressionBaseToken(e Expression) *ExpressionBaseToken {
	return &ExpressionBaseToken{
		e: e,
	}
}
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
	// must add ^ to front for Startwiths matching
	r, err := regexp.Compile("^" + expr)
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
func (s *ExpressionBaseToken) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	token.SetBaseToken()
	return s.e.GrammerParsing(grammer, src, token.GetChildrun()[0])
}
func (s *ExpressionRefer) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	//token.AddData(s.id)
	token.SetChildrun(token.Make(1)...)
	return s.refer.GrammerParsing(grammer, src, token.GetChildrun()[0])
}
func (s *ExpressionPrefix) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	if bytes.HasPrefix(src, []byte(s.prefix)) {
		if rtk, ok := token.(ReferToken); ok{
			rtk.Reference(s)
		}
		token.AddData(s.prefix)
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
	token.AddData(string(res))
	return src[len(res):], nil
}
func (s *ExpressionAnd) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	mk := token.Make(len(s.cond))
	token.SetChildrun(mk...)
	var err error
	for k, v := range s.cond {
		src = RemoveSpace(src)
		//if len(temp) == len(src) && k != 0{
		//	return nil, errors.WithMessage(ErrorGrammaticalDiscrepancy, "And Expression need space for each condition")
		//}
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
	if _, err := s.e.GrammerParsing(grammer, src, token.Make(1)[0]); err == nil {
		return nil, errors.WithMessage(ErrorException, string(src))
	}
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	tk := token.Make(1)[0]
	token.SetChildrun(tk)
	return s.ori.GrammerParsing(grammer, src, tk)
}
func (s *ExpressionMultiple) GrammerParsing(grammer *Grammer, src []byte, token Token) ([]byte, error) {
	var err error
	var temp []byte
	var chr []Token
	for {
		var tk = token.Make(1)[0]
		temp, err = s.e.GrammerParsing(grammer, src, tk)
		if err != nil {
			break
		}
		src = temp
		chr = append(chr, tk)
	}
	if rtk, ok := token.(ReferToken); ok{
		rtk.Reference(s)
	}
	token.SetChildrun(chr...)
	return src, nil
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





func (s *ExpressionBaseToken) String() string{
	return fmt.Sprintf("!%s", s.e.String())
}
func (s *ExpressionRefer) String() string{
	return fmt.Sprintf("@%s", s.id)
}
func (s *ExpressionPrefix) String() string{
	return fmt.Sprintf("'%s'", s.prefix)
}
func (s *ExpressionRegexp) String() string{
	return fmt.Sprintf("regex(%s)", s.original)
}
func (s *ExpressionAnd) String() string{
	return "And"
}
func (s *ExpressionOr) String() string{
	return "Or"
}
func (s *ExpressionExcept) String() string{
	return "Except"
}
func (s *ExpressionMultiple) String() string{
	return "Multiple"
}
func (s *ExpressionPossible) String() string{
	return "Possible"
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
