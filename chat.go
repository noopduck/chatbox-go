package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/noopduck/chatbox/internal/client"
)

// Prompt read standard input
// Exit program if Prompt receives "/bye"
func Prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Comment > ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	text = strings.TrimSuffix(text, "\n")

	if text == "/bye" {
		os.Exit(0)
	}

	return text
}

func main() {
	// Call the llm-client

	config := client.Config{
		APIBaseURL: "http://192.168.10.19",
		APIToken:   "your_api_token",
		Port:       11434,
		Model:      "llama3.1:8b",
	}

	ctx := client.Context{
		Config: config,
	}

	for {
		ctx.Client(Prompt())
	}
}
