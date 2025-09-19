package config_template

const DockerfileConfigTemplate string = `
FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /app
COPY . .

RUN apk update && apk add upx
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "./bin/dist/{{.ProjectName}}" ./cmd

RUN upx --best --lzma ./bin/dist/{{.ProjectName}}

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/bin/dist/{{.ProjectName}} /
ENTRYPOINT ["/{{.ProjectName}}"]
`

const DockerComposeDefaultConfigTemplate string = `services:`

const DockerComposeVolumeConfigTemplate string = `
volumes:`

const DockerComposeAppConfigTemplate string = `
# uncomment this when you want to try your app in container
# app
  # {{.ProjectName}}:
  #  build:
  #    context: .
  #    dockerfile: Dockerfile
  #  container_name: {{.ProjectName}}
  #  restart: unless-stopped
  #  depends_on:
  #    - {{.ProjectName}}_db
  #    - {{.ProjectName}}_mq
  #    - {{.ProjectName}}_cache
  #    - {{.ProjectName}}_storage
  #  ports:
  #    - "8080:8080"
  #   #make sure to mount the config file to /
  #   #change to config.yaml and ENV "production" for production
  #  volumes:
  #    - ./config.local.yaml:/config.local.yaml
  #  environment:
  #    - ENV=${ENV}
  #  command: "/{{.ProjectName}}"
`
