// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"

	"github.com/gnaoh1379/unioffice"
)

type CT_Colors struct {
	// Color Indexes
	IndexedColors *CT_IndexedColors
	// MRU Colors
	MruColors *CT_MRUColors
}

func NewCT_Colors() *CT_Colors {
	ret := &CT_Colors{}
	return ret
}

func (m *CT_Colors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.IndexedColors != nil {
		seindexedColors := xml.StartElement{Name: xml.Name{Local: "ma:indexedColors"}}
		e.EncodeElement(m.IndexedColors, seindexedColors)
	}
	if m.MruColors != nil {
		semruColors := xml.StartElement{Name: xml.Name{Local: "ma:mruColors"}}
		e.EncodeElement(m.MruColors, semruColors)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Colors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Colors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "indexedColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "indexedColors"}:
				m.IndexedColors = NewCT_IndexedColors()
				if err := d.DecodeElement(m.IndexedColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "mruColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "mruColors"}:
				m.MruColors = NewCT_MRUColors()
				if err := d.DecodeElement(m.MruColors, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Colors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Colors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Colors and its children
func (m *CT_Colors) Validate() error {
	return m.ValidateWithPath("CT_Colors")
}

// ValidateWithPath validates the CT_Colors and its children, prefixing error messages with path
func (m *CT_Colors) ValidateWithPath(path string) error {
	if m.IndexedColors != nil {
		if err := m.IndexedColors.ValidateWithPath(path + "/IndexedColors"); err != nil {
			return err
		}
	}
	if m.MruColors != nil {
		if err := m.MruColors.ValidateWithPath(path + "/MruColors"); err != nil {
			return err
		}
	}
	return nil
}