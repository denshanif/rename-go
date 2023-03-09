package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// get the csv file name from the command line
	csvFileName := os.Args[1]
	// get the directory name from the command line
	dirName := os.Args[2]

	// open the csv file
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// create a new csv reader
	csvReader := csv.NewReader(csvFile)

	// create a new slice to hold the new names
	var newNames []string

	// read the csv file
	for {
		// read a line from the csv file
		csvLine, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// append the new name to the new names slice
		newNames = append(newNames, csvLine[0])
	}

	// get the files in the directory
	files, err := filepath.Glob(dirName + "/*.pdf")
	if err != nil {
		log.Fatal(err)
	}

	// rename the files
	for i, file := range files {
		// get the file name
		fileName := filepath.Base(file)

		// get the file extension
		fileExt := filepath.Ext(file)

		// get the file name without the extension
		fileNameNoExt := strings.TrimSuffix(fileName, fileExt)

		// get the new name
		newName := newNames[i]

		// rename the file
		err := os.Rename(dirName+"/"+fileName, dirName+"/"+newName+".pdf")
		if err != nil {
			log.Fatal(err)
		}

		// print the old and new names
		fmt.Println(fileNameNoExt + ".pdf" + " renamed to " + newName + ".pdf")
	}
}