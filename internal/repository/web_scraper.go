package repository

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type WebScraperRepository struct {
	client *http.Client
}

func NewWebScraperRepository() *WebScraperRepository {
	return &WebScraperRepository{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (r *WebScraperRepository) GetHacktownContent() (string, error) {
	resp, err := r.client.Get("https://hacktown.com.br/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(body)

	// Remove scripts e estilos
	content = r.cleanHTML(content)

	// Extrai texto relevante
	return r.extractRelevantText(content), nil
}

func (r *WebScraperRepository) cleanHTML(html string) string {
	// Remove scripts
	scriptRe := regexp.MustCompile(`(?s)<script.*?</script>`)
	html = scriptRe.ReplaceAllString(html, "")

	// Remove estilos
	styleRe := regexp.MustCompile(`(?s)<style.*?</style>`)
	html = styleRe.ReplaceAllString(html, "")

	// Remove tags HTML
	tagRe := regexp.MustCompile(`<[^>]*>`)
	html = tagRe.ReplaceAllString(html, " ")

	// Limpa espaços extras
	spaceRe := regexp.MustCompile(`\s+`)
	html = spaceRe.ReplaceAllString(html, " ")

	return strings.TrimSpace(html)
}

func (r *WebScraperRepository) extractRelevantText(content string) string {
	// Procura por palavras-chave relevantes
	keywords := []string{"hacktown", "2025", "agosto", "santa rita", "sapucaí", "inscrição", "prêmio", "categoria"}

	lines := strings.Split(content, ".")
	var relevantLines []string

	for _, line := range lines {
		line = strings.ToLower(strings.TrimSpace(line))
		if len(line) < 20 || len(line) > 200 {
			continue
		}

		for _, keyword := range keywords {
			if strings.Contains(line, keyword) {
				relevantLines = append(relevantLines, line)
				break
			}
		}

		if len(relevantLines) >= 10 {
			break
		}
	}

	return strings.Join(relevantLines, ". ")
}
