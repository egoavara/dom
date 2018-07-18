package grammer

import (
	"testing"
	"bytes"
	"fmt"
)


func TestEBNFEncode(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	enc := NewEBNFEncoder(buf)
	enc.Encode(g)
	fmt.Println(string(buf.Bytes()))
}