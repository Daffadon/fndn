# fndn

**A CLI scaffolding tool for Go backend projects.**
Bootstrap your Go projects with clean architecture, best practices, and a solid foundation — all generated in seconds.

[![Release](https://img.shields.io/github/v/release/Daffadon/fndn)](https://github.com/Daffadon/fndn/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/Daffadon/fndn)](https://go.dev/)
[![Build](https://github.com/Daffadon/fndn/actions/workflows/releaser.yml/badge.svg)](https://github.com/Daffadon/fndn/actions)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/Daffadon/fndn)

> [!NOTE]
> v0.7.0 is the latest stable version. It generates one application set with its drivers for third-party services. You can generate framework, database, message queue, in-memory store, and object storage configs.

## Prerequisites

Go >= v1.26.8

## Installation

Install the tool to your system with `go install`, or download a binary from the [releases page](https://github.com/Daffadon/fndn/releases):

```bash
go install github.com/daffadon/fndn@latest
```

To pin a version:

```bash
go install github.com/daffadon/fndn@v0.7.0
```

## Get started

```bash
go run github.com/daffadon/fndn@v0.7.0 init .
```

\* `.` generates in the current directory.

Or see all commands:

```bash
go run github.com/daffadon/fndn@v0.7.0 --help
```

> [!NOTE]
> The first generation takes longer than expected, depending on the Go build cache, module cache, and internet speed.

After the project is generated, you can add a config for another tech stack with:

```bash
go run github.com/daffadon/fndn@v0.7.0 generate [command]
```

*command*: framework, database, mq, cache, storage

## Features

- Clean architecture scaffolding
- Customizable tech stack
- Interactive CLI with Bubble Tea
- Docker and containerization ready
- Go modules setup

## The tech stack

Generation runs in default mode or custom mode, with the freedom to choose the `framework`, `database`, `message queue`, `in-memory store`, and `object storage` you need. Default mode uses the first tech stack from each section.

<details>
<summary>Supported tech stacks</summary>

- Framework

![Gin](https://img.shields.io/badge/gin-3997AA?style=for-the-badge&logo=gin&logoColor=white)
![Fiber](https://img.shields.io/badge/fiber-00ACD7?style=for-the-badge&logo=fiber)
![Echo](https://img.shields.io/badge/echo-00AFD1?style=for-the-badge&logo=echo)
![Chi](https://img.shields.io/badge/chi-01933F?style=for-the-badge&logo=chi)
![Gorilla/mux](https://img.shields.io/badge/gorilla/mux-939292?style=for-the-badge&logo=gorilla)

- Database

![Postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![MariaDB](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white)
![ClickHouse](https://img.shields.io/badge/ClickHouse-FFCC01?style=for-the-badge&logo=ClickHouse&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![FerretDB](https://img.shields.io/badge/FerretDB-black?style=for-the-badge&logo=ferretdb&logoColor=white)
![Neo4J](https://img.shields.io/badge/Neo4j-008CC1?style=for-the-badge&logo=neo4j&logoColor=white)

- Message Queue

![Nats](https://img.shields.io/badge/nats-2DACE1?style=for-the-badge&logo=nats&logoColor=white)
![RabbitMQ](https://img.shields.io/badge/rabbitmq-%23FF6600.svg?&style=for-the-badge&logo=rabbitmq&logoColor=white)
![Kafka](https://img.shields.io/badge/Apache_Kafka-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white)
![Amazon SQS](https://img.shields.io/badge/amazon%20sqs-F79114?style=for-the-badge&logoColor=white)

- Cache

![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![Valkey](https://img.shields.io/badge/valkey-2D2471?style=for-the-badge&logo=valkey&logoColor=white)
![Dragonfly](https://img.shields.io/badge/Dragonfly-0B0514?style=for-the-badge&logo=dragonfly&logoColor=white)
![Redict](https://img.shields.io/badge/Redict-800000?style=for-the-badge&logo=Redict&logoColor=white)

- Object Storage

![Rustfs](https://img.shields.io/badge/RustFS-0196D0?style=for-the-badge&logoColor=white)
![Seaweedfs](https://img.shields.io/badge/Seaweedfs-0059AC?style=for-the-badge&logoColor=white)
![Minio](https://img.shields.io/badge/minio-C8324D?style=for-the-badge&logo=minio&logoColor=white)

- Deployment

![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

</details>

## Config Reference

- [All Docker Compose](https://github.com/daffadon/fndn/blob/main/internal/template/common/docker-compose.all.md)
- [Config YAML](https://github.com/daffadon/fndn/blob/main/internal/template/common/all_config.yaml.md)
- [Platform Config File](https://github.com/daffadon/fndn/blob/main/internal/template/common/platform_config_file.md)

## What fndn does for you (v0.\*)

It generates a folder structure that uses clean architecture as reference. If you're not familiar with the scheme, don't worry, let's talk about it.

<details>
<summary>Project Folder Structure</summary>

```bash
project
├── cmd
│   ├── bootstrap
│   │   └── bootstrap.go
│   ├── di
│   │   └── container.go
│   ├── server
│   │   └── server.go
│   └── main.go
├── config
│   ├── cache
│   │   └── redis.go
│   ├── env
│   │   └── env.go
│   ├── logger
│   │   └── zerolog.go
│   ├── mq
│   │   ├── nats-server.conf
│   │   └── nats.go
│   ├── router
│   │   └── http.go
│   └── storage
│       ├── minio.go
│       └── postgresql.go
├── internal
│   ├── domain
│   │   ├── dto
│   │   │   └── todo.go
│   │   ├── handler
│   │   │   ├── http.go
│   │   │   └── todo.go
│   │   ├── repository
│   │   │   └── todo.go
│   │   └── service
│   │       └── todo.go
│   ├── infra
│   │   ├── cache
│   │   │   └── redis.go
│   │   ├── mq
│   │   │   └── jetstream_infra.go
│   │   └── storage
│   │       ├── minio.go
│   │       └── querier.go
│   └── pkg
│       └── .gitkeep
├── script
│   ├── build-binary.sh
│   └── docker-build.sh
├── .air.toml
├── .env.example
├── .gitignore
├── Dockerfile
├── Makefile
├── README.md
├── VERSION
├── config.local.yaml
├── docker-compose.yml
├── go.mod
└── go.sum
```

</details>

The folder structure is grouped by its usage:

1. `cmd`: where the commands exist to run the application. There are several folders for bootstrapping, dependency injection, and constructing the server. `main.go` is the entrypoint for all of those.
2. `config`: stores all the configs; connection to 3rd party services, instantiation of dependencies, configuration for the http server, and certificates for tls. Furthermore you can add more like grpc server config, log emitter, or any other configuration.
3. `internal`: the place where you put your app business logic that should not be exposed. This is a special folder for golang because the module can't be imported from anywhere even when the repository is publicly accessible. [see more](https://go.dev/doc/go1.4#internalpackages)
4. `script`: shell scripts to build the app. There are two scripts, one to build the binary and one to build the docker image.

Several generated files you can change for your app:

1. `.air.toml`: Check your repository readme for special notes if it's not working on windows
2. `.env.example`: check your repository readme for what you should do with this file
3. `Dockerfile`: this generated Dockerfile uses **multistage and distroless**. So in case you want to do something to your containerized app and need a shell, you can change the base image of the second stage.
4. `config.local.yaml`: check your repository readme for what you should do with this file
5. `docker-compose.yml`: this is for production purposes. For development, this file is purposed to run the 3rd party services for your app.

## How to read the code

```bash
                                                 |----> config/*.go (except /env)
(cmd)                                            |
main.go -> bootstrap/bootstrap.go -> di/di.go ---|----> internal/domain/*.go (except /dto)
    |                                            |
    |                                            |----> internal/infra/*.go
    |----> server.go
              |
              |(/internal)
              |
              |---> handler/http.go -> handler/todo.go -> service/todo.go -> repository/todo.go
```

> [!NOTE]
> All of the dependencies are injected in the `cmd/di/container.go`. So, calling the infra in the `repository/todo.go` is not shown.

## Troubleshoot

### Air is not working on wsl

> [!NOTE]
> If you use windows and generate the project using wsl, the hot reload won't work. Better you use the fndn for windows in this case or if its already generated, you can change the .air.toml in `bin and cmd` to become like below and **run air from windows**, not from wsl.
>
> ```yml
> bin = "./tmp/main.exe"
> cmd = "go build -o ./tmp/main.exe ./cmd"
> ```

### go run command can't be stopped on wsl

> [!NOTE]
> In windows environment, sometimes go run command can't be stopped. It's because the compatibility. Just use powershell to run the app and don't use the wsl.

### FerretDB is not working as expected

> [!NOTE]
> Due to limitation, you can't use any database. Instead use, `postgres` database. If you find a similar log with the below log in your postgres db, change the database to `postgres` (i've made this default, but in case you change the database name in docker-compose.yml, change your database).
>
> ```
> /usr/local/bin/docker-entrypoint.sh: running /docker-entrypoint-initdb.d/20-install.sql
> psql:/docker-entrypoint-initdb.d/20-install.sql:1: NOTICE: installing required extension "documentdb_core"
> psql:/docker-entrypoint-initdb.d/20-install.sql:1: NOTICE: installing required extension "pg_cron"
> 2025-09-30 06:23:15.653 UTC [76] ERROR: can only create extension in database postgres
> psql:/docker-entrypoint-initdb.d/20-install.sql:1: ERROR: can only create extension in database postgres
> DETAIL: Jobs must be scheduled from the database configured in cron.database_name, since the pg_cron background worker reads job descriptions from this database.
> HINT: Add cron.database_name = 'database_name' in postgresql.conf to use the current database.
> ```
