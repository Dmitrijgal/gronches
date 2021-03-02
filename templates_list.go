package main

import (
	"encoding/xml"
)

// TemplatesList is structure for saving data which later will be parsed. It countains Template structures.
type TemplatesList struct {
	XMLName  xml.Name   `xml:"DATA"`
	Template []Template `xml:"ROW"`
}

// AppendVariables is TemplatesList method which appends variable field to every template in list.
func (templateList *TemplatesList) AppendVariables() {

	for i, t := range templateList.Template {
		t.AppendVariables()
		templateList.Template[i] = t
	}
}
