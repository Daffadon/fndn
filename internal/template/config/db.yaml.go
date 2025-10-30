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
const ClickHouseYamlConfigTemplate string = `
database:
  sql:
    host: "localhost"
    user: "myusername"
    password: "password"
    port: "9002"
    name: "database_name"
    debug: true
`

const MongoDBYamlConfigTemplate string = `
database:
  nosql:
    protocol: "mongodb"
    host: "localhost"
    user: "myusername"
    password: "password"
    port: "27017"
    name: "database_name"
`

const FerretDBYamlConfigTemplate string = `
database:
  nosql:
    protocol: "mongodb"
    host: "localhost"
    port: "27017"
`

const Neo4JYamlConfigTemplate string = `
database:
  graph:
    protocol: "bolt"
    host: "localhost"
    port: "7687"
    user: "neo4j"
    password: "password"
`
