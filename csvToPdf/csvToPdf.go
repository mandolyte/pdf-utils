package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	input := flag.String("i", "", "CSV file name to sort; default STDIN")
	output := flag.String("o", "", "PDF output file name; required")
	papersize := flag.String("ps", "Letter", "Paper size; default 'Letter'")
	font := flag.String("font", "Arial", "Font 'Arial', 'Courier', or 'Times'; default Arial")
	fontsize := flag.Int("fs", 12, "Font size; default 12")
	minfontsize := flag.Int("minfs", 8, "Mininum font size; default 8")
	orient := flag.String("orient", "portrait", "Either 'portrait' or 'landscape'; default is 'portrait'")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		usage("Help...")
	}

	orientation := gofpdf.OrientationPortrait
	if *orient == "landscape" {
		orientation = gofpdf.OrientationLandscape
	}

	// output PDF filename is required
	if *output == "" {
		usage("Output PDF filename is required")
	}

	csvall := slurpInputCSV(*input)

	pdf := gofpdf.New(orientation, "pt", *papersize, ".")
	pdf.AddPage()

	// make a table of floats to store the string width
	// in current font for each cell
	cells := make([][]float64, len(csvall))
	for n := range cells {
		cells[n] = make([]float64, len(csvall[n]))
	}

	var cols []float64

	for ifs := *fontsize; ifs >= *minfontsize; ifs-- {
		fmt.Printf("Attempt fit using font size: %v", ifs)
		pdf.SetFont(*font, "", float64(ifs))
		em := pdf.GetStringWidth("m")

		// compute the cell widths needed
		for n := range csvall {
			for m, v := range csvall[n] {
				// get string width plus 2 em for padding grid
				cells[n][m] = pdf.GetStringWidth(v) + (2 * em)
			}
		}

		// get max width of each column
		cols = make([]float64, len(cells[0]))
		for j := 0; j < len(cells[0]); j++ {
			for i := 0; i < len(cells); i++ {
				if cells[i][j] > cols[j] {
					cols[j] = cells[i][j]
				}
			}
		}
		wSum := 0.0
		for i := 0; i < len(cols); i++ {
			wSum += cols[i]
		}

		// get the page size
		pagewidth, _ := pdf.GetPageSize()
		// get the margins
		lm, _, rm, _ := pdf.GetMargins()

		// compute the amount of page available
		availablepage := pagewidth - (lm + rm)

		if availablepage > wSum {
			fmt.Println(" ... fits!")
			break
		} else {
			fmt.Printf(" ... too big; available %.2f, need %.2f\n",
				availablepage, wSum)
		}

	}

	// Good to go... make the PDF now.
	makepdf(*output, pdf, csvall, cols)
}

func makepdf(filename string, pdf *gofpdf.Fpdf, csvall [][]string, cols []float64) {
	// get the current fontsize in points
	ptsz, _ := pdf.GetFontSize()
	fill := false

	for row := 0; row < len(csvall); row++ {
		for col := 0; col < len(csvall[row]); col++ {
			// settings for header and body
			if row == 0 {
				pdf.SetFillColor(180, 180, 180)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetDrawColor(128, 0, 0)
				pdf.SetLineWidth(.3)
				pdf.SetFont("", "B", 0)

				pdf.CellFormat(cols[col], ptsz, csvall[row][col],
					"1", 0, "C", true, 0, "")
			} else {
				pdf.SetFillColor(240, 240, 240)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("", "", 0)

				pdf.CellFormat(cols[col], ptsz, csvall[row][col],
					"1", 0, "C", fill, 0, "")
			}
		}
		pdf.Ln(-1)
		fill = !fill
	}

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Fatalf("Pdf.OutputFileAndClose() error on %v:%v", filename, err)
	}

}

func usage(msg string) {
	fmt.Printf("%v\n", msg)
	flag.PrintDefaults()
	os.Exit(0)
}

func slurpInputCSV(filename string) [][]string {
	// open input file
	var r *csv.Reader
	if filename == "" {
		r = csv.NewReader(os.Stdin)
	} else {
		fi, fierr := os.Open(filename)
		if fierr != nil {
			log.Fatal("os.Open() Error:" + fierr.Error())
		}
		defer fi.Close()
		r = csv.NewReader(fi)
	}

	// read into memory
	csvall, raerr := r.ReadAll()
	if raerr != nil {
		log.Fatal("r.ReadAll() Error:" + raerr.Error())
	}
	return csvall
}
