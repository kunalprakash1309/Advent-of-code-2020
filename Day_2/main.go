package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func firstPart(data []string) {
	var sum = 0
	for i := 0; i<len(data); i++ {
		eachElement := strings.Fields(data[i])
		fmt.Println("------------",eachElement,"----------")

		for i := 0 ; i < 1; i++ {

			firstColumn := eachElement[0]
			limit := strings.Split(string(firstColumn), "-")
			min, _ := strconv.Atoi(limit[0])
			max, _ := strconv.Atoi(limit[1])
			fmt.Println(min,"-------", max)
	
			secondColumn := eachElement[1]
			requiredKey := string(secondColumn[0])
			fmt.Println("Required Key:-  ",requiredKey)
	
			thirdColumn := eachElement[2]
			fmt.Println("Given string:- ",thirdColumn)
			fmt.Println("Required count:- ",strings.Count(thirdColumn, requiredKey))
			count := strings.Count(thirdColumn, requiredKey)
	
			if count <= max && count >= min {
				fmt.Print("found \n\n")
				sum++
			} else {
				fmt.Print("not found \n\n")
			}
		}
	}

	fmt.Println("***********  sum is :- ", sum, "   *************")

}

func secondPart(data []string) {
	var sum = 0
	for i := 0; i<len(data); i++ {
		eachElement := strings.Fields(data[i])
		fmt.Println("------------",eachElement,"----------")

		for i := 0 ; i < 1; i++ {

			firstColumn := eachElement[0]
			limit := strings.Split(string(firstColumn), "-")
			min, _ := strconv.Atoi(limit[0])
			max, _ := strconv.Atoi(limit[1])
			fmt.Println(min,"-------", max)
	
			secondColumn := eachElement[1]
			requiredKey := string(secondColumn[0])
			fmt.Println("Required Key:-  ",requiredKey)
	
			thirdColumn := eachElement[2]
			fmt.Println("Given string:- ",thirdColumn)
			fmt.Println("Required count:- ",strings.Count(thirdColumn, requiredKey))
			// count := strings.Count(thirdColumn, requiredKey)
	
			if string(thirdColumn[min-1]) == requiredKey{
				if string(thirdColumn[max-1]) != requiredKey{
					sum++
				}
			}
			if string(thirdColumn[min-1]) != requiredKey{
				if string(thirdColumn[max-1]) == requiredKey{
					sum++
				}
			} else {
				fmt.Println("not found")
			}
			fmt.Print("--------------   end    ---------------\n\n\n\n\n\n")
		}
	}

	fmt.Println("***********  sum is :- ", sum, "   *************")
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error in opening in file", err)
	}

	// converting compelete content to slice of string
	data := strings.Split(string(content), "\n")
	
	// comment in the other function which you not want to get result
	firstPart(data)
	secondPart(data)

}