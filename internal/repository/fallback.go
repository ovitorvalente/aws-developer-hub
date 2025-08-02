package repository

import (
	"strings"

	"hacktown-cli/internal/domain"
)

type FallbackRepository struct{}

func NewFallbackRepository() *FallbackRepository {
	return &FallbackRepository{}
}

func (r *FallbackRepository) GetAnswer(question domain.Question) domain.Answer {
	return domain.Answer{Text: "🤖 Sistema temporariamente indisponível. Tente novamente em alguns instantes ou acesse hacktown.com.br para informações atualizadas."}
}