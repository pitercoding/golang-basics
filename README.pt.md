<h1 align="center">Golang Basics</h1>

<p align="center">
  <strong>Idiomas:</strong><br>
  <a href="README.pt.md">Português</a> |
  <a href="README.md">English</a>
</p>

Golang Basics é um **repositório de aprendizado em Go** que reúne exercícios práticos pequenos e mini-projetos incrementais para estudar os conceitos centrais da linguagem passo a passo.

Atualmente, o repositório inclui:

- programas introdutórios de Go cobrindo sintaxe, loops, funções, ponteiros, arrays, slices, maps, structs, interfaces, composição e tratamento de erros
- exercícios de modularização com pacotes locais e múltiplos `go.mod`
- um exemplo de calculadora separada em pacotes para praticar organização de projeto
- um exemplo simples de API REST com handlers HTTP, use cases, repositories, modelos de request/response e fluxo em memória

## 🎯 Motivação do Projeto

Este repositório foi criado para estudar Go de forma progressiva e prática.

O objetivo principal é manter um único lugar para experimentos que vão da sintaxe básica até um design de aplicação mais estruturado, facilitando revisitar conceitos e comparar exercícios simples com exemplos mais completos.

Ele é especialmente útil para praticar:

- sintaxe de Go e fundamentos da biblioteca padrão
- comportamento de valores vs ponteiros
- slices, maps, structs e interfaces
- organização de pacotes e uso de módulos
- separação de responsabilidades em uma pequena API REST

## ✅ Conteúdo Atual

### 📘 Exercícios de Aprendizado

- exemplos básicos de "hello world" e build
- exercícios de aritmética e cálculo de idade
- loops e `for range`
- funções com parâmetros por cópia e por referência
- prática com ponteiros
- arrays, slices, capacidade, subslices e `make`
- maps e iteração sobre maps
- structs, interfaces e composição
- exercícios de tratamento de erros

### 📦 Prática com Módulos e Pacotes

- `24_mod` para um módulo simples com pacotes internos
- `25_calculadora_modular` para uma calculadora dividida em pacotes
- `go.work` na raiz configurado para trabalhar com exemplos multi-módulo

### 🌐 Exemplo de API REST

A pasta `26_rest_api` contém o exemplo mais completo do repositório até o momento.

Implementação atual:

- bootstrap da API por meio de `cmd/api`
- exemplo simples de client por meio de `cmd/client`
- endpoint `GET /users` para listar usuários
- endpoint `POST /users` para criar usuários
- modelos de request e response para criação de usuário
- repositório de usuários em memória
- camada de use case com validação de e-mail duplicado
- geração de UUID para usuários criados
- respostas JSON para cenários de sucesso e erro

## 🔄 Progressão do Aprendizado

O repositório está organizado como uma progressão, e não como uma aplicação única.

Caminho típico:

1. Começar pelos exercícios isolados nas pastas numeradas
2. Avançar para os exemplos de organização com pacotes e módulos
3. Finalizar com o exemplo de API REST, que aplica esses conceitos em uma estrutura mais próxima de um projeto real

## 🧰 Tecnologias

- Go 1.26.2
- biblioteca padrão do Go
- `github.com/google/uuid`
- Go modules e `go.work`

## ▶️ Como Rodar Localmente

### 1. 📥 Clonar o repositório

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. ▶️ Rodar um exemplo básico

Na raiz do repositório, execute qualquer exemplo numerado:

```bash
go run ./00_hello
```

Você pode substituir `00_hello` por qualquer outra pasta de aprendizado, como `10_for_range` ou `20_map`.

### 3. 🧮 Rodar o exemplo da calculadora modular

```bash
go run ./25_calculadora_modular
```

### 4. 🌐 Rodar o exemplo da API REST

Inicie a API:

```bash
go run ./26_rest_api/cmd/api
```

A API roda em:

- `http://localhost:8080`

### 5. 🧪 Rodar o exemplo do client REST

Com a API rodando em outro terminal:

```bash
go run ./26_rest_api/cmd/client
```

### 6. 🗂️ Observação sobre workspace

Este repositório usa um arquivo `go.work` para os exemplos com múltiplos módulos. Se você adicionar novas pastas com módulo próprio no futuro, inclua-as no `go.work` quando necessário.

## 🔌 Observações da REST API

Endpoints atuais em `26_rest_api`:

- `GET /users`
- `POST /users`

Exemplo de body da requisição:

```json
{
  "name": "Racha Cuca",
  "email": "rc@test.com"
}
```

Comportamento atual:

- os usuários são armazenados em memória
- cada novo usuário recebe um UUID gerado
- e-mails duplicados são rejeitados na camada de use case
- as respostas são serializadas em JSON

## 🧪 Status de Testes

Status atual:

- ainda não existe uma suíte automatizada de testes
- o repositório está focado em aprendizado e experimentação manual por enquanto

Próximo escopo recomendado para testes:

- adicionar testes unitários para a camada de use cases da API REST
- adicionar testes de handlers para `GET /users` e `POST /users`
- adicionar testes de validação para cenários de e-mail duplicado

## 🔮 Próximas Melhorias

### 📦 Conteúdo de Aprendizado

- adicionar mais exercícios com goroutines e channels
- adicionar exemplos de leitura/escrita de arquivos e parsing de JSON
- adicionar mais desafios com interfaces e composição

### 🌐 REST API

- persistir usuários em um banco real em vez de memória
- melhorar o tratamento de status HTTP para erros de negócio
- adicionar validação para nome e e-mail vazios
- adicionar endpoints de atualização e remoção de usuários
- adicionar testes para handlers, repositories e use cases

### 🛠️ Tooling

- adicionar fluxo de lint e formatação
- adicionar CI para build e testes dos exemplos
- documentar cada pasta numerada com pequenas notas de aprendizado

## 📁 Estrutura de Pastas

```text
golang-basics/
|-- 00_hello/                      # Exemplo introdutório
|-- 01_build/                      # Conceitos básicos de build e execução
|-- 02_sum/                        # Aritmética simples
|-- 03_age/                        # Variáveis e cálculos
|-- 04_loops/                      # Prática com loops
|-- 05_birth_year_challenge/       # Pequeno desafio de lógica
|-- 06_functions/                  # Fundamentos de funções
|-- 07_copy_params/                # Comportamento de parâmetros por valor
|-- 08_ref_params/                 # Comportamento de parâmetros por referência
|-- 09_ponteiros/                  # Prática com ponteiros
|-- 10_for_range/                  # Iteração com range
|-- 11_struct/                     # Fundamentos de structs
|-- 12_errors/                     # Prática com tratamento de erros
|-- 13_calculadora/                # Exemplo de calculadora
|-- 14_arrays/                     # Arrays
|-- 15_slices/                     # Slices
|-- 16_slices_from_arrays/         # Slices a partir de arrays
|-- 17_slices_cap/                 # Capacidade de slices
|-- 18_slices_make/                # Criação de slices com make
|-- 19_subslices/                  # Subslices
|-- 20_map/                        # Fundamentos de map
|-- 21_map_iter/                   # Iteração sobre map
|-- 22_interface_01/               # Fundamentos de interface
|-- 23_composition/                # Prática com composição
|-- 24_mod/                        # Exemplo com módulo e pacotes
|-- 25_calculadora_modular/        # Projeto de calculadora modular
|-- 26_rest_api/                   # Projeto de aprendizado com API REST
|   |-- cmd/
|   |   |-- api/                   # Entrypoint da API
|   |   `-- client/                # Exemplo de client HTTP
|   |-- internal/
|   |   |-- handlers/              # Handlers HTTP
|   |   |-- models/                # Modelos de request/response e domínio
|   |   |-- repositories/          # Abstrações de repositório e implementação em memória
|   |   `-- usecases/              # Lógica de negócio
|   |-- go.mod
|   `-- go.sum
|-- go.mod                         # Módulo raiz para exercícios simples
|-- go.work                        # Workspace para exemplos multi-módulo
|-- LICENSE
|-- README.md                      # Documentação (English)
`-- README.pt.md                   # Documentação (Português)
```

## 📄 Licença

Este projeto está licenciado sob a **MIT License**.

## 👤 Autor

**Piter Gomes** - Computer Science Student (6th Semester) & Full-Stack Developer

[Email](mailto:piterg.bio@gmail.com) | [LinkedIn](https://www.linkedin.com/in/piter-gomes-4a39281a1/) | [GitHub](https://github.com/pitercoding) | [Portfolio](https://portfolio-pitergomes.vercel.app/)
