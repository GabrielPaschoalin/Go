# Go - Curso Udemy

Repositório com meus estudos e projetos práticos do curso de Go (Golang) na Udemy. Cada pasta numerada representa um módulo do curso, evoluindo dos fundamentos da linguagem até uma API REST completa.

## Estrutura do repositório

```
1 - Basic/                     Fundamentos da linguagem (arrays, maps, structs, ponteiros, etc.)
2 - functions-deepDive/        Funções: anônimas, recursão, variádicas, funções como valores
3 - Price calculator - project/ Projeto prático de calculadora de preços com I/O de arquivos
4 - Concurrency/                Goroutines e channels
5 - API REST/                   API REST completa com autenticação, em Go + Gin + SQLite
```

### 1 - Basic

Exercícios sobre os conceitos básicos da linguagem:

- `Array` — arrays, slices e maps
- `bank` — simulação simples de conta bancária
- `custom-type` — tipos customizados
- `investment_calculator` / `profit_calculator` — pequenas calculadoras de linha de comando
- `notes-project` — CRUD de notas em memória/arquivo
- `pointers` — ponteiros
- `struct` — structs e métodos
- `Utils` / `modulo07-pratica` — exercícios diversos

### 2 - functions-deepDive

Aprofundamento em funções: funções anônimas, funções como valores (`functinAsValue`), recursão e funções variádicas (`Variadic.go`).

### 3 - Price calculator - project

Projeto que lê preços de um arquivo texto, aplica diferentes taxas de imposto em paralelo (uma goroutine por taxa) e grava o resultado em arquivos JSON de saída (`files/outputPrice_*.json`).

### 4 - Concurrency

Exemplos de concorrência em Go usando goroutines, channels e `select`.

### 5 - API REST

API REST construída com [Gin](https://github.com/gin-gonic/gin) e SQLite (`mattn/go-sqlite3`), com autenticação via JWT.

**Funcionalidades:**

- Cadastro e login de usuários (`/signup`, `/login`) com senha hasheada e token JWT
- CRUD de eventos (`/events`), com rotas de escrita protegidas por autenticação
- Inscrição/cancelamento de inscrição de usuários em eventos

**Rotas:**

| Método | Rota                         | Autenticação | Descrição                        |
|--------|------------------------------|:------------:|-----------------------------------|
| GET    | `/events`                    | não          | Lista todos os eventos            |
| GET    | `/events/:id`                | não          | Detalhes de um evento             |
| POST   | `/events`                    | sim          | Cria um evento                    |
| PUT    | `/events/:id`                | sim          | Atualiza um evento (só o criador) |
| DELETE | `/events/:id`                | sim          | Remove um evento (só o criador)   |
| POST   | `/events/:id/register`       | sim          | Inscreve o usuário no evento      |
| DELETE | `/events/:id/register`       | sim          | Cancela a inscrição do usuário    |
| POST   | `/signup`                    | não          | Cria um novo usuário              |
| POST   | `/login`                     | não          | Autentica e retorna um token JWT  |

**Como rodar:**

```bash
cd "5 - API REST"
go run main.go
```

O servidor sobe em `http://localhost:8080`. O arquivo `api-test/` contém requisições `.http` prontas (compatíveis com a extensão REST Client do VS Code) para testar signup, login, criação e remoção de eventos.

## Requisitos

- Go 1.26+ (algumas pastas usam módulos com versões diferentes do Go, verifique o `go.mod` de cada projeto)

## Como usar

Cada pasta é um módulo Go independente. Entre na pasta desejada e rode com `go run`:

```bash
cd "<pasta do módulo>"
go run .
```

---

Projeto de estudos pessoal, sem garantias de uso em produção.
