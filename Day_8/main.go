package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(content), "\n")
	return data, nil
}

func main() {
	content, err := getInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var index []int

	i := 0
	acc := 0
	for i < len(content) {
		instruction := strings.Fields(content[i])
		if checkRepeat(index, i) {
			break
		}
		if instruction[0] == "nop" {
			// fmt.Println(i, " nop founded")
			fmt.Println(i, "nop")
			i++
		}
		if instruction[0] == "acc" {
			argument := instruction[1]
			// fmt.Println(string(argument))
			if string(argument[0]) == "+" {
				value , _:= strconv.Atoi(string(argument[1:]))
				acc = acc + value
			}
			if string(argument[0]) == "-" {
				value , _:= strconv.Atoi(string(argument[1:]))
				acc = acc - value
				// fmt.Println(value)
			}
			fmt.Println(i, "acc")
			i++
		}
		if instruction[0] == "jmp" {
			argument := instruction[1]
			if string(argument[0]) == "+" {
				value, _ := strconv.Atoi(string(argument[1:]))
				fmt.Println(i, "jmp +")
				i = i + value 
			}
			if string(argument[0]) == "-" {
				value, _ := strconv.Atoi(string(argument[1:]))
				fmt.Println(i, "jmp -")
				i = i - value
			}
		}
		index = append(index, i)
		fmt.Println("Total Value: ", acc)
	}
}

func checkRepeat(s []int, i int) bool {
	for _, v := range s{
		if i == v {
			return true
		}
	}
	return false
}