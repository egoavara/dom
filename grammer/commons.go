package grammer

import "bytes"

var space = " \t\n"

func RemoveSpace(src []byte) []byte {
	return bytes.TrimLeft(src, space)
}