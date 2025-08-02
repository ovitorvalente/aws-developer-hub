package main

import (
	"hacktown-cli/internal/service"
	"hacktown-cli/internal/ui"
)

func main() {
	questionService := service.NewQuestionService()
	cli := ui.NewCLI(questionService)
	cli.Run()
}