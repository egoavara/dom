package svg11

import (
	"encoding/xml"
	"io"
)

type (
	Element interface {
		Parent() Element
		Childrun() []Element
	}
	// https://www.w3.org/TR/SVG11/struct.html#SVGElement
	ElemSVG struct {
		Core
		ConditionalProcess
		DocumentEvent
		GraphicalEvent
		Presentation
		elemSVGInnerData
		//
		childrun []Element
	}
	elemSVGInnerData struct {
		Class                     string      `xml:"class,attr"`
		Style                     string      `xml:"style,attr"`
		ExternalResourcesRequired string      `xml:"externalResourcesRequired,attr"`
		X                         LengthValue `xml:"x,attr"`
		Y                         LengthValue `xml:"y,attr"`
		Width                     LengthValue `xml:"width,attr"`
		Height                    LengthValue `xml:"height,attr"`
		ViewBox                   string      `xml:"viewBox,attr"`
		PreserveAspectRatio       string      `xml:"preserveAspectRatio,attr"`
		ZoomAndPan                string      `xml:"zoomAndPan,attr"`
		Version                   string      `xml:"version,attr"`
		BaseProfile               string      `xml:"baseProfile,attr"`
		ContentScriptType         string      `xml:"contentScriptType,attr"`
		ContentStyleType          string      `xml:"contentStyleType,attr"`
	}

	// https://www.w3.org/TR/SVG11/struct.html#GElement
	ElemGroup struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemGroupInnerData
		//
		parent   Element
		childrun []Element
	}
	elemGroupInnerData struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
	}
	// https://www.w3.org/TR/SVG11/paths.html#PathElement
	ElemPath struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemPathInnerData
		//
		parent   Element
		childrun []Element
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

func NewElement(name xml.Name) Element {
	switch name.Local {
	case "g":
		return new(ElemGroup)
	case "path":
		return new(ElemPath)
	}
	return nil
}

func (s *ElemSVG) Parent() Element {
	return nil
}
func (s *ElemSVG) Childrun() []Element {
	return s.childrun
}
func (s *ElemSVG) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tk xml.Token
	var err error
	for tk, err = d.Token(); err == nil; tk, err = d.Token() {
		switch t := tk.(type) {
		case xml.StartElement:
			elem := NewElement(t.Name)
			if elem != nil{
				err = d.DecodeElement(elem, &t)
				if err != nil {
					return err
				}
				s.childrun = append(s.childrun, elem)
			}
		}
	}
	if err != io.EOF && err != nil {
		return err
	}
	//
	err = d.DecodeElement(&s.Core, &start)
	if err != io.EOF && err != nil {
		return err
	}
	//
	err = d.DecodeElement(&s.ConditionalProcess, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.DocumentEvent, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.GraphicalEvent, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.elemSVGInnerData, &start)
	if err != io.EOF && err != nil {
		return err
	}

	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemGroup) Parent() Element {
	return s.parent
}
func (s *ElemGroup) Childrun() []Element {
	return s.childrun
}
func (s *ElemGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tk xml.Token
	var err error
	for tk, err = d.Token(); err == nil; tk, err = d.Token() {
		switch t := tk.(type) {
		case xml.StartElement:
			elem := NewElement(t.Name)
			if elem != nil{
				err = d.DecodeElement(elem, &t)
				if err != nil {
					return err
				}
				s.childrun = append(s.childrun, elem)
			}
		}
	}
	if err != io.EOF && err != nil {
		return err
	}
	//
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

	err = d.DecodeElement(&s.elemGroupInnerData, &start)
	if err != io.EOF && err != nil {
		return err
	}

	s.Presentation.xmlAttrs(start.Attr...)
	return nil
}

func (s *ElemPath) Parent() Element {
	return s.parent
}
func (s *ElemPath) Childrun() []Element {
	return s.childrun
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
