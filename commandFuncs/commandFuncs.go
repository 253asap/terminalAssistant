package commandfuncs

import (
	"assist/util"
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func ChatGpt() {
	client := openai.NewClient("")
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a friendly virtual penpal named サターニャ who speaks Japanese. You only reply in standard Japanese kana and easy kanji. You reply in english only when asked to.",
		},
	}
	keepChatting := true
	input := bufio.NewScanner(os.Stdin)

	for keepChatting {
		fmt.Print("User: ")
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
