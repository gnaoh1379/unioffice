// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"strconv"

	"github.com/gnaoh1379/unioffice"
)

type OfcShapedefaults struct {
	OfcCT_ShapeDefaults
}

func NewOfcShapedefaults() *OfcShapedefaults {
	ret := &OfcShapedefaults{}
	ret.OfcCT_ShapeDefaults = *NewOfcCT_ShapeDefaults()
	return ret
}

func (m *OfcShapedefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:shapedefaults"
	return m.OfcCT_ShapeDefaults.MarshalXML(e, start)
}

func (m *OfcShapedefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_ShapeDefaults = *NewOfcCT_ShapeDefaults()
	for _, attr := range start.Attr {
		if attr.Name.Local == "spidmax" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SpidmaxAttr = &parsed
			continue
		}
		if attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "stroke" {
			m.StrokeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "fill" {
			m.FillAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcShapedefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				m.Fill = NewFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				m.Stroke = NewStroke()
				if err := d.DecodeElement(m.Stroke, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				m.Textbox = NewTextbox()
				if err := d.DecodeElement(m.Textbox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				m.Shadow = NewShadow()
				if err := d.DecodeElement(m.Shadow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				m.Skew = NewOfcSkew()
				if err := d.DecodeElement(m.Skew, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				m.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(m.Extrusion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				m.Callout = NewOfcCallout()
				if err := d.DecodeElement(m.Callout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				m.Lock = NewOfcLock()
				if err := d.DecodeElement(m.Lock, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "colormru"}:
				m.Colormru = NewOfcCT_ColorMru()
				if err := d.DecodeElement(m.Colormru, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "colormenu"}:
				m.Colormenu = NewOfcCT_ColorMenu()
				if err := d.DecodeElement(m.Colormenu, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on OfcShapedefaults %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcShapedefaults
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcShapedefaults and its children
func (m *OfcShapedefaults) Validate() error {
	return m.ValidateWithPath("OfcShapedefaults")
}

// ValidateWithPath validates the OfcShapedefaults and its children, prefixing error messages with path
func (m *OfcShapedefaults) ValidateWithPath(path string) error {
	if err := m.OfcCT_ShapeDefaults.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}