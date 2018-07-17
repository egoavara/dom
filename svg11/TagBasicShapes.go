package svg11

import (
	"encoding/xml"
	"io"
)

// https://www.w3.org/TR/SVG11/shapes.html
type (
	// https://www.w3.org/TR/SVG11/shapes.html#RectElement
	ElemRect struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemRect
		//
		commonTree
	}
	elemRect struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		X                         string `xml:"x,attr"`
		Y                         string `xml:"y,attr"`
		Width                     string `xml:"width,attr"`
		Height                    string `xml:"height,attr"`
		Rx                        string `xml:"rx,attr"`
		Ry                        string `xml:"ry,attr"`
	}

	// https://www.w3.org/TR/SVG11/shapes.html#CircleElement
	ElemCircle struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemCircle
		//
		commonTree
	}
	elemCircle struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		Cx                        string `xml:"cx,attr"`
		Cy                        string `xml:"cy,attr"`
		R                         string `xml:"r,attr"`
	}

	// https://www.w3.org/TR/SVG11/shapes.html#EllipseElement
	ElemEllipse struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemEllipse
		//
		commonTree
	}
	elemEllipse struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		Cx                        string `xml:"cx,attr"`
		Cy                        string `xml:"cy,attr"`
		Rx                        string `xml:"rx,attr"`
		Ry                        string `xml:"ry,attr"`
	}

	// https://www.w3.org/TR/SVG11/shapes.html#LineElement
	ElemLine struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemLine
		//
		commonTree
	}
	elemLine struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		X1                        string `xml:"x1,attr"`
		Y1                        string `xml:"y1,attr"`
		X2                        string `xml:"x2,attr"`
		Y2                        string `xml:"y2,attr"`
	}

	// https://www.w3.org/TR/SVG11/shapes.html#PolylineElement
	ElemPolyline struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemPolyline
		//
		commonTree
	}
	elemPolyline struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		Points                    string `xml:"points,attr"`
	}

	// https://www.w3.org/TR/SVG11/shapes.html#PolygonElement
	ElemPolygon struct {
		ConditionalProcess
		Core
		GraphicalEvent
		Presentation
		elemPolygon
		//
		commonTree
	}
	elemPolygon struct {
		Class                     string `xml:"class,attr"`
		Style                     string `xml:"style,attr"`
		ExternalResourcesRequired string `xml:"externalResourcesRequired,attr"`
		Transform                 string `xml:"transform,attr"`
		Points                    string `xml:"points,attr"`
	}
)

func (s *ElemRect) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemRect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemRect, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemCircle) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemCircle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemCircle, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemEllipse) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemEllipse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemEllipse, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemLine) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemLine, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemPolyline) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemPolyline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemPolyline, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}

func (s *ElemPolygon) createElement(name xml.Name) Element {
	if e := create(name, createAnimatable, createDescriptive); e != nil {
		return e
	}
	return nil
}
func (s *ElemPolygon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	err = d.DecodeElement(&s.elemPolygon, &start)
	if err != io.EOF && err != nil {
		return err
	}
	s.Presentation.xmlAttrs(start.Attr...)

	return nil
}
