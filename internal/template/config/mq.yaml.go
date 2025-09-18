package config_template

const NatsYamlConfigTemplate string = `
nats:
  protocol: "nats"
  address: "localhost"
  port: "4221"
  credential:
    user: "user"
    password: "password"
  connection_name: "your_service/app_name"
  timeout: 10
`

const JetstreamConfigTemplate string = `
# uncomment if you want to create example stream and 
# and maybe consumer in your app. for example how to create, go to 
# http://https://github.com/micros-template/notification-service/blob/main/cmd/server/subscriber.go
# for the example or nats official documentation
# jetstream:
  # notification:
    # stream:
      # name: "example_stream"
      # description: "this is example stream"
    # subject:
      # global: "example_stream.>"
      # specific: "example_stream.specific"
`

