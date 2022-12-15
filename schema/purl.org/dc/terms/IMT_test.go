// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/gnaoh1379/unioffice/schema/purl.org/dc/terms"
)

func TestIMTConstructor(t *testing.T) {
	v := terms.NewIMT()
	if v == nil {
		t.Errorf("terms.NewIMT must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.IMT should validate: %s", err)
	}
}

func TestIMTMarshalUnmarshal(t *testing.T) {
	v := terms.NewIMT()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewIMT()
	xml.Unmarshal(buf, v2)
}
