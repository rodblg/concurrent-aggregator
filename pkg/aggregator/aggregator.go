package aggregator

import (
	order "concurrent-aggregator/Order"
	profile "concurrent-aggregator/Profile"
	"context"
	"fmt"
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
		timeout:        5 * time.Second, //Default timeout
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

	dctx, cancel := context.WithTimeout(ctx, ua.timeout)
	defer cancel()

	//first we will try to cancel after 8

	newProfile := profile.GetProfile(dctx)
	newOrder := order.GetOrder(dctx)
	for i := 1; i < 10; i++ {

		select {
		case <-dctx.Done():
			fmt.Printf("timeout at iteration %d: %v", i, dctx.Err())
			return
		default:
		}

		log.Printf("Iteration %d User: %v | Orders: %v", i, (<-newProfile).Name, (<-newOrder).Quantity)

		if i == 7 {
			log.Println("WE REACH CANCEL POINT")
			cancel()
			return
		}
	}

}
