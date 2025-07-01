# 💡 SOGE - Sistema de Sugestões de Colaboradores

Projeto web simples e básico para o recebimento e gerenciamento de sugestões enviadas por colaboradores da empresa. Neste projeto eu implementei banco de dados, api e frontend.  

---

## 🚀 Tecnologias Utilizadas

- **Frontend:** Vue.js 3
- **Backend:** Go (Golang)
- **Banco de Dados:** PostgreSQL 16
- **Documentação da API:** Swagger
- **Gerenciamento de Containers:** Docker + Docker Compose

---

## 📦 Requisitos para Executar o Projeto

O projeto roda em docker, então é necessário apenas ele para rodar.

- [Docker](https://www.docker.com/)

---

## ▶️ Como Rodar o Projeto Localmente

docker compose up --build

## 🌐 Acesse os serviços no navegador:

| Serviço               | URL                                                                                  |
| --------------------- | ------------------------------------------------------------------------------------ |
| 🌐 Frontend (Vue.js)  | [http://localhost:5173](http://localhost:5173)                                       |
| ⚙️ API (Go)           | [http://localhost:5000](http://localhost:5000)                                       |
| 📘 Swagger (API Docs) | [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html) |


🧑‍💻 Uso do Frontend
🔘 Envio de Sugestão na Tela Inicial
Basta preencher os campos e apertar no botão "Enviar sugestão".

📋 Gerenciar Sugestões
Ao clicar no botão "Gerenciar Sugestões", você será direcionado para a tela de gerenciamento. Nela, é possível:

🔄 Alterar o Status de uma Sugestão
Clique no botão de status (disponível apenas quando a descrição estiver expandida).

📖 Visualizar ou Ocultar a Descrição
Clique sobre a sugestão para expandir ou ocultar sua descrição.

🔍 Filtros Disponíveis
✅ Filtro por Status
Clique em "Filtrar por status" e escolha uma das opções:

Aberto

Em análise

Implementado

Todos

ℹ️ Importante:
Ao escolher "Todos", o sistema irá ordenar automaticamente as sugestões por status, em vez de aplicar um filtro fixo.

🏢 Filtro por Setor
Clique em "Filtrar por setor" para abrir o campo de busca.
Você pode:

✅ Digitar um setor exato para visualizar apenas sugestões desse setor.

✅ Deixar em branco e clicar na lupa para exibir todas as sugestões ordenadas por setor.



