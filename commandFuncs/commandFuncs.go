package commandfuncs

import (
	"assist/util"
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GptEpal() {

	client := openai.NewClient(util.GetEnvVar("API_TOKEN"))
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a friendly virtual penpal named Satania who speaks Japanese. You only reply in standard Japanese kana and easy kanji. You reply in english only when asked to.",
		},
	}
	keepChatting := true
	input := bufio.NewScanner(os.Stdin)

	for keepChatting {
		fmt.Print("You: ")
		input.Scan()
		message := input.Text()
		if message == "exit" {
			keepChatting = false
			break
		}
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: message})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)
		util.ErrorCheck("Couldnt complete your request to openAI", err)

		gptResp := resp.Choices[0].Message.Content
		fmt.Printf("\nSatania: %s\n\n", gptResp)
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: gptResp})
	}
}

func GptBase() {
	client := openai.NewClient(util.GetEnvVar("API_TOKEN"))
	var messages []openai.ChatCompletionMessage
	keepChatting := true
	input := bufio.NewScanner(os.Stdin)

	for keepChatting {
		fmt.Print("You: ")
		input.Scan()
		message := input.Text()
		if message == "exit" {
			keepChatting = false
			break
		}
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: message})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)
		util.ErrorCheck("Couldnt complete your request to openAI", err)

		gptResp := resp.Choices[0].Message.Content
		fmt.Printf("\nGPT: %s\n\n", gptResp)
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: gptResp})
	}
}
