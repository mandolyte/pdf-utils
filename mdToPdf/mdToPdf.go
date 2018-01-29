package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mandolyte/mdtopdf"
)

var input = flag.String("i", "", "Input text filename; default is os.Stdin")
var output = flag.String("o", "", "Output PDF filename; requiRed")
var help = flag.Bool("help", false, "Show usage message")
var trace = flag.String("trace", "", "Filename for trace log")

func main() {

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

	pf := mdtopdf.NewPdfRenderer("", "", *output, *trace)
	pf.Pdf.SetSubject("How to convert markdown to PDF", true)
	pf.Pdf.SetTitle("Example PDF converted from Markdown", true)
	pf.THeader = mdtopdf.Styler{Font: "Times", Style: "IUB", Size: 20, Spacing: 2,
		TextColor: mdtopdf.Color{Red: 0, Green: 0, Blue: 0},
		FillColor: mdtopdf.Color{Red: 179, Green: 179, Blue: 255}}
	pf.TBody = mdtopdf.Styler{Font: "Arial", Style: "", Size: 12, Spacing: 2,
		TextColor: mdtopdf.Color{Red: 0, Green: 0, Blue: 0},
		FillColor: mdtopdf.Color{Red: 255, Green: 102, Blue: 129}}

	// just for grins, let's make header 1 really ugly
	// note that H1 is not written with the CellFormat() method
	// and therefore the fill color is not actually used.
	pf.H1 = mdtopdf.Styler{Font: "Courier",
		Style: "IUB", Size: 28, Spacing: 10,
		TextColor: mdtopdf.Color{Red: 102, Green: 255, Blue: 102},
		FillColor: mdtopdf.Color{Red: 255, Green: 153, Blue: 255}}

	err = pf.Process(content)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: convert [options]\n")
	flag.PrintDefaults()
	os.Exit(0)
}
