package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"hacktown-cli/internal/domain"
)

type MockAmazonQRepository struct{}

func NewMockAmazonQRepository() *MockAmazonQRepository {
	return &MockAmazonQRepository{}
}

func (r *MockAmazonQRepository) Query(question domain.Question) (domain.Answer, error) {
	// Processa a pergunta diretamente
	q := strings.ToLower(question.Text)
	
	if strings.Contains(q, "quando") || strings.Contains(q, "data") {
		return domain.Answer{Text: "📅 O HackTown 2025 acontece de 31 de julho a 3 de agosto de 2025 em Santa Rita do Sapucaí, MG! É o maior festival de inovação, criatividade e tecnologia da América Latina, com mais de 1.000 atividades."}, nil
	}
	
	if strings.Contains(q, "onde") || strings.Contains(q, "local") {
		return domain.Answer{Text: "📍 O HackTown 2025 será realizado em Santa Rita do Sapucaí, Minas Gerais"}, nil
	}
	
	if strings.Contains(q, "aws") || strings.Contains(q, "amazon") {
		return domain.Answer{Text: "🚀 A AWS estará presente com o 'Vibe Coding Dojo' de 31/07 a 03/08 (10:00-12:30) na ETE FMC, oferecendo experiências práticas com Amazon Q Developer. Também haverá Workshop Amazon Bedrock no dia 03/08 no INATEL!"}, nil
	}
	
	if strings.Contains(q, "programação") || strings.Contains(q, "atividade") {
		return domain.Answer{Text: "📅 O HackTown 2025 tem programação completa de 31/07 a 03/08! Inclui palestras sobre IA, workshops de blockchain, design thinking, empreendedorismo sustentável e muito mais. Confira a programação completa no site oficial!"}, nil
	}
	
	return domain.Answer{Text: "🤖 O HackTown 2025 é o maior festival de inovação da América Latina, com mais de 1.000 atividades de 31 de julho a 3 de agosto em Santa Rita do Sapucaí, MG!"}, nil
}

func (r *MockAmazonQRepository) generateResponse(question string, hacktown map[string]interface{}) domain.Answer {
	q := strings.ToLower(question)
	
	if strings.Contains(q, "quando") || strings.Contains(q, "data") {
		dataInicio := hacktown["data_inicio"].(string)
		dataFim := hacktown["data_fim"].(string)
		local := hacktown["local"].(string)
		descricao := hacktown["descricao"].(string)
		
		return domain.Answer{
			Text: fmt.Sprintf("📅 O HackTown 2025 acontece de 31 de julho a 3 de agosto de 2025 em %s!\n\n%s", local, descricao),
		}
	}
	
	if strings.Contains(q, "onde") || strings.Contains(q, "local") {
		local := hacktown["local"].(string)
		return domain.Answer{Text: fmt.Sprintf("📍 O HackTown 2025 será realizado em %s", local)}
	}
	
	if strings.Contains(q, "aws") || strings.Contains(q, "amazon") {
		return domain.Answer{
			Text: "🚀 A AWS estará presente com o 'Vibe Coding Dojo' de 31/07 a 03/08 (10:00-12:30) na ETE FMC, oferecendo experiências práticas com Amazon Q Developer. Também haverá Workshop Amazon Bedrock no dia 03/08 no INATEL!",
		}
	}
	
	return domain.Answer{Text: "🤖 O HackTown 2025 é o maior festival de inovação da América Latina, com mais de 1.000 atividades de 31 de julho a 3 de agosto em Santa Rita do Sapucaí, MG!"}