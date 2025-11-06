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

## Usage

To start development, you can start by copying the `.env.example` and change the name to `.env`. Make sure that you fill the variable the same with value that you used to connect from your apps in the `config.local.yaml`. For example:

```bash
ENV="production"

# db env
# spesial for neo4j, the db user is neo4j and databasename for database name

DB_USER=myusername
# DB_USER=neo4j
# DB_NAME=databasename
DB_PASSWORD=password
DB_NAME=database_name

# nats env
MQ_USER=user
MQ_PASSWORD=password

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
> If you use windows and generate the project using wsl, the hot reload won't work. better you use the fndn for windows in this case or if its already generated, you can change the .air.toml in `bin and cmd` part to become like below and **run air from windows**, not from wsl.
>
> ```yml
> bin = "./tmp/main.exe"
> cmd = "go build -o ./tmp/main.exe ./cmd"
> ```

## Production

1. Generate **self-signed certificate** using makefile command

```bash
make cert-gen
```

2. Use file named `config.yaml` with the same structure as `config.local.yaml` for production
3. Change port number to **443** in config.yaml
4. Uncomment your **app service** in docker compose
5. Make sure the ENV is **production**
6. Re-run docker compose up command

```bash
docker compose up -d
```
