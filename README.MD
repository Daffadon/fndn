# fndn

**A CLI scaffolding tool for Go backend projects.**  
Bootstrap your Go projects with clean architecture, best practices, and a solid foundation — all generated in seconds.

---

## 🚀 Features
<!-- - 🔧 Configurable tech stack options -->

- 🏗️ Clean architecture scaffolding
- 🖧 Organized project structure
- ⚡ Interactive CLI with Bubble Tea
- 🐳 Docker & containerization ready
- 📦 Go modules setup

---

## 🔥 Get started

```bash
go run github.com/daffadon/fndn@0.0.1 init .
```

\* . generate in current directory

OR

```bash
go run github.com/daffadon/fndn@0.0.1 --help
```

to see how can you use the tools

## 📦 Installation

If you want to install the tools to your system, you can either using go install or download the binary in available release:

```bash
go install github.com/your-username/fndn@latest
```

## 🛠️ Default techstack

currently, the generation is in default mode (later will be customizable). the generated techstack are:

![](https://img.shields.io/badge/gin-3997AA?style=for-the-badge&logo=gin&logoColor=white)
![](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![](https://img.shields.io/badge/nats-2DACE1?style=for-the-badge&logo=nats&logoColor=white)
![](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![](https://img.shields.io/badge/minio-C8324D?style=for-the-badge&logo=nats&logoColor=white)
![](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

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
4. `script`: this is shell script for build the app. there are thow scripts, one for build the binary and one for build the docker image.

for the files, there are several files that is generated and you can change for your app:

1. `.air.toml`: Check your repository readme for special notes if it's not working on windows
2. `.env.example`: check your repository readme for what should you do to this file
3. `Dockerfile`: this is generated Dockerfile that use **multistage and distroless**. so in case you want to do something to your containerized app and need a shell, you can change the base image of the second stage.
4. `config.local.yaml`: check your repository readme for what should you do to this file
5. `docker-compose.yml`: this is for production purpose. for development, this file is purposed to run the 3rd party for your app.

## 🔀 how to the code

```bash
                                                 |----> config/*.go (except /env)
(cmd)                                            |
main.go -> bootstrap/bootstrap.go -> di/di.go ---|
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