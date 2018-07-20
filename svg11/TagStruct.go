package svg11

import (
	"encoding/xml"
	"github.com/iamGreedy/dom/svg11/presentation"
	"io"
)

// https://www.w3.org/TR/SVG11/struct.html
type (
	// https://www.w3.org/TR/SVG11/struct.html#SVGElement
	ElemSVG struct {
		Core
		ConditionalProcess
		DocumentEvent
		GraphicalEvent
		Presentation
		elemSVGInnerData
		//
		commonTree
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
		commonTree
	}
	elemGroupInnerData struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
	}

	// https://www.w3.org/TR/SVG11/struct.html#DefsElement
	ElemDefine struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemDefineInnerData
		//
		commonTree
	}
	elemDefineInnerData struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
	}

	// https://www.w3.org/TR/SVG11/struct.html#DescriptionAndTitleElements
	ElemDesc struct {
		Core
		elemDesc
		//
		parent Element
	}
	elemDesc struct {
		Class string `xml:"class,attr"`
		Style string `xml:"style,attr"`
		Text  string `xml:",chardata"`
	}

	// https://www.w3.org/TR/SVG11/struct.html#DescriptionAndTitleElements
	ElemTitle struct {
		Core
		elemTitle
		//
		parent Element
	}
	elemTitle struct {
		Class string `xml:"class,attr"`
		Style string `xml:"style,attr"`
		Text  string `xml:",chardata"`
	}

	// TODO : https://www.w3.org/TR/SVG11/struct.html#SymbolElement
	// TODO : https://www.w3.org/TR/SVG11/struct.html#UseElement
	// TODO : https://www.w3.org/TR/SVG11/struct.html#ImageElement
	// TODO : https://www.w3.org/TR/SVG11/struct.html#SwitchElement

)

func (s *ElemSVG) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive, createShape, createStructural, createGradient); e != nil {
		return e
	}
	switch name.Local {
	case "a":
		//TODO
		return nil
	case "altGlyphDef":
		//TODO
		return nil
	case "clipPath":
		//TODO
		return nil
	case "color-profile":
		//TODO
		return nil
	case "cursor":
		//TODO
		return nil
	case "filter":
		//TODO
		return nil
	case "font":
		//TODO
		return nil
	case "font-face":
		//TODO
		return nil
	case "foreignObject":
		//TODO
		return nil
	case "image":
		//TODO
		return nil
	case "marker":
		//TODO
		return nil
	case "mask":
		//TODO
		return nil
	case "pattern":
		//TODO
		return nil
	case "script":
		//TODO
		return nil
	case "style":
		//TODO
		return nil
	case "switch":
		//TODO
		return nil
	case "text":
		//TODO
		return nil
	case "view":
		//TODO
		return nil
	}
	return nil
}
func (s *ElemSVG) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tk xml.Token
	var err error
	for tk, err = d.Token(); err == nil; tk, err = d.Token() {
		switch t := tk.(type) {
		case xml.StartElement:
			elem := s.createElement(t.Name)
			if elem != nil {
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
func (s *ElemSVG) String() string {

	res := NewFuncStyle("svg")
	if s.Core.ID != nil {
		res.Map("id", s.Core.ID.String())
	}
	if len(s.Version) > 0 {
		res.Map("version", s.Version)
	}
	if s.X.Value != 0 {
		res.Map("x", s.X.String())
	}
	if s.Y.Value != 0 {
		res.Map("y", s.Y.String())
	}
	if s.Width.Value != 0 {
		res.Map("width", s.Width.String())
	}
	if s.Height.Value != 0 {
		res.Map("height", s.Height.String())
	}
	if len(s.ViewBox) > 0 {
		res.Map("viewBox", s.ViewBox)
	}
	s.Presentation.alignedarglist(res)
	return res.Build()
}

func (s *ElemGroup) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive, createShape, createStructural, createGradient); e != nil {
		return e
	}
	switch name.Local {
	case "a":
		//TODO
		return nil
	case "altGlyphDef":
		//TODO
		return nil
	case "clipPath":
		//TODO
		return nil
	case "color-profile":
		//TODO
		return nil
	case "cursor":
		//TODO
		return nil
	case "filter":
		//TODO
		return nil
	case "font":
		//TODO
		return nil
	case "font-face":
		//TODO
		return nil
	case "foreignObject":
		//TODO
		return nil
	case "image":
		//TODO
		return nil
	case "marker":
		//TODO
		return nil
	case "mask":
		//TODO
		return nil
	case "pattern":
		//TODO
		return nil
	case "script":
		//TODO
		return nil
	case "style":
		//TODO
		return nil
	case "switch":
		//TODO
		return nil
	case "text":
		//TODO
		return nil
	case "view":
		//TODO
		return nil
	}
	return nil
}
func (s *ElemGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tk xml.Token
	var err error
	for tk, err = d.Token(); err == nil; tk, err = d.Token() {
		switch t := tk.(type) {
		case xml.StartElement:
			elem := s.createElement(t.Name)
			if elem != nil {
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

func (s *ElemDefine) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive, createShape, createStructural, createGradient); e != nil {
		return e
	}
	switch name.Local {
	case "a":
		//TODO
		return nil
	case "altGlyphDef":
		//TODO
		return nil
	case "clipPath":
		//TODO
		return nil
	case "color-profile":
		//TODO
		return nil
	case "cursor":
		//TODO
		return nil
	case "filter":
		//TODO
		return nil
	case "font":
		//TODO
		return nil
	case "font-face":
		//TODO
		return nil
	case "foreignObject":
		//TODO
		return nil
	case "image":
		//TODO
		return nil
	case "marker":
		//TODO
		return nil
	case "mask":
		//TODO
		return nil
	case "pattern":
		//TODO
		return nil
	case "script":
		//TODO
		return nil
	case "style":
		//TODO
		return nil
	case "switch":
		//TODO
		return nil
	case "text":
		//TODO
		return nil
	case "view":
		//TODO
		return nil
	}
	return nil
}
func (s *ElemDefine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tk xml.Token
	var err error
	for tk, err = d.Token(); err == nil; tk, err = d.Token() {
		switch t := tk.(type) {
		case xml.StartElement:
			elem := s.createElement(t.Name)
			if elem != nil {
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

	err = d.DecodeElement(&s.elemDefineInnerData, &start)
	if err != io.EOF && err != nil {
		return err
	}

	//
	s.Presentation = make(Presentation)
	s.Presentation[presentation.Display] = presentation.DisplayNone
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemDesc) Parent() Element {
	return s.parent
}
func (s *ElemDesc) Childrun() []Element {
	return nil
}
func (s *ElemDesc) createElement(name xml.Name) Element {
	return nil
}
func (s *ElemDesc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	if err != io.EOF && err != nil {
		return err
	}

	//
	err = d.DecodeElement(&s.Core, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.elemDesc, &start)
	if err != io.EOF && err != nil {
		return err
	}

	return nil
}

func (s *ElemTitle) Parent() Element {
	return s.parent
}
func (s *ElemTitle) Childrun() []Element {
	return nil
}
func (s *ElemTitle) createElement(name xml.Name) Element {
	return nil
}
func (s *ElemTitle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	if err != io.EOF && err != nil {
		return err
	}

	//
	err = d.DecodeElement(&s.Core, &start)
	if err != io.EOF && err != nil {
		return err
	}

	err = d.DecodeElement(&s.elemTitle, &start)
	if err != io.EOF && err != nil {
		return err
	}

	return nil
}
