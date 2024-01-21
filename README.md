# XSIS ACADEMY TEST Backend Developer
## _Service Movie Documentation

## Features

- List Movie
- Detail Movie
- Add Movie
- Update Movie
- Delete Movie

## Tech & Dependencies

Technology used to build servicesy:

- [Golang 1.20] - Golang Fiber Framework
- [Visual Studio Code] - Visual Studio Code
- [API Documentation] - Swagger Open API 3.0
- [MariaDB] - MariaDB Latest Version
- [DilingerIO] - ReadMe Tools
- [Redis] - Redis 7.0.12

## Installation

Install the dependencies and dev Dependencies and start the server.

```sh
git clone https://github.com/reynaldojosuatogatorop/xsis-academy-test-service-movie.git
cd xsis-academy-test-service-movie
go mod tidy
Open config.yaml then adjust the Database settings, Path Migrate, Asset URL according to your device settings
go run app/main.go -c config.yaml
```