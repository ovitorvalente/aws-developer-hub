<!-- markdownlint-disable MD040 -->
<!-- markdownlint-disable MD029 -->

# Hacktown 2025 CLI

Backend CLI em Go com integração Google Gemini para responder perguntas sobre o Hacktown 2025 em tempo real.

## Arquitetura

Projeto estruturado seguindo Clean Code e princípios SOLID:

```
├── cmd/                    # Ponto de entrada da aplicação
├── internal/
│   ├── domain/            # Entidades e interfaces de negócio
│   ├── service/           # Lógica de negócio
│   ├── repository/        # Acesso a dados externos
│   └── ui/                # Interface do usuário
├── go.mod
└── README.md
```

## Configuração

1. Configure a variável de ambiente no arquivo .env:

```bash
GEMINI_API_KEY=sua_chave_aqui
```

2. Execute o CLI:

```bash
go run cmd/main.go
```

## Funcionalidades

- 🔍 **Web Scraping**: Extrai informações atuais do hacktown.com.br
- 🤖 **Integração Google Gemini**: Processa conteúdo real com IA
- 💾 **Cache Inteligente**: Cache de 5 minutos para otimizar consultas
- 🎨 **Interface Visual**: CLI colorido e formatado
- 🔄 **Sistema Fallback**: Respostas locais quando site/API indisponível
- ⚡ **Timeout**: 10 segundos para consultas
- 📅 **Dados Estruturados**: Base de conhecimento com programação completa
- 🤖 **Respostas Dinâmicas**: IA gera respostas personalizadas e naturais

## Princípios Aplicados

- **Single Responsibility**: Cada módulo tem uma responsabilidade específica
- **Open/Closed**: Extensível para novos repositórios de dados
- **Liskov Substitution**: Interfaces bem definidas
- **Interface Segregation**: Interfaces pequenas e específicas
- **Dependency Inversion**: Dependências injetadas via interfaces

## Exemplo de uso

```
╔══════════════════════════════════════════════════════════════╗
║                    🚀 HACKTOWN 2025 CLI                    ║
║              Assistente Inteligente com Google Gemini      ║
╚══════════════════════════════════════════════════════════════╝
💡 Digite sua pergunta sobre o Hacktown 2025 ou 'sair' para encerrar

❯ quando será o hacktown?
🔍 Consultando dados do Hacktown 2025...
🤖 Processando com Google Gemini...
┌─ Resposta:
│ 📅 O HackTown 2025 acontece de 31 de julho a 3 de agosto de 2025 em Santa Rita do Sapucaí, MG!
│ É o maior festival de inovação, criatividade e tecnologia da América Latina, com mais de 1.000 atividades.
└────────────────────────────────────────────────────────────────
```
