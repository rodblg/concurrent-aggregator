package aggregator

import (
	"context"
	"log/slog"
	"time"
)

type UserAggregator struct {
	logger  slog.Logger
	timeout time.Duration
}

// Option is a functional option type that allows us to configure the Client.
type Option func(*UserAggregator)

//We will use a creation options pattern for our constructor

func NewUserAggregator(options ...Option) *UserAggregator {
	UserAggregator := &UserAggregator{
		logger:  slog.Logger{},
		timeout: 30 * time.Second, //Default timeout

	}

	for _, opt := range options {
		opt(UserAggregator)
	}

	return UserAggregator
}

func WithTimeouts(timeout time.Duration) Option {
	return func(ua *UserAggregator) {
		ua.timeout = timeout
	}

}

// func NewUserAggregator(logger slog.Logger) *UserAggregator {
// 	return &UserAggregator{
// 		logger: logger,
// 	}
// }

func (ua *UserAggregator) Aggregate(ctx context.Context, id int) {

}
