// Package pdf produces a single PDF from a list of websites. If the url has an ! in it, the page header is set
// to the text after the !
package pdf

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func ToPDF(pages []string, PDFname string) error {
	// Create new PDF generator
	pdfg, err := wk.NewPDFGenerator()
	if err != nil {
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wk.OrientationPortrait)

	for ind := 0; ind < len(pages); ind++ {
		url, header := pages[ind], ""

		// if there's a ! in the url, that means a header has been specified
		if strings.Contains(url, "!") {
			splt := strings.SplitN(url, "!", 2)
			url, header = splt[0], splt[1]
		}

		page := wk.NewPage(url)

		page.FooterRight.Set("[page]")
		page.FooterFontSize.Set(10)
		page.Zoom.Set(0.95)
		if header != "" {
			page.HeaderCenter.Set(header)
		}

		pdfg.AddPage(page)
	}

	// Create PDF
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	return pdfg.WriteFile(PDFname)
}

// GetWorkOrder reads the urls from fileName
func GetWorkOrder(fileName string) (workOrder []string, err error) {
	var handle *os.File
	if handle, err = os.Open(fileName); err != nil {
		return nil, err
	}
	defer func() { _ = handle.Close() }()

	rdr := bufio.NewReader(handle)
	for {
		var line string
		if line, err = rdr.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		workOrder = append(workOrder, strings.ReplaceAll(line, "\n", ""))
	}

	return workOrder, nil
}
