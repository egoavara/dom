package svg11

import (
	"encoding/xml"
	"io"
)

func Parse(reader io.Reader) (doc *Document, err error) {
	doc = new(Document)
	doc.Root = new(ElemSVG)

	//
	dec := xml.NewDecoder(reader)
	//
	err = dec.Decode(doc.Root)
	if err != io.EOF && err != nil {
		return nil, err
	}
	//
	return doc, nil
}