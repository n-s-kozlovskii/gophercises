package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	filePathFlag := flag.String("csv", "questions.csv", "path to file with questions in format 'question,answer'")
	limit := flag.Int("limit", 30, "time limit of quiz in seconds")
	flag.Parse()

	score := 0
	scanner := bufio.NewReader(os.Stdin)
	file, err := os.Open(*filePathFlag)
	if err != nil {
		log.Fatalf("failed to open file: %s", *filePathFlag)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	questions := make([]question, len(lines))
	if err != nil {
		log.Fatalf("failed to load questions from %s", *filePathFlag)
	}
	for i, l := range lines {
		questions[i] = question{question: l[0], answer: strings.TrimSpace(l[1])}
	}

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	go func() {
		defer os.Exit(0)
		<-timer.C
		fmt.Printf("\nPassed %d seconds, you're done;\nscore is %d out of %d\n", limit, score, len(questions))
	}()

	for _, q := range questions {
		fmt.Printf("%s? ", q.question)
		userAnswer, err := scanner.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(userAnswer) == q.answer {
			score = score + 1
		}
	}
	fmt.Printf("your score is %d out %d\n", score, len(questions))
}
