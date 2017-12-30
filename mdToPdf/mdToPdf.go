package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	bf "gopkg.in/russross/blackfriday.v2"
)

var input = flag.String("i", "", "Input text filename; default is os.Stdin")
var output = flag.String("o", "", "Output PDF filename; required")
var debug = flag.Bool("debug", false, "Output debug info; default false")
var help = flag.Bool("help", false, "Show usage message")

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

	//sContent := string(content)
	//sContent = strings.Replace(sContent, "\t", strings.Repeat(" ", *tabwidth), -1)

	pf := NewPdfRenderer()

	_ = bf.Run(content, bf.WithRenderer(pf))

	err = pf.pdf.OutputFileAndClose(*output)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}

	//	pdf := gofpdf.New(gofpdf.OrientationPortrait, "pt", *papersize, ".")

}

// Styler is the struct to capture the styling features for text
type Styler struct {
	Font    string
	Style   string
	Size    float64
	Spacing float64
}

// PdfRenderer is the struct to manage conversion of a markdown object
// to PDF format.
type PdfRenderer struct {
	pdf                *gofpdf.Fpdf
	Orientation, units string
	Papersize, fontdir string

	// current settings
	current Styler

	// normal text
	Normal Styler

	// backticked text
	Backtick Styler

	// headings
	H1 Styler
	H2 Styler
	H3 Styler
	H4 Styler
	H5 Styler
	H6 Styler
}

// NewPdfRenderer creates and configures an PdfRenderer object,
// which satisfies the Renderer interface.
func NewPdfRenderer() *PdfRenderer {
	pdfr := new(PdfRenderer)

	// Global things
	pdfr.Orientation = "portrait"
	pdfr.units = "pt"
	pdfr.Papersize = "A4"
	pdfr.fontdir = "."

	// Normal Text
	pdfr.Normal = Styler{Font: "Arial", Style: "", Size: 12, Spacing: 5}

	// Backticked text
	pdfr.Backtick = Styler{Font: "Courier", Style: "", Size: 12, Spacing: 5}

	// Headings
	pdfr.H1 = Styler{Font: "Arial", Style: "b", Size: 24, Spacing: 12}
	pdfr.H2 = Styler{Font: "Arial", Style: "b", Size: 22, Spacing: 11}
	pdfr.H3 = Styler{Font: "Arial", Style: "b", Size: 20, Spacing: 10}
	pdfr.H4 = Styler{Font: "Arial", Style: "b", Size: 18, Spacing: 9}
	pdfr.H5 = Styler{Font: "Arial", Style: "b", Size: 16, Spacing: 8}
	pdfr.H6 = Styler{Font: "Arial", Style: "b", Size: 14, Spacing: 6}

	pdfr.pdf = gofpdf.New(pdfr.Orientation, pdfr.units,
		pdfr.Papersize, pdfr.fontdir)
	pdfr.pdf.AddPage()
	return pdfr
}

func (r *PdfRenderer) setFont(s Styler) {
	r.pdf.SetFont(s.Font, s.Style, s.Size)
}

func (r *PdfRenderer) write(s Styler, t string) {
	r.pdf.Write(s.Size+s.Spacing, t)
}

// RenderNode is a default renderer of a single node of a syntax tree. For
// block nodes it will be called twice: first time with entering=true, second
// time with entering=false, so that it could know when it's working on an open
// tag and when on close. It writes the result to w.
//
// The return value is a way to tell the calling walker to adjust its walk
// pattern: e.g. it can terminate the traversal by returning Terminate. Or it
// can ask the walker to skip a subtree of this node by returning SkipChildren.
// The typical behavior is to return GoToNext, which asks for the usual
// traversal to the next node.
// (above taken verbatim from the blackfriday v2 package)
func (r *PdfRenderer) RenderNode(w io.Writer, node *bf.Node, entering bool) bf.WalkStatus {
	switch node.Type {
	case bf.Text:
		s := strings.Replace(string(node.Literal), "\n", " ", -1)
		dbg("Text", s)
		r.setFont(r.current)
		r.write(r.current, s)
	case bf.Softbreak:
		dbg("Softbreak", "Not handled")
	case bf.Hardbreak:
		dbg("Hardbreak", "Not handled")
	case bf.Emph:
		if entering {
			dbg("Emph (entering)", "Processing")
			r.current.Style += "i"
		} else {
			dbg("Emph (leaving)", "Processing")
			r.current.Style = strings.Replace(r.current.Style, "i", "", -1)
		}
	case bf.Strong:
		if entering {
			dbg("Strong (entering)", "Processing")
			r.current.Style += "b"
		} else {
			dbg("Strong (leaving)", "Processing")
			r.current.Style = strings.Replace(r.current.Style, "b", "", -1)
		}
	case bf.Del:
		if entering {
			dbg("DEL (entering)", "Not handled")
		} else {
			dbg("DEL (leaving)", "Not handled")
		}
	case bf.HTMLSpan:
		dbg("HTMLSpan", "Not handled")
	case bf.Link:
		// mark it but don't link it if it is not a safe link: no smartypants
		//dest := node.LinkData.Destination
		if entering {
			dbg("Link (entering)", "Not handled")
		} else {
			dbg("Link (leaving)", "Not handled")
		}
	case bf.Image:
		if entering {
			dbg("Image (entering)", "Not handled")
		} else {
			dbg("Image (leaving)", "Not handled")
		}
	case bf.Code:
		dbg("Code", "Processing")
		r.setFont(r.Backtick)
		r.write(r.Backtick, string(node.Literal))
	case bf.Document:
		dbg("Document", "Processing")
		//break
	case bf.Paragraph:
		if entering {
			dbg("Paragraph (entering)", "Processing")
			r.current = r.Normal
			r.cr()
		} else {
			dbg("Paragraph (leaving)", "Processing")
			r.pdf.Ln(r.current.Size)
		}
	case bf.BlockQuote:
		if entering {
			dbg("BlockQuote (entering)", "Not handled")
		} else {
			dbg("BlockQuote (leaving)", "Not handled")
		}
	case bf.HTMLBlock:
		dbg("HTMLBlock", "Not handled")
	case bf.Heading:
		if entering {
			r.cr()
			switch node.HeadingData.Level {
			case 1:
				dbg("Heading (1, entering)", "Processing")
				r.current = r.H1
			case 2:
				dbg("Heading (2, entering)", "Processing")
				r.current = r.H2
			case 3:
				dbg("Heading (3, entering)", "Processing")
				r.current = r.H3
			case 4:
				dbg("Heading (4, entering)", "Processing")
				r.current = r.H4
			case 5:
				dbg("Heading (5, entering)", "Processing")
				r.current = r.H5
			case 6:
				dbg("Heading (6, entering)", "Processing")
				r.current = r.H6
			}
		} else {
			dbg("Heading (leaving)", "Processing")
			r.current = r.Normal
			r.cr()
		}
	case bf.HorizontalRule:
		dbg("HorizontalRule", "Not handled")
	case bf.List:
		/*
			openTag := ulTag
			closeTag := ulCloseTag
			if node.ListFlags&ListTypeOrdered != 0 {
				openTag = olTag
				closeTag = olCloseTag
			}
			if node.ListFlags&ListTypeDefinition != 0 {
				openTag = dlTag
				closeTag = dlCloseTag
			}
		*/
		if entering {
			dbg("List (entering)", "Not handled")
		} else {
			dbg("List (leaving)", "Not handled")
		}
	case bf.Item:
		/*
			openTag := liTag
			closeTag := liCloseTag
			if node.ListFlags&ListTypeDefinition != 0 {
				openTag = ddTag
				closeTag = ddCloseTag
			}
			if node.ListFlags&ListTypeTerm != 0 {
				openTag = dtTag
				closeTag = dtCloseTag
			}
		*/
		if entering {
			dbg("Item (entering)", "Not handled")
		} else {
			dbg("Item (leaving)", "Not handled")
		}
	case bf.CodeBlock:
		dbg("Codeblock", "Processing")
		r.cr()
		r.pdf.SetFillColor(200, 220, 255)
		r.setFont(r.Backtick)
		lines := strings.Split(string(node.Literal), "\n")
		for n := range lines {
			r.pdf.CellFormat(0, r.Backtick.Size,
				lines[n], "", 1, "LT", true, 0, "")
		}

	case bf.Table:
		if entering {
			dbg("Table (entering)", "Not handled")
		} else {
			dbg("Table (leaving)", "Not handled")
		}
	case bf.TableCell:
		/*
			openTag := tdTag
			closeTag := tdCloseTag
			if node.IsHeader {
				openTag = thTag
				closeTag = thCloseTag
			}
		*/
		if entering {
			dbg("TableCell (entering)", "Not handled")
		} else {
			dbg("TableCell (leaving)", "Not handled")
		}
	case bf.TableHead:
		if entering {
			dbg("TableHead (entering)", "Not handled")
		} else {
			dbg("TableHead (leaving)", "Not handled")
		}
	case bf.TableBody:
		if entering {
			dbg("TableBody (entering)", "Not handled")
		} else {
			dbg("TableBody (leaving)", "Not handled")
		}
	case bf.TableRow:
		if entering {
			dbg("TableRow (entering)", "Not handled")
		} else {
			dbg("TableRow (leaving)", "Not handled")
		}
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return bf.GoToNext
}

// RenderHeader writes HTML document preamble and TOC if requested.
func (r *PdfRenderer) RenderHeader(w io.Writer, ast *bf.Node) {
	dbg("RenderHeader", "Not handled")
}

// RenderFooter writes HTML document footer.
func (r *PdfRenderer) RenderFooter(w io.Writer, ast *bf.Node) {
	dbg("RenderFooter", "Not handled")
}

func (r *PdfRenderer) cr() {
	r.pdf.Ln(r.current.Size + r.current.Spacing)
}

// Helper functions
func dbg(source, msg string) {
	if *debug {
		fmt.Printf("[%v] %v\n", source, msg)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: mdToPdf [options]\n")
	flag.PrintDefaults()
	os.Exit(0)
}
