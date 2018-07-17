package presentation

type DisplayValue uint8

const (
	DisplayInline           DisplayValue = iota
	DisplayBlock            DisplayValue = iota
	DisplayListitem         DisplayValue = iota
	DisplayRunin            DisplayValue = iota
	DisplayCompact          DisplayValue = iota
	DisplayMarker           DisplayValue = iota
	DisplayTable            DisplayValue = iota
	DisplayInlinetable      DisplayValue = iota
	DisplayTablerowgroup    DisplayValue = iota
	DisplayTableheadergroup DisplayValue = iota
	DisplayTablefootergroup DisplayValue = iota
	DisplayTablerow         DisplayValue = iota
	DisplayTablecolumngroup DisplayValue = iota
	DisplayTablecolumn      DisplayValue = iota
	DisplayTablecell        DisplayValue = iota
	DisplayTablecaption     DisplayValue = iota
	DisplayNone             DisplayValue = iota
	DisplayInherit          DisplayValue = iota
)
