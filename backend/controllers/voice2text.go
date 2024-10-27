package controllers

import (
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func Voice2Text() {
	client := openai.NewClient("sk-1234567890abcdef1234567890abcdef")
	log.Println(client)
}
