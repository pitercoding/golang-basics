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

Todas as lições, exercícios e a maioria dos projetos de estudo são pequenos programas Go com seu próprio `main.go`.

### 1. Clonar o repositório

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. Rodar uma lição a partir do módulo principal

```bash
go run ./lessons/00_hello
```

### 3. Rodar um exercício por nível e tema

Os exercícios seguem o roadmap em `exercises/go-exercises.md`:

```text
exercises/<level>/<topic>/<exercise_folder>
```

Exemplo:

```bash
go run ./exercises/level-1-fundamentals/1.1-basics/001_hello_world
```

### 4. Rodar um projeto por nível

Os projetos alinhados ao roadmap estão agrupados por nível de estudo:

```bash
go run ./projects/level-1/01_hello_user_cli
```

Os projetos antigos ficam preservados em `projects/legacy`. Alguns deles usam o próprio `go.mod`, então rode esses casos de dentro do diretório específico quando necessário:

```bash
cd projects/legacy/04_rest_api
go run ./cmd/api
```

## Árvore do Projeto

```text
golang-basics/
|-- exercises/
|   |-- go-exercises.md
|   |-- level-1-fundamentals/
|   |   |-- 1.1-basics/
|   |   |-- 1.2-strings/
|   |   |-- 1.3-math/
|   |   |-- 1.4-conditionals/
|   |   `-- 1.5-slices/
|   |-- level-2-flow-control/
|   |   |-- 2.1-loops/
|   |   |-- 2.2-slices-intermediate/
|   |   `-- 2.3-maps/
|   |-- level-3-core-go/
|   |   |-- 3.1-functions/
|   |   |-- 3.2-structs/
|   |   `-- 3.3-structured-logic/
|   `-- level-4-data-structures-and-algorithms/
|       `-- 4.1-advanced-slices/
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
|   |-- level-1/
|   |   |-- 01_hello_user_cli/
|   |   |-- 02_age_calculator/
|   |   |-- 03_unit_converter/
|   |   |-- 04_grade_calculator/
|   |   `-- 05_bmi_calculator/
|   |-- level-2/
|   |   |-- 01_cli_calculator/
|   |   |-- 02_terminal_todo_list/
|   |   `-- 03_password_generator/
|   |-- level-3/
|   |   |-- 01_file_reader_application/
|   |   |-- 02_simple_login_system/
|   |   |-- 03_in_memory_crud_system/
|   |   |-- 04_inventory_management_system/
|   |   `-- 05_banking_system_simulator/
|   `-- legacy/
|       |-- 01_calculadora/
|       |-- 02_mod/
|       |-- 03_calculadora_modular/
|       `-- 04_rest_api/
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
