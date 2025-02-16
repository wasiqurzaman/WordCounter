package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines") // Defining a boolean flag -l to count lines instead of words

	flag.Parse() // Parsing the flags provided by the user

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r) // A scanner is used to read text from a Reader (such as files)

	if !countLines {
		// If the countLines flag is not set, we want to count words
		scanner.Split(bufio.ScanWords) // Define the scanner split type to words (default is split by lines)
	}

	wc := 0 // counter

	for scanner.Scan() {
		wc++ // for every word scanned increase the counter by 1
	}

	return wc
}
