# Go + Gin + 3rd Party + Clean Architecture 🤯

```bash
 __            _
/ _|          | |
| |_ _ __   __| |_ __
|  _| '_ \ / _' | '_ \
| | | | | | (_| | | | |
|_| |_| |_|\__,_|_| |_|
```

This project is generated with [fndn](https://github.com/Daffadon/fndn), the foundation for modern golang project that easy to modify, extend, and learn. to learn more, visit our github at [here](https://github.com/Daffadon/fndn).

## Techstack

Here are generated configuration and 3rd party that can be used immediately.

![](https://img.shields.io/badge/gin-3997AA?style=for-the-badge&logo=gin&logoColor=white)
![](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![](https://img.shields.io/badge/nats-2DACE1?style=for-the-badge&logo=nats&logoColor=white)
![](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![](https://img.shields.io/badge/minio-C8324D?style=for-the-badge&logo=nats&logoColor=white)
![](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

with optional, you can also generate [air](https://github.com/air-verse/air) for hot reload in development when prompted in the **fndn init** command.

## Usage

To start development, you can start by copying the **.env.example** and change the name to **.env**. Make sure that you fill the variable the same with value that you used to connect from your apps in the **config.local.example**. For example:

```bash
ENV="production"

# db env
DB_USER=myusername
DB_PASSWORD=password
DB_NAME=database_name

# nats env
NATS_USER=user
NATS_PASSWORD=password

# redis env
REDIS_PASSWORD=password

# minio env
MINIO_ROOT_USER=ROOTUSER
MINIO_ROOT_PASSWORD=CHANGEME123
```

After its been set, run the third party via docker command in the root of your project

```bash
docker compose up -d
```

OR via Makefile

```bash
make dev-start
```

then, run your app by using air or go run command

```bash
air
```

```bash
go run cmd/main.go
```

OR via Makefile

```bash
make run
```

> [!NOTE]
> If you use windows and generate the project using wsl, the hot reload won't work. better you use the fndn for windows instead of linux in this case or if its already generated, you can change the .air.toml in **cmd part** to become like below and **run air from windows**, not from wsl.
>
> ```yml
> cmd = "go build -o ./tmp/main.exe ./cmd"
> ```

## Production

1. Generate **self-signed certificate** using makefile command

```bash
make cert-gen
```

2. Use file named **config.yaml** for production
3. Change port number to **443** in config.yaml
4. Uncomment your **app service** in docker compose
5. Re-run docker compose up command

```bash
docker compose up -d
```