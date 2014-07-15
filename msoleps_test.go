package msoleps

import (
	"os"
	"testing"
)

var (
	testDocSum = "test/DocumentSummaryInformation"
	testSum    = "test/SummaryInformation"
	testSum1   = "test/SummaryInformation1"
)

func testFile(t *testing.T, path string) {
	file, _ := os.Open(path)
	defer file.Close()
	doc, err := New(file)
	if err != nil {
		t.Errorf("Error opening file; Returns error: ", err)
	}
	for prop, _ := doc.Next(); entry != nil; entry, _ = doc.Next() {

	}

}

func TestDocSum(t *testing.T) {
	testFile(t, testDocSum)
}

func TestSum(t *testing.T) {
	testFile(t, testSum)
}

func TestSum1(t *testing.T) {
	testFile(t, testSum1)
}
