package config_template

const ServerYamlConfigTemplate string = `
server:
  cors:
    allow_origins: "http://localhost"
    allow_methods: "GET,POST,PUT,DELETE,OPTIONS,PATCH"
    allow_headers: "Content-Type,Authorization,X-Requested-With,X-CSRF-Token,Accept,Origin,Cache-Control,X-File-Name,X-File-Type,X-File-Size"
    expose_headers: "Content-Length,Content-Range"
    max_age: 86400
    allow_credential: false
`
