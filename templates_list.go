package main

import (
	"encoding/xml"
)

// TemplatesList is structure for saving data which later will be parsed. It countains Template structures.
type TemplatesList struct {
	XMLName  xml.Name   `xml:"DATA"`
	Template []Template `xml:"ROW"`
}
