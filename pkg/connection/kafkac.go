package connection

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/riicarus/loveshop/conf"
	"github.com/segmentio/kafka-go"
)

type KafkaHanler[T interface{}] struct {
	Addr          string
	ConsumerGroup string
	Topic         string
}

func NewKakfaHandler[T interface{}](groupId string, topic string) *KafkaHanler[T] {
	return &KafkaHanler[T]{
		Addr:          conf.ServiceConf.Kafka.Addr,
		ConsumerGroup: groupId,
		Topic:         topic,
	}
}

func (h *KafkaHanler[T]) Write(key string, msg T) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP(h.Addr),
		Topic:    h.Topic,
		Balancer: &kafka.Hash{},
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	m := kafka.Message{
		Value: msgBytes,
		Key: []byte(key),
	}
	if err := w.WriteMessages(context.Background(), m); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (h *KafkaHanler[T]) Fetch() (*T, func() error, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{h.Addr},
		GroupID:        h.ConsumerGroup,
		Topic:          h.Topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})

	ctx := context.Background()

	t := new(T)
	m, err := r.FetchMessage(ctx)
	if err != nil {
		return t, nil, err
	}

	if err := json.Unmarshal([]byte(m.Value), t); err != nil {
		return t, nil, err
	}

	commit := func() error {
		defer r.Close()

		if err := r.CommitMessages(ctx, m); err != nil {
			return err
		}

		return nil
	}

	return t, commit, nil
}

func (h *KafkaHanler[T]) Read(msgChan chan T) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{h.Addr},
		GroupID:        h.ConsumerGroup,
		Topic:          h.Topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})
	defer r.Close()

	ctx := context.Background()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("kafka read err: ", err)
		}

		fmt.Printf("partition: %d, offset: %d, key:%s, msg: %s", m.Partition, m.Offset, m.Key, m.Value)

		t := new(T)
		if err := json.Unmarshal([]byte(m.Value), t); err != nil {
			fmt.Println("unmarshal kafka msg err: ", err)
		}

		msgChan <- *t
	}
}
