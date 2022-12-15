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
	"fmt"

	"github.com/gnaoh1379/unioffice"
	"github.com/gnaoh1379/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_Document struct {
	ConformanceAttr sharedTypes.ST_ConformanceClass
	// Document Background
	Background *CT_Background
	Body       *CT_Body
}

func NewCT_Document() *CT_Document {
	ret := &CT_Document{}
	ret.ConformanceAttr = sharedTypes.ST_ConformanceClass(1)
	return ret
}

func (m *CT_Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ConformanceAttr.MarshalXMLAttr(xml.Name{Local: "w:conformance"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	if m.Body != nil {
		sebody := xml.StartElement{Name: xml.Name{Local: "w:body"}}
		e.EncodeElement(m.Body, sebody)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Document) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ConformanceAttr = sharedTypes.ST_ConformanceClass(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Document:
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "body"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "body"}:
				m.Body = NewCT_Body()
				if err := d.DecodeElement(m.Body, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Document %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Document
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Document and its children
func (m *CT_Document) Validate() error {
	return m.ValidateWithPath("CT_Document")
}

// ValidateWithPath validates the CT_Document and its children, prefixing error messages with path
func (m *CT_Document) ValidateWithPath(path string) error {
	if m.ConformanceAttr == sharedTypes.ST_ConformanceClassUnset {
		return fmt.Errorf("%s/ConformanceAttr is a mandatory field", path)
	}
	if err := m.ConformanceAttr.ValidateWithPath(path + "/ConformanceAttr"); err != nil {
		return err
	}
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	if m.Body != nil {
		if err := m.Body.ValidateWithPath(path + "/Body"); err != nil {
			return err
		}
	}
	return nil
}
