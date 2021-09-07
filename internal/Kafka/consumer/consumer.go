package consumer

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/ozonva/ova-food-api/internal/logger"
)

func Subscribe(ctx context.Context, topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logger.GlobalLogger.Error().Err(err).Msg("Error retrieving partitionList ")
	}
	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		partitionCons, err := consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			logger.GlobalLogger.Error().Err(err).Msg("Error retrieving partitionList ")
		}
		go func(ctx context.Context, partitionCons sarama.PartitionConsumer) {
			for {
				select {
				case message := <-partitionCons.Messages():
					logger.GlobalLogger.Info().Msgf("Message %s was readed from Kafka by consumer", string(message.Value))
				case <-ctx.Done():
					logger.GlobalLogger.Warn().Msgf("Kafka was stopped")
					return
				}

			}
		}(ctx, partitionCons)
	}
}
