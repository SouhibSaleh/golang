package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func readFromTerminal() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "-1" {
			break
		}
		fmt.Println(input.Text())
		counts[input.Text()]++

	}
	for line, arg := range counts {
		fmt.Printf("%s : %d \n", line, arg)
	}
}

func scanningFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readFromFile() {
	file, err := os.Open("data/test.txt")

	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		os.Create("data/test.txt")
		return
	}
	defer file.Close()
	scanningFile(file)
}

func readUsingIOUtil() {
	var all, err = io.ReadAll(
		io.Reader(
			bytes.NewReader(
				[]byte("test.txt"))))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(all))
	}

}
