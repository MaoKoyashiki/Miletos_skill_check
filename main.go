package main

import (
	"encoding/json"
	"fmt"
	"log"

	"skill_check/parser"
)

const fileName = "./test.conf"

func main() {
	fmt.Println("=== 出力 ===")
	printAsJSON(fileName)
}

func printAsJSON(input string) {
	result, err := parser.ParseFile(input)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}
	fmt.Println(string(jsonBytes))
}
