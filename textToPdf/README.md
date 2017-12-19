# TextToPdf

This simple example demonstrates how to create a minimal PDF 
file using an input text file.

To show help:
```
$ go run textToPdf.go -help
Help Message

Usage: textToPdf [options]
Note: font and line height are 'points', 1/72 inch
  -font string
    	Font 'Arial', 'Courier', or 'Times'; default Arial (default "Arial")
  -fs float
    	Font size; default 12 (default 12)
  -help
    	Show usage message
  -i string
    	Input text filename; default is os.Stdin
  -lh float
    	Line height; default 15 (default 15)
  -o string
    	Output PDF filename; required
  -ps string
    	Paper size; default 'Letter' (default "Letter")
  -style string
    	Style: B for bold, U for underline, I for Italics or any combo thereof
  -tabwidth int
    	Number of spaces for a tab character; default is 4 (default 4)
$ 
```

## Example 1: create a PDF of the source
Using the Courier font:
```
go run textToPdf.go -i textToPdf.go -font Courier -o $HOME/data/textToPdf.pdf

```

## Example 2: cram as much as possible onto the page
Using 8 point font and 8 point line height:
```
go run textToPdf.go -lh 8 -fs 8 -i textToPdf.go -font Courier -o $HOME/data/textToPdf.pdf
```