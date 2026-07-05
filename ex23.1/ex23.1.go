package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return " ", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, err
}
func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("fail to create file")
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("fail to read file")
			return
		}
	}
	fmt.Println("file contents:", line)
}
