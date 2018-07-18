package grammer

import (
	"fmt"
	"strings"
	"testing"
)

var g *Grammer
func init() {
	temp := NewGrammerBuilder()
	temp.AddExpression("Digits", MustExpressionRegexp(`[0-9]+`))
	temp.AddExpression("IntegerLiteral", NewExpressionRefer("Digits"))
	temp.AddExpression("DecimalLiteral",
		NewExpressionOr(
			NewExpressionAnd(NewExpressionPrefix("."), NewExpressionRefer("Digits")),
			NewExpressionAnd(NewExpressionRefer("Digits"), NewExpressionPrefix("."), MustExpressionRegexp(`[0-9]*`)),
		),
	)
	var err error
	g, err = temp.Build()
	if err != nil {
		panic(err)
	}
}

func TestElemPrefix(t *testing.T) {
	var testset = []string{
		"s",
		"$,rgae",
		"$a",
		"123",
		"512.2",
		`"Hello?"`,
		"$a32[a]",
	}

	for _, test := range testset {
		fmt.Println("==================================================")
		fmt.Printf("Test %s : ", test)
		tk := NewDefaultToken()
		_, err := g.Tokenize("DecimalLiteral", test, tk)
		if err != nil {
			fmt.Printf("Error : '%s'\n", err.Error())
			continue
		}
		fmt.Println()
		Recur(tk, 0)
	}

}
func Recur(tk Token, depth int) {
	fmt.Print(strings.Repeat("    ", depth))
	if temp := tk.GetData(); len(temp) == 0 {
		fmt.Println("<no data>")
	} else {
		fmt.Println(temp)
	}

	for _, child := range tk.GetChildrun() {
		Recur(child, depth+1)
	}
}
