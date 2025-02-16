# Gemini Go

## Project Overview

Gemini Go is an AI-based response generator written in Go. The project uses Google's Generative AI to generate meme captions based on user-provided prompts. The main entry point of the application is `main.go`, which starts an HTTP server. The AI functionality is implemented in `ai/client.go`, and the HTTP request handler is defined in `handlers/request-handler.go`. Environment variables are managed in `lib/utils.go`, and dependencies are listed in `go.mod`.

## Setup

To set up the project, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/ojuss/gemini-go.git
   cd gemini-go
   ```

2. Install the dependencies:
   ```sh
   go mod tidy
   ```

3. Create a `.env` file in the root directory and add your Gemini API key:
   ```sh
   GEMINI_API_KEY=your_api_key_here
   ```

4. Run the project:
   ```sh
   go run main.go
   ```

## Usage

To use the API endpoint, send a POST request to `http://localhost:8080/ai` with a JSON payload containing the prompt. For example:
```sh
curl -X POST http://localhost:8080/ai -d '{"prompt": "funny cat"}' -H "Content-Type: application/json"
```

## Dependencies

The main dependencies used in this project are:
- [github.com/google/generative-ai-go](https://github.com/google/generative-ai-go)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [google.golang.org/api](https://pkg.go.dev/google.golang.org/api)
