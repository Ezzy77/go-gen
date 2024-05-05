# go-gen CLI Tool

## Overview

`go-gen` is a CLI tool for generating Restful API backends in Golang. It provides a convenient way to create a new Golang project with options to choose preferred routing frameworks.

## Installation

To install `go-gen`, you need to have Go installed on your system. Then, run the following command:

```bash
go install github.com/Ezzy77/go-gen@latest
```

## Usage

Once installed, you can use the `go-gen` command to create a new Golang project. Here's how to use it:

```bash
go-gen create --name [project_name]
```

Replace `[project_name]` with the name of your project.

## Options

- `--name, -n`: Specifies the name of the project to be created.

## Interactive Mode

If you run the `go-gen create` command without specifying the project name, it will prompt you to enter the project name and select options for the routing framework and database.

```bash
go-gen create
```

Follow the prompts to select the desired options.

## Example

```bash
go-gen create --name myproject
```

This command will create a new directory named `myproject` in your home directory, and generate a Golang project with the selected options.

## Selected Options

After selecting the routing framework and database, `go-gen` will print the selected options:

- Routing Framework: [Selected Routing Framework]
- Database: [Selected Database]

## Generated Files

`go-gen` generates the following files and directories in your project directory:

- `main.go`: Entry point of your application.
- `handlers/`: Directory containing HTTP request handlers.
- `models/`: Directory for database models (if selected).
- `utils/`: Directory for utility functions.
- `go.mod` and `go.sum`: Go module files.

## Dependencies
