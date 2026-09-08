package pricing

import (
	"github.com/Izaiaspertrelly/apple-store-cli/internal/asc"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"
)

// SetAvailabilityClientFactory replaces the ASC client factory for availability tests.
// It returns a restore function to reset the previous handler.
func SetAvailabilityClientFactory(fn func() (*asc.Client, error)) func() {
	previous := pricingAvailabilityClientFactory
	if fn == nil {
		pricingAvailabilityClientFactory = shared.GetASCClient
	} else {
		pricingAvailabilityClientFactory = fn
	}
	return func() {
		pricingAvailabilityClientFactory = previous
	}
}
