package eimg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// TestSetParameters tests SetPerameters().
func TestSetParameters(t *testing.T) {
	cases := []struct {
		name     string
		rootDir  string
		fromExt  string
		toExt    string
		expected []string
	}{
		{name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: "", expected: []string{"test/documents", "jpeg", "png"}},
		{name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: "", expected: []string{"test/img", "gif", "png"}},
		{name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif", expected: []string{"test/img", "jpeg", "gif"}},
		{name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg", expected: []string{"test/img", "gif", "jpeg"}},
		{name: "invalid path", rootDir: "test/test", fromExt: "", toExt: "", expected: []string{"Name: invalid path\nDescription: This path is invalid\nHint: Check if the path exists\nDebug: "}},
	}

	defer func() {
		if _, err := os.Stat("test"); err == nil {
			if err := os.RemoveAll("test"); err != nil {
				t.Errorf("Failed to remove test")
			}
		}
	}()
	CopyFilesRec(t, "testdata", "test")

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

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

			eimg := New()
			if err := eimg.SetParameters(); err != nil {
				if err.Error() != c.expected[0] {
					t.Errorf("failed to set parameter:\n%s\n%s\n", err.Error(), c.expected[0])
				}
			} else if eimg.RootDir != c.expected[0] {
				t.Errorf("RootDir=> Actual: %s, Expected: %s", eimg.RootDir, c.expected[0])
			} else if eimg.FromExt != c.expected[1] {
				t.Errorf("FromExt=> Actual: %s, Expected: %s", eimg.FromExt, c.expected[1])

			} else if eimg.ToExt != c.expected[2] {
				t.Errorf("ToExt=> Actual: %s, Expected: %s", eimg.ToExt, c.expected[2])

			}

			os.Args = []string{}
		})

	}
}

// TestEncodeFile tests EncodeFile()
func TestEncodeFile(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
		fromExt  string
		toExt    string
		expected string
	}{
		{name: "invalid file", filePath: ".", fromExt: "txt", toExt: "", expected: "Name: failed to convert image object\nDescription: Failed to Convert image object\nHint: Check the specified formats\nDebug: ."},
		{name: "invalid path", filePath: "test/test", fromExt: "", toExt: "", expected: "Name: invalid path\nDescription: This path is invalid\nHint: Check if the path exists\nDebug: test/test"},
		{name: "check png", filePath: "test/img/green.jpeg", fromExt: "jpeg", toExt: "png", expected: "test/img/green.png"},
		{name: "check jpg", filePath: "test/img/blue.gif", fromExt: "gif", toExt: "jpeg", expected: "test/img/blue.jpeg"},
		{name: "check gif", filePath: "test/img/red.png", fromExt: "png", toExt: "gif", expected: "test/img/red.gif"},
	}

	eimg := New()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

			defer func() {
				if err := os.RemoveAll("test"); err != nil {
					t.Errorf("Failed to remove test")
				}
			}()
			CopyFilesRec(t, "testdata", "test")

			eimg.FromExt = c.fromExt
			eimg.ToExt = c.toExt
			if err := eimg.EncodeFile(c.filePath); err != nil {
				if err.Error() != c.expected {
					t.Errorf("failed encodeFile():\n%s\n", err.Error())
				}
			} else if _, err := os.Stat(c.expected); err != nil {
				t.Errorf("%s: %s", c.expected, err)
			}
		})
	}
}

// TestConvertExtension tests ConvertExtension()
func TConvertExtension(t *testing.T) {
	cases := []struct {
		name     string
		rootDir  string
		fromExt  string
		toExt    string
		expected []string
	}{
		{name: "invalid file", rootDir: ".", fromExt: "txt", toExt: "", expected: []string{"Name: failed to convert image object\nDescription: Failed to Convert image object\nHint: test/documents/fuga.txt\nDebug: image: unknown format"}},
		{name: "invalid path", rootDir: "test/test", fromExt: "jpeg", toExt: "png", expected: []string{"Name: invalid path\nDescription: This path is invalid\nHint: Check if the path exists\nDebug: Name: invalid path\nDescription: This path is invalid\nHint: Check if the path exists\nDebug: open test/test: no such file or directory"}},
		{name: "check png", rootDir: "test", fromExt: "jpeg", toExt: "png", expected: []string{"test/img/green.png"}},
		{name: "check jpg", rootDir: "test", fromExt: "gif", toExt: "jpeg", expected: []string{"test/img/blue.jpeg"}},
		{name: "check gif", rootDir: "test", fromExt: "png", toExt: "gif", expected: []string{"test/img/red.gif", "test/white.gif"}},
	}

	eimg := New()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			defer func() {
				if err := os.RemoveAll("test"); err != nil {
					t.Errorf("Failed to remove test")
				}
			}()
			CopyFilesRec(t, "testdata", "test")

			eimg.RootDir = c.rootDir
			eimg.FromExt = c.fromExt
			eimg.ToExt = c.toExt

			if err := eimg.ConvertExtension(); err != nil {
				if err.Error() != c.expected[0] {
					t.Errorf("failed ConvertExtension():\n%s\n", err.Error())
				}
			} else {
				for _, filePath := range c.expected {
					if _, err := os.Stat(filePath); err != nil {
						t.Errorf("%s: %s", filePath, err)
					}
				}
			}
		})
	}
}

// CopyDir copies files recursively
func CopyFilesRec(t *testing.T, filePath string, destRootPath string) {
	t.Helper()

	// Get filePaths recursively
	filePaths, err := GetFilePathsRec(filePath)
	if err != nil {
		t.Errorf("Failed to GetFilePathsRec(): %s", err)
	}

	for _, srcFilePath := range filePaths {
		func() {
			// open srcFile
			srcFile, err := os.Open(srcFilePath)
			if err != nil {
				t.Errorf("Failed to open file: %s", srcFilePath)
			}
			defer func() {
				cerr := srcFile.Close()
				if cerr != nil {
					fmt.Fprintf(os.Stderr, "Failed to close file: %s\n", srcFilePath)
				}
			}()

			// destFile should not exist
			// if exists, overwrites it
			destFilePath := destRootPath + srcFilePath[len(filePath):]

			// craete dir if not exist
			if _, err := os.Stat(filepath.Dir(destFilePath)); err != nil {
				if err := os.MkdirAll(filepath.Dir(destFilePath), 0755); err != nil {
					t.Errorf("Failed to create dir: %s\n", err)
				}
			}

			// create file
			destFile, err := os.Create(destFilePath)
			if err != nil {
				t.Errorf("Failed to create file: %s\n", destFilePath)
			}
			defer func() {
				cerr := destFile.Close()
				if cerr != nil {
					fmt.Fprintf(os.Stderr, "Failed to close file: %s\n", destFilePath)
				}
			}()

			// copy file
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				t.Errorf("Failed to copy file: %s\n", srcFilePath)
			}
		}()
	}

}
