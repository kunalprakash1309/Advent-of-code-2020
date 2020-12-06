package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

func getInput(filename string) ([]string, error){
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n\n"), nil
}


func main() {
	data, err := getInput("input.txt")
	if err != nil {
		fmt.Println("error in opening the file", err)
	}

	// fmt.Println(data)
	fmt.Println("<<<<<<<<----->>>>>>>>")
	fmt.Println(data[0])
}