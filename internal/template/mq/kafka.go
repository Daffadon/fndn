package mq_template

const KafkaConfigTemplate string = `
package mq

import(
	"github.com/spf13/viper"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func NewKafkaConnection() *kafka.Conn {
	// Build Kafka broker address
	addr := fmt.Sprintf("%s:%s", viper.GetString("kafka.address"), viper.GetString("kafka.port"))

	// Create SASL mechanism if auth is enabled
	var dialer *kafka.Dialer
	if viper.GetBool("kafka.auth.enabled") {
		mechanism := plain.Mechanism{
			Username: viper.GetString("kafka.auth.username"),
			Password: viper.GetString("kafka.auth.password"),
		}

		dialer = &kafka.Dialer{
			Timeout:       viper.GetDuration("kafka.timeout") * time.Second,
			SASLMechanism: mechanism,
		}
	} else {
		dialer = &kafka.Dialer{
			Timeout: viper.GetDuration("kafka.timeout") * time.Second,
		}
	}

	// Try connecting to the broker
	conn, err := dialer.DialContext(context.Background(), "tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Kafka broker: %v", err))
	}

	return conn
}
`

const DockerComposeKafkaConfigTemplate string = `
# kafka
  {{.ProjectName}}_mq:
    image: apache/kafka:3.9.1
    container_name: {{.ProjectName}}_mq
    ports:
      - "9092:9092"
    environment:
      # --- KRaft mode ---
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: broker,controller
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@{{.ProjectName}}_mq:9093
      CLUSTER_ID: secure-cluster-id

      # --- Listeners ---
      KAFKA_LISTENERS: SASL_PLAINTEXT://:9092,CONTROLLER://:9093
      KAFKA_ADVERTISED_LISTENERS: SASL_PLAINTEXT://localhost:9092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,SASL_PLAINTEXT:SASL_PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: SASL_PLAINTEXT

      # --- SASL/PLAIN Authentication ---
      KAFKA_SASL_ENABLED_MECHANISMS: PLAIN
      KAFKA_SASL_MECHANISM_INTER_BROKER_PROTOCOL: PLAIN
      KAFKA_OPTS: "-Djava.security.auth.login.config=/etc/kafka/jaas.conf"

      # --- Misc ---
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"

    volumes:
      - {{.ProjectName}}_kafka_data:/var/lib/kafka/data
      - ./config/mq/jaas.conf:/etc/kafka/jaas.conf
`

const DockerComposeKafkaVolumeTemplate string = `
  {{.ProjectName}}_kafka_data: {}`

const KafkaConfigFileTemplate string = `
KafkaServer {
  org.apache.kafka.common.security.plain.PlainLoginModule required
  username="admin"
  password="admin-secret"
  user_admin="admin-secret"
  user_app="app-secret";
};
`
