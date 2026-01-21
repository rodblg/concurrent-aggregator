package main

import (
	"concurrent-aggregator/pkg/aggregator"
	"context"
	"time"
)

func main() {

	userAgg := aggregator.NewUserAggregator(
		aggregator.WithTimeouts(10 * time.Second),
	)

	ctx := context.TODO()

	userAgg.Aggregate(ctx, 0)

}
