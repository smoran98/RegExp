package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main{

	var responses = []string {

		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!",

	}


	for i := 0; i < length(responses); i++ {
		fmt.Println(responses[i])
		fmt.Println(ElizaResponse(responses[i]))
	}

	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// http.HandleFunc("/user-input", handleEliza)
	// http.ListenAndServe(":8080", nil)

	ElizaResponse()
}


func ElizaResponse (input string) string {
	var responses = []string{
		"I'm not sure what you're trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	
	// Returns random choice of responses
	return choice(responses)

}

// Returns random string from list
func choice(input []string) string {
	randIndex := rand.Intn(length(input))
	return input[randIndex]
}