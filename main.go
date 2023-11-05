package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func maxCalories(fname string, elves *[]int) {

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
			*elves = append(*elves, total_calories)
			total_calories = 0
			continue
		}
		c, err := strconv.ParseInt(val, 0, 64)
		if err != nil {
			log.Fatalf(err.Error())
		}
		total_calories += int(c)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Reading : %s", err.Error())
	}

	if total_calories != 0 {
		//append this to the total calories
		*elves = append(*elves, total_calories)
		total_calories = 0
	}

}

func main() {

	foodElves := make([]int, 0)
	maxCalories("adventofcode.com_2022_day_1_input.txt", &foodElves)

	slices.Sort(foodElves)
	// should be 69912
	maxiElve := slices.Max(foodElves)
	fmt.Printf(" Maximum Elve is %d \n", maxiElve)

	// sorted elves
	fmt.Printf("Sorted Elves %v\n", foodElves)

	slices.Reverse(foodElves)
	// top 3
	fmt.Printf("Top 3 sorted %v\n", foodElves[:3])

	//sum of top3
	sum := 0
	for _, v := range foodElves[:3] {
		sum += v
	}
	fmt.Printf("Total sum is  %d\n", sum)

}
