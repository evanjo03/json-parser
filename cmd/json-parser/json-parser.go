package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName, err := getFileName()
	if err != nil {
		log.Fatal(err)
	}
	readFile(fileName)
}

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		if rune, size, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("%s [%d]\n", string(rune), size)
		}
	}

	return nil
}

func getFileName() (string, error) {
	args := os.Args[1:]
	if len(args) < 1 {
		return "", fmt.Errorf("no file name provided, please provide a file name as an argument")
	}
	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments, please provide only one file name as an argument")
	}
	return args[0], nil
}
