package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numSafe int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var vals []int

		s_vals := strings.Split(scanner.Text(), " ")

		for i := 0; i < len(s_vals); i++ {
			if tval, err := strconv.Atoi(s_vals[i]); err == nil {
				vals = append(vals, tval)
			}
		}

		if is_safe(vals) {
			numSafe++
		}
	}

	fmt.Println("Number of safe levels: ", numSafe)

}

func is_safe(vals []int) bool {
	// checks to see if the values are "safe" according to the following rules
	//	for the entire set of values:
	//
	// 1. The distance between the two values is less than or equal to 3
	// 2. The difference between the two values stays positive or negative
	//

	isPositive := false
	isNegative := false

	dist := vals[1] - vals[0]
	if dist < -3 || dist > 3 {
		return false
	}

	// get the initial positivity or negativity of the set of values
	if dist < 0 {
		isNegative = true
	} else if dist > 0 {
		isPositive = true
	} else {
		return false
	}

	for i := 1; i < len(vals)-1; i++ {
		dist := vals[i+1] - vals[i]
		if dist < -3 || dist > 3 {
			return false
		} else if dist == 0 {
			return false
		} else {
			if dist > 0 && isNegative {
				return false
			} else if dist < 0 && isPositive {
				return false
			}
		}
	}
	return true
}
