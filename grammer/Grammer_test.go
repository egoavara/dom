package grammer

//var g *Grammer
//func init()  {
//	temp := GrammerBuilder()
//	temp.AddExpression(&ExpressionAnd{
//		nameBase:nameBase{"Variable"},
//		cond:[]Expression{
//			&ExpressionPrefix{
//				nameBase:nameBase{"Dollar"},
//				prefix:"$",
//			},
//			MustExpressionRegexp("Digit", `[a-zA-Z0-9_]*`),
//			&ExpressionPrefix{
//				nameBase:nameBase{"BracketOpen"},
//				prefix:"[",
//			},
//			MustExpressionRegexp("Digit", `[a-zA-Z0-9_]*`),
//			&ExpressionPrefix{
//				nameBase:nameBase{"BracketClose"},
//				prefix:"]",
//			},
//		},
//	})
//}
//
//func TestElemPrefix(t *testing.T) {
//	es := &
//
//	var testset = []string{
//		"s",
//		"$,rgae",
//		"$a",
//		"$a32",
//		"$a32[1",
//		"$a32[1]",
//		"$a32[a]",
//	}
//	var tk NamedToken
//	for _, test := range testset {
//		fmt.Println("==================================================")
//		fmt.Printf("Test %s : ", test)
//		left, err := es.GrammerParsing(nil, []byte(test), &tk)
//		if err != nil {
//			fmt.Printf("Error : '%s'", err.Error())
//		}else {
//			fmt.Printf("Left : '%s'\n", string(left))
//			tk.RecursivePrint(os.Stdout)
//		}
//		fmt.Println()
//
//	}
//}