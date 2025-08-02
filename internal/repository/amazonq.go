package repository

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"hacktown-cli/internal/domain"
)

type QRequest struct {
	Message string `json:"message"`
}

type QResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type AmazonQRepository struct {
	client  *http.Client
	baseURL string
	token   string
}

func NewAmazonQRepository() *AmazonQRepository {
	return &AmazonQRepository{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: "https://api.amazonq.aws/chat",
		token:   os.Getenv("AMAZON_Q_TOKEN"),
	}
}

func (r *AmazonQRepository) Query(question domain.Question) (domain.Answer, error) {
	reqBody := QRequest{Message: question.Text}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return domain.Answer{}, err
	}

	req, err := http.NewRequest("POST", r.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return domain.Answer{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.token)

	resp, err := r.client.Do(req)
	if err != nil {
		return domain.Answer{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Answer{}, err
	}

	var qResp QResponse
	if err := json.Unmarshal(body, &qResp); err != nil {
		return domain.Answer{}, err
	}

	if qResp.Error != "" {
		return domain.Answer{}, err
	}

	return domain.Answer{Text: qResp.Response}, nil
}