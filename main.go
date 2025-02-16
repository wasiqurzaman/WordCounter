package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r) // A scanner is used to read text from a Reader (such as files)

	scanner.Split(bufio.ScanWords) // Define the scanner split type to words (default is split by lines)

	wc := 0 // counter

	for scanner.Scan() {
		wc++ // for every word scanned increase the counter by 1
	}

	return wc
}
