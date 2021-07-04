package loc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type JSLOC struct {
	Filepath              string
	CommentsCount         int
	BlanksCount           int
	CodeCount             int
	BlockComments         int
	SingleLineComments    int
	isBlockCommentRunning bool
}

// TODO we can implement one more function which will recursively look of js files and call the EstimateLOCInJS for each file

// Updates the blank counter
func (jsloc *JSLOC) UpdateBlankCounter(trimmedLine string) bool {
	lenOfTrimmedLine := len(trimmedLine)
	// blank comment updater
	if lenOfTrimmedLine == 0 { // isBlockCommentStarted != true // ALSO calculates the blank in block comment
		// fmt.Println("Blank comment found")
		jsloc.BlanksCount++
		// continue
		return true
	}
	return false
}

// Updates the single line comment counter
func (jsloc *JSLOC) UpdateSingleLinCommentCounter(trimmedLine string) bool {
	if !jsloc.isBlockCommentRunning && fmt.Sprintf("%c", trimmedLine[0]) == "/" && fmt.Sprintf("%c", trimmedLine[1]) == "/" {
		jsloc.CommentsCount++
	}
	return false
}

// Checks if block comment running
func (jsloc *JSLOC) CheckIfBlockCommentRunning(trimmedLine string) bool {
	lenOfTrimmedLine := len(trimmedLine)
	if lenOfTrimmedLine == 0 && jsloc.isBlockCommentRunning { // isBlockCommentStarted != true // ALSO calculates the blank in block comment
		return true
	}
	return false
}

// checks if block comment is ending in current line
func (jsloc *JSLOC) CheckIfBlockCommentEnds(trimmedLine string) bool {
	lenOfTrimmedLine := len(trimmedLine)
	if jsloc.isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-2]) != "*" && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-1]) != "/" {
		jsloc.BlockComments++
		return true
	}
	return false
}

// checks if first block comment line
func (jsloc *JSLOC) CheckIfFirstBlockCommentLine(trimmedLine string) bool {
	lenOfTrimmedLine := len(trimmedLine)

	// finding first block comment line
	if !jsloc.isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[0]) == "/" && fmt.Sprintf("%c", trimmedLine[1]) == "*" {
		// fmt.Println("Block comment starts")
		jsloc.BlockComments++
		jsloc.isBlockCommentRunning = true
		return true
	}
	return false
}

// checks if last block comment line
func (jsloc *JSLOC) CheckIfLastBlockCommentLine(trimmedLine string) bool {
	lenOfTrimmedLine := len(trimmedLine)
	// finding last block comment line
	if jsloc.isBlockCommentRunning && lenOfTrimmedLine >= 2 && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-2]) == "*" && fmt.Sprintf("%c", trimmedLine[lenOfTrimmedLine-1]) == "/" {
		jsloc.BlockComments++
		// fmt.Println("Block comment ends")
		jsloc.isBlockCommentRunning = false
		return true
	}
	return false
}

// estemates the line of code by running on each line of file
func (jsloc *JSLOC) EstimateLOCInJS() {

	filePath := strings.TrimSpace(jsloc.Filepath)
	// fmt.Println("File path is", filePath)

	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	jsloc.isBlockCommentRunning = false

	for fileScanner.Scan() {

		fileLineText := fileScanner.Text()
		trimmedLine := strings.TrimSpace(fileLineText)

		// lenOfTrimmedLine := len(trimmedLine)
		// fmt.Println("trimmedLine", trimmedLine)
		// fmt.Println("lenOfTrimmedLine", lenOfTrimmedLine)

		if jsloc.UpdateBlankCounter(trimmedLine) {
			continue
		}

		if jsloc.UpdateSingleLinCommentCounter(trimmedLine) {
			continue
		}

		if jsloc.CheckIfBlockCommentRunning(trimmedLine) {
			continue
		}

		if jsloc.CheckIfBlockCommentEnds(trimmedLine) {
			continue
		}

		if jsloc.CheckIfFirstBlockCommentLine(trimmedLine) {
			continue
		}

		if jsloc.CheckIfLastBlockCommentLine(trimmedLine) {
			continue
		}

		// else update the code line counter
		jsloc.CodeCount++
	}

	// fmt.Println("print the line of LOC", jsloc.CodeCount)
	readFile.Close()
	fmt.Println("Finished js file line of code counter")
}
