package config_template

const PostresqlYamlConfigTemplate string = `
database:
  host: "localhost"
  user: "myusername"
  password: "password"
  port: "5432"
  name: "database_name"
  sslmode: "disable"
`
