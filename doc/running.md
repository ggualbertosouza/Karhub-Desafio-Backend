# Rodando a aplicação

Este documento descreve como rodar a aplicação em diferentes ambientes.

---

## Configuração Inicial

A aplicação utiliza arquivos de ambiente separados para **desenvolvimento** e **produção**.

### Arquivos disponíveis

- `.env.dev` → configuração para desenvolvimento local
- `.env.prod` → configuração para execução via Docker Compose (produção)

Antes de rodar a aplicação, é necessário **copiar um desses arquivos para `.env`**, que é o arquivo efetivamente carregado pela aplicação.

---

## Pré-requisitos

- Desenvolvimento local
    - Go 1.25+
    - Make

- Docker / Produção
    - Docker
    - Docker Compose
    - Make

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

## Produção — Desafio da Vaga (Docker Compose)

Este é o modo recomendado para avaliação do desafio, pois reproduz um ambiente próximo ao de produção.

**Passo a passo**

1. Faça uma copia do arquivo `.env.prod`:
```bash
cp .env.prod .env
```

2. Suba o ambiente de produção
```bash
make prod-up
```

3. Após a aplicação estar rodando, execute o script de seed:
```bash
chmod +x seed.sh
./seed.sh
```

**Importante:**
O script seed.sh executa chamadas HTTP para a API, portanto a aplicação precisa estar rodando antes de utilizá-lo.

Verifique a aplicação,a API estará disponível em:
`http://localhost:3001`

## Encerrando o Ambiente de Produção

Para parar e remover os containers:
```bash
make prod-down
```