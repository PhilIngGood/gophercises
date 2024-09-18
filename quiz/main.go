package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//const filename = "problems.csv"
//const timeDuration = 5

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	score    int
}

func main() {

	filename := flag.String("filename", "problems.csv", "CSV File that conatins quiz questions")
	timeLimit := flag.Int("limit", 30, "Time Limit for each question")
	flag.Parse()

	quiz := quiz{}
	populateQuiz(&quiz, *filename)

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	defer timer.Stop()

	quizOver := false

	go func() {
		<-timer.C
		fmt.Println("Time's up !")
		quizOver = true
	}()

	for index, problem := range quiz.problems {
		fmt.Printf("Question %v: %v\n", index+1, problem.question)
		var answer string
		fmt.Print("Answer: ")
		fmt.Scan(&answer)

		if quizOver {
			break
		}

		if answer == problem.answer {
			quiz.score += 1
		}
	}
	fmt.Printf("You answered %v questions correclty out of %v", quiz.score, len(quiz.problems))

}

func newProblem(q string, a string) problem {
	return problem{
		question: q,
		answer:   a,
	}
}

func populateQuiz(q *quiz, f string) {
	file, err := os.Open(f)

	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		p := newProblem(line[0], line[1])
		q.problems = append(q.problems, p)
	}

}
