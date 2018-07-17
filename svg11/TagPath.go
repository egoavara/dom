package svg11

import (
	"encoding/xml"
	"io"
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
		commonTree
	}
	elemPathInnerData struct {
		Class                     string  `xml:"class,attr"`
		Style                     string  `xml:"style,attr"`
		ExternalResourcesRequired string  `xml:"externalResourcesRequired,attr"`
		Transform                 string  `xml:"transform,attr"`
		Data                      string  `xml:"d,attr"`
		PathLength                float32 `xml:"pathLength,attr"`
	}
)

func (s *ElemPath) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive);e != nil{
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

	return nil
}