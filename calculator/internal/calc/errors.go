package calc

import (
	"log"
	"os"
)

func usage() {
	log.Fatalf("Usage: %s <number> <+ | - | * | /> <number>", os.Args[0])
}