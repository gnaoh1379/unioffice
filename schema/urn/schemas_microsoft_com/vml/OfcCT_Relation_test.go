// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gnaoh1379/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RelationConstructor(t *testing.T) {
	v := vml.NewOfcCT_Relation()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Relation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Relation should validate: %s", err)
	}
}

func TestOfcCT_RelationMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Relation()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Relation()
	xml.Unmarshal(buf, v2)
}
