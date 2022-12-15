// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/gnaoh1379/unioffice/schema/soo/ofc/math"
)

func TestCT_MathPrChoiceConstructor(t *testing.T) {
	v := math.NewCT_MathPrChoice()
	if v == nil {
		t.Errorf("math.NewCT_MathPrChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MathPrChoice should validate: %s", err)
	}
}

func TestCT_MathPrChoiceMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MathPrChoice()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MathPrChoice()
	xml.Unmarshal(buf, v2)
}
