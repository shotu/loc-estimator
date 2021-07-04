package main

import (
	"strings"
	"testing"

	loc "github.com/shotu/loc-estimator/loc"
)

func TestCalJSLOC(t *testing.T) {

	filepath := "testfile2.js"
	jsloc := &loc.JSLOC{
		Filepath: strings.TrimSpace(filepath),
	}

	CalJSLOC(jsloc)

	exptectedBlockCommentCount := 11
	actualBlockCommentCount := jsloc.BlockComments + jsloc.BlanksCount

	if exptectedBlockCommentCount != actualBlockCommentCount {
		t.Errorf("Expected %v but got %v", exptectedBlockCommentCount, actualBlockCommentCount)
	}
	exptectedCommentsCount := 2
	actualCommentCount := jsloc.CommentsCount

	if exptectedCommentsCount != actualCommentCount {
		t.Errorf("Expected %v but got %v", exptectedCommentsCount, actualCommentCount)
	}

	exptectedCodeCount := 2
	actualCodeCount := jsloc.CodeCount

	if exptectedCodeCount != actualCodeCount {
		t.Errorf("Expected %v but got %v", exptectedCodeCount, actualCodeCount)
	}

}
