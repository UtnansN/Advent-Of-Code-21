package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	x, y := 0, 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		switch tokens[0] {
		case "forward":
			val, _ := strconv.Atoi(tokens[1])
			x += val
		case "up":
			val, _ := strconv.Atoi(tokens[1])
			y -= val
		case "down":
			val, _ := strconv.Atoi(tokens[1])
			y += val
		default:
			break
		}
	}
	fmt.Println(x * y)
}
