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

		// 3) "I AM"
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",

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
func pronouns(input string) string {
	sent := strings.Fields(input)

	// Map of string to string key-value pairs, to substitute pronouns
 	// List of substitutions adapted from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
	pronouns := Map[string]string{
		"i":      "you",
		"me":     "you",
		"you":    "I",
 		"was":    "were",
		"your":   "my",
 		"yours":  "mine",
		"my":     "your",
 		"i'd":    "you would",
 		"i've":   "you have",
 		"i'll":   "you will",
 		"are":    "am",
 		"you've": "I have",
 		"you'll": "I will",
 				
	}

	for index, word := range sent {
		if value, ok := pronouns[strings.ToLower(word)];
		ok{
			sent[index] = value;
		}
	}
	return strings.Join(sent, " ")
}


func ElizaResponse (input string) string {
	var responses = []string{
		"I'm not sure what you're trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	
	// 2) Recognise Father
	regex, _ := RegExpress.Compile("father")

	// Match String' for Father
	if regex.matStr(input) {
		return "Why don’t you tell me more about your father?"
	}
	else{
		re := RegExpress.MustCompile(`(?i)i(?:'|\sa)?m (.*)`)
	}

	if re.matStr(input) {
		return re.ReplaceAll(input, "How do you know you are $1?")
	}

	capture = pronouns(capture)

	// Returns random choice of responses
	return choice(responses)

}

// Returns random string from list
func choice(input []string) string {
	randIndex := rand.Intn(length(input))
	return input[randIndex]
}