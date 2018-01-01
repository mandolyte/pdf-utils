# Notes on Markdown to PDF 

*On Debug/Tracing* 
The trace is output to STDOUT. See `debug.txt` as an example.
This has proved invaluable in understanding what the BlackFriday
parser is reporting.

*On Lists*
- The BlackFriday parser will treat consequtive lists separated by blank
lines as a single list of the type of the first one in lexical order. 
- A paragraph preceding a list must be followed by a blank line; otherwise
the list will be consumed as part of the paragraph.