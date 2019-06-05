package msgbus

import (
	"strings"

	"github.com/Shopify/sarama"
)

type SyncProducer struct {
	producer sarama.SyncProducer
}

func NewSyncProducer(kafkaUrl string) (*SyncProducer, error) {
	sp := &SyncProducer{}

	config := sarama.NewConfig()
	//	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_0_0_0

	var err error

	if sp.producer, err = sarama.NewSyncProducer(strings.Split(kafkaUrl, ","), config); err != nil {
		return nil, err
	}

	return sp, nil
}

func (sp *SyncProducer) SendMessage(topic string, data []byte) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err := sp.producer.SendMessage(message)

	return err
}

func (sp *SyncProducer) Close() error {
	return sp.producer.Close()
}
