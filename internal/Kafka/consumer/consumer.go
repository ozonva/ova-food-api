package consumer

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

func Subscribe(ctx context.Context, topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving partitionList ")
	}
	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		partitionCons, err := consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			log.Error().Err(err).Msg("Error retrieving partitionList ")
		}
		go func(ctx context.Context, partitionCons sarama.PartitionConsumer) {
			for {
				select {
				case message := <-partitionCons.Messages():
					log.Info().Msgf("Message %s was read from Kafka by consumer", string(message.Value))
				case <-ctx.Done():
					log.Info().Msg("Kafka consumer stopped")
					return
				}

			}
		}(ctx, partitionCons)
	}
}
