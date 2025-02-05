package ai

import (
	"context"
	
	"log"

	"gemini-go/lib"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Ai(prompt string) (string, error) {
	utils.LoadEnv()
	ctx := context.Background()
	// Access your API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GetApiKey()))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")
	response, err := model.GenerateContent(ctx, genai.Text("you're a meme generator so generate only 1 caption for meme on "+prompt))

	if err != nil {
		return "", err
	}

	if response != nil && len(response.Candidates) > 0 {
		return response.Candidates[0].Content.Part[0], nil
	}

	return "No response from AI", nil
}

