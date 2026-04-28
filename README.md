<h1 align="center">Golang Basics</h1>

<p align="center">
  <strong>Languages:</strong><br>
  <a href="README.pt.md">Portuguese</a> |
  <a href="README.md">English</a>
</p>

Golang Basics is a **Go learning repository** that brings together small hands-on exercises and incremental mini-projects to practice core language concepts step by step.

The repository currently includes:

- introductory Go programs covering syntax, loops, functions, pointers, arrays, slices, maps, structs, interfaces, composition, and error handling
- modularization exercises with local packages and multiple `go.mod` setups
- a calculator example split into packages to practice project organization
- a simple REST API example with HTTP handlers, use cases, repositories, request/response models, and an in-memory data flow

## 🎯 Project Motivation

This repository was built to study Go in a progressive and practical way.

The main goal is to keep a single place for experiments that move from basic syntax to more structured application design, making it easier to revisit concepts and compare simpler exercises with more complete examples.

It is especially useful to practice:

- Go syntax and standard library fundamentals
- value vs pointer behavior
- slices, maps, structs, and interfaces
- package organization and module usage
- separation of concerns in a small REST API

## ✅ Current Content

### 📘 Learning Exercises

- basic "hello world" and build examples
- arithmetic and age calculation exercises
- loops and `for range`
- functions with copy and reference parameters
- pointer practice
- arrays, slices, capacities, subslices, and `make`
- maps and map iteration
- structs, interfaces, and composition
- error handling exercises

### 📦 Module and Package Practice

- `24_mod` for a simple module with internal packages
- `25_calculadora_modular` for a calculator split into packages
- root `go.work` configured to work with multi-module examples

### 🌐 REST API Example

The `26_rest_api` folder contains the most complete example in the repository so far.

Current implementation:

- API bootstrap through `cmd/api`
- simple client request example through `cmd/client`
- `GET /users` endpoint to list users
- `POST /users` endpoint to create users
- request and response models for user creation
- in-memory user repository
- use case layer with duplicate email validation
- UUID generation for created users
- JSON responses for success and error cases

## 🔄 Learning Progression

The repository is organized as a progression rather than a single app.

Typical path:

1. Start with isolated language exercises in numbered folders
2. Move to package and module organization examples
3. Finish with the REST API example that applies those concepts in a more realistic structure

## 🧰 Technologies

- Go 1.26.2
- Go standard library
- `github.com/google/uuid`
- Go modules and `go.work`

## ▶️ How to Run Locally

### 1. 📥 Clone the repository

```bash
git clone https://github.com/pitercoding/golang-basics.git
cd golang-basics
```

### 2. ▶️ Run a basic example

From the repository root, run any numbered example:

```bash
go run ./00_hello
```

You can replace `00_hello` with any other learning folder such as `10_for_range` or `20_map`.

### 3. 🧮 Run the modular calculator example

```bash
go run ./25_calculadora_modular
```

### 4. 🌐 Run the REST API example

Start the API:

```bash
go run ./26_rest_api/cmd/api
```

The API runs on:

- `http://localhost:8080`

### 5. 🧪 Run the REST client example

With the API running in another terminal:

```bash
go run ./26_rest_api/cmd/client
```

### 6. 🗂️ Workspace note

This repository uses a `go.work` file for the multi-module examples. If you add new module folders later, include them in `go.work` when needed.

## 🔌 REST API Notes

Current endpoints in `26_rest_api`:

- `GET /users`
- `POST /users`

Example request body:

```json
{
  "name": "Racha Cuca",
  "email": "rc@test.com"
}
```

Current behavior:

- users are stored in memory
- each new user receives a generated UUID
- duplicate emails are rejected in the use case layer
- responses are encoded as JSON

## 🧪 Testing Status

Current status:

- no automated test suite has been added yet
- the repository is focused on learning and manual experimentation for now

Recommended next test scope:

- add unit tests for the REST API use case layer
- add handler tests for `GET /users` and `POST /users`
- add validation tests for duplicate email scenarios

## 🔮 Next Improvements

### 📦 Learning Content

- add more exercises around goroutines and channels
- add examples for file I/O and JSON parsing
- add more interface and composition challenges

### 🌐 REST API

- persist users in a real database instead of memory
- improve HTTP status handling for business errors
- add request validation for empty name and email
- add update and delete user endpoints
- add tests for handlers, repositories, and use cases

### 🛠️ Tooling

- add linting and formatting workflow
- add CI for building and testing examples
- document each numbered folder with short learning notes

## 📁 Folder Structure

```text
golang-basics/
|-- 00_hello/                      # Introductory example
|-- 01_build/                      # Build and execution basics
|-- 02_sum/                        # Simple arithmetic
|-- 03_age/                        # Variables and calculations
|-- 04_loops/                      # Loop practice
|-- 05_birth_year_challenge/       # Small logic challenge
|-- 06_functions/                  # Functions basics
|-- 07_copy_params/                # Value parameter behavior
|-- 08_ref_params/                 # Reference parameter behavior
|-- 09_ponteiros/                  # Pointer practice
|-- 10_for_range/                  # Range iteration
|-- 11_struct/                     # Struct basics
|-- 12_errors/                     # Error handling practice
|-- 13_calculadora/                # Calculator example
|-- 14_arrays/                     # Arrays
|-- 15_slices/                     # Slices
|-- 16_slices_from_arrays/         # Slices from arrays
|-- 17_slices_cap/                 # Slice capacity
|-- 18_slices_make/                # Slice creation with make
|-- 19_subslices/                  # Subslices
|-- 20_map/                        # Map basics
|-- 21_map_iter/                   # Map iteration
|-- 22_interface_01/               # Interface basics
|-- 23_composition/                # Composition practice
|-- 24_mod/                        # Module and package example
|-- 25_calculadora_modular/        # Modular calculator project
|-- 26_rest_api/                   # REST API learning project
|   |-- cmd/
|   |   |-- api/                   # API entrypoint
|   |   `-- client/                # HTTP client example
|   |-- internal/
|   |   |-- handlers/              # HTTP handlers
|   |   |-- models/                # Request/response and domain models
|   |   |-- repositories/          # Repository abstractions and memory implementation
|   |   `-- usecases/              # Business logic
|   |-- go.mod
|   `-- go.sum
|-- go.mod                         # Root module for simple exercises
|-- go.work                        # Workspace for multi-module examples
|-- LICENSE
|-- README.md                      # Documentation (English)
`-- README.pt.md                   # Documentation (Portuguese)
```

## 📄 License

This project is licensed under the **MIT License**.

## 👤 Author

**Piter Gomes** - Computer Science Student (6th Semester) & Full-Stack Developer

[Email](mailto:piterg.bio@gmail.com) | [LinkedIn](https://www.linkedin.com/in/piter-gomes-4a39281a1/) | [GitHub](https://github.com/pitercoding) | [Portfolio](https://portfolio-pitergomes.vercel.app/)
