package infra_template

const JetstreamInfraTemplate string = `
package mq_infra

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
