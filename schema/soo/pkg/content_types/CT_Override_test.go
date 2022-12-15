// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package content_types_test

import (
	"encoding/xml"
	"testing"

	"github.com/gnaoh1379/unioffice/schema/soo/pkg/content_types"
)

func TestCT_OverrideConstructor(t *testing.T) {
	v := content_types.NewCT_Override()
	if v == nil {
		t.Errorf("content_types.NewCT_Override must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.CT_Override should validate: %s", err)
	}
}

func TestCT_OverrideMarshalUnmarshal(t *testing.T) {
	v := content_types.NewCT_Override()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewCT_Override()
	xml.Unmarshal(buf, v2)
}
