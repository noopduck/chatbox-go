# Chatbox

A simple command-line chat interface written in Go that interacts with a local LLM API.

## Features

- Interactive command-line interface
- Streaming responses from the LLM
- Conversation context maintenance
- Simple exit command (`/bye`)

## Prerequisites

- Go 1.x
- Access to LLM API endpoint (default: `http://192.168.10.19:11434/api/chat`)
- LLM model: llama3.1:8b

## Usage

1. Clone the repository:
```bash
git clone https://github.com/noopduck/chatbox-go.git
cd chatbox-go
```

2. Run the application:
```bash
go run chat.go
```

3. Start chatting! Type your messages at the prompt:
```
Comment > 
```

4. To exit the application, type `/bye`

## Project Structure

- `chat.go` - Main application entry point and CLI interface
- `internal/client/client.go` - LLM API client implementation

## License

This project is licensed under the MIT License:

```
MIT License

Copyright (c) 2025 noopduck

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
