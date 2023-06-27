package main

import (
	"flag"

	"github.com/invertedv/pdf"
)

func main() {
	var (
		workOrder []string
		err       error
	)

	workOrderFile := flag.String("list", "", "string")
	outPDF := flag.String("out", "", "string")

	if workOrder, err = pdf.GetWorkOrder(*workOrderFile); err != nil {
		panic(err)
	}

	if e := pdf.ToPDF(workOrder, *outPDF); e != nil {
		panic(e)
	}
}
