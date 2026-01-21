package profile

import (
	"context"
	"math/rand/v2"
)

//we need a profile function that returns a profile name
//we will mock this by randomly select a name from a list

var listName = []string{
	"Emma", "Liam", "Olivia", "Noah", "Ava",
	"Ethan", "Sophia", "Mason", "Isabella", "William",
	"Mia", "James", "Charlotte", "Benjamin", "Amelia",
	"Lucas", "Harper", "Henry", "Evelyn", "Alexander",
	"Abigail", "Michael", "Emily", "Daniel", "Elizabeth",
	"Matthew", "Sofia", "Jackson", "Avery", "Sebastian",
}

func GetProfile(
	ctx context.Context,
) <-chan Profile {

	profile := make(chan Profile)

	go func() {
		defer close(profile)
		for {
			select {
			case <-ctx.Done():
				return
			case profile <- Profile{Name: listName[rand.IntN(len(listName))]}:
			}
		}
	}()
	return profile
}
