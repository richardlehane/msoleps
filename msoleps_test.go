package msoleps

import (
	"fmt"
	"os"
	"testing"
)

var (
	testDocSum = "test/DocumentSummaryInformation"
	testSum    = "test/SummaryInformation"
	testSum1   = "test/SummaryInformation1"
)

func testFile(t *testing.T, path string) *Reader {
	file, _ := os.Open(path)
	defer file.Close()
	doc, err := NewFrom(file)
	if err != nil {
		t.Errorf("Error opening file; Returns error: ", err)
	}
	for _, prop := range doc.Property {
		fmt.Printf("%s: %s\n", prop.Name, prop)
	}
	return doc
}

func TestDocSum(t *testing.T) {
	doc := testFile(t, testDocSum)
	if len(doc.Property) != 12 {
		t.Error("Expecting 12 properties, got %d", len(doc.Property))
	}
}

func TestSum(t *testing.T) {
	doc := testFile(t, testSum)
	if len(doc.Property) != 17 {
		t.Error("Expecting 17 properties, got %d", len(doc.Property))
	}
}

func TestSum1(t *testing.T) {
	doc := testFile(t, testSum1)
	if len(doc.Property) != 3 {
		t.Error("Expecting 3 properties, got %d", len(doc.Property))
	}
}
