# 🚀 HTTP Server Projeto Korp

Projeto de infraestrutura observável com **Go + Docker + Prometheus + Grafana + Ansible**, simulando um ambiente de produção com monitoramento completo de métricas HTTP.

---

# 📌 Visão Geral

Este projeto implementa:

- API HTTP em Go
- Coleta de métricas com Prometheus
- Dashboards no Grafana
- Infraestrutura containerizada com Docker
- Provisionamento automatizado com Ansible

---

# 🏗️ Arquitetura

Cliente → HTTP Server (Go) → /metrics → Prometheus → Grafana (Dashboard)

---

# ⚙️ Tecnologias utilizadas

- Go (HTTP Server)
- Docker & Docker Compose
- Prometheus
- Grafana
- Ansible
- Nginx (proxy opcional)

---

# 📡 Endpoints da aplicação

## Health check

GET /health

Resposta:
{
  "status": "UP"
}

---

## Endpoint principal

GET /projeto-korp

Resposta:
{
  "nome": "Projeto Korp",
  "horario": "2026-06-01T00:00:00Z"
}

---

## Métricas Prometheus

GET /metrics

Exemplo:

http_requests_total{endpoint="/health"} 3  
http_requests_total{endpoint="/projeto-korp"} 4  

---

# 📊 Dashboard Grafana

O dashboard inclui:

- Total de requests HTTP
- Requests por endpoint
- Requests por segundo (RPS)

---

# 🚀 Como executar o projeto

## 1. Subir ambiente com Ansible

ansible-playbook -i ansible/inventory ansible/playbook.yml --ask-become-pass

---

## 2. Subir manualmente com Docker (opcional)

docker compose up -d

---

# 🌐 Acessos

- HTTP Server: http://localhost:8080  
- Prometheus: http://localhost:9090  
- Grafana: http://localhost:3000 (admin/admin)

---

# 📈 Métricas implementadas

- http_requests_total (Counter)
  - Label: endpoint
  - Conta requisições por rota

---

# 🧠 Conceitos aplicados

- Observabilidade (Metrics + Dashboards)
- Prometheus scraping
- Instrumentação em Go
- Infraestrutura como código (Ansible)
- Containers Docker
- Visualização com Grafana

---

# 🔧 Estrutura do projeto

http-server-projeto-korp/
├── ansible/
├── grafana/
│   ├── dashboards/
│   └── provisioning/
├── prometheus/
├── docker-compose.yml
├── main.go
└── README.md

---

# 👨‍💻 Autor

Hérik Fernandes

---

# 📌 Observações finais

- O dashboard é provisionado automaticamente pelo Grafana
- As métricas são expostas em /metrics
- Toda infraestrutura roda em containers Docker
