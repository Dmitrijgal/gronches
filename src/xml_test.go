package src

import (
	"os"
	"testing"
)

func BenchmarkReadCsv(t *testing.B) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()

	_, _ = ReadXML(f)

}

func ReadTest(t *testing.T) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()

	_, _ = ReadXML(f)
}
