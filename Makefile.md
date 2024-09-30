# Makefile Targets Documentation

The `Makefile` provides a set of commands to streamline development tasks for this Go project. Below is a list of available targets with their descriptions.

## How to Use

To execute any of the targets, use the following command:

```bash
make [target]
```

For example, to build the project, run:

```bash
make build
```

## Available Targets

| Target          | Description                                               |
|-----------------|-----------------------------------------------------------|
| `help`          | Display this help message                                  |
| `build`         | Build the Go application                                   |
| `run`           | Build and run the Go application                           |
| `test`          | Run tests                                                  |
| `clean`         | Clean up generated files                                   |
| `generate`      | Generate the Prisma client                                 |
| `migrate-dev`   | Apply Prisma migrations for development (requires `DATABASE_URL`) |
| `migrate-deploy`| Apply Prisma migrations for production (requires `DATABASE_URL`)  |
| `setup`         | Set up the project (download dependencies and generate Prisma client) |
| `start`         | Generate Prisma client and run the application             |
| `env`           | Display current environment variables                      |
| `lint`          | Run linting on the Go code                                 |
| `fmt`           | Format the Go code                                         |
| `check`         | Run formatting, linting, and tests                         |

## Setting Up the Project

To set up the project for the first time, run:

```bash
make setup
```

This command will download all Go dependencies and generate the Prisma client.

## Applying Migrations

### For Development
```bash
make migrate-dev
```

### For Production
```bash
make migrate-deploy
```

Make sure to set `DATABASE_URL` before running migrations:

```bash
export DATABASE_URL="your-database-url"
```

## Code Quality Checks

To run formatting, linting, and tests in one go, use:

```bash
make check
```
