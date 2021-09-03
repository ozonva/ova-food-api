package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

func Subscribe(topic string, consumer sarama.Consumer) {
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
		go func(partitionCons sarama.PartitionConsumer) {
			for message := range partitionCons.Messages() {
				log.Info().Msgf("Message %s was readed from Kafka by consumer", string(message.Value))
			}
		}(partitionCons)
	}
}