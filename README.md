# Go RESTful API

[![GoDoc](https://godoc.org/github.com/Fox658/api_go_exercise?status.png)](http://godoc.org/github.com/Fox658/api_go_exercise)
[![Code Coverage](https://codecov.io/gh/Fox658/api_go_exercise/branch/master/graph/badge.svg)](https://codecov.io/gh/Fox658/api_go_exercise)
[![Go Report](https://goreportcard.com/badge/github.com/Fox658/api_go_exercise)](https://goreportcard.com/report/github.com/Fox658/api_go_exercise)

This starter kit is designed to get you up and running with a project structure optimized for developing
RESTful API services in Go. It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
It encourages writing clean and idiomatic Go code.

The kit provides the following features right out of the box:

- RESTful endpoints in the widely accepted format
- Environment dependent application configuration management
- Structured logging with contextual information
- Error handling with proper error response generation
- Data validation

The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted.

- Routing: [chi](https://github.com/go-chi/chi)
- Logging: [Logrus](https://github.com/sirupsen/logrus")

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/Fox658/api_go_exercise.git

cd api_go_exercise

# run the RESTful API server
go run cmd\api\main.go

```

At this time, you have a RESTful API server running at `http://127.0.0.1:8000`. It provides the following endpoints:

- `POST /account/coins/?username=`: authenticates a user and generates a JSON

Try the URL `http://localhost:8080/account/coins/?username=username` in postman, and you should see something like `"{"Code":400,"Message":"Invalid username or taken."}"` displayed.

## Project Layout

The starter kit uses the following project layout:

```
.
├── api
│   └── api.go                      the API server structs and erro writing
├── cmd                             main applications of the project
│   └── api
│       └── main.go                 the API server application
├── internal
│   ├── handlers                    handlers for api and get coin balance
│   │   ├── api.go                  package handlers for the http api
│   │   └── get_coin_balance.go     handles coin balance
│   ├── middleware
│   │   └── authorization.go        authentication feature
│   └── tools
│       ├── database.go             database interface
│       └── mockdb.go               mock database
├── go.mod                          tracks modules imported
└── go.sum                          details modules applied
```

The top level directories `cmd`, `internal` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example,

Within each feature package, code are organized in layers (API, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## Common Development Tasks

This section describes some common development tasks using this starter kit.

```

```
