package resources

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/omeid/go-resources/live"
	"github.com/omeid/go-resources/testdata/generated"
)

//go:generate go build -o testdata/resources github.com/omeid/go-resources/cmd/resources
//go:generate testdata/resources -package generated -declare -output testdata/generated/store_prod.go testdata/*.txt testdata/*.sql testdata/sub-a

var testCases = []struct {
	name    string
	snippet string
}{
	{name: "test.txt", snippet: "this is test.txt"},
	{name: "patrick.txt", snippet: "no, this is patrick!"},
	{name: "query.sql", snippet: `drop table "files";`},
	{name: "sub-a/test2.txt", snippet: "this is test2.txt"},
	{name: "sub-a/sub-b/patrick2.txt", snippet: "no, this is patrick2!"},
}

func TestLive(t *testing.T) {
	liveDir := live.Dir("./testdata")

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := liveDir.Open(tt.name)

			if err != nil {
				t.Fatalf("expected no error opening live file, got %v", err)
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf("expected no error reading live file, got %v", err)
			}

			if !strings.Contains(string(content), tt.snippet) {
				t.Errorf("expected to find snippet %q in live file", tt.snippet)
			}
		})
	}
}

func TestGenerated(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := generated.FS.Open("testdata/" + tt.name)

			if err != nil {
				t.Fatalf("expected no error opening generated file, got %v", err)
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf("expected no error reading generated file, got %v", err)
			}

			if !strings.Contains(string(content), tt.snippet) {
				t.Errorf("expected to find snippet %q in generated file", tt.snippet)
			}
		})
	}
}
