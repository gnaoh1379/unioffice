// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/gnaoh1379/unioffice"
)

type CT_CxnList struct {
	Cxn []*CT_Cxn
}

func NewCT_CxnList() *CT_CxnList {
	ret := &CT_CxnList{}
	return ret
}

func (m *CT_CxnList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Cxn != nil {
		secxn := xml.StartElement{Name: xml.Name{Local: "cxn"}}
		for _, c := range m.Cxn {
			e.EncodeElement(c, secxn)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CxnList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CxnList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "cxn"}:
				tmp := NewCT_Cxn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cxn = append(m.Cxn, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_CxnList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CxnList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CxnList and its children
func (m *CT_CxnList) Validate() error {
	return m.ValidateWithPath("CT_CxnList")
}

// ValidateWithPath validates the CT_CxnList and its children, prefixing error messages with path
func (m *CT_CxnList) ValidateWithPath(path string) error {
	for i, v := range m.Cxn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cxn[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
