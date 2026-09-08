package xcodecloud

import (
	"github.com/Izaiaspertrelly/apple-store-cli/internal/asc"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"
)

func printXcodeCloudDoctorResult(result *asc.XcodeCloudDoctorResult, output string, pretty bool) error {
	return shared.PrintOutput(result, output, pretty)
}
