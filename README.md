# 🚀 GoLang Practice Project

This repository serves as a practical exercise for learning and applying the fundamentals of the **Go** (Golang) programming language, focusing specifically on package management and third-party dependency handling.

---

## 🔧 Core Go Commands for Package Management

The following commands were used throughout the development of this project to manage dependencies, initiate modules, and run the code.

| Command | Description | Example Use Case |
| :--- | :--- | :--- |
| `go mod init <module_name>` | **Initializes a new module** (project) and creates a `go.mod` file in the current directory. This is the first step for any new Go project that uses modules. | `go mod init pack` |
| `go get <package_path>` | **Installs a required third-party package** and adds it to the `go.mod` file. | `go get github.com/fatih/color` |
| `go mod tidy` | **Cleans up dependencies** by adding missing dependencies and removing unused ones from the `go.mod` and `go.sum` files. This is a common practice after writing new code or importing packages. | `go mod tidy` |
| `go mod vendor` | **Creates a local copy** of all necessary dependencies inside a `vendor/` directory. This is useful for building reproducible projects without relying on the network at build time. | `go mod vendor` |
| `go run <path_to_main_file>` | **Compiles and runs** the specified main package file. | `go run .\19.packages\main.go` |