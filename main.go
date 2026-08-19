package main

import (
	"fmt"
	"os"
)

// json_processor - JSON manipulation toolkit
func json_processor(path string) {
	fmt.Println("========================================")
	fmt.Println("  JSON-Processor")
	fmt.Println("  JSON manipulation toolkit")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	json_processor(path)
}
