package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var input []int

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		input = append(input, x)
	}

	count := -1
	prev := input[0]
	for i := 2; i < len(input); i++ {
		sum := input[i] + input[i-1] + input[i-2]
		if sum > prev {
			count++
		}
		prev = sum
	}
	fmt.Println(count)
}
