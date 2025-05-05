package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

//go 连接kafka
// mysql-monitor-agent的指标主要上报到kafka平台 这里主要演示golang如何基于kafka go sdk 上传数据

type Metric struct {
	connection *kafka.Conn
}

func NewInstance() *Metric {
	return &Metric{}
}

func (m *Metric) Connect(ctx context.Context, net string, url string, topic string) error {
	connection, err := kafka.DialLeader(ctx, net, url, topic, 0)
	if err != nil {
		return err
	}
	m.connection = connection
	return nil
}

func (m *Metric) Close() error {
	if m.connection == nil {
		return nil
	}
	if err := m.connection.Close(); err != nil {
		return err
	}
	m.connection = nil
	return nil
}

func (m *Metric) WriteMessage(message []byte) error {
	_, err := m.connection.WriteMessages(kafka.Message{
		Time:  time.Now().UTC(),
		Value: message,
	})
	return err
}

func TestKafka() {
	kafkaInstance := NewInstance()
	_ = kafkaInstance.Connect(context.Background(), "tcp", "192.168.2.88:9092", "test-topic")
	_ = kafkaInstance.WriteMessage([]byte("hello kafka!03"))
	kafkaInstance.Close()
}
