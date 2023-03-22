package main

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {

	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"0.0.0.0:9876"}),
	)
	_ = p.Start()
	defer p.Shutdown()

	// c, _ := rocketmq.NewPushConsumer(
	// 	consumer.WithNamespace("namespace"),
	// 	consumer.WithGroupName("group"),
	// 	consumer.WithNameServer([]string{"0.0.0.0:9876"}),
	// )

	// _ = c.Subscribe("test-topic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	// 	for _, msg := range msgs {
	// 		fmt.Printf("Received message: %v\n", msg)
	// 	}
	// 	return consumer.ConsumeSuccess, nil
	// })

	// _ = c.Start()
	// defer c.Shutdown()

	_, _ = p.SendSync(context.Background(),
		primitive.NewMessage("test-topic", []byte("Hello RocketMQ!")))
}
