package svg11

import (
	"encoding/xml"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/dom/svg11/lengthtype"
	"github.com/iamGreedy/psvg"
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
		Class                     string          `xml:"class,attr"`
		Style                     string          `xml:"style,attr"`
		ExternalResourcesRequired string          `xml:"externalResourcesRequired,attr"`
		Transform                 string          `xml:"transform,attr"`
		X                         lengthtype.Type `xml:"x,attr"`
		Y                         lengthtype.Type `xml:"y,attr"`
		Width                     lengthtype.Type `xml:"width,attr"`
		Height                    lengthtype.Type `xml:"height,attr"`
		Rx                        lengthtype.Type `xml:"rx,attr"`
		Ry                        lengthtype.Type `xml:"ry,attr"`
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
		Class                     string          `xml:"class,attr"`
		Style                     string          `xml:"style,attr"`
		ExternalResourcesRequired string          `xml:"externalResourcesRequired,attr"`
		Transform                 string          `xml:"transform,attr"`
		Cx                        lengthtype.Type `xml:"cx,attr"`
		Cy                        lengthtype.Type `xml:"cy,attr"`
		R                         lengthtype.Type `xml:"r,attr"`
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
		Class                     string          `xml:"class,attr"`
		Style                     string          `xml:"style,attr"`
		ExternalResourcesRequired string          `xml:"externalResourcesRequired,attr"`
		Transform                 string          `xml:"transform,attr"`
		Cx                        lengthtype.Type `xml:"cx,attr"`
		Cy                        lengthtype.Type `xml:"cy,attr"`
		Rx                        lengthtype.Type `xml:"rx,attr"`
		Ry                        lengthtype.Type `xml:"ry,attr"`
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
		Class                     string          `xml:"class,attr"`
		Style                     string          `xml:"style,attr"`
		ExternalResourcesRequired string          `xml:"externalResourcesRequired,attr"`
		Transform                 string          `xml:"transform,attr"`
		X1                        lengthtype.Type `xml:"x1,attr"`
		Y1                        lengthtype.Type `xml:"y1,attr"`
		X2                        lengthtype.Type `xml:"x2,attr"`
		Y2                        lengthtype.Type `xml:"y2,attr"`
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


const circle = 1 - 0.552284749831


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
func (s *ElemRect) String() string {
	res := NewFuncStyle("Rect")
	if s.X.Value != 0 {
		res.Map("x", s.X.String())
	}
	if s.Y.Value != 0 {
		res.Map("y", s.Y.String())
	}
	if s.Rx.Value != 0 {
		res.Map("rx", s.Rx.String())
	}
	if s.Ry.Value != 0 {
		res.Map("ry", s.Ry.String())
	}
	if s.Width.Value != 0 {
		res.Map("width", s.Width.String())
	}
	if s.Height.Value != 0 {
		res.Map("height", s.Height.String())
	}

	s.Presentation.alignedarglist(res)
	return res.Build()
}
func (s *ElemRect) Paths() []psvg.Elem {
	var x, y, w, h, rx, ry float32
	switch s.X.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		x = s.X.Value
	}
	switch s.Y.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		y = s.Y.Value
	}
	switch s.Width.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		w = s.Width.Value
	}
	switch s.Height.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		h = s.Height.Value
	}
	switch s.Rx.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		rx = s.Rx.Value
	}
	switch s.Ry.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		ry = s.Ry.Value
	}

	if rx == 0 || ry == 0 {
		return []psvg.Elem{
			psvg.MoveToAbs{To: mgl32.Vec2{x, y}},
			psvg.LineToAbs{To: mgl32.Vec2{x + w, y}},
			psvg.LineToAbs{To: mgl32.Vec2{x + w, y + h}},
			psvg.LineToAbs{To: mgl32.Vec2{x, y + h}},
			psvg.ClosePath{},
		}
	}
	if rx > w/2 {
		rx = w / 2
	}
	if ry > h/2 {
		ry = h / 2
	}
	return []psvg.Elem{
		psvg.MoveToAbs{
			To: mgl32.Vec2{x, y + ry}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x, y + ry * circle},
			P1: mgl32.Vec2{x + rx * circle, y},
			To: mgl32.Vec2{x + rx, y}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + w - rx, y}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + w - rx * circle, y},
			P1: mgl32.Vec2{x + w, y + ry * circle},
			To: mgl32.Vec2{x + w, y + ry}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + w, y + h - ry}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + w, y + h - ry * circle},
			P1: mgl32.Vec2{x + w - rx * circle, y + h},
			To: mgl32.Vec2{x + w - rx, y + h}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + rx, y + h}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + rx * circle, y + h},
			P1: mgl32.Vec2{x, y + h - ry * circle},
			To: mgl32.Vec2{x, y + h - ry}},
		psvg.ClosePath{},
	}
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
func (s *ElemCircle) String() string {
	res := NewFuncStyle("circle")
	if s.Cx.Value != 0 {
		res.Map("cx", s.Cx.String())
	}
	if s.Cy.Value != 0 {
		res.Map("cy", s.Cy.String())
	}
	if s.R.Value != 0 {
		res.Map("r", s.R.String())
	}

	s.Presentation.alignedarglist(res)
	return res.Build()
}
func (s *ElemCircle) Paths() []psvg.Elem {
	var x, y, r float32
	switch s.Cx.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		x = s.Cx.Value
	}
	switch s.Cy.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		y = s.Cy.Value
	}
	switch s.R.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		r = s.R.Value
	}
	l := r * 2
	return []psvg.Elem{
		psvg.MoveToAbs{
			To: mgl32.Vec2{x, y + r}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x, y + r * circle},
			P1: mgl32.Vec2{x + r * circle, y},
			To: mgl32.Vec2{x + r, y}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + l - r, y}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + l - r * circle, y},
			P1: mgl32.Vec2{x + l, y + r * circle},
			To: mgl32.Vec2{x + l, y + r}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + l, y + l - r}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + l, y + l - r * circle},
			P1: mgl32.Vec2{x + l - r * circle, y + l},
			To: mgl32.Vec2{x + l - r, y + l}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + r, y + l}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + r * circle, y + l},
			P1: mgl32.Vec2{x, y + l - r * circle},
			To: mgl32.Vec2{x, y + l - r}},
		psvg.ClosePath{},
	}
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
func (s *ElemEllipse) String() string {
	res := NewFuncStyle("Ellipse")
	if s.Cx.Value != 0 {
		res.Map("cx", s.Cx.String())
	}
	if s.Cy.Value != 0 {
		res.Map("cy", s.Cy.String())
	}
	if s.Rx.Value != 0 {
		res.Map("rx", s.Rx.String())
	}
	if s.Ry.Value != 0 {
		res.Map("ry", s.Ry.String())
	}

	s.Presentation.alignedarglist(res)
	return res.Build()
}
func (s *ElemEllipse) Paths() []psvg.Elem {
	var x, y, rx, ry float32
	switch s.Cx.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		x = s.Cx.Value
	}
	switch s.Cy.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		y = s.Cy.Value
	}
	switch s.Rx.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		rx = s.Rx.Value
	}
	switch s.Ry.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		ry = s.Ry.Value
	}
	w, h := 2 * rx, 2 * ry
	return []psvg.Elem{
		psvg.MoveToAbs{
			To: mgl32.Vec2{x, y + ry}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x, y + ry * circle},
			P1: mgl32.Vec2{x + rx * circle, y},
			To: mgl32.Vec2{x + rx, y}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + w - rx, y}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + w - rx * circle, y},
			P1: mgl32.Vec2{x + w, y + ry * circle},
			To: mgl32.Vec2{x + w, y + ry}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + w, y + h - ry}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + w, y + h - ry * circle},
			P1: mgl32.Vec2{x + w - rx * circle, y + h},
			To: mgl32.Vec2{x + w - rx, y + h}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x + rx, y + h}},
		psvg.CurveToCubicAbs{
			P0: mgl32.Vec2{x + rx * circle, y + h},
			P1: mgl32.Vec2{x, y + h - ry * circle},
			To: mgl32.Vec2{x, y + h - ry}},
		psvg.ClosePath{},
	}
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
func (s *ElemLine) String() string {
	res := NewFuncStyle("Line")
	res.Map("x1", s.X1.String())
	res.Map("y1", s.Y1.String())
	res.Map("x2", s.X2.String())
	res.Map("y2", s.Y2.String())
	s.Presentation.alignedarglist(res)
	return res.Build()
}
func (s *ElemLine) Paths() []psvg.Elem {
	var x1, y1, x2, y2 float32
	switch s.X1.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		x1 = s.X1.Value
	}
	switch s.X2.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		x2 = s.X2.Value
	}


	switch s.Y1.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		y1 = s.Y1.Value
	}
	switch s.Y2.Unit {
	case lengthtype.NUMBER:
		fallthrough
	case lengthtype.PX:
		y2 = s.Y2.Value
	}


	return []psvg.Elem{
		psvg.MoveToAbs{
			To: mgl32.Vec2{x1, y1}},
		psvg.LineToAbs{
			To: mgl32.Vec2{x2, y2}},
	}
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
func (s *ElemPolyline) String() string {
	return "Polyline"
}
// TODO  : ElemPolyline implements
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
func (s *ElemPolygon) String() string {
	return "Polygon"
}
