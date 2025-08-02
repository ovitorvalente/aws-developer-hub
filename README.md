# Hacktown 2025 CLI

Backend CLI em Go com integração Amazon Q para responder perguntas sobre o Hacktown 2025 em tempo real.

## Configuração

1. Configure a variável de ambiente com seu token do Amazon Q:
```bash
export AMAZON_Q_TOKEN="seu_token_aqui"
```

2. Execute o CLI:
```bash
go run main.go
```

## Funcionalidades

- 🤖 **Integração Amazon Q**: Consultas em tempo real via API
- 🎨 **Interface Visual**: CLI colorido e formatado
- 🔄 **Sistema Fallback**: Respostas locais quando API não disponível
- ⚡ **Respostas Rápidas**: Timeout de 10 segundos

## Exemplo de uso

```
╔══════════════════════════════════════════════════════════════╗
║                    🚀 HACKTOWN 2025 CLI                    ║
║              Assistente Inteligente com Amazon Q           ║
╚══════════════════════════════════════════════════════════════╝
💡 Digite sua pergunta sobre o Hacktown 2025 ou 'sair' para encerrar

❯ quando será o hacktown?
🤔 Consultando Amazon Q...
┌─ Resposta:
│ 📅 O Hacktown 2025 acontecerá de 11 a 13 de setembro de 2025
└────────────────────────────────────────────────────────────────
```

## Perguntas Suportadas

- Data e horários do evento
- Local e endereço
- Processo de inscrição
- Prêmios e premiação
- Categorias de competição
- Formato do hackathon
- Público-alvo
- Patrocinadores e parceiros

## Configuração da API Amazon Q

Para usar a integração completa com Amazon Q:

1. Obtenha acesso à API do Amazon Q
2. Configure suas credenciais AWS
3. Defina a variável de ambiente `AMAZON_Q_TOKEN`
4. A URL da API será configurada automaticamente

> **Nota**: O sistema funciona com respostas locais mesmo sem a configuração da API.