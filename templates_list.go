package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
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

// ReadUnmarshalXML is function to read xml file.
func ReadUnmarshalXML(r io.Reader) (template TemplatesList, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return template, err
	}
	xml.Unmarshal(data, &template)
	return template, nil
}
