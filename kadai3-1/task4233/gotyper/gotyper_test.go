package gotyper

import (
	"fmt"
	"os"
	"reflect"
	"testing"
"strconv"
)

// TestSetParameters tests SetPerameters().
func TestSetParameters(t *testing.T) {
	cases := []struct {
		name             string
		time             int
		wordlist         string
		expectedTime     int
		expectedWordPath string
		expectedWords    []string
	}{
		{name: "no arguments", time: 0, wordlist: "", expectedTime: 30, expectedWordPath: "wordlist.txt", expectedWords: []string{"alpha", "bravo", "charlie"}},
		{name: "time only", time: 60, wordlist: "", expectedTime: 60, expectedWordPath: "wordlist.txt", expectedWords: []string{"alpha", "bravo", "charlie"}},
		{name: "wordlist only", time: 0, wordlist: "wordlist.txt", expectedTime: 30, expectedWordPath: "wordlist.txt", expectedWords: []string{"alpha", "bravo", "charlie"}},
		{name: "all arguments", time: 60, wordlist: "wordlist.txt", expectedTime: 60, expectedWordPath: "wordlist.txt", expectedWords: []string{"alpha", "bravo", "charlie"}},
		{name: "invalid path", time: 0, wordlist: "wrongpath.txt", expectedTime: 30, expectedWordPath: "failed to SetParameters(): Name: invalid path\nDescription: This path is invalid\nHint: Check if the path exists\nDebug: wrongpath.txt", expectedWords: []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

			os.Args = append(os.Args, "go")
			if c.time != 0 {
				os.Args = append(os.Args, "-time="+strconv.Itoa(c.time))
			}
			if c.wordlist != "" {
				os.Args = append(os.Args, "-wordlist="+c.wordlist)
			}

			gotyper := New()
			if err := gotyper.SetParameters(); err != nil {
				if err.Error() != c.expectedWordPath {
					t.Errorf("failed to set parameter:\n%s\n%s\n", err.Error(), c.expectedWordPath)
				}
			} else if gotyper.TimeLimitSec != c.expectedTime {
				t.Errorf("TimeLimitSec=> Actual: %d, Expected: %d", gotyper.TimeLimitSec, c.expectedTime)
			} else if !reflect.DeepEqual(gotyper.Words, c.expectedWords) {
				t.Errorf("Words=> Actual: %s, Expected: %s", gotyper.Words, c.expectedWords)
			}
			os.Args = []string{}
		})

	}
}
