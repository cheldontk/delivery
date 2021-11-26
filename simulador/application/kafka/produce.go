package kafka

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cheldontk/delivery/application/route"
	"github.com/cheldontk/delivery/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPosition()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("kafkaProduceTopic"), producer)
		//time.Sleep(time.Millisecond * 500)
	}

}
