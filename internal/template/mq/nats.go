package mq_template

const NatsConfigTemplate string = `
package mq

import(
	"github.com/spf13/viper"
	"github.com/nats-io/nats.go"
)

func NewNatsConnection() *nats.Conn {
	addr := fmt.Sprintf("%s://%s:%s", viper.GetString("nats.protocol"), viper.GetString("nats.address"), viper.GetString("nats.port"))
	nc, err := nats.Connect(addr,
		nats.UserInfo(viper.GetString("nats.credential.user"), viper.GetString("nats.credential.password")),
		nats.Name(viper.GetString("nats.connetion_name")),
		nats.Timeout(viper.GetDuration("nats.timeout")*time.Second),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(viper.GetDuration("nats.timeout")*time.Second),
	)
	if err != nil {
		panic("failed to connect to nats server")
	}
	return nc
}
`

const DockerComposeNatsConfigTemplate string = `
# nats
  {{.ProjectName}}_mq:
    image: nats:2.11.9
    container_name: {{.ProjectName}}_mq
    restart: unless-stopped
    volumes:
      - ./config/mq/nats-server.conf:/etc/nats/nats.conf
      - {{.ProjectName}}_nats_data:/nats:rw
      - {{.ProjectName}}_nats_jetstream-data:/jetstream/data:rw
    environment:
      - NATS_USER=${NATS_USER}
      - NATS_PASSWORD=${NATS_PASSWORD}
    ports:
      -	"8081:8081"
    command: "-c /etc/nats/nats.conf --name nats -p 4221"
`

const DockerComposeNatsVolumeTemplate string = `
  {{.ProjectName}}_nats_data: {}
  {{.ProjectName}}_nats_jetstream-data: {}`

const NatsConfigFileTemplate string = `
port 4221

authorization {
  user: $NATS_USER
  password: $NATS_PASSWORD
}

jetstream {
  store_dir: /jetstream/data
  max_mem: 1GiB
  max_file: 100GiB
}

websocket {
  port: 8081
  no_tls: true
}
`
