package broker

import "github.com/nsqio/go-nsq"

type Consumer struct {
	consumer    *nsq.Consumer
	consumeFunc func([]byte)
}

func NewConsumer(topic string, consumeFunc func([]byte)) (*Consumer, error) {
	var (
		messageConsumer = &Consumer{
			consumeFunc: consumeFunc,
		}
		config = nsq.NewConfig()
	)

	consumer, err := nsq.NewConsumer(topic, "devel", config)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(messageConsumer)

	messageConsumer.consumer = consumer

	return messageConsumer, nil
}

func (c *Consumer) Stop() {
	c.consumer.Stop()
}

func (c *Consumer) HandleMessage(msg *nsq.Message) error {
	defer msg.Finish()
	c.consumeFunc(msg.Body)
	return nil
}
