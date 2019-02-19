package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inPlace bool
	var sourceFileName, outFileName string
	var currentPadding, padIncrement int
	var sourceFile, outFile *os.File
	var err error

	/* Define the available command line args */
	flag.StringVar(&sourceFileName, "f", "httpd.conf", "The config file to format.")
	flag.BoolVar(&inPlace, "i", false, "Edit file in place, backs up the orginal file with the .bak extention then overwrites the original.")
	flag.IntVar(&padIncrement, "p", 4, "The number of characters to pad each indentation level with.")

	/* Aaand parse 'em */
	flag.Parse()

	/* Special case, if they pass in one arg and it's not the -i flag assume it's file to process */
	if len(os.Args) == 2 && os.Args[1] != "-i" {
		sourceFileName = os.Args[1]
	}

	/* Make sure the source file even exists */
	_, err = os.Stat(sourceFileName)
	if os.IsNotExist(err) {
		fmt.Println("Error opening file " + sourceFileName + ".  File does not exist.")
		os.Exit(1)
	}

	/* If the inPlace flag is passed, back up the original file and use that as the source */
	if inPlace {
		/* Let's make sure the backup file doesn't already exist first */
		_, err = os.Stat(sourceFileName + ".bak")

		if !os.IsNotExist(err) {
			fmt.Println("Error, the backup file " + sourceFileName + ".bak already exists, please rename or delete it")
			os.Exit(1)
		}

		fmt.Print("Renaming " + sourceFileName + " to " + sourceFileName + ".bak...")
		os.Rename(sourceFileName, sourceFileName+".bak")
		fmt.Println("done.")

		outFileName = sourceFileName
		sourceFileName = sourceFileName + ".bak"
	}

	sourceFile, err = os.Open(sourceFileName)

	if err != nil {
		fmt.Println("An error occursed opening the file.")
		fmt.Println("Error was:", err.Error())
		os.Exit(1)
	}
	defer sourceFile.Close()

	if inPlace {
		outFile, err = os.Create(outFileName)
		if err != nil {
			fmt.Println("An error occured creating the new file.")
			fmt.Println("Error was:", err.Error())
			os.Exit(1)
		}
	}
	defer outFile.Close()

	inScanner := bufio.NewScanner(sourceFile)

	if inPlace {
		fmt.Print("Processing...")
	}
	for inScanner.Scan() {
		output := strings.TrimLeft(inScanner.Text(), " \t")

		/* Is this a closing tag? */
		if strings.HasPrefix(output, "</") {
			currentPadding -= padIncrement
		}

		for i := 0; i < currentPadding; i++ {
			if inPlace {
				outFile.WriteString(" ")
			} else {
				fmt.Print(" ")
			}
		}
		if inPlace {
			outFile.WriteString(output + "\n")
		} else {
			fmt.Println(output)
		}

		/* Does it start with "<"?  Is it a closing or opening tag? */
		if strings.HasPrefix(output, "<") {
			if !strings.HasPrefix(output, "</") {
				currentPadding += padIncrement
			}
		}
	}

	if inPlace {
		fmt.Println("done.")
	}
}
