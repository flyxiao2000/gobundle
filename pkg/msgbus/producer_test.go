package msgbus

import (
	"testing"
)

func TestKafkaSyncProducer(t *testing.T) {
	sp, err := NewSyncProducer("192.168.38.129:9092")
	if err != nil {
		t.Errorf("create sync producer error, %v", err)
	}

	if err = sp.SendMessage("JA_FS_ImportTableReport", []byte("Hello World!")); err != nil {
		t.Errorf("send message error, topic: %s, %v", "JA_FS_ImportTableReport", err)
	}

	if err := sp.Close(); err != nil {
		t.Errorf("close sync producer error %v", err)
	}
}
