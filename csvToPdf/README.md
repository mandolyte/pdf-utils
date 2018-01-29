# csvToPdf

This example demonstrates how to create a minimal PDF 
file using an input CSV file. The CSV is done as a table with
bold headers and alternating colors for the rows.

The code will fit the CSV file to the page size by starting at the
given font size and adjusting down until the minimum font size is 
reached. To do this does the following:

1. Given a paper size and an orientation (landscape or portrait) with a starting font size and a minimum font size, then ...
2. Based on current font: 
    - Find the maximum cell width in each column and sum these up as a "worst case" width
    - Find the available page width by subtracting the left and right margins from the paper size width
3. Then subtract and see if there is room to fit the table
4. If there is room, then create the PDF; otherwise, decrement the font size and try to fit again
5. If it reaches the minimum font size and still can't fit the table to the page, it shows an error.

To show help:
```
$ go run csvToPdf.go -help
Help...
  -font string
        Font 'Arial', 'Courier', or 'Times'; default Arial (default "Arial")
  -fs int
        Font size; default 12 (default 12)
  -help
        Show help message
  -i string
        CSV file name to sort; default STDIN
  -minfs int
        Mininum font size; default 8 (default 8)
  -o string
        PDF output file name; required
  -orient string
        Either 'portrait' or 'landscape'; default is 'portrait' (default "portrait")
  -ps string
        Paper size; default 'Letter' (default "Letter")
$
```

$ go run csvToPdf.go -o cb.pdf -orient landscape -i cincy_breweries.csv -ps a4 -minfs 6
Attempt fit using font size: 12 ... too big; available 785.19, need 1524.43
Attempt fit using font size: 11 ... too big; available 785.19, need 1397.40
Attempt fit using font size: 10 ... too big; available 785.19, need 1270.36
Attempt fit using font size: 9 ... too big; available 785.19, need 1143.32
Attempt fit using font size: 8 ... too big; available 785.19, need 1016.29
Attempt fit using font size: 7 ... too big; available 785.19, need 889.25
Attempt fit using font size: 6 ... fits!
$

$ go run csvToPdf.go -o cb.pdf -orient landscape -i cincy_breweries.csv  -ps a4 -minfs 10
Attempt fit using font size: 12 ... too big; available 785.19, need 1524.43
Attempt fit using font size: 11 ... too big; available 785.19, need 1397.40
Attempt fit using font size: 10 ... too big; available 785.19, need 1270.36
2018/01/19 12:33:53 Unable to fit CSV onto page using provided constraints
exit status 1
$