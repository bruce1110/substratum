package queue

import (
	"context"
	"time"
)

var (
	impl Service
)

// Return the service implementor.
func Implementor() Service {
	return impl
}

// Register service implementor.
func RegisterImplementor(s Service) {
	impl = s
}

type Service interface {
	// Publish writes a message body to the specified queue.
	Publish(queue string, content []byte, opts ...PublishOption) error
	// Subscribe consumes the messages of the specified queue and topic.
	Subscribe(queue, topic string, handler MessageHandler, opts ...SubscribeOption) error
}

type PublishOption func(*PublishOptions)

var EmptyPublishOptions = func() *PublishOptions {
	return &PublishOptions{
		Context: context.Background(),
	}
}

type PublishOptions struct {
	context.Context
	Delay time.Duration
}

func WithPublishContext(ctx context.Context) PublishOption {
	return func(opts *PublishOptions) {
		opts.Context = ctx
	}
}

func WithPublishDelay(delay time.Duration) PublishOption {
	return func(opts *PublishOptions) {
		opts.Delay = delay
	}
}

type SubscribeOption func(*SubscribeOptions)

type SubscribeOptions struct {
	context.Context
	Concurrency int
	MaxRetry    int
	Idempotent  Idempotent
}

var EmptySubscribeOptions = func() *SubscribeOptions {
	return &SubscribeOptions{
		Context:     context.Background(),
		Concurrency: 1,
		MaxRetry:    3,
		Idempotent:  IdempotentImplementor(),
	}
}

func WithConsumeContext(ctx context.Context) SubscribeOption {
	return func(opts *SubscribeOptions) {
		opts.Context = ctx
	}
}

func WithConsumeConcurrency(concurrency int) SubscribeOption {
	return func(opts *SubscribeOptions) {
		opts.Concurrency = concurrency
	}
}

func WithConsumeRetry(retry int) SubscribeOption {
	return func(opts *SubscribeOptions) {
		opts.MaxRetry = retry
	}
}

func WithConsumeIdempotent(impl Idempotent) SubscribeOption {
	return func(opts *SubscribeOptions) {
		opts.Idempotent = impl
	}
}
