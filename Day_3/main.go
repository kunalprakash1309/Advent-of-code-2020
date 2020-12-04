package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func calculateTress(data1 []string, moveX, moveY int) int {
	var counter = 0

	data := make([][]string, len(data1))
	for i := 0; i < len(data1); i++ {
		data[i] = strings.Fields(data1[i])
	}

	for currentX, currentY := 0, 0; currentY < len(data); currentX, currentY = currentX+moveX, currentY+moveY {

		if currentX >= (len(data[currentY][0])) {
			currentX = currentX % (len(data[currentY][0]))
		}

		position := data[currentY][0][currentX]

		if string(position) == "#" {
			counter++
		}
	}
	return counter
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error in opening the file")
	}
	data1 := strings.Split(string(content), "\n")

	result1 := calculateTress(data1, 3, 1)
	fmt.Println("First Part result ---> ", result1)

	multiply := calculateTress(data1, 1, 1) * calculateTress(data1, 3, 1) * calculateTress(data1, 5, 1) * calculateTress(data1, 7, 1) * calculateTress(data1, 1, 2)
	fmt.Println("Second Part Result ---> ", multiply)

}
