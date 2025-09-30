package config_template

const PostresqlYamlConfigTemplate string = `
database:
  sql:
    protocol: "postgresql"
    host: "localhost"
    user: "myusername"
    password: "password"
    port: "5432"
    name: "database_name"
    sslmode: "disable"
`

const MariaDBYamlConfigTemplate string = `
database:
  sql:
    host: "localhost"
    user: "myusername"
    password: "password"
    port: "3306"
    name: "database_name"
`

const MongoDBYamlConigTemplate string = `
database:
  nosql:
    protocol: "mongodb"
    host: "localhost"
    user: "myusername"
    password: "password"
    port: "27017"
    name: "database_name"
`
