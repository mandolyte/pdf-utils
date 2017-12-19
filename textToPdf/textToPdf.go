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
	tabwidth := flag.Int("tabwidth", 4, "Number of spaces for a tab character; default is 4")
	help := flag.Bool("help", false, "Show usage message")

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
	sContent = strings.Replace(sContent, "\t", strings.Repeat(" ", *tabwidth), -1)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "pt", *papersize, ".")
	pdf.AddPage()
	pdf.SetFont(*font, *style, *fontsize)
	pdf.Write(*lineheight, sContent)
	err = pdf.OutputFileAndClose(*output)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: textToPdf [options]\n")
	fmt.Println("Note: font and line height are 'points', 1/72 inch")
	flag.PrintDefaults()
	os.Exit(0)
}
