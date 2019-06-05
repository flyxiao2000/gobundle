package msgbus

import (
	"fmt"
	"testing"
	"time"
)

type MsgHandler struct {
}

func (m *MsgHandler) Handler(topic string, data []byte) {
	fmt.Println("consumer*********** topic: ", topic, "data: ", string(data))
}

func TestKafkaConsumer(t *testing.T) {
	c, err := NewConsumer(":9092")
	if err != nil {
		t.Errorf("create kafka consumer error, %v", err)
	}

	if err := c.ConsumeTopic("helloworld", "FA_DA_UploadTableReport", &MsgHandler{}); err != nil {
		t.Errorf("consume topic error, %v", err)
	}

	time.Sleep(10 * time.Second)

	if err := c.Close(); err != nil {
		t.Errorf("close kafka consumer error, %v", err)
	}

}
