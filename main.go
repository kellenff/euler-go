package main

import (
	"fmt"
	"github.com/kellenff/euler-go/problem_1"
	"os"
	"strconv"
)

func main() {
	problem := os.Args[1]
	args := os.Args[2:]

	switch problem {
	case "problem_1":
		limit, _ := strconv.ParseUint(args[0], 10, 64)
		fmt.Println(problem_1.Solve(uint(limit)))
	}
}
