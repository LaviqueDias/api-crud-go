# 📦 API CRUD GO

API RESTful desenvolvida em Go para operações de CRUD de usuários, com autenticação JWT, Swagger para documentação e conexão com banco de dados MySQL.

## 📚 Descrição

Este projeto é um exemplo de implementação de uma API estruturada em Golang, com as seguintes funcionalidades principais:

- Cadastro, login e autenticação de usuários
- Proteção de rotas com middleware JWT
- Persistência de dados com MySQL
- Estrutura modular com separação por camadas (controller, service, repository)
- Documentação interativa com Swagger

## 🚀 Tecnologias Utilizadas

- **Golang 1.21+**
- **Gin Gonic** — Framework web leve e rápido
- **MySQL** — Banco de dados relacional
- **JWT** — Autenticação baseada em tokens
- **Swaggo** — Geração de documentação Swagger automática
- **Zap Logger** — Logging estruturado
- **godotenv** — Gerenciamento de variáveis de ambiente

## 📁 Estrutura do Projeto

## ⚙️ Como Rodar o Projeto

### Pré-requisitos

- Go instalado (https://golang.org/dl/)
- MySQL em execução (ex: porta `3307`)
- `migrate` instalado para rodar as migrations:

```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Adicione o binary ao PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## 1. Clone o repositório

```bash
git clone https://github.com/seuusuario/api-crud-go.git
cd api-crud-go
```

## 2. Configure o `.env`

Crie um arquivo `.env` com o seguinte conteúdo:

```bash
DB_USER=seu_user
DB_PASSWORD=sua_senha
DB_HOST=localhost
DB_PORT=sua_porta
DB_NAME=nome_do_banco

SECRET=sua_chave_secreta
```

## 3. Rode as Migrations

```bash
migrate -path ./migrations -database "mysql://user:pass@tcp(host:port)/dbname" up
```

## 4. Inicie a aplicação

```bash
go run .
```

> A API estará disponível em:
> 📍 `http://localhost:8080`

> A documentação Swagger:
> 📄 `http://localhost:8080/swagger/index.html`

## 🧪 Endpoints

### 🔓 Rotas Públicas

- `POST /api/user/login` — Login e geração do token JWT via cookie

### 🔒 Rotas Privadas

As rotas protegidas devem conter o cookie `Authorization` com o JWT gerado no login.

- `POST /api/user/` — Cadastro de usuário

- `GET /api/user/` — Buscar todos os usuários

- `GET /api/user/email/:userEmail` — Buscar usuário por e-mail

- `GET /api/user/validate` — Valida o token enviado no cookie da requisição

- `PUT /api/user/:userId` — Atualizar dados de um usuário específico

- `DELETE /api/user/:userId` — Remover usuário por ID

- `POST /api/user/logout` — Logout do usuário (remove o cookie Authorization)

## 🧪 Testes e Navegação

Você pode testar os endpoints utilizando:

- 🧪 [Postman](https://www.postman.com/)
- 🔍 [Insomnia](https://insomnia.rest/)
- 📄 Swagger UI: [`http://localhost:8080/swagger/index.html`](http://localhost:8080/swagger/index.html)

## 👨‍💻 Autor

Desenvolvido por Lavique Dias  
📧 lavidoliveira@gmail.com  
💼 [LinkedIn](https://www.linkedin.com/in/lavique-dias-5852572b4/) — [GitHub](https://github.com/LaviqueDias)
