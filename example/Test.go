package main

import (
	"github.com/iamGreedy/dom/grammer"

	"github.com/iamGreedy/dom/xpath/compiler"
	"testing"
	"fmt"
)

func main() {
	var test = `count(/root/user[login='user1' and name='User 1' and profile[value='admin'] and profile[value='operator']])`
	var res testing.BenchmarkResult
	res = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N;i++{
			compiler.GrammerXPath.Tokenize("XPath", test, grammer.NewReferenceToken())
		}
	})
	fmt.Println(res)
	//compiler.GrammerXPath.Tokenize("XPath", test, grammer.NewReferenceToken())
}
