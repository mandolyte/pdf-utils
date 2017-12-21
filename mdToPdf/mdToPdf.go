package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

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
	_ = bf.Run(content, bf.WithRenderer(pf))

	err = pf.pdf.OutputFileAndClose(*output)
	if err != nil {
		log.Fatalf("pdf.OutputFileAndClose() error:%v", err)
	}

	//	pdf := gofpdf.New(gofpdf.OrientationPortrait, "pt", *papersize, ".")

}

// PdfRenderer is the struct to manage conversion of a markdown object
// to PDF format.
type PdfRenderer struct {
	pdf                                    *gofpdf.Fpdf
	orientation, units, papersize, fontdir string
	font, style                            string
	fontsize, lineheight                   float64
	currentsize                            float64
	// headings
	H1Font       string
	H1Style      string
	H1Size       float64
	H1Lineheight float64
}

// NewPdfRenderer creates and configures an PdfRenderer object,
// which satisfies the Renderer interface.
func NewPdfRenderer() *PdfRenderer {
	pdfr := new(PdfRenderer)
	pdfr.orientation = "portrait"
	pdfr.units = "pt"
	pdfr.papersize = "A4"
	pdfr.fontdir = "."
	pdfr.font = "Arial"
	pdfr.fontsize = 12
	pdfr.lineheight = pdfr.fontsize + 2
	pdfr.style = ""
	pdfr.H1Font = "times"
	pdfr.H1Size = 20
	pdfr.H1Lineheight = pdfr.H1Size + 2
	pdfr.H1Style = "b"
	pdfr.pdf = gofpdf.New(pdfr.orientation, pdfr.units, pdfr.papersize, pdfr.fontdir)
	pdfr.pdf.AddPage()
	return pdfr
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
		r.pdf.Write(r.currentsize+2, string(node.Literal))
	case bf.Softbreak:
		panic("Not handled: bf.Softbreak")
	case bf.Hardbreak:
		panic("Not handled: bf.Hardbreak")
	case bf.Emph:
		if entering {
			r.pdf.SetFont(r.font, "i", r.fontsize)
		} else {
			r.pdf.SetFont(r.font, r.style, r.fontsize)
		}
	case bf.Strong:
		if entering {
			panic("Not handled: bf.Strong (entering)")
		} else {
			panic("Not handled: bf.Strong (!entering)")
		}
	case bf.Del:
		if entering {
			panic("Not handled: bf.Del (entering)")
		} else {
			panic("Not handled: bf.Del (!entering)")
		}
	case bf.HTMLSpan:
		panic("Not handled: bf.HTMLSpan")
	case bf.Link:
		// mark it but don't link it if it is not a safe link: no smartypants
		//dest := node.LinkData.Destination
		if entering {
			panic("Not handled: bf.Link (entering)")
		} else {
			panic("Not handled: bf.Link (!entering)")
		}
	case bf.Image:
		if entering {
			panic("Not handled: bf.Image (entering)")
		} else {
			panic("Not handled: bf.Image (!entering)")
		}
	case bf.Code:
		panic("Not handled: bf.Code")
		//r.out(w, codeTag)
		//escapeHTML(w, node.Literal)
		//r.out(w, codeCloseTag)
	case bf.Document:
		break
	case bf.Paragraph:
		if entering {
			r.pdf.SetFont(r.font, r.style, r.fontsize)
			r.currentsize = r.fontsize
		} else {
			r.cr()
			r.pdf.SetFont(r.font, r.style, r.fontsize)
			r.currentsize = r.fontsize
		}
	case bf.BlockQuote:
		if entering {
			panic("Not handled: bf.BlockQuote (entering)")
		} else {
			panic("Not handled: bf.BlockQuote (!entering)")
		}
	case bf.HTMLBlock:
		panic("Not handled: bf.HTMLBlock")
	case bf.Heading:
		//openTag, closeTag := headingTagsFromLevel(node.Level)
		if entering {
			r.pdf.SetFont(r.H1Font, r.H1Style, r.H1Size)
			r.currentsize = r.H1Size
		} else {
			r.cr()
			r.pdf.SetFont(r.font, r.style, r.fontsize)
			r.currentsize = r.fontsize
		}
	case bf.HorizontalRule:
		//r.cr(w)
		//r.outHRTag(w)
		//r.cr(w)
		panic("Not handled: bf.HorizontalRule")
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
			panic("Not handled: bf.List (entering)")
		} else {
			panic("Not handled: bf.List (!entering)")
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
			panic("Not handled: bf.Item (entering)")
		} else {
			panic("Not handled: bf.Item (!entering)")
		}
	case bf.CodeBlock:
		panic("Not handled: bf.CodeBlock")
	case bf.Table:
		if entering {
			panic("Not handled: bf.Table (entering)")
		} else {
			panic("Not handled: bf.Table (!entering)")
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
			panic("Not handled: bf.TableCell (entering)")
		} else {
			panic("Not handled: bf.TableCell (!entering)")
		}
	case bf.TableHead:
		if entering {
			panic("Not handled: bf.TableHead (entering)")
		} else {
			panic("Not handled: bf.TableHead (!entering)")
		}
	case bf.TableBody:
		if entering {
			panic("Not handled: bf.TableBody (entering)")
		} else {
			panic("Not handled: bf.TableBody (!entering)")
		}
	case bf.TableRow:
		if entering {
			panic("Not handled: bf.TableRow (entering)")
		} else {
			panic("Not handled: bf.TableRow (!entering)")
		}
	default:
		panic("Unknown node type " + node.Type.String())
	}
	return bf.GoToNext
}

// RenderHeader writes HTML document preamble and TOC if requested.
func (r *PdfRenderer) RenderHeader(w io.Writer, ast *bf.Node) {
	dumpNode(ast)
	//panic("Not handled: RenderHeader")
}

// RenderFooter writes HTML document footer.
func (r *PdfRenderer) RenderFooter(w io.Writer, ast *bf.Node) {
	dumpNode(ast)
	//panic("Not handled: RenderFooter")
}

func (r *PdfRenderer) cr() {
	r.pdf.Write(r.currentsize+2, "\n")
}

// Helper functions
func dumpNode(n *bf.Node) {
	fmt.Printf("Node Dump:\n%v\n", n)
}



func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: mdToPdf [options]\n")
	flag.PrintDefaults()
	os.Exit(0)
}
