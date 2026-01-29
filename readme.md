# Karhub Backend Challenge

Este repositório contém a implementação do backend desenvolvida como parte de um **teste técnico para a vaga de Desenvolvedor Backend**.

---

## 📌 Visão Geral

---

## 🧰 Tecnologias Utilizadas

* **Go (Golang)** — linguagem principal
* **Gin** — framework HTTP
* **Docker** — build e execução em produção
* **Distroless** — imagem mínima para produção
* **Air** — hot reload em desenvolvimento
* **Makefile** — automação de comandos
* **godotenv** — carregamento de variáveis de ambiente em desenvolvimento

---

## 📁 Estrutura do Projeto

```text
├── docker
│   ├── app
│   │   └── Dockerfile
│   └── postgres              # (reservado para uso futuro)
├── src
│   ├── cmd
│   │   └── main.go            # Entry point da aplicação
│   ├── server
│   │   ├── config             # Configuração de ambiente
│   │   │   ├── config.go
│   │   │   └── env.go
│   │   ├── health.go          # Health check
│   │   ├── router.go          # Rotas HTTP
│   │   └── server.go          # HTTP server e lifecycle
│   └── internal               # Código interno (reservado para evolução)
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Configuração por Variáveis de Ambiente

A aplicação segue o princípio de **configuração via environment variables (12-factor app)**.

### Variáveis suportadas atualmente

| Variável | Descrição             | Obrigatória | Default     |
| -------- | --------------------- | ----------- | ----------- |
| APP_ENV  | Ambiente da aplicação | Sim         | development |
| APP_HOST | Host do servidor      | Não         | 0.0.0.0     |
| APP_PORT | Porta do servidor     | Sim         | —           |

> Em desenvolvimento, as variáveis podem ser definidas em um arquivo `.env`.

---

## 🚀 Como Rodar o Projeto

Esta seção descreve os **modelos de execução disponíveis**, partindo de um **setup padrão comum a todos os ambientes** e depois detalhando cada forma de rodar a aplicação.

---

### 🔰 Setup Inicial (comum a todos os ambientes)

Independentemente do ambiente (desenvolvimento, testes locais ou produção), o primeiro passo é configurar as variáveis de ambiente.

1. Crie o arquivo `.env` a partir do exemplo:

```bash
cp .env.example .env
```

2. Ajuste as variáveis conforme necessário:

```env
APP_ENV=development
APP_PORT=3001
APP_HOST=0.0.0.0
```

> Este setup é reutilizado por **todos os modos de execução** descritos abaixo.

---

### 📋 Pré-requisitos

Dependendo do modo de execução, os pré-requisitos variam:

* **Desenvolvimento local**: Go 1.25+, Make
* **Execução via Docker**: Docker, Docker Compose e Make
* **Produção**: Docker e Make

---

### 🔧 Desenvolvimento Local (com Hot Reload)

Este modo é recomendado para desenvolvimento ativo, utilizando **hot reload** com o `air`.

**Pré-requisitos:** Go 1.25+, Make

1. Instale o Air (uma única vez):

```bash
go install github.com/air-verse/air@latest
```

2. Inicie a aplicação:

```bash
make run-dev
```

A aplicação será recompilada e reiniciada automaticamente a cada alteração no código.

---

### 🧪 Desenvolvimento Local (sem Hot Reload)

Este modo utiliza diretamente o runtime do Go, sem hot reload.

**Pré-requisitos:** Go 1.25+

```bash
go run ./src/cmd/main.go
```

---

### 🐳 Execução Local via Docker

Este modo permite rodar a aplicação **apenas com Docker**, sem necessidade de ter Go instalado na máquina.

**Pré-requisitos:** Docker, Docker Compose e Make

Após realizar o setup inicial do `.env`, execute:

```bash
make build
make run-prod
```

A aplicação será executada em container Docker, utilizando a imagem gerada localmente.

---

### 🚀 Produção

Após realizar o **setup inicial** das variáveis de ambiente, a execução em produção é feita com um único comando:

```bash
make run-prod
```

#### O que o `run-prod` faz

O comando `run-prod` executa as seguintes etapas:

1. Interrompe e remove containers existentes com o mesmo nome
2. Realiza o build da imagem Docker (multi-stage)
3. Executa o container em background
4. Injeta as variáveis de ambiente a partir do arquivo `.env`
5. Expõe a porta configurada da aplicação

Este fluxo garante uma execução **reproduzível e consistente**, alinhada com um ambiente de produção.

---

## 🧠 Arquitetura e Decisões

* **Configuração centralizada**: todas as variáveis de ambiente são carregadas e validadas no startup
* **Fail fast**: a aplicação não inicia se variáveis obrigatórias estiverem ausentes
* **Separação de responsabilidades**:

  * `cmd/` → bootstrap da aplicação
  * `server/` → HTTP, rotas e lifecycle
  * `config/` → leitura e validação de ambiente
* **Ambientes bem definidos**: development vs production

---
