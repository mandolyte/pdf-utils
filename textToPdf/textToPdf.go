package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	papersize := flag.String("ps", "Letter", "Paper size; default 'Letter'")
	font := flag.String("font", "Arial", "Font 'Arial', 'Courier', or 'Times'; default Arial")
	style := flag.String("style", "", "Style: B for bold, U for underline, I for Italics or any combo thereof")
	fontsize := flag.Float64("fs", 12, "Font size; default 12")
	help := flag.Bool("help", false, "Show usage message")

	input := flag.String("i", "README.md", "Input text filename; default 'README.md'")
	output := flag.String("o", "README.pdf", "Output PDF filename; default 'README.pdf'")

	flag.Parse()

	if *help {
		usage("Help Message")
		os.Exit(0)
	}

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", *papersize, ".")
	pdf.AddPage()
	pdf.SetFont(*font, *style, *fontsize)
	pdf.Write(5, string(content))
	err = pdf.OutputFileAndClose(*output)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: textToPdf [options]\n")
	flag.PrintDefaults()
}
