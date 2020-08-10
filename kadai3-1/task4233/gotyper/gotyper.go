/*
Package gotyper provides a typing game.
It writes an english word on stdout. After that, it reads one line from stdin.
A user repeats that within the time limit, and it provides result.

e.g.)
gotyper -time=60 -wordlist=./wordlist

*/
package gotyper

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	version = "0.0.1"
	msg     = "gotyper v" + version + ", provides a typing game.\n"
)

var (
	timeLimit    = flag.Int("time", 30, "time limit(.sec)")
	wordlistPath = flag.String("wordlist", "wordlist.txt", "file for wordlist")
)

// Gotyper struct is the manager for game.
type Gotyper struct {
	Score int
	// TimeLimitSec is 30sec by default
	TimeLimitSec int
	// Words prepared manually
	Words []string
}

// New makes Gotyper instance.
// Call this function if you use this package.
func New() *Gotyper {
	return &Gotyper{
		Score:        0,
		TimeLimitSec: 30,
	}
}

// Run starts typeing game.
func (gotyper *Gotyper) Run() error {
	if err := gotyper.SetParameters(); err != nil {
		return err
	}

	return nil
}

// SetParameters sets parameters for game
func (gotyper *Gotyper) SetParameters() error {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t./gotyper [-time=timeLimit] -wordlist=./wordlist.txt\n")
		os.Exit(1)
	}

	// init Parameters for test
	*timeLimit = 30
	*wordlistPath = "./wordlist.txt"

	// parse information
	flag.Parse()

	// set information
	if timeLimit != nil && *timeLimit != 30 {
		gotyper.TimeLimitSec = *timeLimit
	}
	if err := gotyper.SetWords(*wordlistPath); err != nil {
		return fmt.Errorf("failed to SetParameters(): %w", err)
	}

	return nil
}

// SetWords sets wordlists
// Maxwords is 1024 words because it's enough I think
func (gotyper *Gotyper) SetWords(wordlistPath string) error {
	if _, err := os.Stat(wordlistPath); err != nil {
		return ErrInvalidPath.WrapErr(err).WithDebug(wordlistPath)
	}

	wordFile, err := os.Open(filepath.Clean(wordlistPath))
	if err != nil {
		return ErrInvalidPath.WrapErr(err).WithDebug(wordlistPath)
	}
	defer func() {
		cerr := wordFile.Close()
		if cerr != nil {
			fmt.Fprintf(os.Stderr, "failed to close file: %s\n", wordlistPath)
		}
	}()

	s := bufio.NewScanner(wordFile)
	var line int = 0
	for s.Scan() && line < 1024 {
		gotyper.Words = append(gotyper.Words, s.Text())
		line++
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("failed to SetWords(): %w", err)
	}

	return nil
}
