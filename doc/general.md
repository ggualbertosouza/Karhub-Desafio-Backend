# Karhub Backend Challenge — General Overview

## Sobre a Aplicação

Este projeto implementa um **microserviço para gerenciamento de estilos de cerveja**.  
Além do CRUD completo, a aplicação possui um endpoint que, dada uma temperatura, retorna o **estilo de cerveja mais adequado**, seguindo regras de negócio específicas.


---

## Regra de Negócio Principal

Dada uma temperatura de entrada, o sistema seleciona o estilo de cerveja cuja **média entre temperatura mínima e máxima** seja a mais próxima do valor informado.

Critérios:
1. Menor distância entre a média do estilo e a temperatura informada
2. Em caso de empate, ordenação alfabética pelo nome do estilo

---

## ⚠️ Nota sobre Spotify

No momento do desenvolvimento, a **API do Spotify não permitia a criação de novas aplicações** no dashboard oficial, impedindo a obtenção de credenciais válidas.

Por esse motivo:
- A integração com o Spotify **não foi ativada**

---

## Cache em Memória (BeerStyleCache)

Para melhorar a performance do endpoint de seleção por temperatura, foi implementado um **cache em memória específico para os estilos de cerveja**.

### Motivação

- Endpoint `/beerstyles/temperature` é read-heavy
- Estilos de cerveja mudam pouco
- Evitar acesso desnecessário ao banco de dados

### Funcionamento

- No startup da aplicação, os estilos são carregados do banco
- Os dados ficam armazenados em memória como uma lista tipada
- O cálculo do estilo ideal ocorre sem acessar o banco
- O cache pode ser inicializado, repopulado ou invalidado

### Benefícios

- Redução de latência
- Menor carga no banco
- Código simples e explícito
- Complexidade O(n), com n pequeno

---

## Decisões Arquiteturais

- **Separação clara de responsabilidades**
  - `domain` → entidades e regras
  - `handlers` → orquestração
  - `infra` → banco, cache e integrações
- **Fail fast**: aplicação não inicia sem variáveis obrigatórias
- **Configuração por variáveis de ambiente (12-factor app)**

---
