package unioffice_test

import (
	"github.com/gnaoh1379/unioffice/document"
	"github.com/gnaoh1379/unioffice/spreadsheet"
)

func Example_document() {
	// see the github.com/gnaoh1379/unioffice/document documentation or _examples/document
	// for more examples
	doc := document.New()
	defer doc.Close()
	doc.AddParagraph().AddRun().AddText("Hello World!")
	doc.SaveToFile("document.docx")
}

func Example_spreadsheeet() {
	// see the github.com/gnaoh1379/unioffice/spreadsheet documentation or _examples/spreadsheet
	// for more examples
	ss := spreadsheet.New()
	defer ss.Close()
	sheet := ss.AddSheet()
	sheet.AddRow().AddCell().SetString("Hello")
	sheet.Cell("B1").SetString("World!")
	ss.SaveToFile("workbook.xlsx")
}
