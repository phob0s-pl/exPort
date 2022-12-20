package broker

import (
	"bytes"
	"encoding/gob"

	"github.com/nsqio/go-nsq"
	"github.com/phob0s-pl/exPort/domain"
)

type Publisher struct {
	producer *nsq.Producer
}

func (p *Publisher) Stop() {
	p.producer.Stop()
}

func NewPublisher() (*Publisher, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("nsqd:4150", config)
	if err != nil {
		return nil, err
	}

	return &Publisher{producer: producer}, nil
}

func (p *Publisher) PublishPortStore(port *domain.Port) error {
	var (
		buff    bytes.Buffer
		encoder = gob.NewEncoder(&buff)
	)

	if err := encoder.Encode(port); err != nil {
		return err
	}

	return p.producer.Publish(domain.TopicPortStore, buff.Bytes())
}

func (p *Publisher) PublishPortProcess(fileName string) error {
	return p.producer.Publish(domain.TopicPortProcess, []byte(fileName))
}

func (p *Publisher) PublishPortGetRequest(key string, requestID domain.RequestID) error {
	var (
		buff    bytes.Buffer
		encoder = gob.NewEncoder(&buff)
	)

	if err := encoder.Encode(&domain.PortRequest{
		Key:       key,
		RequestID: requestID,
	}); err != nil {
		return err
	}

	return p.producer.Publish(domain.TopicPortGetRequest, buff.Bytes())
}

func (p *Publisher) PublishPortGetResponse(requestID domain.RequestID, port *domain.Port) error {
	var (
		buff    bytes.Buffer
		encoder = gob.NewEncoder(&buff)
	)

	if err := encoder.Encode(&domain.PortResponse{
		Port:      port,
		RequestID: requestID,
	}); err != nil {
		return err
	}

	return p.producer.Publish(domain.TopicPortGetRequest, buff.Bytes())
}
