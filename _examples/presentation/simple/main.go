// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"

	"github.com/gnaoh1379/unioffice/color"
	"github.com/gnaoh1379/unioffice/measurement"
	"github.com/gnaoh1379/unioffice/schema/soo/dml"

	"github.com/gnaoh1379/unioffice/presentation"
)

func main() {
	ppt := presentation.New()
	defer ppt.Close()
	for i := 0; i < 5; i++ {
		slide := ppt.AddSlide()

		tb := slide.AddTextBox()
		tb.Properties().SetGeometry(dml.ST_ShapeTypeStar10)

		tb.Properties().SetWidth(3 * measurement.Inch)
		pos := measurement.Distance(i) * measurement.Inch
		tb.Properties().SetPosition(pos, pos)

		tb.Properties().SetSolidFill(color.AliceBlue)
		tb.Properties().LineProperties().SetSolidFill(color.Blue)

		p := tb.AddParagraph()
		p.Properties().SetAlign(dml.ST_TextAlignTypeCtr)

		r := p.AddRun()
		r.SetText("gooxml")
		r.Properties().SetSize(24 * measurement.Point)

	}
	if err := ppt.Validate(); err != nil {
		log.Fatal(err)
	}
	ppt.SaveToFile("simple.pptx")
}
