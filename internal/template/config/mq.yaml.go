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
# uncomment if you want to create example_stream
# and maybe consumer in your app. for example how to create, go to 
# http://https://github.com/micros-template/notification-service/blob/main/cmd/server/subscriber.go
# or nats official documentation
# jetstream:
  # notification:
    # stream:
      # name: "example_stream"
      # description: "this is example stream"
    # subject:
      # global: "example_stream.>"
      # specific: "example_stream.specific"
`

const RabbitYamlConfigTemplate string = `
rabbitmq:
  protocol: amqp
  credential:
    user: user
    password: password
  address: localhost
  port: 5672
  heartbeat: 10
`

const KafkaYamlConfigTemplate string = `
kafka:
  address: "localhost"
  port: "9092"
  timeout: 10

  auth:
    enabled: true
    username: "app"
    password: "app-secret"
`
