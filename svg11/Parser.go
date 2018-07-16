package svg11

import (
	"encoding/xml"
	"io"
)

func Parse(reader io.Reader) (root *ElemSVG, err error) {
	root = new(ElemSVG)
	//
	dec := xml.NewDecoder(reader)
	//
	err = dec.Decode(root)
	if err != io.EOF && err != nil {
		return nil, err
	}
	//
	return root, nil
}