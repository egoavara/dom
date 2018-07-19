package compiler

import (
	"testing"
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
	`ws:gHel/./World[@id="iam"]/text()`,
	`/bookstore/book[price>35]/title`,
	`/root/user[login='user1' and name='User 1' and profile[value='admin'] and profile[value='operator']]`,
	`count(/root/user[login='user1' and name='User 1' and profile[value='admin'] and profile[value='operator']])`,
}
var benchmarkset = testset[len(testset) - 1]

func TestSingleGrammer(t *testing.T) {
	const index = 6
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

func RecurDataonly(tk grammer.ReferToken) {
	dt := tk.GetData()
	if len(dt) > 0{
		fmt.Print(tk.GetExpression())
		fmt.Print(" :: '", tk.GetData(), "'")
		fmt.Println()
	}
	for _, child := range tk.GetChildrun() {
		RecurDataonly(child.(grammer.ReferToken))
	}
}

// History
// [windows:386]
// 		- Intel(R) Core(TM) i3-6100
// 		- Intel(R) HD Graphics 530
//		- Intel(R) 100 Series Chipset Family (H110)
// [windows:amd64]
// 		- Intel(R) Core(TM) i5-6600
// 		- NVIDIA GeForce GTX 960
//		- Intel(R) 100 Series/C230 Series Chipset Family (B150)
// 			= 898049ns
func BenchmarkGrammer(b *testing.B) {
	b.Run(benchmarkset, func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			b.StartTimer()
			_, err := GrammerXPath.Tokenize("XPath", benchmarkset, grammer.NewReferenceToken())
			if err != nil {
				b.Error(err)
			}
			b.StopTimer()
		}
	})
}

func TestBenchmarkResult(t *testing.T) {
	tk, err := GrammerXPath.Tokenize("XPath", benchmarkset, grammer.NewReferenceToken())
	if err != nil {
		t.Error(err)
	}
	//Recur(tk.(grammer.ReferToken), nil)
	RecurDataonly(tk.(grammer.ReferToken))
}