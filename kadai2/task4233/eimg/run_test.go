package eimg

import (
	"fmt"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		rootDir  string
		fromExt  string
		toExt    string
		expected string
	}{
		{name: "invalid file", rootDir: ".", fromExt: "txt", toExt: "", expected: "Name: failed to convert image object\nDescription: Failed to Convert image object\nHint: Check the specified formats\nDebug: test/documents/fuga.txt"},
		{name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: "", expected: ""},
		{name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: "", expected: ""},
		{name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif", expected: ""},
		{name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg", expected: ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

			// set arguments
			os.Args = append(os.Args, "go")
			if c.fromExt != "" {
				os.Args = append(os.Args, "-f="+c.fromExt)
			}
			if c.toExt != "" {
				os.Args = append(os.Args, "-t="+c.toExt)
			}
			if c.rootDir != "" {
				os.Args = append(os.Args, c.rootDir)
			}

			defer func() {
				if err := os.RemoveAll("test"); err != nil {
					t.Errorf("Failed to remove test")
				}
			}()
			CopyFilesRec(t, "testdata", "test")

			eimg := New()
			if err := eimg.Run(); err != nil {
				if err.Error() != c.expected {
					t.Errorf("failed Run():\n%s\n", err.Error())
				}
			}

			os.Args = []string{}
		})
	}
}
