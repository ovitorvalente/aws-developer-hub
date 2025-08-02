package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hacktown-cli/internal/domain"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type CLI struct {
	questionService domain.QuestionService
	scanner         *bufio.Scanner
}

func NewCLI(questionService domain.QuestionService) *CLI {
	return &CLI{
		questionService: questionService,
		scanner:         bufio.NewScanner(os.Stdin),
	}
}

func (c *CLI) Run() {
	c.printHeader()

	for {
		c.printPrompt()
		if !c.scanner.Scan() {
			break
		}

		input := strings.TrimSpace(c.scanner.Text())

		if c.isExitCommand(input) {
			c.printGoodbye()
			break
		}

		if input == "" {
			continue
		}

		c.processQuestion(input)
	}
}

func (c *CLI) processQuestion(input string) {
	c.printThinking()
	question := domain.Question{Text: input}
	answer := c.questionService.ProcessQuestion(question)
	c.printAnswer(answer.Text)
}

func (c *CLI) isExitCommand(input string) bool {
	lower := strings.ToLower(input)
	return lower == "sair" || lower == "exit"
}

func (c *CLI) printHeader() {
	fmt.Printf("%s╔══════════════════════════════════════════════════════════════╗%s\n", colorCyan, colorReset)
	fmt.Printf("%s║%s                    🚀 HACKTOWN CLI 2025                    %s║%s\n", colorCyan, colorYellow, colorCyan, colorReset)
	fmt.Printf("%s║%s              Assistente Inteligente com Google Gemini      %s║%s\n", colorCyan, colorWhite, colorCyan, colorReset)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════╝%s\n", colorCyan, colorReset)
	fmt.Printf("%s💡 Digite sua pergunta sobre o Hacktown 2025 ou 'sair' para encerrar%s\n\n", colorGreen, colorReset)
}

func (c *CLI) printPrompt() {
	fmt.Printf("%s❯%s ", colorBlue, colorReset)
}

func (c *CLI) printThinking() {
	fmt.Printf("%s🔍 Consultando dados do Hacktown 2025...%s\n", colorYellow, colorReset)
	fmt.Printf("%s🤖 Processando com Google Gemini...%s\n", colorYellow, colorReset)
}

func (c *CLI) printAnswer(answer string) {
	fmt.Printf("%s┌─ Resposta:%s\n", colorGreen, colorReset)
	fmt.Printf("%s│%s %s\n", colorGreen, colorReset, answer)
	fmt.Printf("%s└────────────────────────────────────────────────────────────────%s\n\n", colorGreen, colorReset)
}

func (c *CLI) printGoodbye() {
	fmt.Printf("%s\n👋 Obrigado por usar o Hacktown CLI 2025! Até logo!%s\n", colorPurple, colorReset)
}
