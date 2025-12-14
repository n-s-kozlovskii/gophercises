package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type question struct {
	question string
	answer   string
}

func main() {

	score := 0
	questions := make([]question, 0, 100)
	scanner := bufio.NewReader(os.Stdin)
	file, err := os.Open("example.csv")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, question{record[0], record[1]})
	}
	for _, q := range questions {
		fmt.Printf("%s? ", q.question)
		userAnswer, err := scanner.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(userAnswer) == strings.TrimSpace(q.answer) {
			score = score + 1
		}
	}
	fmt.Printf("your score is %d\n", score)
}
