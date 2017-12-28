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

func main() {
	input := flag.String("i", "", "Input text filename; default is os.Stdin")
	output := flag.String("o", "", "Output PDF filename; required")
	help := flag.Bool("help", false, "Show usage message")

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
	pf.H1.Spacing = 24
	pf.H2.Spacing = 22
	pf.H3.Spacing = 20
	pf.H4.Spacing = 5
	pf.H5.Spacing = 3
	pf.H6.Spacing = 0

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
	pdfr.Normal = Styler{Font: "Arial", Style: "", Size: 12, Spacing: 3}

	// Headings
	pdfr.H1 = Styler{Font: "Arial", Style: "b", Size: 24, Spacing: 5}
	pdfr.H2 = Styler{Font: "Arial", Style: "b", Size: 22, Spacing: 5}
	pdfr.H3 = Styler{Font: "Arial", Style: "b", Size: 20, Spacing: 5}
	pdfr.H4 = Styler{Font: "Arial", Style: "b", Size: 18, Spacing: 5}
	pdfr.H5 = Styler{Font: "Arial", Style: "b", Size: 16, Spacing: 5}
	pdfr.H6 = Styler{Font: "Arial", Style: "b", Size: 14, Spacing: 5}

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
		r.setFont(r.current)
		r.write(r.current, string(node.Literal))
	case bf.Softbreak:
		dbg("Not handled: bf.Softbreak")
	case bf.Hardbreak:
		dbg("Not handled: bf.Hardbreak")
	case bf.Emph:
		if entering {
			r.current.Style += "i"
		} else {
			r.current.Style = strings.Replace(r.current.Style, "i", "", -1)
		}
	case bf.Strong:
		if entering {
			r.current.Style += "b"
		} else {
			r.current.Style = strings.Replace(r.current.Style, "b", "", -1)
		}
	case bf.Del:
		if entering {
			dbg("Not handled: bf.Del (entering)")
		} else {
			dbg("Not handled: bf.Del (!entering)")
		}
	case bf.HTMLSpan:
		dbg("Not handled: bf.HTMLSpan")
	case bf.Link:
		// mark it but don't link it if it is not a safe link: no smartypants
		//dest := node.LinkData.Destination
		if entering {
			dbg("Not handled: bf.Link (entering)")
		} else {
			dbg("Not handled: bf.Link (!entering)")
		}
	case bf.Image:
		if entering {
			dbg("Not handled: bf.Image (entering)")
		} else {
			dbg("Not handled: bf.Image (!entering)")
		}
	case bf.Code:
		dbg("Not handled: bf.Code")
		//r.out(w, codeTag)
		//escapeHTML(w, node.Literal)
		//r.out(w, codeCloseTag)
	case bf.Document:
		break
	case bf.Paragraph:
		if entering {
			r.setFont(r.Normal)
		} else {
			r.cr()
		}
	case bf.BlockQuote:
		if entering {
			dbg("Not handled: bf.BlockQuote (entering)")
		} else {
			dbg("Not handled: bf.BlockQuote (!entering)")
		}
	case bf.HTMLBlock:
		dbg("Not handled: bf.HTMLBlock")
	case bf.Heading:
		//openTag, closeTag := headingTagsFromLevel(node.Level)
		//dbg(fmt.Sprintf("Heading level is %v", node.HeadingData.Level))
		if entering {
			switch node.HeadingData.Level {
			case 1:
				r.current = r.H1
			case 2:
				r.current = r.H2
			case 3:
				r.current = r.H3
			case 4:
				r.current = r.H4
			case 5:
				r.current = r.H5
			case 6:
				r.current = r.H6
			}
		} else {
			r.cr()
			r.current = r.Normal
		}
	case bf.HorizontalRule:
		//r.cr(w)
		//r.outHRTag(w)
		//r.cr(w)
		dbg("Not handled: bf.HorizontalRule")
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
			dbg("Not handled: bf.List (entering)")
		} else {
			dbg("Not handled: bf.List (!entering)")
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
			dbg("Not handled: bf.Item (entering)")
		} else {
			dbg("Not handled: bf.Item (!entering)")
		}
	case bf.CodeBlock:
		dbg("Not handled: bf.CodeBlock")
	case bf.Table:
		if entering {
			dbg("Not handled: bf.Table (entering)")
		} else {
			dbg("Not handled: bf.Table (!entering)")
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
			dbg("Not handled: bf.TableCell (entering)")
		} else {
			dbg("Not handled: bf.TableCell (!entering)")
		}
	case bf.TableHead:
		if entering {
			dbg("Not handled: bf.TableHead (entering)")
		} else {
			dbg("Not handled: bf.TableHead (!entering)")
		}
	case bf.TableBody:
		if entering {
			dbg("Not handled: bf.TableBody (entering)")
		} else {
			dbg("Not handled: bf.TableBody (!entering)")
		}
	case bf.TableRow:
		if entering {
			dbg("Not handled: bf.TableRow (entering)")
		} else {
			dbg("Not handled: bf.TableRow (!entering)")
		}
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return bf.GoToNext
}

// RenderHeader writes HTML document preamble and TOC if requested.
func (r *PdfRenderer) RenderHeader(w io.Writer, ast *bf.Node) {
	dbg("Not handled: RenderHeader")
}

// RenderFooter writes HTML document footer.
func (r *PdfRenderer) RenderFooter(w io.Writer, ast *bf.Node) {
	dbg("Not handled: RenderFooter")
}

func (r *PdfRenderer) cr() {
	r.write(r.current, "\n")
}

// Helper functions
func dbg(msg string) {
	fmt.Println(msg)
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: mdToPdf [options]\n")
	flag.PrintDefaults()
	os.Exit(0)
}
