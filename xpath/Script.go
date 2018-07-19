package xpath

type (
	// PathExpr
	Worker struct {
		Works []Work
	}
	Work interface {
		ToExpression() string
	}
)
type (
	// AxisStepExpr
	AxisStep struct {}

)

