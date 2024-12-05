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
	var numSafeDampened int = 0

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

		if is_safe(vals, 0) {
			numSafe++
		}

		if is_safe(vals, 1) {
			numSafeDampened++
		}
	}

	fmt.Println("Number of safe levels: ", numSafe)
	fmt.Println("Number of safe levels with dampening: ", numSafeDampened)

}

func is_safe(vals []int, fault_tol int) bool {
	// checks to see if the values are "safe" according to the following rules
	//	for the entire set of values:
	//
	// 1. The distance between the two values is less than or equal to 3
	// 2. The difference between the two values stays positive or negative
	//

	stateSet := false
	isPositive := false
	isNegative := false
	fault_count := 0

	for i := 0; i < len(vals)-1; i++ {
		dist := vals[i+1] - vals[i]
		sign := get_sign(dist)

		if !stateSet {
			stateSet = true
			if sign == 1 {
				isPositive = true
			} else if sign == -1 {
				isNegative = true
			} else {
				fault_count++
				stateSet = false
			}
		}
		if stateSet {
			if dist < -3 || dist > 3 {
				fault_count++
			} else if dist == 0 {
				fault_count++
			} else {
				if dist > 0 && isNegative {
					fault_count++
				} else if dist < 0 && isPositive {
					fault_count++
				}
			}
		}
		if fault_count > fault_tol {
			return false
		}
	}
	return true
}

func get_sign(val int) int {
	if val < 0 {
		return -1
	} else if val > 0 {
		return 1
	}
	return 0
}
