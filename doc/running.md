# Rodando a aplicação

Este documento descreve como rodar a aplicação em diferentes ambientes.

---

## Configuração Inicial

1. Crie o arquivo `.env` a partir do exemplo:

```bash
cp .env.example .env
```

2. Ajuste as variáveis:

```bash
APP_ENV=development
APP_PORT=3001
APP_HOST=0.0.0.0

DB_HOST=localhost
DB_PORT=5432
DB_NAME=beershop
DB_USER=beershop
DB_PASSWORD=beershop
```

---
## Pré-requisitos

Desenvolvimento local: Go 1.25+, Make
Docker: Docker, Docker Compose, Make

---

## Desenvolvimento Local (Hot Reload)

Instale o Air:
```go
go install github.com/air-verse/air@latest
```

Execute:
```bash
make run-dev
```

## Desenvolvimento Local (Sem Hot Reload)
```go
go run ./src/cmd/main.go
```

## Execução com Docker

```bash
make build
make run-prod
```

A aplicação será executada em container Docker, utilizando imagem multi-stage com Distroless.