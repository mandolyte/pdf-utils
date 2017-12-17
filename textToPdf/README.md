# HelloWorld

This simple examples demonstrates how to create a minimal PDF 
file using an input text file.

To show help:
```
$ go run textToPdf.go -help
Help Message

Usage: textToPdf [options]
  -font string
    	Font 'Arial', 'Courier', or 'Times'; default Arial (default "Arial")
  -fs float
    	Font size; default 12 (default 12)
  -help
    	Show usage message
  -i string
    	Input text filename; default 'README.md' (default "README.md")
  -o string
    	Output PDF filename; default 'README.pdf' (default "README.pdf")
  -ps string
    	Paper size; default 'Letter' (default "Letter")
  -style string
    	Style: B for bold, U for underline, I for Italics or any combo thereof
$ 
```

Example: create a PDF of the source using the Courier font:
```
$ go run textToPdf.go -i textToPdf.go -font Courier -o $HOME/data/textToPdf.pdf
$ 
```