// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/gnaoh1379/unioffice"
)

type EG_AreaChartShared struct {
	Grouping   *CT_Grouping
	VaryColors *CT_Boolean
	Ser        []*CT_AreaSer
	DLbls      *CT_DLbls
	DropLines  *CT_ChartLines
}

func NewEG_AreaChartShared() *EG_AreaChartShared {
	ret := &EG_AreaChartShared{}
	return ret
}

func (m *EG_AreaChartShared) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Grouping != nil {
		segrouping := xml.StartElement{Name: xml.Name{Local: "c:grouping"}}
		e.EncodeElement(m.Grouping, segrouping)
	}
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "c:varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.DropLines != nil {
		sedropLines := xml.StartElement{Name: xml.Name{Local: "c:dropLines"}}
		e.EncodeElement(m.DropLines, sedropLines)
	}
	return nil
}

func (m *EG_AreaChartShared) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_AreaChartShared:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "grouping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "grouping"}:
				m.Grouping = NewCT_Grouping()
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_AreaSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dropLines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dropLines"}:
				m.DropLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_AreaChartShared %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_AreaChartShared
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_AreaChartShared and its children
func (m *EG_AreaChartShared) Validate() error {
	return m.ValidateWithPath("EG_AreaChartShared")
}

// ValidateWithPath validates the EG_AreaChartShared and its children, prefixing error messages with path
func (m *EG_AreaChartShared) ValidateWithPath(path string) error {
	if m.Grouping != nil {
		if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
			return err
		}
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.DropLines != nil {
		if err := m.DropLines.ValidateWithPath(path + "/DropLines"); err != nil {
			return err
		}
	}
	return nil
}