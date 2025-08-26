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

	ctx := client.Context{}

	for {
		ctx.Client(Prompt())
	}
}
