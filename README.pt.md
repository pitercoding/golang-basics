<h1 align="center">Golang Basics</h1>

<p align="center">
  <strong>Idiomas:</strong><br>
  <a href="README.pt.md">Portugues</a> |
  <a href="README.md">English</a>
</p>

**Golang Basics** é um repositório de estudo para aprender Go de forma prática e progressiva.

O projeto está organizado em três grandes áreas:

- `lessons`: exemplos curtos para aprender os conceitos centrais da linguagem
- `exercises`: prática guiada organizada por nível e tema
- `projects`: aplicações maiores para consolidar o que foi estudado

O objetivo é manter um único lugar para experimentação, repetição e evolução gradual durante os estudos de Golang.

## Sobre o Repositório

Este repositório foi pensado para apoiar o estudo de Go desde os fundamentos até uma organização de código mais estruturada.

Em vez de apresentar apenas exemplos isolados ou apenas aplicações completas, o repositório combina:

- lições focadas em conceitos
- exercícios práticos
- projetos para consolidação

Isso facilita estudar, revisitar assuntos anteriores e evoluir da sintaxe básica para pequenas estruturas mais próximas do mundo real.

## Organização do Repositório

### `lessons`

A pasta `lessons` contém exemplos pequenos e diretos, cada um focado em um conceito por vez.

Os temas atuais incluem:

- sintaxe básica e execução
- loops e funções
- ponteiros e tratamento de erros
- arrays, slices e maps
- structs, interfaces e composição

### `exercises`

A pasta `exercises` é a área de prática do repositório.

Ela está organizada por nível de aprendizado e por tema, com cada exercício vivendo em sua própria pasta e arquivo `main.go`.

O roadmap atual de exercícios está documentado em `exercises/go-exercises.md`.

### `projects`

A pasta `projects` reúne exemplos mais amplos, que conectam vários conceitos de Go dentro do mesmo código.

Essa área é útil para praticar:

- separação de pacotes
- organização por módulos
- reutilização de código
- estrutura de aplicação

## Fluxo de Estudo Sugerido

Se você estiver usando este repositório como trilha de aprendizado, uma ordem simples é:

1. Começar por `lessons`
2. Reforçar o tema em `exercises`
3. Aplicar os conceitos em `projects`

Essa estrutura ajuda a transformar teoria em repetição e repetição em implementação.

## Como Rodar

### 1. Clonar o repositório

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. Rodar uma lição

```bash
go run ./lessons/00_hello
```

### 3. Rodar um exercício

```bash
go run ./exercises/level1/01_setup_and_language_basics/001_hello_world
```

### 4. Rodar um projeto

```bash
go run ./projects/01_calculadora
```

Algumas pastas em `projects` usam o próprio `go.mod`, então também é possível executá-las de dentro do diretório específico quando necessário.

## Árvore do Projeto

```text
golang-basics/
|-- exercises/
|   |-- go-exercises.md
|   `-- level1/
|       |-- 01_setup_and_language_basics/
|       |   |-- 001_hello_world/
|       |   |-- 002_print_name_age/
|       |   |-- 003_variables_declaration/
|       |   |-- 004_read_user_input/
|       |   |-- ...
|       `-- 02_string_basics/
|           |-- 001_concatenate_strings/
|           |-- 002_reverse_string/
|           |-- 003_count_vowels/
|           |-- 004_palindrome_check/
|           |-- ...
|-- lessons/
|   |-- 00_hello/
|   |-- 01_build/
|   |-- 02_sum/
|   |-- 03_age/
|   |-- 04_loops/
|   |-- 05_birth_year_challenge/
|   |-- 06_functions/
|   |-- 07_copy_params/
|   |-- 08_ref_params/
|   |-- 09_ponteiros/
|   |-- 10_for_range/
|   |-- 11_struct/
|   |-- 12_errors/
|   |-- 14_arrays/
|   |-- 15_slices/
|   |-- 16_slices_from_arrays/
|   |-- 17_slices_cap/
|   |-- 18_slices_make/
|   |-- 19_subslices/
|   |-- 20_map/
|   |-- 21_map_iter/
|   |-- 22_interface_01/
|   `-- 23_composition/
|-- projects/
|   |-- 01_calculadora/
|   |-- 02_mod/
|   |   |-- operacao/
|   |   `-- saudacao/
|   |-- 03_calculadora_modular/
|   |   |-- operacao/
|   |   `-- runner/
|   `-- 04_rest_api/
|       |-- cmd/
|       |   |-- api/
|       |   `-- client/
|       |-- internal/
|       |   |-- handlers/
|       |   |-- models/
|       |   |-- repositories/
|       |   |   `-- users/
|       |   `-- usecases/
|       `-- pkg/
|-- go.mod
|-- LICENSE
|-- README.md
`-- README.pt.md
```

## Licença

Este projeto está licenciado sob a **MIT License**.

## Autor

**Piter Gomes** - Computer Science Student (6th Semester) & Full-Stack Developer

[Email](mailto:piterg.bio@gmail.com) | [LinkedIn](https://www.linkedin.com/in/piter-gomes-4a39281a1/) | [GitHub](https://github.com/pitercoding) | [Portfolio](https://portfolio-pitergomes.vercel.app/)
