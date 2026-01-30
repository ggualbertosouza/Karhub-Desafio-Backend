# Testes de Carga

Este documento descreve os testes de carga realizados na aplicação, incluindo o contexto de execução, metodologia, resultados e interpretação.

---

## Contexto Geral

Os testes de carga foram executados **localmente**, com a aplicação rodando em ambiente **production-like** utilizando **Docker Compose**, e o gerador de carga (**k6**) executado diretamente na máquina host.

Essa abordagem foi escolhida para:
- evitar latência adicional entre containers
- simular o comportamento de clientes reais
- manter o ambiente de testes simples e reproduzível

**Objetivo dos testes**  
Os testes não têm como objetivo medir a latência absoluta máxima da aplicação, mas sim:
- validar comportamento sob carga
- avaliar estabilidade
- confirmar o impacto positivo do cache em memória
- observar previsibilidade da latência

---

## Ferramenta Utilizada

- **k6** — ferramenta de testes de carga amplamente utilizada para APIs HTTP
- Execução local via CLI
- Scripts escritos em JavaScript

---

## 🖥️ Ambiente de Execução

Os testes foram executados no seguinte ambiente:

### Sistema
Kernel: Arch Linux
CPU: AMD Ryzen 5 5500
Ram: 16gb

Docker version 29.2.0
Docker Compose version 5.0.2

---

## Teste de carga constante
Estabelecer um baseline de performance do endpoint principal da aplicação sob carga estável e contínua.

Esse teste responde à pergunta:
`“Como a aplicação se comporta em um cenário de uso normal, com múltiplos clientes acessando simultaneamente?”`

**Endpoint testado**
`GET /beerstyles/temperature`

**Payload**
```json
{
  "temperature": -7
}
```

#### Configuração do teste

Usuários virtuais (VUs): 10
Duração: 30 segundos
Tipo de carga: constante
Intervalo entre requisições: 0.5s
Requisições por usuário: ~2 req/s
Throughput total esperado: ~20 req/s

#### Resultado obtido
```
http_req_duration..............: avg=544.22µs min=299.44µs med=529.35µs max=956.83µs p(90)=664.72µs p(95)=724.63µs
http_req_failed................: 0.00% 0 out of 600
http_reqs......................: 600   19.94 req/s
```

**Latência**
Latência média: ~0.54 ms
p95: ~0.72 ms
Máximo observado: < 1 ms

**Estabilidade**
0% de falhas HTTP
Nenhum timeout
Nenhum erro sob carga constante

**Throughput**
~20 requisições por segundo sustentadas
Throughput alinhado com a configuração do teste