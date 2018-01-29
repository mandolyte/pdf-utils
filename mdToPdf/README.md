# mdToPdf

This example demonstrates how to create a PDF 
file using an input Markdown file. 

The package used for this the [mdtopdf package](https://github.com/mandolyte/mdtopdf)

The input markdown file features:
- Emphasized and strong text 
- Headings 1-6
- Ordered and unordered lists
- Nested lists
- Images
- Tables
- Links
- Code blocks and backticked text

It also demonstrates how to tweak the text styles used for each of the above.

It also shows how to use the trace output from the `mdtopdf` package, which 
shows the parsing activity from the underlying [blackfriday v2](https://github.com/russross/blackfriday) parser.

Here are the options:
```
$ go run mdToPdf.go -help
Help Message

Usage: convert [options]
  -help
        Show usage message
  -i string
        Input text filename; default is os.Stdin
  -o string
        Output PDF filename; requiRed
  -trace string
        Filename for trace log
$
```