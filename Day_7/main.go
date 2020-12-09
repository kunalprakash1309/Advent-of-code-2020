package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)



func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(content), "\n")
	return data, nil
}

func setData(data []string){
	var relation map[string]map[string]int
	var parent string
	var child map[string]int
	for i, v := range data {
		if i == 0 || i == 1 {
			parent = parent + string(v)
		}

		
		// if i == 4 || i == 5 || i == 6 {
			
		// }
		// if i == 8 || i == 9 || i == 10 {
		// 	child = append(child, string(v))
		// }
	}
	fmt.Println(parent)
	fmt.Println(child)
	relation[parent] = child
}


func main() {
	content, err := getInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, statement:= range content{
		data := strings.Fields(statement)
		setData(data)
		fmt.Println("XXXXXXXXXXXXXXXXXXXXX")
	}
}