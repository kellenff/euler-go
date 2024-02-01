package main

import (
	"fmt"
	"github.com/kellenff/euler-go/problem_1"
	"github.com/kellenff/euler-go/problem_2"
	"github.com/kellenff/euler-go/problem_3"
	"github.com/kellenff/euler-go/problem_4"
	"github.com/kellenff/euler-go/problem_5"
	"github.com/kellenff/euler-go/problem_6"
	"github.com/kellenff/euler-go/problem_7"
	"github.com/kellenff/euler-go/problem_8"
	"github.com/kellenff/euler-go/problem_9"
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
	case "problem_2":
		limit, _ := strconv.ParseUint(args[0], 10, 64)
		fmt.Println(problem_2.Solve(limit))
	case "problem_3":
		n, _ := strconv.ParseUint(args[0], 10, 64)
		fmt.Println(problem_3.Solve(n))
	case "problem_4":
		fmt.Println(problem_4.Solve())
	case "problem_5":
		fmt.Println(problem_5.Solve())
	case "problem_6":
		fmt.Println(problem_6.Solve())
	case "problem_7":
		fmt.Println(problem_7.Solve())
	case "problem_8":
		fmt.Println(problem_8.Solve())
	case "problem_9":
		fmt.Println(problem_9.Solve())
	}
}
