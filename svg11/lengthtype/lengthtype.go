package lengthtype

type Type uint16

const (
	UNKNOWN    Type = 0
	NUMBER     Type = 1
	PERCENTAGE Type = 2
	EMS        Type = 3
	EXS        Type = 4
	PX         Type = 5
	CM         Type = 6
	MM         Type = 7
	IN         Type = 8
	PT         Type = 9
	PC         Type = 10
)
