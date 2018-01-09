package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	papersize := flag.String("ps", "Letter", "Paper size; default 'Letter'")
	font := flag.String("font", "Arial", "Font 'Arial', 'Courier', or 'Times'; default Arial")
	style := flag.String("style", "", "Style: B for bold, U for underline, I for Italics or any combo thereof")
	fontsize := flag.Float64("fs", 12, "Font size; default 12")
	lineheight := flag.Float64("lh", 15, "Line height; default 15")
	help := flag.Bool("help", false, "Show usage message")
	compressed := flag.Bool("compressed", false, "Don't add blank line between items")

	input := flag.String("i", "", "Input text filename; default is os.Stdin")
	output := flag.String("o", "", "Output PDF filename; required")

	flag.Parse()

	if *help {
		usage("Help Message")
	}

	if *output == "" {
		usage("Output PDF filename is required")
	}

	// get text for PDF
	var content []byte
	var err error
	if *input == "" {
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		content, err = ioutil.ReadFile(*input)
		if err != nil {
			log.Fatal(err)
		}
	}

	sContent := string(content)
	sContent = strings.Replace(sContent, "\r\n", "\n", -1)
	items := strings.Split(sContent, "\n")

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "pt", *papersize, ".")
	pdf.AddPage()
	pdf.SetFont(*font, *style, *fontsize)
	em := pdf.GetStringWidth("m")

	pdf.Write(*lineheight, "This is a demonstration of creating a numbered list.\n\n")

	// lists are normally indented
	pdf.SetLeftMargin(3 * em)

	itemNum := 0
	for _, item := range items {
		if item == "" {
			continue
		}
		itemNum++
		// output the number of the item, aligned right
		pdf.CellFormat(3*em, *lineheight,
			fmt.Sprintf("%v. ", itemNum),
			"", 0, "RB", false, 0, "")
		pdf.MultiCell(0, *lineheight, item, "", "", false)
		if !*compressed {
			pdf.Ln(-1)
		}
	}

	pdf.Write(*lineheight, "\nEnd of demonstration.")

	err = pdf.OutputFileAndClose(*output)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: textToList [options]\n")
	fmt.Println("Note: font and line height are 'points', 1/72 inch")
	fmt.Println("This example takes a text file of paragraphs and")
	fmt.Println("demonstrates how to create a numbered list in PDF format")
	fmt.Println("Blank lines are skipped!")
	flag.PrintDefaults()
	os.Exit(0)
}

/*

pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")

*/
