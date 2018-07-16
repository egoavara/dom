package svg11

import "net/url"

// https://www.w3.org/TR/SVG11/struct.html

type (
	Document struct {
		Title    string
		Referrer string
		Domain   string
		URL      url.URL
		Root Element
	}
)
