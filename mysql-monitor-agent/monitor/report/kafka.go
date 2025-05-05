package report

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

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
