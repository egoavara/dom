package grammer

import (
	"github.com/pkg/errors"
	"fmt"
)

type Grammer struct {
	gram map[string]Expression
}

var (
	ErrorNotNilable = errors.New("Not nil allow")
	ErrorRemainSrc  = errors.New("Source is remain")
)
func (s *Grammer ) Tokenize(expr string, src string, tk Token) (Token, error) {
	if tk == nil{
		return nil, errors.WithMessage(ErrorNotNilable, "Token")
	}
	if e, ok := s.gram[expr]; ok{
		left, err := e.GrammerParsing(s, []byte(src), tk)
		if err != nil {
			return nil, err
		}
		left = RemoveSpace(left)
		if len(left) > 0{
			return nil, errors.WithMessage(ErrorRemainSrc, string(left))
		}
		return tk, nil
	}
	return nil, errors.WithMessage(ErrorNoReference, fmt.Sprintf("'%s' is not exist", expr))
}
func (s *Grammer ) Get(expr string) Expression {
	return s.gram[expr]
}
type GrammerBuilder struct {
	cache map[string]Expression
}
func NewGrammerBuilder() *GrammerBuilder {
	return &GrammerBuilder{
		cache:make(map[string]Expression),
	}
}
var (
	ErrorNamespaceConfliction = errors.New("Namespace already has that name")
)
func (s GrammerBuilder) AddExpression(name string, expression Expression) error {
	if _, ok := s.cache[name]; ok{
		return ErrorNamespaceConfliction
	}
	s.cache[name] = expression
	return nil
}
type PreGrammerBuild interface {
	PreGrammerBuild(g *Grammer) error
}
func (s GrammerBuilder) Build() (*Grammer, error) {
	res := &Grammer{
		gram:s.cache,
	}
	for _, v := range res.gram {
		err := recursiveDo(v, func(tk Expression) error {
			if vpgb,ok := tk.(PreGrammerBuild);ok{
				err := vpgb.PreGrammerBuild(res)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
func recursiveDo(expr Expression, fn func(tk Expression) error) error{
	err := fn(expr)
	if err != nil {
		return err
	}
	if hiexpr, ok := expr.(HaveInnerExpressions);ok{
		for _, iexpr := range hiexpr.InnerExpressions() {
			err := recursiveDo(iexpr, fn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}