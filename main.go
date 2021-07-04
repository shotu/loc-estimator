package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// type LOCer interface{
// 	GetLOC
// }

type JSLOC struct {
	filepath           string
	estimateLOCInJS    int64
	CommentsCount      int64
	BlanksCount        int64
	CodeCount          int64
	BlockComments      int64
	SingleLineComments int64
}

func (jsloc *JSLOC) EstimateLOCInJS() {
	filePath := strings.TrimSpace(jsloc.filepath)

	fmt.Println("File path is", filePath)

	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	isBlockCommentRunning := false

	for fileScanner.Scan() {

		fileLineText := fileScanner.Text()
		trimmedLine := strings.TrimSpace(fileLineText)
		// trimmedLine := fileLineText
		// check if blank line
		lenOfTrimmedLine := len(trimmedLine)
		fmt.Println("trimmedLine", trimmedLine)
		fmt.Println("lenOfTrimmedLine", lenOfTrimmedLine)
		if lenOfTrimmedLine == 0 { // isBlockCommentStarted != true // ALSO calculates the blank in block comment
			fmt.Println("Blank comment found")
			jsloc.BlanksCount++
			continue
		}

		// check for single line comments // some comment
		if !isBlockCommentRunning && fmt.Sprintf("%c", trimmedLine[0]) == "/" && fmt.Sprintf("%c", trimmedLine[1]) == "/" {
			jsloc.CommentsCount++
		}

		// if block comment running continue after seeing blank line
		if lenOfTrimmedLine == 0 && isBlockCommentRunning { // isBlockCommentStarted != true // ALSO calculates the blank in block comment
			continue
		}

		if isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-2]) != "*" && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-1]) != "/" {
			jsloc.BlockComments++
			continue
		}
		// finding first block comment line
		if !isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[0]) == "/" && fmt.Sprintf("%c", trimmedLine[1]) == "*" {
			fmt.Println("Block comment starts")
			jsloc.BlockComments++
			isBlockCommentRunning = true
			continue
		}

		// finding last block comment line
		if isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-2]) == "*" && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-1]) == "/" {
			jsloc.BlockComments++
			fmt.Println("Block comment ends")
			isBlockCommentRunning = false
			continue
		}

		jsloc.CodeCount++
		fmt.Println("fileLineText:", trimmedLine)
		fmt.Println("count of char", len(trimmedLine))
	}

	fmt.Println("print the line of LOC", jsloc.CodeCount)
	readFile.Close()
	fmt.Println("jsloc", jsloc)
}

func GetFileType(filePath string) string {
	return "JS"
}

func main() {

	reader := bufio.NewReader(os.Stdin) // reading file name, from userinput, make sure file is present in current working directory
	fmt.Print("Enter filepath:")
	filepath, _ := reader.ReadString('\n')
	fmt.Println(filepath)
	fileType := GetFileType(filepath)

	// case file type:
	if fileType == "JS" {
		jsloc := &JSLOC{
			filepath: strings.TrimSpace(filepath),
		}
		jsloc.EstimateLOCInJS()
		fmt.Println("Blank comment count          :", jsloc.BlockComments+jsloc.BlanksCount)
		fmt.Println("CommentsCount comment count  :", jsloc.CommentsCount)
		fmt.Println("Code comment count           :", jsloc.CodeCount)
		fmt.Println("Total count                  :", jsloc.BlockComments+jsloc.BlanksCount+jsloc.CodeCount+jsloc.CommentsCount)
	}
}
