package svg11

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"github.com/pkg/errors"
)


type Name struct {
	Space string
	Local string
}

func (s *Name) UnmarshalXMLAttr(attr xml.Attr) error {
	splt := strings.Split(attr.Value, ":")
	switch len(splt) {
	case 1:
		s.Local = splt[0]
	case 2:
		s.Space = splt[0]
		s.Local = splt[1]
	default:
		return errors.New("Invalid Name")
	}
	return nil
}

func (s Name) String() string {
	if len(s.Space) > 0{
		return s.Space + ":" + s.Local
	}
	return s.Local
}
