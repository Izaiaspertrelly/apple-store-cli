package subscriptions

import (
	"flag"

	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"
)

func requiredPositiveIntegerUsageError(fs *flag.FlagSet, name string) error {
	parameter := "--" + name
	if flagWasProvided(fs, name) {
		return shared.InvalidValueUsageError(parameter)
	}
	return shared.MissingRequiredUsageError(parameter)
}
