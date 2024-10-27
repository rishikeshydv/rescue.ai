package controllers

import (
	"log"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func Voice2Text() {
	client := openai.NewClient(
		option.WithAPIKey(""),
	)
	log.Println(client)
}
