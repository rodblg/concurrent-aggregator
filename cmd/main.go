package main

import (
	"concurrent-aggregator/pkg/aggregator"
	"context"
	"time"
)

func main() {

	userAgg := aggregator.NewUserAggregator(
		aggregator.WithTimeouts(3 * time.Second),
	)

	ctx := context.Background()

	userAgg.Aggregate(ctx, 0)

}
