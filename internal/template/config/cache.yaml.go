package config_template

const RedisYamlConfigTemplate string = `
redis:
  address: "localhost"
  port: "6379"
  client_name: "your_service/app_name"
  password: "password"
`

const ValkeyYamlConfigTemplate string = `
valkey:
  address: "localhost"
  port: "6379"
  username: "username"
  password: "password"
`
