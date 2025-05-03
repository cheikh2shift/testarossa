package testarossa

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func GenerateTests(apikey, prompt string) (string, error) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apikey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("error performing query: %s", err)
	}

	return result.Text(), nil
}
