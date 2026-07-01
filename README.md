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

The current exercises and projects roadmap is documented in `go-exercises-and-projects-roadmap.md`.

### `projects`

The `projects` folder groups broader examples that connect multiple Go concepts in the same codebase.

Projects are organized by study level from `level-1` to `level-6`, with older experiments preserved in `projects/legacy`.

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

All lessons, exercises, and most study projects are small Go programs with their own `main.go`.

### 1. Clone the repository

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. Run a lesson from the root module

```bash
go run ./lessons/00_hello
```

### 3. Run an exercise by level and topic

Exercises follow the roadmap in `go-exercises-and-projects-roadmap.md`:

```text
exercises/<level>/<topic>/<exercise_folder>
```

Example:

```bash
go run ./exercises/level-1-fundamentals/1.1-basics/001_hello_world
```

### 4. Run a project by level

Projects that match `go-exercises-and-projects-roadmap.md` are grouped by study level:

```bash
go run ./projects/level-1/01_hello_user_cli
```

Legacy projects are kept in `projects/legacy`. Some of them have their own `go.mod`, so run those from inside the specific project directory when needed:

```bash
cd projects/legacy/04_rest_api
go run ./cmd/api
```

## Project Tree

```text
golang-basics/
|-- go-exercises-and-projects-roadmap.md
|-- exercises/
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
|   |-- level-4-data-structures-and-algorithms/
|   |   |-- 4.1-advanced-slices/
|   |   |-- 4.2-advanced-maps/
|   |   `-- 4.3-algorithms/
|   `-- level-5-concurrency/
|       `-- 5.1-goroutines-and-channels/
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
|   |-- level-4/
|   |   |-- 01_tic_tac_toe_game/
|   |   |-- 02_contact_manager_with_search/
|   |   `-- 03_text_analyzer/
|   |-- level-5/
|   |   |-- 01_concurrent_file_processor/
|   |   |-- 02_worker_pool_task_runner/
|   |   `-- 03_concurrent_log_analyzer/
|   |-- level-6/
|   |   |-- 01_basic_http_server_api/
|   |   |-- 02_rest_api_crud/
|   |   `-- 03_user_authentication_api/
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

## License

This project is licensed under the **MIT License**.

## Author

**Piter Gomes** - Computer Science Student (6th Semester) & Full-Stack Developer

[Email](mailto:piterg.bio@gmail.com) | [LinkedIn](https://www.linkedin.com/in/piter-gomes-4a39281a1/) | [GitHub](https://github.com/pitercoding) | [Portfolio](https://portfolio-pitergomes.vercel.app/)
