package producer

import (
	"fmt"

	"github.com/ozonva/ova-food-api/internal/logger"

	"github.com/Shopify/sarama"
)

type producer struct {
	foodProducer sarama.SyncProducer
	topic        string
}

func (p producer) Send(msg Message) error {
	sMsg, preparedMsg := prepareMessage(p.topic, msg)

	partition, offset, err := p.foodProducer.SendMessage(preparedMsg)
	if err != nil {
		logger.GlobalLogger.Warn().Msgf("sending msg error: %v", err.Error())
	} else {
		logger.GlobalLogger.Info().Msgf("message %s was send to partition: %d, offset: %d", sMsg, partition, offset)
	}
	return nil
}

func NewProducer(brokers []string, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	prod, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logger.GlobalLogger.Error().Err(err).Msgf("kafka producer creation error")
		return nil, err
	}
	return &producer{prod, topic}, nil
}

func prepareMessage(topic string, message Message) (string, *sarama.ProducerMessage) {
	var sMsg string
	switch message.CmdType {
	case CREATE:
		sMsg = fmt.Sprint("CREATE: ", message.Info)
	case UPDATE:
		sMsg = fmt.Sprint("UPDATE: ", message.Info)
	case DELETE:
		sMsg = fmt.Sprint("DELETE: ", message.Info)
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(sMsg),
	}

	return sMsg, msg
}
