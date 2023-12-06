package customerimporter

import (
	"os"
	"testing"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func TestSaveDomainsToCSV(t *testing.T) {
	domainCountSlice := countSlice{
		{"example.com", 3},
		{"test.com", 1},
		{"demo.com", 2},
	}
	filename := "test_example.csv"
	saveDomainsToCSV(domainCountSlice, filename)
	created := fileExists(filename)
	if !created {
		t.Errorf("Output files was not created.")
	}
	defer func() {
		if created {
			os.Remove(filename)
		}
	}()

	content, _ := os.ReadFile(filename)
	strContent := string(content)
	expectedContent := "domain,count\nexample.com,3\ntest.com,1\ndemo.com,2\n"

	if expectedContent != strContent {
		t.Errorf("Expected file output \n'%s', but got \n'%s'", expectedContent, strContent)
	}

}
