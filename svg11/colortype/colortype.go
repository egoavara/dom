package colortype

type Type uint16

const (
	UNKNOWN           Type = 0
	RGBCOLOR          Type = 1
	RGBCOLOR_ICCCOLOR Type = 2
	CURRENTCOLOR      Type = 3
)
