package domain

// Question representa uma pergunta do usuário
type Question struct {
	Text string
}

// Answer representa uma resposta do sistema
type Answer struct {
	Text string
}

// QuestionService define o contrato para processamento de perguntas
type QuestionService interface {
	ProcessQuestion(question Question) Answer
}