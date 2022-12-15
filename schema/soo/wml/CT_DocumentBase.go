// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"

	"github.com/gnaoh1379/unioffice"
)

type CT_DocumentBase struct {
	// Document Background
	Background *CT_Background
}

func NewCT_DocumentBase() *CT_DocumentBase {
	ret := &CT_DocumentBase{}
	return ret
}

func (m *CT_DocumentBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocumentBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocumentBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "background"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "background"}:
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_DocumentBase %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocumentBase
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocumentBase and its children
func (m *CT_DocumentBase) Validate() error {
	return m.ValidateWithPath("CT_DocumentBase")
}

// ValidateWithPath validates the CT_DocumentBase and its children, prefixing error messages with path
func (m *CT_DocumentBase) ValidateWithPath(path string) error {
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	return nil
}
