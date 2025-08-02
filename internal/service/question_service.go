package service

import (
	"fmt"

	"hacktown-cli/internal/domain"
	"hacktown-cli/internal/repository"
)

type QuestionServiceImpl struct {
	geminiRepo       *repository.GeminiRepository
	fallbackRepo     *repository.FallbackRepository
	webScraperRepo   *repository.WebScraperRepository
	cacheRepo        *repository.CacheRepository
	hacktownDataRepo *repository.HacktownDataRepository
}

func NewQuestionService() domain.QuestionService {
	return &QuestionServiceImpl{
		geminiRepo:       repository.NewGeminiRepository(),
		fallbackRepo:     repository.NewFallbackRepository(),
		webScraperRepo:   repository.NewWebScraperRepository(),
		cacheRepo:        repository.NewCacheRepository(),
		hacktownDataRepo: repository.NewHacktownDataRepository(),
	}
}

func (s *QuestionServiceImpl) ProcessQuestion(question domain.Question) domain.Answer {
	// Obtém dados estruturados do Hacktown (fonte confiável)
	hacktownData := s.hacktownDataRepo.GetHacktownInfo()

	// Cria contexto priorizando dados estruturados
	contextualQuestion := domain.Question{
		Text: fmt.Sprintf(`Você é um assistente especializado no HackTown 2025. Use EXCLUSIVAMENTE as informações oficiais abaixo para responder.

DADOS OFICIAIS DO HACKTOWN 2025 (FONTE CONFIÁVEL):
%s

INSTRUÇÕES IMPORTANTES:
- SEMPRE use as datas dos dados oficiais: 31 de julho a 3 de agosto de 2025
- NUNCA mencione setembro - o evento é em julho/agosto
- Responda em português brasileiro de forma natural e concisa
- Use emojis quando apropriado
- Para datas específicas, consulte o campo "data_inicio" e "data_fim"
- Para programação, use as atividades listadas por data
- Se não encontrar informação nos dados oficiais, diga que não tem essa informação

PERGUNTA: %s`, hacktownData, question.Text),
	}

	// Consulta Google Gemini com contexto rico
	answer, err := s.geminiRepo.Query(contextualQuestion)
	if err != nil {
		return s.fallbackRepo.GetAnswer(question)
	}

	return answer
}