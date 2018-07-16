package svg11

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"strings"
	"fmt"
	"strconv"
)

type LengthValue struct {
	Value float32
	Unit LengthUnit
}


type LengthUnit uint8

const (
	LengthUnitPx   LengthUnit = iota
	LengthUnitCm   LengthUnit = iota
	LengthUnitInch LengthUnit = iota
)

func (s *LengthValue) String() string {
	switch s.Unit {
	case LengthUnitPx:
		return fmt.Sprintf("%fpx", s.Value)
	case LengthUnitCm:
		return fmt.Sprintf("%fcm", s.Value)
	case LengthUnitInch:
		return fmt.Sprintf("%finch", s.Value)
	}
	return fmt.Sprintf("Unknown(%f)", s.Value)
}

func (s *LengthValue) parse(temp string) error {
	var unit string
	temp = strings.TrimSpace(temp)
	temp = strings.ToLower(temp)
	//
	for i := len(temp)-1; i >= 0; i--{
		code := temp[i]
		if !('a' < code && code < 'z'){
			if i != len(temp)-1{
				unit = string(temp[i + 1:])
			}
			temp = string(temp[:i + 1])
			break
		}
	}
	//
	f32, err := strconv.ParseFloat(temp, 32)
	if err != nil {
		return err
	}
	switch unit {
	case "px":
		s.Value = float32(f32)
		s.Unit = LengthUnitPx
	case "cm":
		s.Value = float32(f32)
		s.Unit = LengthUnitCm
	case "inch":
		s.Value = float32(f32)
		s.Unit = LengthUnitInch
	default:
		return errors.New("Invalid Unit Type '"+ unit +"'")
	}
	//
	return nil
}
func (s *LengthValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var temp string
	err := d.DecodeElement(&temp, &start)
	if err != nil {
		return err
	}
	//
	return s.parse(temp)
}
func (s *LengthValue) UnmarshalXMLAttr(attr xml.Attr) error {
	return s.parse(attr.Value)
}
