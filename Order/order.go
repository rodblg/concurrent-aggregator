package order

import (
	"context"
	"math/rand/v2"
)

// This struct will mock the json response of the system, to simplicity we will only get the order number
type Order struct {
	Quantity int
}

func GetOrder(
	ctx context.Context,
) <-chan Order {

	order := make(chan Order)

	go func() {
		defer close(order)

		for {
			select {
			case <-ctx.Done():
				return
			case order <- Order{Quantity: rand.IntN(15)}:
			}
		}
	}()
	return order

}
