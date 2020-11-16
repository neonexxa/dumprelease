package main

import (
	"log"
	"os"
	"path/filepath"
	pdf "github.com/adrg/go-wkhtmltopdf"
)

func main() {
	var argsGen_pathFile string
	if len(os.Args) == 1 {
		log.Fatal("No path given, Please specify path.")
		return
	}
	if argsGen_pathFile = os.Args[1]; argsGen_pathFile == "" {
		log.Fatal("No path given, Please specify path.")
		return
	}
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

	// Content Generation
	middlefiles, err := FilePathWalkDir(argsGen_pathFile + "/middle/")
	if err != nil {
		panic(err)
	}
	for _, midfile := range middlefiles {
		objectContentPage, err := pdf.NewObject(midfile)
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

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

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
