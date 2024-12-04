package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var leftList []int
	var rightList []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), "   ")
		leftVal := conv_to_int(strings.Trim(pairs[0], " "))
		rightVal := conv_to_int(strings.Trim(pairs[1], " "))
		leftList = append(leftList, leftVal)
		rightList = append(rightList, rightVal)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	//check if the lists are the same size
	if len(leftList) != len(rightList) {
		log.Fatal("Lists are not the same size")
	}

	//iterate through the both lists and find the distance between the two values
	var totalSum int = 0
	for i := 0; i < len(leftList); i++ {
		dist := rightList[i] - leftList[i]
		if dist < 0 {
			dist = dist * -1
		}
		totalSum += dist
	}

	fmt.Println("Total sum is: ", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func conv_to_int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
