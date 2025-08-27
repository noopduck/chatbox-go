package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Config struct {
	APIBaseURL string
	APIToken   string
	Port       int
	Model      string
}

type Context struct {
	previousMessages string
	Config           Config
}

func (c *Context) Client(text string) {
	// create the payload
	req := Request{
		Model: c.Config.Model,
		Messages: []Message{
			{
				Role:    "user",
				Content: c.previousMessages + " " + text,
			},
		},
		Stream: true,
	}

	fmt.Println(c.previousMessages)

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var connection = c.Config.APIBaseURL + ":" + fmt.Sprint(c.Config.Port) + "/api/chat"

	resp, err := http.Post(connection, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 response %d\n", resp.StatusCode)
		return
	}

	// Define the Message struct matching the response
	type Message struct {
		Model     string `json:"model"`
		CreatedAt string `json:"created_at"`
		Message   struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Done bool `json:"done"`
	}

	// Use a scanner to read the response line by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		var message Message
		err := json.Unmarshal(line, &message)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}
		// Print each token's content without extra newlines if needed
		fmt.Print(message.Message.Content)
		c.previousMessages += message.Message.Content
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
	}
	fmt.Println() // add newline after complete response
}
