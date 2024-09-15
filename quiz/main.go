package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const filename = "problems.csv"

func main() {
	content := readCSV(filename)

	var correctAnswers int

	for i, v := range content {
		askQuestion(i, v[0])
		if parseAnswer(v[1]) {
			correctAnswers += 1
		}
	}

	log.Printf("You answered correclty to %v out of %v questions\n", correctAnswers, len(content))
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("file not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	content, err := reader.ReadAll()

	if err != nil {
		log.Println("error reading file")
	}
	return content
}

func askQuestion(index int, value string) {
	fmt.Printf("Question %v: %v\n", index+1, value)
}

func parseAnswer(value string) bool {
	var answer string
	fmt.Print("Anwser: ")
	fmt.Scan(&answer)

	return answer == value
}

// TODO : add timer & trim answers
