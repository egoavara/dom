package grammer

import (
	"bytes"
	"github.com/pkg/errors"
	"regexp"
)

type Grammer struct {
	gram map[string]Expression
}










type grammerBuilder struct {
	cache map[string]Expression
}

func GrammerBuilder() *grammerBuilder {
	return &grammerBuilder{
		cache:make(map[string]Expression),
	}
}
var (
	ErrorNamespaceConfliction = errors.New("Namespace already has that name")
)
func (s grammerBuilder) AddExpression(expression Expression) error {
	if _, ok := s.cache[expression.GetName()]; ok{
		return ErrorNamespaceConfliction
	}
	s.cache[expression.GetName()] = expression
	return nil
}
func (s grammerBuilder) Build() *Grammer {
	return &Grammer{
		gram:s.cache,
	}
}