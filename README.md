<h1 align="center">Golang Basics</h1>

<p align="center">
  <strong>Languages:</strong><br>
  <a href="README.pt.md">Portuguese</a> |
  <a href="README.md">English</a>
</p>

Golang Basics is a study repository for learning Go in a practical and progressive way.

The project is organized around three main areas:

- `lessons`: short examples to learn core language concepts
- `exercises`: guided practice grouped by level and topic
- `projects`: larger applications used to apply what was studied

The goal is to keep a single place for experimentation, repetition, and gradual progress while studying Golang.

## About This Repository

This repository is meant to support Go study from fundamentals to more structured code organization.

Instead of presenting only isolated examples or only full applications, the repository combines:

- concept-focused lessons
- hands-on exercises
- practical projects for consolidation

That makes it easier to study, revisit previous topics, and evolve from syntax basics to small real-world structures.

## Repository Organization

### `lessons`

The `lessons` folder contains small, direct examples focused on one concept at a time.

Topics currently include:

- basic syntax and execution
- loops and functions
- pointers and error handling
- arrays, slices, and maps
- structs, interfaces, and composition

### `exercises`

The `exercises` folder is the practice area of the repository.

It is organized by learning level and topic, with each exercise living in its own folder and `main.go` file.

The current roadmap is documented in `exercises/go-exercises.md`.

### `projects`

The `projects` folder groups broader examples that connect multiple Go concepts in the same codebase.

This area is useful for practicing:

- package separation
- module organization
- code reuse
- application structure

## Suggested Study Flow

If you are using this repository as a learning path, a simple order is:

1. Start with `lessons`
2. Reinforce the topic in `exercises`
3. Apply the concepts in `projects`

This structure helps turn theory into repetition and repetition into implementation.

## How To Run

### 1. Clone the repository

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. Run a lesson

```bash
go run ./lessons/00_hello
```

### 3. Run an exercise

```bash
go run ./exercises/level1/01_setup_and_language_basics/001_hello_world
```

### 4. Run a project

```bash
go run ./projects/01_calculadora
```

Some project folders use their own `go.mod`, so you can also run them from inside the specific directory when needed.

## Project Tree

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

## License

This project is licensed under the **MIT License**.

## Author

**Piter Gomes** - Computer Science Student (6th Semester) & Full-Stack Developer

[Email](mailto:piterg.bio@gmail.com) | [LinkedIn](https://www.linkedin.com/in/piter-gomes-4a39281a1/) | [GitHub](https://github.com/pitercoding) | [Portfolio](https://portfolio-pitergomes.vercel.app/)
