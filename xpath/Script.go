package xpath

import "encoding/xml"

type OrAndExpr struct {

}
type (
	// PathExpr
	PathExpr struct {
		Works []StepExpression
	}
	StepNav interface {
		Next() StepExpression
		Prev() StepExpression
		NextN(n int) []StepExpression
		PrevN(n int) []StepExpression
		SkipN(n int)
	}
	StepExpression interface {

	}

)
type (
	// ^/
	RootStep struct {
		Name xml.Name
	}
	// Move
	MoveStep struct {}
	// Recurive
	RecuriveStep struct {}
	// name
	AxisStep struct {
		Name xml.Name
	}
	// ..
	ReverseAxisStep struct {}
	// .
	CurrentAxisStep struct {}
	// @
	AttribeStep struct {
		Name xml.Name
	}



)

