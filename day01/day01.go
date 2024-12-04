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

		// TODO: look into splitting the string more universally
		pairs := strings.Split(scanner.Text(), "   ")
		leftVal := conv_to_int(strings.Trim(pairs[0], " "))
		rightVal := conv_to_int(strings.Trim(pairs[1], " "))
		leftList = append(leftList, leftVal)
		rightList = append(rightList, rightVal)
	}

	// make sure the lists are the same size
	if len(leftList) != len(rightList) {
		log.Fatal("Lists are not the same size")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalDist := calc_distance(leftList, rightList)
	fmt.Println("Total Distance is: ", totalDist)

	totalSim := calc_similarity(leftList, rightList)
	fmt.Println("Total Similarity is: ", totalSim)

}

func conv_to_int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func calc_distance(leftList []int, rightList []int) int {
	// make sure to "match up" the smallest of the left with the smallest of
	// the right
	sort.Ints(leftList)
	sort.Ints(rightList)

	// iterate through the both lists and find the distance between the
	// two values which is the positive difference of the two values
	var totalDist int = 0
	for i := 0; i < len(leftList); i++ {
		dist := rightList[i] - leftList[i]
		if dist < 0 {
			dist = dist * -1
		}
		totalDist += dist
	}

	return totalDist
}

func calc_similarity(leftList []int, rightList []int) int {

	count_map := count_instances(rightList)

	totSimularity := 0
	for i := 0; i < len(leftList); i++ {
		totSimularity += leftList[i] * count_map[leftList[i]]
	}

	return totSimularity
}

func count_instances(list []int) map[int]int {
	counts := make(map[int]int)
	for _, val := range list {
		counts[val]++
	}
	return counts
}
