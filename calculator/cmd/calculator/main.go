package main

import (
	"fmt"
	"log"
	"os"

	"github.com/turman17/GoLang-progress/calculator/internal/calc"
)

func usage() {
	log.Fatalf("Usage: %s <number> <+ | - | * | /> <number>", os.Args[0])
}

func main() {
	if len(os.Args) != 4 {
		usage()
	}
	args := os.Args[1:] // [num1, op, num2]

	nums, op, err := calc.ParseArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	res, err := calc.Evaluate(nums, op)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %.3f\n", res)
}