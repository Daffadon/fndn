package mq_template

const AmazonSQSConfigTemplate string = `
package mq

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/spf13/viper"
)

func NewMQConnection(ctx context.Context) *sqs.Client {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			viper.GetString("sqs.auth.access_key_id"), 			// AWS Access Key ID
			viper.GetString("sqs.auth.access_secret_key"),	// AWS Secret Access Key
			"", 																						// session token default empty unless using temporary creds
		)),
		config.WithRegion("sqs.auth.region"))
	if err != nil {
		log.Fatalf("failed to connect to amazon account: %v", err)
	}
	sqsClient := sqs.NewFromConfig(cfg)
	return sqsClient
}
`
