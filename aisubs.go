package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/publicsuffix"
)

func generateSubdomains(subdomain, domain, apiKey string, amount string) ([]string, error) {
	client := openai.NewClient(apiKey)

	for {
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: "You are a helpful assistant that generates new subdomains based on a given subdomain.",
					}, {
						Role:    openai.ChatMessageRoleUser,
						Content: "Generate " + amount + " subdomains similar to " + subdomain + ". Only reply with the bare subdomains, each on a new line without additional text.",
					},
				},
			},
		)

		if err != nil {
			if strings.Contains(err.Error(), "Rate limit exceeded") || strings.Contains(err.Error(), "That model is currently overloaded with other requests") {
				fmt.Println("Rate limit exceeded. Sleeping for 20 seconds...")
				time.Sleep(20 * time.Second)
				continue
			}

			if strings.Contains(err.Error(), "You exceeded your current quota") {
				fmt.Println("You exceeded your current quota, please check your plan and billing details. For more information on this error, read the docs: https://platform.openai.com/docs/guides/error-codes/api-errors")
			}

			return nil, fmt.Errorf("error generating subdomains: %w", err)
		}

		choices := resp.Choices
		if len(choices) == 0 {
			return nil, fmt.Errorf("no subdomains generated")
		}

		aiGeneratedSubdomains := make([]string, 0, len(choices[0].Message.Content))
		lines := strings.Split(choices[0].Message.Content, "\n")

		for _, line := range lines {
			if !strings.Contains(line, " ") {
				continue
			}

			if !strings.Contains(line, ".") {
				continue
			}

			parts := strings.Split(line, " ")
			result := parts[1]

			domainparts := strings.Split(result, ".")
			subdomain := domainparts[0]

			mainDomain, err := extractMainDomain(domain)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			aiGeneratedSubdomains = append(aiGeneratedSubdomains, fmt.Sprintf("%s.%s", subdomain, mainDomain))
		}

		return aiGeneratedSubdomains, nil
	}
}

func extractMainDomain(domain string) (string, error) {
	mainDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", err
	}
	return mainDomain, nil
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: cat subdomains.txt | go run main.go --apikey <OpenAI API Key> --amount 5 | httpx -ip -sc -cl -title -silent")
		fmt.Println("Usage: echo www.domain.com | go run main.go --apikey <OpenAI API Key> --amount 5")
		return
	}

	apiKey := os.Args[2]
	amount := os.Args[4]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "*") {
			continue
		}

		parts := strings.SplitN(line, ".", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Expected <subdomain>.<domain>.")
			continue
		}

		subdomain, domain := parts[0], parts[1]

		newSubdomains, err := generateSubdomains(subdomain, domain, apiKey, amount)
		if err != nil {
			fmt.Printf("Error generating subdomains: %v\n", err)
			continue
		}

		fmt.Println(strings.Join(newSubdomains, "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
