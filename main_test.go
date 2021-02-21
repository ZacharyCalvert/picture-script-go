package main

import "testing"

func TestFileExtraction(t *testing.T) {
	if "t.jpg" != pathToFname("t.jpg") {
		t.Errorf("Problem trying to parse no slash filename")
	}
	if "t.jpg" != pathToFname("/this/is/a/sub/t.jpg") {
		t.Errorf("Problem trying to parse slash filename")
	}
}
