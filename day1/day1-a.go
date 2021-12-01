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
	prev := 0
	for _, line := range input {
		if line > prev {
			count++
		}
		prev = line
	}
	fmt.Println(count)
}
