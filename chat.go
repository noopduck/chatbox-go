package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/noopduck/chatbox/internal/client"
)

// Input prompt, read standard input
// Exit program if Prompt receives "/bye"
func Prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Comment > ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	if strings.Compare(text, "/bye") == 0 {
		os.Exit(0)
	}

	return text
}

func main() {
	// Call the llm-client
	fmt.Println("OOOKEY")

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
