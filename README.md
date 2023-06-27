## pdf

Essentially, this a wrapper to github.com/SebastiaanKlippert/go-wkhtmltopdf which is a Go
wrapper to wkhtmltopdf.  The input is a list of url's, which are combined into a single pdf.

The one wrinkle is that the header can be specified for each url. The header text is preceded by an !
at the end of the url.

"https://google.com!Hi I am Google"

This could all be done in a shell script, but this is a bit cleaner.