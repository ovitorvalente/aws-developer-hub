package repository

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"hacktown-cli/internal/domain"
)

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

type GeminiCandidate struct {
	Content GeminiContent `json:"content"`
}

type GeminiRepository struct {
	client *http.Client
	apiKey string
}

func NewGeminiRepository() *GeminiRepository {
	return &GeminiRepository{
		client: &http.Client{Timeout: 10 * time.Second},
		apiKey: loadEnvVar("GEMINI_API_KEY"),
	}
}

func loadEnvVar(key string) string {
	file, err := os.Open(".env")
	if err != nil {
		return os.Getenv(key)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			return strings.TrimPrefix(line, key+"=")
		}
	}
	return os.Getenv(key)
}

func (r *GeminiRepository) Query(question domain.Question) (domain.Answer, error) {
	reqBody := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{Text: question.Text},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return domain.Answer{}, err
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return domain.Answer{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", r.apiKey)

	resp, err := r.client.Do(req)
	if err != nil {
		return domain.Answer{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Answer{}, err
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return domain.Answer{}, err
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return domain.Answer{}, err
	}

	return domain.Answer{Text: geminiResp.Candidates[0].Content.Parts[0].Text}, nil
}