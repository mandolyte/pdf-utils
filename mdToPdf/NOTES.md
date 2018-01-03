# Notes on Markdown to PDF 

*Limitations*
At present nested (self) elements are handled, such as a blockquote 
within a blockquote. This will require a push/pop stack or some
such technique to handle the recursive nature of self-nested elements.

*On Debug/Tracing* 
The trace is output to `debug.txt`.
This has proved invaluable in understanding what the BlackFriday
parser is reporting.

*On Lists*
- The BlackFriday parser will treat consequtive lists separated by blank
lines as a single list of the type of the first one in lexical order. 
- A paragraph preceding a list must be followed by a blank line; otherwise
the list will be consumed as part of the paragraph.