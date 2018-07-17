package svg11

import (
	"testing"
	"os"
	"io"
	"strings"
	"reflect"
)

func TestParse(t *testing.T) {
	f, err := os.Open("../example/Go-gopher-Vector/svg/gopher01c.svg")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	//
	root, err := Parse(f)
	if err != io.EOF && err != nil {
		t.Error(err)
		return
	}


	Recur(t, root, 0)
}
func Recur(t *testing.T, element Element, depth int){
	v := reflect.ValueOf(element)
	tp := v.Type()
	t.Log(strings.Repeat("    ", depth),":", tp)
	for _, child := range element.Childrun() {
		Recur(t, child, depth + 1)
	}
}