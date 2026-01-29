# Rodando a aplicação

Este documento descreve como rodar a aplicação em diferentes ambientes.

---

## Configuração Inicial

1. Crie o arquivo `.env` a partir do exemplo:

```bash
cp .env.example .env
```

2. Ajuste as variáveis:

### Configuração das Variáveis de Ambiente
#### Desenvolvimento Local (fora do Docker)

Quando a aplicação é executada diretamente na máquina local, o banco de dados é acessado via localhost:

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

**Por quê?**
Neste cenário, a aplicação roda fora de containers e se conecta ao banco utilizando a interface de loopback da máquina local.

#### Produção / Docker Compose
Quando a aplicação é executada via Docker Compose, a comunicação entre os serviços ocorre através da network interna do Docker.

Nesse caso, o host do banco não deve ser localhost, e sim o nome do serviço definido no compose.prod.yml.

```bash
APP_ENV=production
APP_PORT=3001
APP_HOST=0.0.0.0

DB_HOST=postgres
DB_PORT=5432
DB_NAME=beershop
DB_USER=beershop
DB_PASSWORD=beershop
```

**Por quê?**
Dentro do Docker, cada container possui seu próprio localhost.
O nome postgres é resolvido automaticamente pelo Docker DNS e aponta para o container do banco.

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

Configure corretamente o .env para produção.
Certifique-se de que a variável do banco está definida como:
```bash
DB_HOST=postgres
```

Suba o ambiente de produção
```bash
make prod-up
```

Esse comando irá:

1. Subir o banco de dados PostgreSQL
2. Executar o script init.sql na primeira inicialização
3. Buildar a imagem da aplicação
4. Subir a API já conectada ao banco

#### Popular os estilos de cerveja

Após a aplicação estar rodando, execute o script de seed:
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