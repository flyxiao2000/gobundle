package msgbus

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	wg       sync.WaitGroup
	consumer sarama.Consumer
	client   sarama.Client

	sync.Mutex

	offsetManagerList     map[string]sarama.OffsetManager
	PartitionConsumerList []sarama.PartitionConsumer
}

//Attention: resouce conflict
type ConsumeHandler interface {
	Handler(topic string, data []byte)
}

func NewConsumer(kafkaUrl string) (*Consumer, error) {
	c := &Consumer{
		offsetManagerList:     make(map[string]sarama.OffsetManager),
		PartitionConsumerList: make([]sarama.PartitionConsumer, 0),
	}

	var err error
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 2 * time.Second
	config.Version = sarama.V2_0_0_0

	fmt.Printf("kafka url %v\n", strings.Split(kafkaUrl, ","))

	if c.client, err = sarama.NewClient(strings.Split(kafkaUrl, ","), config); err != nil {
		return nil, err
	}

	if c.consumer, err = sarama.NewConsumerFromClient(c.client); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Consumer) ConsumeTopic(group, topic string, handler ConsumeHandler) error {
	c.Lock()
	defer c.Unlock()

	tmpStr := fmt.Sprintf("%s:%s", group, topic)
	if _, ok := c.offsetManagerList[tmpStr]; ok {
		return fmt.Errorf("already add to cosume list")
	}

	partitionList, err := c.consumer.Partitions(topic)
	if err != nil {
		return err
	}
	offsetManager, err := sarama.NewOffsetManagerFromClient(group, c.client)
	if err != nil {
		return err
	}
	c.offsetManagerList[tmpStr] = offsetManager

	for _, partition := range partitionList {
		if pom, err := offsetManager.ManagePartition(topic, partition); err != nil {
			return err
		} else {
			offset, _ := pom.NextOffset()
			if offset == sarama.OffsetNewest {
				offset = sarama.OffsetOldest
			}
			fmt.Printf("consume group: %s, topic: %s, offset: %v\n", group, topic, offset)

			if pc, err := c.consumer.ConsumePartition(topic, partition, offset); err == nil {
				c.wg.Add(1)
				c.PartitionConsumerList = append(c.PartitionConsumerList, pc)
				go func(pc sarama.PartitionConsumer) {
					defer c.wg.Done()
					defer pom.Close()

					for msg := range pc.Messages() {
						handler.Handler(msg.Topic, msg.Value)
						pom.MarkOffset(msg.Offset+1, "")
					}
				}(pc)
			} else {
				return err
			}
		}
	}

	return err
}

func (c *Consumer) Close() error {
	for _, pc := range c.PartitionConsumerList {
		pc.Close()
	}

	fmt.Println("close kafka consumer partition, sync wait group is waiting")
	c.wg.Wait()
	fmt.Println("close kafka consumer partition, sync wait group is done")

	for _, om := range c.offsetManagerList {
		om.Close()
	}

	if err := c.consumer.Close(); err != nil {
		return err
	}

	return nil
}
