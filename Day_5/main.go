package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func handleInput(fileName string) ([]string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	data := strings.Fields(string(content))
	return data, err

}

func findRow(data string) int{
	min, max := 0, 128
	data = data[:len(data)-3]
	for i:=0; i<len(data);i++ {
		if string(data[i]) == "B"{
			min = (min + max + 1)/2
		} 
		if string(data[i]) == "F"{
			max = (min + max + 1)/2
			max = max-1
		}
	}
	return min
}

func findColumn(data string) int{
	min, max := 0, 7
	data = data[len(data)-3:]
	for i:=0; i<len(data);i++ {
		if string(data[i]) == "R"{
			min = (min + max + 1)/2
		} 
		if string(data[i]) == "L"{
			max = (min + max + 1)/2
			max = max-1
		}
	}
    return min
}

func main() {
	var highestId int
	index := 0
	data, err := handleInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//data = []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	// fmt.Println(data)
	for i:=0; i<len(data); i++ {
		seatID := (findRow(data[i])*8) + findColumn(data[i])
		if seatID == 1 {
			fmt.Println(seatID)
		}
		if seatID > highestId {
			highestId = seatID
			index = i
		}
	}

	fmt.Println("Highest seat ID:- ", highestId)
	fmt.Println("Index", index)
}