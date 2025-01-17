package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "topic_vip"
)

// producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // commit space
		StartOffset:    kafka.FirstOffset, // get first value
	})
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(string(jsonBody)),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)

	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) // "consumer-group-"

	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)

	defer reader.Close()

	fmt.Printf("Consumer(%d) ATC:\n", id)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Printf("Consumer(%d) error: %v", id, err)
		}
		fmt.Printf("Consumer(%d), topic %v, partition, offset %v,time: %d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	r.Run(":8999")
}
