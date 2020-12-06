package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
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

func secondPart(data []int) {
	sort.Ints(data)
	for i:=0; i<len(data);i++ {
		if i + 35 != data[i]{
			fmt.Println("My previous seat:= ",data[i-1])
			fmt.Println("My next seat:= ",data[i])
			fmt.Printf("So my seat is in between these:= %d", (data[i-1] + data[i])/2)
			break
		}
	}
}

func main() {
	var highestId int
	data, err := handleInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var array = []int{}
	for i:=0; i<len(data); i++ {
		seatID := (findRow(data[i])*8) + findColumn(data[i])
		array = append(array, seatID)
		if seatID == 1 {
			fmt.Println(seatID)
		}
		if seatID > highestId {
			highestId = seatID
		}
	}

	fmt.Println("Highest seat ID:- ", highestId)
	secondPart(array)
}