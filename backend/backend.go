package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	// "strings"
	// "time"
	"html"


)

type Question struct{
	Category 		string 		`json:"category"`
	Type 			string 		 `json:"type"`
	Difficulty		string 		`json:"difficulty"`
	Question		string 		`json:"question"`
	CorrectAnswer	string 		`json:"correct_answers"`
	IncorrectAnswers []string	`json:"incorrect_answers"`

}

type ApiResponse struct {
	ResponseCode int `json:"response_code"`
	Results []Question `json:"results"`
}

func fetchQuestions (amount int) ([]Question, error){
	url := fmt.Sprintf("https://opentdb.com/api.php?amount=%d", amount)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err 
	}
	defer resp.Body.Close()

	var apiResp ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp.Results, nil

}



func main() {
	var amount_1 int
    fmt.Println("Welcome to the Trivia CLI Game!")
	fmt.Println("How many questions would you like?")
	fmt.Scan(&amount_1)
	fmt.Println("Fetching questions...")

	questions, err := fetchQuestions(amount_1)
	if err != nil {
		fmt.Println("Failed to fetch trivia questions: ", err)
		os.Exit(1)
	}

	// score := 0

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, html.UnescapeString(q.Question))
		options := append(q.IncorrectAnswers, q.CorrectAnswer)

		fmt.Println(options)
	}

	
	
}