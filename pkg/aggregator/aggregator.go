package aggregator

import (
	order "concurrent-aggregator/Order"
	profile "concurrent-aggregator/Profile"
	"context"
	"log"
	"log/slog"
	"time"
)

type UserAggregator struct {
	logger  slog.Logger
	timeout time.Duration

	profileService profile.Profile
	orderService   order.Order
}

// Option is a functional option type that allows us to configure the Client.
type Option func(*UserAggregator)

//We will use a creation options pattern for our constructor

func NewUserAggregator(options ...Option) *UserAggregator {
	UserAggregator := &UserAggregator{
		logger:         slog.Logger{},
		timeout:        30 * time.Second, //Default timeout
		profileService: profile.Profile{},
		orderService:   order.Order{},
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

func (ua *UserAggregator) Aggregate(ctx context.Context, id int) {

	dctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//first we will try to cancel after 8

	newProfile := profile.GetProfile(dctx)
	newOrder := order.GetOrder(dctx)
	for i := 1; i < 10; i++ {

		log.Printf("Iteration %d User: %v | Orders: %v", i, (<-newProfile).Name, (<-newOrder).Quantity)

		if i == 7 {
			log.Println("WE REACH CANCEL POINT")
			cancel()
			break
		}
	}

}
