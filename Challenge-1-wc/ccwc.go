package main

// Importing necessary packages for file I/O, buffered reading, logging, and string manipulation.
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ReadFile reads all content from the given reader (file or stdin) and returns the byte count along with any error encountered.
func ReadFile(reader io.Reader) (int, error) {
	content, err := io.ReadAll(reader)
	count := len(content) // Counting the bytes in the read content
	return count, err
}

// ScanLines reads lines from the given reader (file or stdin) and returns the line count along with any error encountered.
func ScanLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	count := 0
	for scanner.Scan() {
		count++ // Increment count for each line
	}
	return count, scanner.Err()
}

// ScanWords reads words from the given reader (file or stdin) using a scanner split function for words,
// and returns the word count along with any error encountered.
func ScanWords(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords) // Setting the scanner to split input into words
	wordCounter := 0
	for scanner.Scan() {
		wordCounter++ // Increment count for each word
	}
	return wordCounter, scanner.Err()
}

// main is the entry point of the program. It processes command-line arguments to perform byte, line, or word counts on input.
func main() {
	var byteCount, lineCount, wordCount int
	var err error

	// Processing command-line arguments to perform the requested operation.
	switch os.Args[1] {
	case "-c": // Byte count
		if len(os.Args) == 3 { // If a filename is provided
			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()
			byteCount, err = ReadFile(file)
		} else { // Read from stdin if no filename is provided
			byteCount, err = ReadFile(os.Stdin)
		}
		fmt.Printf("Size of file is  %d \n", byteCount)

	case "-l": // Line count
		if len(os.Args) == 3 {
			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()
			lineCount, err = ScanLines(file)
		} else {
			lineCount, err = ScanLines(os.Stdin)
		}
		fmt.Printf("The number of lines are %d \n", lineCount)

	case "-w": // Word count
		if len(os.Args) == 3 {
			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()
			wordCount, err = ScanWords(file)
		} else {
			wordCount, err = ScanWords(os.Stdin)
		}
		fmt.Printf("The number of words are %d \n", wordCount)

	case "-m": // Combined metrics: bytes, lines, and words
		if len(os.Args) == 3 { // Handling file input
			filePath := os.Args[2]
			// Repeating file opening and closing for each metric to reset read position
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
				return
			}
			byteCount, err = ReadFile(file)
			file.Close()

			file, err = os.Open(filePath)
			lineCount, err = ScanLines(file)
			file.Close()

			file, err = os.Open(filePath)
			wordCount, err = ScanWords(file)
			file.Close()
		} else { // Handling stdin by reading once and analyzing the content for all metrics
			content, err := io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			byteCount = len(content) // Direct byte count from content length

			contentString := string(content) // Converting byte slice to string for further analysis
			lineCount, _ = ScanLines(strings.NewReader(contentString))
			wordCount, _ = ScanWords(strings.NewReader(contentString))
		}
		fmt.Printf("Bytes: %d, Lines: %d, Words: %d\n", byteCount, lineCount, wordCount)

	default:
		fmt.Printf("Not yet implemented")
	}

	if err != nil {
		log.Fatal(err) // Logging any errors encountered during processing
	}
}
