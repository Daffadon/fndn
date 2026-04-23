# fndn

**A CLI scaffolding tool for Go backend projects.**  
Bootstrap your Go projects with clean architecture, best practices, and a solid foundation — all generated in seconds.

---

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/Daffadon/fndn)

## 🚀 Features

- 🏗️ Clean architecture scaffolding
- 🛠️ Customizable set of techstack
- ⚡ Interactive CLI with Bubble Tea
- 🐳 Docker & containerization ready
- 📦 Go modules setup

---

> [!NOTE]
> v0.7.0 now is stable, latest version, and generating one set of application with its driver to 3rd party. in this version, you can generate different framework, database, message queue, in-memory store, and object storage.
## 📋 Prerequisites

Go Version: >= v1.25.0

## 🔥 Get started

```bash
go run github.com/daffadon/fndn@v0.7.0 init .
```

\* . generate in current directory

OR

```bash
go run github.com/daffadon/fndn@v0.7.0 --help
```

to see how can you use the tools

> [!NOTE]
> For The first time, it will take longer than expected to generate depends on go cache, go modcache, and your internet speed.

After the project is generated, you can also generate a config for another techstack provided by using below command.

```bash
go run github.com/daffadon/fndn@v0.7.0 generate [command]
```

*command*: framework, database, mq, cache, storage

## 📦 Installation

If you want to install the tools to your system, you can either using go install or download the binary in available release:

```bash
go install github.com/daffadon/fndn@v0.7.0
```

## 🛠️ The techstack

Currently, the generation are in default mode and custom mode with a freedom to choose `framework`, `database`, `message queue`, `in-memory store`, `object storage` that you need. the `Default` mode is using the first techstack from each section. the generated techstack that you can use are:

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
![Amazon SQS](https://img.shields.io/badge/amazon%20sqs-F79114?style=for-the-badge&logo=amazon%20sqs&logoColor=whit)

- Cache

![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![Valkey](https://img.shields.io/badge/valkey-2D2471?style=for-the-badge&logo=valkey&logoColor=white)
![Dragonfly](https://img.shields.io/badge/Dragonfly-0B0514?style=for-the-badge&logo=dragonfly&logoColor=white)
![Redict](https://img.shields.io/badge/Redict-800000?style=for-the-badge&logo=Redict&logoColor=white)

- Object Storage

![Rustfs](https://img.shields.io/badge/RustFS-0196D0?style=for-the-badge&logo=RustFS&logoColor=white)
![Seaweedfs](https://img.shields.io/badge/Seaweedfs-0059AC?style=for-the-badge&logo=SeaweedFS&logoColor=white)
![Minio](https://img.shields.io/badge/minio-C8324D?style=for-the-badge&logo=nats&logoColor=white)

- Deployment

![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

## 📃 Config Reference

- [All Docker Compose](https://github.com/daffadon/fndn/blob/main/internal/template/common/docker-compose.all.md)
- [Config YAML](https://github.com/daffadon/fndn/blob/main/internal/template/common/all_config.yaml.md)
- [Platform Config File](https://github.com/daffadon/fndn/blob/main/internal/template/common/platform_config_file.md)

## 👌 What fndn do for you (v0.\*)

it's generating the folder structure that use clean architecture as reference. If you're not familiar with the scheme, don't worry, let's talk about it.

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

1. `cmd`: where the command is exist to running the application. there are several folder, which is for bootstraping, dependency injection, construct the server. `main.go` is the entrypoint for all of those.
2. `config`: store all the configs; connection to 3rd party, instantiation of an dependency, configuration for http server, and certificate for tls. furthermore you can add more like grpc server config, log emitter, or any other configuration.
3. `internal`: the place where you put on your app logic business that is not should be exposed. this is special folder for golang cause the module can't be imported from anywhere even when the repository is publicly accessible. [see more](https://go.dev/doc/go1.4#internalpackages)
4. `script`: this is shell script for build the app. there are two scripts, one for build the binary and one for build the docker image.

for the files, there are several files that is generated and you can change for your app:

1. `.air.toml`: Check your repository readme for special notes if it's not working on windows
2. `.env.example`: check your repository readme for what should you do to this file
3. `Dockerfile`: this is generated Dockerfile that use **multistage and distroless**. so in case you want to do something to your containerized app and need a shell, you can change the base image of the second stage.
4. `config.local.yaml`: check your repository readme for what should you do to this file
5. `docker-compose.yml`: this is for production purpose. for development, this file is purposed to run the 3rd party for your app.

## 🔀 how to read the code

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
> All of dependencies are injected in the `cmd/di/di.go`. So, calling to the infra in the `repository/todo.go` is not drawed.

## 🚨 Troubleshoot

### Air is not working on wsl

> [!NOTE]
> If you use windows and generate the project using wsl, the hot reload won't work. better you use the fndn for windows in this case or if its already generated, you can change the .air.toml in `bin and cmd` to become like below and **run air from windows**, not from wsl.
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
> due to limitation, you can't use any database. instead use, `postgres` database. if you find similar log with below log in your postgres db, change the database to `postgres` (i've made this default, but in case you change the database name in docker-compose.yml, change your database).
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
