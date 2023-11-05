package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "dictionnary/words.txt"

	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
