package lengthtype

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/xml"
	"errors"
)

type Unit uint16

const (
	UNKNOWN    Unit = 0
	NUMBER     Unit = 1
	PERCENTAGE Unit = 2
	EMS        Unit = 3
	EXS        Unit = 4
	PX         Unit = 5
	CM         Unit = 6
	MM         Unit = 7
	IN         Unit = 8
	PT         Unit = 9
	PC         Unit = 10
)

type Type struct {
	Value float32
	Unit  Unit
}

func (s Type) String() string {
	switch s.Unit {
	default:
		return "Unknown"
	case NUMBER:
		return fmt.Sprintf("%f", s.Value)
	case PERCENTAGE:
		return fmt.Sprintf("%f%%", s.Value)
	case PX:
		return fmt.Sprintf("%fpx", s.Value)
	case CM:
		return fmt.Sprintf("%fcm", s.Value)
	case MM:
		return fmt.Sprintf("%fmm", s.Value)
	case IN:
		return fmt.Sprintf("%fin", s.Value)
	}
}
func (s *Type) parse(temp string) error {
	var unit string
	temp = strings.TrimSpace(temp)
	temp = strings.ToLower(temp)
	//
	for i := len(temp) - 1; i >= 0; i-- {
		code := temp[i]
		if !(('a' < code && code < 'z') || code == '%') {
			if i != len(temp)-1 {
				unit = string(temp[i+1:])
			}
			temp = string(temp[:i+1])
			break
		}
	}
	//
	f32, err := strconv.ParseFloat(temp, 32)
	if err != nil {
		return err
	}
	switch unit {
	case "":
		s.Value = float32(f32)
		s.Unit = NUMBER
	case "px":
		s.Value = float32(f32)
		s.Unit = PX
	case "cm":
		s.Value = float32(f32)
		s.Unit = CM
	case "mm":
		s.Value = float32(f32)
		s.Unit = MM
	case "in":
		s.Value = float32(f32)
		s.Unit = IN
	default:
		return errors.New("Invalid Unit Unit '" + unit + "'")
	}
	//
	return nil
}
func (s *Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var temp string
	err := d.DecodeElement(&temp, &start)
	if err != nil {
		return err
	}
	//
	return s.parse(temp)
}
func (s *Type) UnmarshalXMLAttr(attr xml.Attr) error {
	return s.parse(attr.Value)
}
