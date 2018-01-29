# TextToList

This simple example demonstrates how to create a numbered list of paragraphs provided from an input text file..

To show help:
```
$ go run textToList.go -help
Help Message

Usage: textToList [options]
Note: font and line height are 'points', 1/72 inch
This example takes a text file of paragraphs and
demonstrates how to create a numbered list in PDF format
Blank lines are skipped!
  -compressed
        Don't add blank line between items
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
$
```

Example using 10 point italicized Times font, with compressed option
```
go run textToList.go -i test.txt -o test.pdf -lh 12 -fs 10 -font Times -style I -compressed 
```