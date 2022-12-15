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
	"fmt"
	"strconv"

	"github.com/gnaoh1379/unioffice/schema/soo/ofc/sharedTypes"
)

type AG_OfficeShapeAttributes struct {
	SptAttr            *float32
	ConnectortypeAttr  OfcST_ConnectorType
	BwmodeAttr         OfcST_BWMode
	BwpureAttr         OfcST_BWMode
	BwnormalAttr       OfcST_BWMode
	ForcedashAttr      sharedTypes.ST_TrueFalse
	OleiconAttr        sharedTypes.ST_TrueFalse
	OleAttr            sharedTypes.ST_TrueFalseBlank
	PreferrelativeAttr sharedTypes.ST_TrueFalse
	CliptowrapAttr     sharedTypes.ST_TrueFalse
	ClipAttr           sharedTypes.ST_TrueFalse
}

func NewAG_OfficeShapeAttributes() *AG_OfficeShapeAttributes {
	ret := &AG_OfficeShapeAttributes{}
	return ret
}

func (m *AG_OfficeShapeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spt"},
			Value: fmt.Sprintf("%v", *m.SptAttr)})
	}
	if m.ConnectortypeAttr != OfcST_ConnectorTypeUnset {
		attr, err := m.ConnectortypeAttr.MarshalXMLAttr(xml.Name{Local: "connectortype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwmodeAttr != OfcST_BWModeUnset {
		attr, err := m.BwmodeAttr.MarshalXMLAttr(xml.Name{Local: "bwmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwpureAttr != OfcST_BWModeUnset {
		attr, err := m.BwpureAttr.MarshalXMLAttr(xml.Name{Local: "bwpure"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwnormalAttr != OfcST_BWModeUnset {
		attr, err := m.BwnormalAttr.MarshalXMLAttr(xml.Name{Local: "bwnormal"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForcedashAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ForcedashAttr.MarshalXMLAttr(xml.Name{Local: "forcedash"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleiconAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OleiconAttr.MarshalXMLAttr(xml.Name{Local: "oleicon"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.OleAttr.MarshalXMLAttr(xml.Name{Local: "ole"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PreferrelativeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PreferrelativeAttr.MarshalXMLAttr(xml.Name{Local: "preferrelative"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CliptowrapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.CliptowrapAttr.MarshalXMLAttr(xml.Name{Local: "cliptowrap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ClipAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ClipAttr.MarshalXMLAttr(xml.Name{Local: "clip"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}

func (m *AG_OfficeShapeAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spt" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.SptAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connectortype" {
			m.ConnectortypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwmode" {
			m.BwmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwpure" {
			m.BwpureAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwnormal" {
			m.BwnormalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oleicon" {
			m.OleiconAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "ole" {
			m.OleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "preferrelative" {
			m.PreferrelativeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "cliptowrap" {
			m.CliptowrapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "clip" {
			m.ClipAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_OfficeShapeAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_OfficeShapeAttributes and its children
func (m *AG_OfficeShapeAttributes) Validate() error {
	return m.ValidateWithPath("AG_OfficeShapeAttributes")
}

// ValidateWithPath validates the AG_OfficeShapeAttributes and its children, prefixing error messages with path
func (m *AG_OfficeShapeAttributes) ValidateWithPath(path string) error {
	if err := m.ConnectortypeAttr.ValidateWithPath(path + "/ConnectortypeAttr"); err != nil {
		return err
	}
	if err := m.BwmodeAttr.ValidateWithPath(path + "/BwmodeAttr"); err != nil {
		return err
	}
	if err := m.BwpureAttr.ValidateWithPath(path + "/BwpureAttr"); err != nil {
		return err
	}
	if err := m.BwnormalAttr.ValidateWithPath(path + "/BwnormalAttr"); err != nil {
		return err
	}
	if err := m.ForcedashAttr.ValidateWithPath(path + "/ForcedashAttr"); err != nil {
		return err
	}
	if err := m.OleiconAttr.ValidateWithPath(path + "/OleiconAttr"); err != nil {
		return err
	}
	if err := m.OleAttr.ValidateWithPath(path + "/OleAttr"); err != nil {
		return err
	}
	if err := m.PreferrelativeAttr.ValidateWithPath(path + "/PreferrelativeAttr"); err != nil {
		return err
	}
	if err := m.CliptowrapAttr.ValidateWithPath(path + "/CliptowrapAttr"); err != nil {
		return err
	}
	if err := m.ClipAttr.ValidateWithPath(path + "/ClipAttr"); err != nil {
		return err
	}
	return nil
}
