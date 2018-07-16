package fullspec

import (
	"github.com/dom/svg11/lengthtype"
	"github.com/dom/svg11/angletype"
	"github.com/dom/svg11/colortype"
	"image/color"
)

// https://www.w3.org/TR/SVG11/types.html
type (
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGElement
	Element struct {
		ID       string
		XMLBase  string
		Owner    SVGElement
		Viewport *Element
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedBoolean
	AnimatedBoolean struct {
		Base bool
		Anim bool
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedString
	AnimatedString struct {
		Base string
		Anim string
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGStringList
	StringList struct {
		data []string
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedEnumeration
	AnimatedEnumeration struct {
		Base uint16
		Anim uint16
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedInteger
	AnimatedInteger struct {
		Base int64
		Anim int64
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGNumber
	Number float32
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedNumber
	AnimatedNumber struct {
		Base float32
		Anim float32
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGNumberList
	NumberList struct {
		data []Number
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedNumberList
	AnimatedNumberList struct {
		Base NumberList
		Anim NumberList
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGLength
	Length struct {
		Type lengthtype.Type
		Value float32
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedLength
	AnimatedLength struct {
		Base Length
		Anim Length
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGLengthList
	LengthList struct {
		data []Length
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedLengthList
	AnimatedLengthList struct {
		Base LengthList
		Anim LengthList
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAngle
	Angle struct {
		Type angletype.Type
		Value float32
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedAngle
	AnimatedAngle struct {
		Base Angle
		Anim Angle
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGColor
	Color struct {
		Type colortype.Type
		// deprecated, IICColor
		Value color.RGBA
	}
	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGICCColor
	// deprecated, don't implement

	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGRect
	Rect struct{
		X, Y float32
		W, H float32
	}

	// https://www.w3.org/TR/SVG11/types.html#InterfaceSVGAnimatedRect
	AnimatedRect struct{
		Base Rect
		Anim Rect
	}

	//

)

func (s *StringList) Length() int {
	return len(s.data)
}
func (s *StringList) Initialize(n string) string {
	for i := range s.data {
		s.data[i] = n
	}
	return n
}
func (s *StringList) Get(i int) string {
	return s.data[i]
}
func (s *StringList) InsertBefore(n string, i int) string {
	front := s.data[:i]
	back := s.data[i:]
	s.data = append(front, append([]string{n}, back...)...)
	return n
}
func (s *StringList) Replace(n string, i int) string {
	s.data[i] = n
	return n
}
func (s *StringList) Remove(i int) string {
	r := s.data[i]
	s.data = append(s.data[:i], s.data[i+1:]...)
	return r
}
func (s *StringList) Append(n string) string {
	s.data = append(s.data, n)
	return n
}

func (s *NumberList) Length() int {
	return len(s.data)
}
func (s *NumberList) Initialize(n Number) Number {
	for i := range s.data {
		s.data[i] = n
	}
	return n
}
func (s *NumberList) Get(i int) Number {
	return s.data[i]
}
func (s *NumberList) InsertBefore(n Number, i int) Number {
	front := s.data[:i]
	back := s.data[i:]
	s.data = append(front, append([]Number{n}, back...)...)
	return n
}
func (s *NumberList) Replace(n Number, i int) Number {
	s.data[i] = n
	return n
}
func (s *NumberList) Remove(i int) Number {
	r := s.data[i]
	s.data = append(s.data[:i], s.data[i+1:]...)
	return r
}
func (s *NumberList) Append(n Number) Number {
	s.data = append(s.data, n)
	return n
}

func (s *Length ) Specified() float32{
	panic("implement me")
}
func (s *Length ) String() string {
	panic("implement me")
}
func (s *Length ) Convert(to lengthtype.Type) string {
	panic("implement me")
}

func (s *Angle ) Specified() float32{
	panic("implement me")
}
func (s *Angle ) String() string {
	panic("implement me")
}
func (s *Angle ) Convert(to lengthtype.Type) string {
	panic("implement me")
}