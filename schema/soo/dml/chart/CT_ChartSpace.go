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

	"github.com/gnaoh1379/unioffice"
	"github.com/gnaoh1379/unioffice/schema/soo/dml"
)

type CT_ChartSpace struct {
	Date1904       *CT_Boolean
	Lang           *CT_TextLanguageID
	RoundedCorners *CT_Boolean
	Style          *CT_Style
	ClrMapOvr      *dml.CT_ColorMapping
	PivotSource    *CT_PivotSource
	Protection     *CT_Protection
	Chart          *CT_Chart
	SpPr           *dml.CT_ShapeProperties
	TxPr           *dml.CT_TextBody
	ExternalData   *CT_ExternalData
	PrintSettings  *CT_PrintSettings
	UserShapes     *CT_RelId
	ExtLst         *CT_ExtensionList
}

func NewCT_ChartSpace() *CT_ChartSpace {
	ret := &CT_ChartSpace{}
	ret.Chart = NewCT_Chart()
	return ret
}

func (m *CT_ChartSpace) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Date1904 != nil {
		sedate1904 := xml.StartElement{Name: xml.Name{Local: "c:date1904"}}
		e.EncodeElement(m.Date1904, sedate1904)
	}
	if m.Lang != nil {
		selang := xml.StartElement{Name: xml.Name{Local: "c:lang"}}
		e.EncodeElement(m.Lang, selang)
	}
	if m.RoundedCorners != nil {
		seroundedCorners := xml.StartElement{Name: xml.Name{Local: "c:roundedCorners"}}
		e.EncodeElement(m.RoundedCorners, seroundedCorners)
	}
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "c:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ClrMapOvr != nil {
		seclrMapOvr := xml.StartElement{Name: xml.Name{Local: "c:clrMapOvr"}}
		e.EncodeElement(m.ClrMapOvr, seclrMapOvr)
	}
	if m.PivotSource != nil {
		sepivotSource := xml.StartElement{Name: xml.Name{Local: "c:pivotSource"}}
		e.EncodeElement(m.PivotSource, sepivotSource)
	}
	if m.Protection != nil {
		seprotection := xml.StartElement{Name: xml.Name{Local: "c:protection"}}
		e.EncodeElement(m.Protection, seprotection)
	}
	sechart := xml.StartElement{Name: xml.Name{Local: "c:chart"}}
	e.EncodeElement(m.Chart, sechart)
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.ExternalData != nil {
		seexternalData := xml.StartElement{Name: xml.Name{Local: "c:externalData"}}
		e.EncodeElement(m.ExternalData, seexternalData)
	}
	if m.PrintSettings != nil {
		seprintSettings := xml.StartElement{Name: xml.Name{Local: "c:printSettings"}}
		e.EncodeElement(m.PrintSettings, seprintSettings)
	}
	if m.UserShapes != nil {
		seuserShapes := xml.StartElement{Name: xml.Name{Local: "c:userShapes"}}
		e.EncodeElement(m.UserShapes, seuserShapes)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartSpace) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Chart = NewCT_Chart()
lCT_ChartSpace:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "date1904"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "date1904"}:
				m.Date1904 = NewCT_Boolean()
				if err := d.DecodeElement(m.Date1904, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "lang"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "lang"}:
				m.Lang = NewCT_TextLanguageID()
				if err := d.DecodeElement(m.Lang, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "roundedCorners"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "roundedCorners"}:
				m.RoundedCorners = NewCT_Boolean()
				if err := d.DecodeElement(m.RoundedCorners, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "style"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "style"}:
				m.Style = NewCT_Style()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "clrMapOvr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "clrMapOvr"}:
				m.ClrMapOvr = dml.NewCT_ColorMapping()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pivotSource"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pivotSource"}:
				m.PivotSource = NewCT_PivotSource()
				if err := d.DecodeElement(m.PivotSource, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "protection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "protection"}:
				m.Protection = NewCT_Protection()
				if err := d.DecodeElement(m.Protection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "chart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "chart"}:
				if err := d.DecodeElement(m.Chart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "externalData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "externalData"}:
				m.ExternalData = NewCT_ExternalData()
				if err := d.DecodeElement(m.ExternalData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "printSettings"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "printSettings"}:
				m.PrintSettings = NewCT_PrintSettings()
				if err := d.DecodeElement(m.PrintSettings, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "userShapes"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "userShapes"}:
				m.UserShapes = NewCT_RelId()
				if err := d.DecodeElement(m.UserShapes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ChartSpace %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartSpace
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartSpace and its children
func (m *CT_ChartSpace) Validate() error {
	return m.ValidateWithPath("CT_ChartSpace")
}

// ValidateWithPath validates the CT_ChartSpace and its children, prefixing error messages with path
func (m *CT_ChartSpace) ValidateWithPath(path string) error {
	if m.Date1904 != nil {
		if err := m.Date1904.ValidateWithPath(path + "/Date1904"); err != nil {
			return err
		}
	}
	if m.Lang != nil {
		if err := m.Lang.ValidateWithPath(path + "/Lang"); err != nil {
			return err
		}
	}
	if m.RoundedCorners != nil {
		if err := m.RoundedCorners.ValidateWithPath(path + "/RoundedCorners"); err != nil {
			return err
		}
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ClrMapOvr != nil {
		if err := m.ClrMapOvr.ValidateWithPath(path + "/ClrMapOvr"); err != nil {
			return err
		}
	}
	if m.PivotSource != nil {
		if err := m.PivotSource.ValidateWithPath(path + "/PivotSource"); err != nil {
			return err
		}
	}
	if m.Protection != nil {
		if err := m.Protection.ValidateWithPath(path + "/Protection"); err != nil {
			return err
		}
	}
	if err := m.Chart.ValidateWithPath(path + "/Chart"); err != nil {
		return err
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if m.ExternalData != nil {
		if err := m.ExternalData.ValidateWithPath(path + "/ExternalData"); err != nil {
			return err
		}
	}
	if m.PrintSettings != nil {
		if err := m.PrintSettings.ValidateWithPath(path + "/PrintSettings"); err != nil {
			return err
		}
	}
	if m.UserShapes != nil {
		if err := m.UserShapes.ValidateWithPath(path + "/UserShapes"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
