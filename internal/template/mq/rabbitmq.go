package mq_template

const RabbitMQConfigTemplate string = `
package mq

import (
	"github.com/spf13/viper"
	amqp "github.com/rabbitmq/amqp091-go"
)
	
func NewRabbitMQConnection() *amqp.Connection {
	addr := fmt.Sprintf("%s://%s:%s@%s:%s/",
		viper.GetString("rabbitmq.protocol"),
		viper.GetString("rabbitmq.credential.user"),
		viper.GetString("rabbitmq.credential.password"),
		viper.GetString("rabbitmq.address"),
		viper.GetString("rabbitmq.port"),
	)
	conn, err := amqp.DialConfig(addr, amqp.Config{
		Heartbeat: viper.GetDuration("rabbitmq.heartbeat") * time.Second,
		Locale:    "en_US",
	})
	if err != nil {
		panic("failed to connect to rabbitmq server")
	}
	return conn
}
`

const DockerComposeRabbitMQConfigTemplate string = `
# rabbitmq
  {{.ProjectName}}_mq:
		image: rabbitmq:4.1.4-alpine
    container_name: {{.ProjectName}}_mq
    ports:
      - "5672:5672"    # RabbitMQ main port
      - "15672:15672"  # Management UI
		volumes:
      - {{.ProjectName}}_rabbitmq_data:/var/lib/rabbitmq
      - ./config/mq/definition.json:/etc/rabbitmq/definitions.json:ro
    environment:
      RABBITMQ_DEFAULT_USER: ${MQ_USER}
      RABBITMQ_DEFAULT_PASS: ${MQ_PASSWORD}
      RABBITMQ_DEFINITIONS_FILE: /etc/rabbitmq/definitions.json
`

const DockerComposeRabbitVolumeTemplate string = `
  {{.ProjectName}}_rabbitmq_data: {}`
const RabbitMQConfigFileTemplate string = `
{
  "users": [
    {
      "name": "user",
      "password": "password",
      "tags": "administrator"
    }
  ],
  "vhosts": [
    {
      "name": "/"
    }
  ],
  "permissions": [
    {
      "user": "guest",
      "vhost": "/",
      "configure": ".*",
      "write": ".*",
      "read": ".*"
    }
  ],
  "queues": [
    {
      "name": "my_queue",
      "vhost": "/",
      "durable": true,
      "auto_delete": false,
      "arguments": {}
    }
  ],
  "exchanges": [
    {
      "name": "my_exchange",
      "vhost": "/",
      "type": "direct",
      "durable": true,
      "auto_delete": false,
      "internal": false,
      "arguments": {}
    }
  ],
  "bindings": [
    {
      "source": "my_exchange",
      "vhost": "/",
      "destination": "my_queue",
      "destination_type": "queue",
      "routing_key": "my_routing_key",
      "arguments": {}
    }
  ]
}
`
