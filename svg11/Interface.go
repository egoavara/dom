package svg11

import "github.com/iamGreedy/psvg"


type Shape interface {
	Paths() []psvg.Elem
}
