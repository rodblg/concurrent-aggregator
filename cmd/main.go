package concurrentaggregator

import (
	"pkg/aggregator/aggregator"
	"time"
)

func main() {

	userAgg := aggregator.NewUserAggregator(
		aggregator.WithTimeouts(10 * time.Second),
	)

	// Get Profiles

	// Get Orders

	/*
		You need to fetch these in parallel to reduce latency.
		However, if either fails, or if the global timeout is reached,
		the entire operation must abort immediately to save resources.
	*/
}
