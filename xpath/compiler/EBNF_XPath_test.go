package compiler

import (
	"testing"
	"os"
	"github.com/iamGreedy/dom/grammer"
	"fmt"
	"strings"
)

var testset = []string{
	`/Hello/World`,
	`/Hello/World/@id`,
	`/Hello/World/text()`,
	`/child::Hello/World/text()`,
	`/child::Hello/World[@id]/text()`,
	`/child::Hello/World[@id="iam"]/text()`,
}

func TestGrammerOut(t *testing.T)  {
	f, err := os.Create("xpath.ebnf")
	if err != nil {
		panic(err)
	}
	enc := grammer.NewEBNFEncoder(f)
	enc.Encode(GrammerXPath)
}
func TestSingleGrammer(t *testing.T) {
	const index = 5
	tk, err := GrammerXPath.Tokenize("XPath", testset[index], grammer.NewReferenceToken())
	if err != nil {
		t.Error(err)
		return
	}
	Recur(tk.(grammer.ReferToken), nil)
}
func TestGrammer(t *testing.T)  {
	var name = t.Name() + "/"
	for _, test := range testset {
		t.Run(test, func(t *testing.T) {
			test := strings.TrimPrefix(t.Name(), name)
			t.Log(test)
			_, err := GrammerXPath.Tokenize("XPath", test, grammer.NewReferenceToken())
			if err != nil {
				t.Error(err)
				return
			}
			t.Log("SUCCESS")
		})
	}
}
func Recur(tk grammer.ReferToken, exprs []grammer.Expression) {
	dt := tk.GetData()
	if len(dt) > 0{
		for _, v := range exprs {
			fmt.Print(v, "/")
		}
		fmt.Print(tk.GetExpression())
		fmt.Print(" :: '", tk.GetData(), "'")
		fmt.Println()
	}
	for _, child := range tk.GetChildrun() {
		Recur(child.(grammer.ReferToken), append(exprs, tk.GetExpression()))
	}
}

// History
// [windows:386]
// 		- Intel(R) Core(TM) i3-6100
// 		- Intel(R) HD Graphics 530
//		- Intel(R) 100 Series Chipset Family (H110)
func BenchmarkGrammer(b *testing.B) {
	var name = b.Name() + "/"
	for _, test := range testset {
		b.Run(test, func(b *testing.B) {
			b.StopTimer()
			test := strings.TrimPrefix(b.Name(), name)
			b.Log(test)
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				_, err := GrammerXPath.Tokenize("XPath", test, grammer.NewReferenceToken())
				if err != nil {
					b.Error(err)
				}
				b.StopTimer()
			}
		})
	}
}