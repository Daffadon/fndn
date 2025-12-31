package config_template

const MinioYamlConfigTemplate string = `
minio:
  host: "localhost"
  port: 9000
  credential:
    user: "ROOTUSER"
    password: "CHANGEME123"
`

const RustfsYamlConfigTemplate string = `
rustfs:
  host: "localhost"
  port: 9000
  credential:
    user: "ROOTUSER"
    password: "CHANGEME123"
`

const SeaweedfsYamlConfigTemplate string = `
seaweedfs:
  host: "localhost"
  port: 9000
  credential:
    user: "ROOTUSER"
    password: "CHANGEME123"
`
