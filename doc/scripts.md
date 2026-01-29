# Scripts do Projeto

Este documento descreve os scripts auxiliares disponíveis no projeto.

---

## init.sql

### O que é

O arquivo `init.sql` contém o **script de criação das tabelas principais** do banco de dados.

### Quando ele roda

- Executado automaticamente **na primeira inicialização do banco**
- Utilizado pelo container PostgreSQL via Docker
- Executado apenas se o volume do banco ainda não existir

### Objetivo

- Criar a tabela de estilos de cerveja
---

## seed.sh

### O que é

O script `seed.sh` popula o banco de dados com os **estilos de cerveja iniciais**, conforme especificado no desafio.

### Importante

**A aplicação precisa estar rodando** para utilizar o script, pois ele executa chamadas HTTP para a API.

### Como rodar

```bash
chmod +x seed.sh
./seed.sh
```

Ou definindo a URL da API:
```bash
APP_BASE_URL=http://localhost:3001 ./seed.sh
```

**O que o script faz:**
    - Executa requisições POST /beerstyles
    - Cria todos os estilos definidos no desafio
    - Não aborta em caso de erro
    - Exibe o status HTTP de cada criação