package link

import "io"

type link struct {
	Href string
	Text string
}

var r io.Reader