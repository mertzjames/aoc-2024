package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sumTotal := 0

	r, _ := regexp.Compile(`mul\((\d*,\d*)\)`)
	mul_matches := r.FindAllStringSubmatch(string(file), -1)

	for i := 0; i < len(mul_matches); i++ {
		mul_vals := strings.Split(mul_matches[i][1], ",")
		left, _ := strconv.Atoi(mul_vals[0])
		right, _ := strconv.Atoi(mul_vals[1])

		sumTotal += left * right

	}

	fmt.Println("Total is: ", sumTotal)
}
