package main

import (
	"bufio"
	"os"
)

func main() {
	f := bufio.NewScanner(os.Stdin)
	// var path :=
	// f, err := os.Open(path)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// defer f.Close()
	// var s []string
	// f.Scan()
	// fileScanner := bufio.NewScanner(f.Text)
	// for fileScanner.Scan() {
	// 	s = append(s, fileScanner.Text())
	// }
	f.Split(bufio.ScanWords)
	wordCounter := 0
	for f.Scan() {
		wordCounter += 1
	}
	print(wordCounter)
}
