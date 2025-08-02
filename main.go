package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type QRequest struct {
	Message string `json:"message"`
}

type QResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func main() {
	printHeader()
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		printPrompt()
		if !scanner.Scan() {
			break
		}
		
		question := strings.TrimSpace(scanner.Text())
		
		if strings.ToLower(question) == "sair" || strings.ToLower(question) == "exit" {
			printGoodbye()
			break
		}
		
		if question == "" {
			continue
		}
		
		printThinking()
		answer := queryAmazonQ(question)
		printAnswer(answer)
	}
}

func printHeader() {
	fmt.Printf("%s╔══════════════════════════════════════════════════════════════╗%s\n", colorCyan, colorReset)
	fmt.Printf("%s║%s                    🚀 HACKTOWN 2025 CLI                    %s║%s\n", colorCyan, colorYellow, colorCyan, colorReset)
	fmt.Printf("%s║%s              Assistente Inteligente com Amazon Q           %s║%s\n", colorCyan, colorWhite, colorCyan, colorReset)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════╝%s\n", colorCyan, colorReset)
	fmt.Printf("%s💡 Digite sua pergunta sobre o Hacktown 2025 ou 'sair' para encerrar%s\n\n", colorGreen, colorReset)
}

func printPrompt() {
	fmt.Printf("%s❯%s ", colorBlue, colorReset)
}

func printThinking() {
	fmt.Printf("%s🤔 Consultando Amazon Q...%s\n", colorYellow, colorReset)
}

func printAnswer(answer string) {
	fmt.Printf("%s┌─ Resposta:%s\n", colorGreen, colorReset)
	fmt.Printf("%s│%s %s\n", colorGreen, colorReset, answer)
	fmt.Printf("%s└────────────────────────────────────────────────────────────────%s\n\n", colorGreen, colorReset)
}

func printGoodbye() {
	fmt.Printf("%s\n👋 Obrigado por usar o Hacktown 2025 CLI! Até logo!%s\n", colorPurple, colorReset)
}

func queryAmazonQ(question string) string {
	// Contexto específico do Hacktown 2025
	contextualQuestion := fmt.Sprintf("Sobre o Hacktown 2025 (maior hackathon do interior do Brasil, em Santa Rita do Sapucaí, MG): %s", question)
	
	reqBody := QRequest{Message: contextualQuestion}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "❌ Erro ao processar pergunta"
	}
	
	// Simula chamada para Amazon Q (substitua pela URL real da API)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", "https://api.amazonq.aws/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		return getFallbackAnswer(question)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("AMAZON_Q_TOKEN"))
	
	resp, err := client.Do(req)
	if err != nil {
		return getFallbackAnswer(question)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return getFallbackAnswer(question)
	}
	
	var qResp QResponse
	if err := json.Unmarshal(body, &qResp); err != nil {
		return getFallbackAnswer(question)
	}
	
	if qResp.Error != "" {
		return getFallbackAnswer(question)
	}
	
	return qResp.Response
}

func getFallbackAnswer(question string) string {
	// Respostas locais baseadas no site oficial hacktown.com.br
	q := strings.ToLower(question)
	
	if strings.Contains(q, "quando") || strings.Contains(q, "data") {
		return "📅 O Hacktown 2025 acontece em setembro de 2025 em Santa Rita do Sapucaí, MG"
	}
	if strings.Contains(q, "onde") || strings.Contains(q, "local") {
		return "📍 Santa Rita do Sapucaí, Minas Gerais - Vale da Eletrônica"
	}
	if strings.Contains(q, "inscrição") || strings.Contains(q, "inscrever") {
		return "📝 Inscrições e mais informações em hacktown.com.br"
	}
	if strings.Contains(q, "prêmio") || strings.Contains(q, "premiação") {
		return "🏆 Hacktown oferece premiação para os melhores projetos - detalhes em hacktown.com.br"
	}
	if strings.Contains(q, "categoria") || strings.Contains(q, "tema") {
		return "🎯 Hacktown abrange diversas categorias de inovação tecnológica"
	}
	if strings.Contains(q, "hacktown") && strings.Contains(q, "que") {
		return "🚀 Hacktown é o maior hackathon do interior do Brasil, realizado em Santa Rita do Sapucaí, MG"
	}
	if strings.Contains(q, "participar") || strings.Contains(q, "público") {
		return "👥 Aberto para desenvolvedores, designers, empreendedores e entusiastas de tecnologia"
	}
	
	return "🤖 Para informações atualizadas sobre o Hacktown 2025, acesse hacktown.com.br ou pergunte sobre: data, local, inscrições, premiação, categorias."
}