package kafkaproducer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
	"trood-test/env"
	eventdispatcher "trood-test/internal/event_dispatcher"
	"trood-test/kafka"

	"github.com/IBM/sarama"
)





type kafkaProducer struct {
	producer sarama.AsyncProducer
	logger   *slog.Logger
}

type UnknownQuestionEvent struct {
	ChatID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func NewMSKAccessTokenProvider(env string, cfg *env.Kafka) *MSKAccessTokenProvider {
    return &MSKAccessTokenProvider{
        cfg: cfg,
        env: env,
    }
}

type MSKAccessTokenProvider struct {
    env string
	cfg *env.Kafka
}

func (m *MSKAccessTokenProvider) Token() (*sarama.AccessToken, error) {
    const op = "kafka.konsumer.Token"

	if m.cfg.Region == "" {
        return nil, fmt.Errorf("%s: region must be specified", op)
    }

    if m.env == "dev" {
        config := &sarama.Config{}
        config.Net.SASL.Enable = true
        config.Net.SASL.User = m.cfg.User
        config.Net.SASL.Password = m.cfg.Pass
        config.Net.SASL.Handshake = true
        config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
        config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { 
            return &kafka.XDGSCRAMClient{HashGeneratorFcn: kafka.SHA512} 
        }
        
        return &sarama.AccessToken{
            Token: fmt.Sprintf("%s:%s", m.cfg.User, m.cfg.Pass),
        }, nil
    }
	
	return &sarama.AccessToken{}, nil
}


func NewKafkaProducer(env string, cfg env.Kafka, log *slog.Logger) (eventdispatcher.EventDispatcher, error) {
    config := sarama.NewConfig()
    
    if env == "local" {
        config.Net.SASL.Enable = false
        config.Net.TLS.Enable = false
    } else {
        config.Net.SASL.Enable = true
        config.Net.SASL.Mechanism = sarama.SASLTypeOAuth
        config.Net.SASL.TokenProvider = &MSKAccessTokenProvider{env: env, cfg: &cfg}
        
        tlsConfig := &tls.Config{}
        config.Net.TLS.Enable = true
        config.Net.TLS.Config = tlsConfig
    }

    config.Producer.Return.Successes = true
    config.Producer.Return.Errors = true
    config.Version = sarama.V2_8_0_0 // укажите вашу версию

    producer, err := sarama.NewAsyncProducer(cfg.Brokers, config)
    if err != nil {
        return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
    }

    return &kafkaProducer{
        producer: producer,
        logger:   log,
    }, nil
}

func (p *kafkaProducer) Dispatch(ctx context.Context, event eventdispatcher.Event) error {

	dataBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal UnknownQuestionEvent: %w", err)
	}

	message := &sarama.ProducerMessage{
		Topic: event.GetNamespace(),
		Key:   sarama.StringEncoder(fmt.Sprintf("%d", event.GetAggregateID())),
		Value: sarama.ByteEncoder(dataBytes),
	}

	select {
	case p.producer.Input() <- message:
		p.logger.Info("Message enqueued successfully", slog.Int64("user_id", event.GetAggregateID()))
	case <-ctx.Done():
		return fmt.Errorf("failed to send message: context canceled: %w", ctx.Err())
	}

	return nil
}

func (p *kafkaProducer) Close() error {
	p.producer.AsyncClose()
	return nil
}