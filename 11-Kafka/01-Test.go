package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	reader *kafka.Reader
	topic  = "user_click"
)

// 生产消息
func writeKafka(ctx context.Context) {
	write := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer write.Close()

	for i := 0; i < 3; i++ {
		if err := write.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("大")},
			kafka.Message{Key: []byte("2"), Value: []byte("大")},
			kafka.Message{Key: []byte("3"), Value: []byte("大")},
			kafka.Message{Key: []byte("4"), Value: []byte("大")},
			kafka.Message{Key: []byte("1"), Value: []byte("大")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Println(err)
			}
		} else {
			break
		}
	}

}

func readKafka(ctx context.Context) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})
	//defer reader.Close()
	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		}

	}
}

// 需要监听信息2和15、当收到信号时关闭reader
func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println(sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	writeKafka(ctx)

	go listenSignal()
	readKafka(ctx)
}
