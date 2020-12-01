package main

import (
	"io/ioutil"
	"fmt"
	_ "os"
	"strings"
	"strconv"
)

func findPairs(data []int) {

	fmt.Println("------------  1st Part  ---------------")

	for i := 0 ; i < len(data); i++ {
		secondElement := 2020 - data[i]
		for j := i+1; j < len(data); j++ {
			if secondElement == data[j] {
				fmt.Println("found that combination", data[i], data[j])
				fmt.Printf("sum is %d \n", data[i] + data[j])
				fmt.Printf("prdouct is %d \n", data[i]*data[j])
			}	
		}
	}

}

func findTriplets(data []int) {

	fmt.Println("------------  2nd Part  ---------------")

	for i := 0 ; i < len(data); i++ {
		secondElement := 2020 - data[i]
		for j := i+1; j < len(data); j++ {
			thirdElement := secondElement - data[j]
			for k := i + 2 ; k < len(data); k++ {
				if thirdElement == data[k] {

					fmt.Println("found that combination", data[i], data[j], data[k])
					fmt.Printf("sum is %d \n", data[i] + data[j] + data[k])
					fmt.Printf("prdouct is %d \n", data[i]*data[j]*data[k])

				}
			}
		}
	}

}

func main() {
	
	// Read file
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error in opening file", err)
	}

	// convert []byte to []strings
	output := strings.Fields(string(content))
	data := make([]int, 200)

	// convert []strings to []int
	for i := 0 ; i < len(data); i++ {
		intData, _ := strconv.Atoi(output[i])
		data[i] = intData
	}

	findPairs(data)
	findTriplets(data)


}