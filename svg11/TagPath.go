package svg11

import (
	"encoding/xml"
	"github.com/iamGreedy/psvg"
	"io"
	"strings"
)

// https://www.w3.org/TR/SVG11/paths.html
type (
	// https://www.w3.org/TR/SVG11/paths.html#PathElement
	ElemPath struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemPathInnerData
		//
		Path []psvg.Elem
		commonTree
	}
	elemPathInnerData struct {
		Class                     string  `xml:"class,attr"`
		Style                     string  `xml:"style,attr"`
		ExternalResourcesRequired string  `xml:"externalResourcesRequired,attr"`
		Transform                 string  `xml:"transform,attr"`
		PathLength                float32 `xml:"pathLength,attr"`
		//
		Data string `xml:"d,attr"`
	}
)

func (s *ElemPath) Paths() []psvg.Elem {
	return s.Path
}
func (s *ElemPath) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemPath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	err = d.DecodeElement(&s.Core, &start)
	if err != io.EOF && err != nil {
		return err
	}
	//
	err = d.DecodeElement(&s.ConditionalProcess, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.GraphicalEvent, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.elemPathInnerData, &start)
	if err != io.EOF && err != nil {
		return err
	}

	s.Presentation.xmlAttrs(start.Attr...)

	parser := psvg.NewParser(strings.NewReader(s.Data))
	for e:= parser.Next(); e != nil; e = parser.Next(){
		switch ee := e.(type) {
		case *psvg.UnknownError:
			return ee.Err
		default:
			s.Path = append(s.Path, ee)
		}
	}
	return nil
}
func (s *ElemPath ) String() string {
	const markmax = 6
	res := NewFuncStyle("Path")
	s.Presentation.alignedarglist(res)
	elems := NewListStyle()
	for i, p := range s.Path {
		if i > markmax{
			elems.SkipMark()
			break
		}
		elems.Append(p.String())
	}
	if elems.Length() > 0{
		res.List(elems.Build())
	}
	return res.Build()
}