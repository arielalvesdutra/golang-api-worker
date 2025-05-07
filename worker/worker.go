package main

import (
	"fmt"
	"os"
	"time"
	"worker/usecase"
	"worker/model"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/avrov2"
)


func instantiateConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("\nConsumer Kafka instanciado...")

	return c
}

func subscribeToTopics(c *kafka.Consumer) {
	topic := "criacao-conta"
	err := c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nInscrito no tópico: %s\n", topic)
}

func consumeMessages(c *kafka.Consumer) {
	run := true

	for run {

		url := "http://localhost:8081"
		msg, err := c.ReadMessage(time.Second)

		if err == nil {

			client, err := schemaregistry.NewClient(schemaregistry.NewConfigWithAuthentication(url, "nil", "nil"))

			if err != nil {
				fmt.Printf("\nFalha ao criar o client de schema registry: %s\n", err)
				os.Exit(1)
			}

			deser, err := avrov2.NewDeserializer(client, serde.ValueSerde, avrov2.NewDeserializerConfig())

			if err != nil {
				fmt.Printf("\nFalha ao criar deserializer avro: %s\n", err)
				os.Exit(1)
			}

			conta := model.ContaAvro{}
			err = deser.DeserializeInto(*msg.TopicPartition.Topic, msg.Value, &conta)

			if err != nil {
				fmt.Println("\n Falha ao tentar deserializar mensagem Kafka: ", err)
				os.Exit(1)
			}

			usecase.PersistirConta(conta)
		}

	}
}

func main() {
	c := instantiateConsumer()
	subscribeToTopics(c)
	consumeMessages(c)
	c.Close()
}
