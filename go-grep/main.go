package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"strings"
)

func main() {
	if !isFromPipe() || !hasStringSearch() {
		fmt.Println("Usage: cat <file> | grep-go \"string-to-search\"")
		os.Exit(1)
	}

	SearchString(os.Args[1], os.Stdin, os.Stdout)
}

func SearchString(search string, r io.Reader, w io.Writer) error {
	red := color.New(color.FgRed, color.Bold)
	green := color.New(color.FgGreen)

	scanner := bufio.NewScanner(bufio.NewReader(r))
	lineCount := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, search) {
			newString := strings.ReplaceAll(line, search, red.Sprintf(search))
			fmt.Fprintln(w, green.Sprintf("%v", lineCount), newString)
		}

		lineCount++
	}

	return nil
}

func isFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func hasStringSearch() bool {
	return len(os.Args) > 1 && os.Args[1] != ""
}
