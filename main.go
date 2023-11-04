package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func maxInt(vals []int) int {
	var max int

	for i, val := range vals {
		if i == 0 || val > max {
			max = val
		}
	}
	return max
}

func maxCalories(calories []string) []int {

	elves := make([]int, 0)

	total_calories := 0

	for _, calorie := range calories {

		if calorie == "" {
			// append total calories to the elves
			elves = append(elves, total_calories)
			total_calories = 0
			continue
		}

		c, err := strconv.ParseInt(calorie, 0, 64)
		if err != nil {
			log.Fatalf(err.Error())
		}
		total_calories += int(c)

	}

	if total_calories != 0 {
		//append this to the total calories
		elves = append(elves, total_calories)
		total_calories = 0
	}

	fmt.Printf("Elves %v\n", elves)
	return elves
}

func maxCaloriesFromFile(fname string) []int {

	elves := make([]int, 0)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total_calories := 0

	for scanner.Scan() {

		val := scanner.Text()

		if val == "" {
			elves = append(elves, total_calories)
			total_calories = 0
			continue
		}

		calorie := string(val)
		c, err := strconv.ParseInt(calorie, 0, 64)
		if err != nil {
			log.Fatalf(err.Error())
		}
		total_calories += int(c)

		//fmt.Printf("Val is %s is prefix %v\n", val, isPrefix)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Reading standard input: %s", err.Error())
	}

	if total_calories != 0 {
		fmt.Println("Last calories", total_calories)
		//append this to the total calories
		elves = append(elves, total_calories)
		total_calories = 0
	}

	fmt.Printf("Elves %v\n", elves)

	return elves
}

func main() {

	// day one

	cals := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	elves := maxCalories(cals)
	max := maxInt(elves)

	fmt.Printf("Maximum elve is %d \n", max)

	//day one real

	foodElves := maxCaloriesFromFile("adventofcode.com_2022_day_1_input.txt")

	maxElve := maxInt(foodElves)

	// should be 69912
	fmt.Printf(" Maximum elve is %d \n", maxElve)

	//inputStr := string(inputs)
	//fmt.Println(inputStr)

}
