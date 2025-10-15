package infra_template

const JetstreamInfraTemplate string = `
package mq_infra

import (
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
) 

type (
	JetstreamInfra interface {
		CreateOrUpdateNewConsumer(ctx context.Context, streamName string, jsConfig *jetstream.ConsumerConfig) (jetstream.Consumer, error)
		CreateOrUpdateNewStream(ctx context.Context, jsConfig *jetstream.StreamConfig) error
		Publish(ctx context.Context, subject string, payload []byte) (*jetstream.PubAck, error)
	}
	jetstreamInfra struct {
		js     jetstream.JetStream
		logger zerolog.Logger
	}
)

func NewJetstreamInfra(logger zerolog.Logger, js jetstream.JetStream) JetstreamInfra {
	return &jetstreamInfra{
		logger: logger,
		js:     js,
	}
}

func (n *jetstreamInfra) CreateOrUpdateNewConsumer(ctx context.Context, streamName string, jsConfig *jetstream.ConsumerConfig) (jetstream.Consumer, error) {
	stream, err := n.js.Stream(ctx, streamName)
	if err != nil {
		n.logger.Fatal().Err(err).Msg("failed to get stream")
	}
	cons, err := stream.CreateOrUpdateConsumer(ctx, *jsConfig)
	if err != nil {
		n.logger.Error().Err(err).Msg("failed to create or update consumer")
		return nil, err
	}
	return cons, nil
}

func (n *jetstreamInfra) CreateOrUpdateNewStream(ctx context.Context, jsConfig *jetstream.StreamConfig) error {
	_, err := n.js.CreateOrUpdateStream(ctx, *jsConfig)
	if err != nil {
		n.logger.Error().Err(err).Msg("error create or update the stream")
		return err
	}
	return nil
}

func (n *jetstreamInfra) Publish(ctx context.Context, subject string, payload []byte) (*jetstream.PubAck, error) {
	ack, err := n.js.Publish(ctx, subject, payload)
	if err != nil {
		return nil, err
	}
	return ack, nil
}
`

const RabbitMQInfraTemplate string = `
package mq_infra

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type (
	RabbitMQInfra interface {
		Publish(ctx context.Context, exchange, routingKey string, payload []byte) error
	}
	rabbitMQInfra struct {
		conn   *amqp.Connection
		ch     *amqp.Channel
		logger zerolog.Logger
	}
)

func NewRabbitMQInfra(logger zerolog.Logger, conn *amqp.Connection) (RabbitMQInfra, error) {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error().Err(err).Msg("failed to create channel")
		return nil, err
	}
	return &rabbitMQInfra{
		logger: logger,
		conn:   conn,
		ch:     ch,
	}, nil
}

func (r *rabbitMQInfra) Publish(ctx context.Context, exchange, routingKey string, payload []byte) error {
	err := r.ch.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		},
	)
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to publish message")
		return err
	}
	return nil
}
`

const KafkaMQInfraTemplate string = `
package mq_infra

import (
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

type (
	KafkaInfra interface {
		CreateTopic(ctx context.Context, topic string, partitions int, replicationFactor int) error
		Publish(ctx context.Context, topic string, payload []byte) error
		Consume(ctx context.Context, topic, groupID string, handler func([]byte) error) error
	}

	kafkaInfra struct {
		conn   *kafka.Conn
		logger zerolog.Logger
	}
)

func NewKafkaInfra(logger zerolog.Logger, conn *kafka.Conn) KafkaInfra {
	return &kafkaInfra{
		logger: logger,
		conn:   conn,
	}
}

// --- Create topic (idempotent) ---
func (k *kafkaInfra) CreateTopic(ctx context.Context, topic string, partitions int, replicationFactor int) error {
	controller, err := k.conn.Controller()
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to get controller")
		return err
	}

	controllerAddr := controller.Host + ":" + strconv.Itoa(controller.Port)
	controllerConn, err := kafka.DialContext(ctx, "tcp", controllerAddr)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to connect to controller")
		return err
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	})
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to create topic")
		return err
	}
	k.logger.Info().Str("topic", topic).Msg("topic created or already exists")
	return nil
}

// --- Publish (similar to JetStream Publish) ---
func (k *kafkaInfra) Publish(ctx context.Context, topic string, payload []byte) error {
	writer := &kafka.Writer{
		Addr:         k.conn.RemoteAddr(),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}
	defer writer.Close()

	err := writer.WriteMessages(ctx, kafka.Message{
		Value: payload,
		Time:  time.Now(),
	})
	if err != nil {
		k.logger.Error().Err(err).Str("topic", topic).Msg("failed to publish message")
		return err
	}
	k.logger.Debug().Str("topic", topic).Msg("message published successfully")
	return nil
}

// --- Consume (similar to JetStream consumer) ---
func (k *kafkaInfra) Consume(ctx context.Context, topic, groupID string, handler func([]byte) error) error {
	brokerAddr := k.conn.RemoteAddr().String()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddr},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			k.logger.Error().Err(err).Msg("failed to fetch message")
			return err
		}

		if err := handler(m.Value); err != nil {
			k.logger.Error().Err(err).Msg("handler failed, skipping commit")
			continue
		}

		if err := reader.CommitMessages(ctx, m); err != nil {
			k.logger.Error().Err(err).Msg("failed to commit message")
		}
	}
}
`
