package producer

import (
	rabbitmq "event_management/pkg/constant/rabbitMq"
	"event_management/pkg/workers/consumer"
	handlerconsumer "event_management/pkg/workers/handler_consumer"
	"event_management/pkg/workers/payload"
)

func StartWorker() {
	go consumer.StartImageConsumer(rabbitmq.SendEventImageQueueName, &handlerconsumer.EventImageHandler{}, func() any { return &payload.ImageUploadPayload{} })
	select {}
}
