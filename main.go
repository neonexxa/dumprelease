// package main

// import (
// 	// u "./pdfGenerator"
// 	"https://gitlab.com/ofo-tech/pdfGenerator/-/tree/master/pdfGenerator"
// 	"fmt"
// )

// func main() {

// 	r := u.NewRequestPdf("")

// 	//html template path
// 	templatePath := "templates/sample.html"

// 	//path for download pdf
// 	outputPath := "storage/example.pdf"

// 	//html template data
// 	templateData := struct {
// 		Title       string
// 		Description string
// 		Company     string
// 		Contact     string
// 		Country     string
// 	}{
// 		Title:       "HTML to PDF generator",
// 		Description: "This is the simple HTML to PDF file.",
// 		Company:     "Jhon Lewis",
// 		Contact:     "Maria Anders",
// 		Country:     "Germany",
// 	}

// 	if err := r.ParseTemplate(templatePath, templateData); err == nil {
// 		ok, _ := r.GeneratePDF(outputPath)
// 		fmt.Println(ok, "pdf generated successfully")
// 	} else {
// 		fmt.Println(err)
// 	}
// }

package main

import (
	"log"
	"os"
	"strconv"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

func main() {
	argsGen_pathFile := os.Args[1]
	argsGen_chunksize := os.Args[2]
	argsGen_chunksizeInt, err := strconv.Atoi(argsGen_chunksize)
	pdf.Init()
	defer pdf.Destroy()

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	// Create object from file.
	objectCoverPage, err := pdf.NewObject(argsGen_pathFile + "/front/0.html")
	if err != nil {
		log.Fatal(err)
	}
	objectCoverPage.LoadImages = true
	// Add created objects to the converter.
	converter.Add(objectCoverPage)

	objectCoverPage2, err := pdf.NewObject(argsGen_pathFile + "/front/1.html")
	if err != nil {
		log.Fatal(err)
	}
	converter.Add(objectCoverPage2)
	for i := 1; i < argsGen_chunksizeInt; i++ {
		objectContentPage, err := pdf.NewObject(argsGen_pathFile + "/middle/" + strconv.Itoa(i) + " - 0.html")
		if err != nil {
			log.Fatal(err)
		}
		converter.Add(objectContentPage)
	}

	// Create object from file.
	objectBackPage, err := pdf.NewObject(argsGen_pathFile + "/back/0.html")
	if err != nil {
		log.Fatal(err)
	}
	converter.Add(objectBackPage)
	// // Create object from URL.
	// object2, err := pdf.NewObject("https://google.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If Want To Display header & footer
	// objectCoverPage.Header.DisplaySeparator = true
	// objectCoverPage.Footer.ContentLeft = "[date]"
	// objectCoverPage.Footer.ContentCenter = "Sample footer information"
	// objectCoverPage.Footer.ContentRight = "[page]"
	// objectCoverPage.Footer.DisplaySeparator = true

	// Create object from reader.
	// inFile, err := os.Open("templates/cover.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer inFile.Close()

	// object3, err := pdf.NewObjectFromReader(inFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// object3.Zoom = 1.5
	// object3.TOC.Title = "Table of Contents"

	// converter.Add(objectCoverPage2)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Portrait
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create(argsGen_pathFile + "/merged.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
}
