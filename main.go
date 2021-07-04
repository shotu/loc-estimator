package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	loc "github.com/shotu/loc-estimator/loc"
)

// type LOCer interface{
// 	GetLOC
// }

func GetFileType(filePath string) string {
	return "JS"
}

func CalJSLOC(jsloc *loc.JSLOC) {
	jsloc.EstimateLOCInJS()
	fmt.Println("Blank comment count          :", jsloc.BlockComments+jsloc.BlanksCount)
	fmt.Println("CommentsCount count          :", jsloc.CommentsCount)
	fmt.Println("Code comment count           :", jsloc.CodeCount)
	fmt.Println("Total count                  :", jsloc.BlockComments+jsloc.BlanksCount+jsloc.CodeCount+jsloc.CommentsCount)
}

func main() {

	reader := bufio.NewReader(os.Stdin) // reading file name, from userinput, make sure file is present in current working directory
	fmt.Print("Enter filepath:")
	filepath, _ := reader.ReadString('\n')
	fmt.Println(filepath)
	fileType := GetFileType(filepath)

	// case file type:
	if fileType == "JS" {
		jsloc := &loc.JSLOC{
			Filepath: strings.TrimSpace(filepath),
		}
		CalJSLOC(jsloc)
	}
}
