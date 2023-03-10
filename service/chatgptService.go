package service

import (
	"github.com/gin-gonic/gin"
	// gogpt "github.com/sashabaranov/go-gpt"
)

// module "github.com/sashabaranov/go-gpt3"

func getContent(c *gin.Context) string {
	// c := gogpt.NewClient("abc")
	// ctx := context.Background()

	// req := gogpt.CompletionRequest{
	// 	Model:     gogpt.GPT3Ada,
	// 	MaxTokens: 5,
	// 	Prompt:    "随便说说",
	// }
	// resp, err := c.CreateCompletion(ctx, req)

	// fmt.Println(resp.Choices[0].Text)
	return "abc"
}
