package storekit

import (
	"strconv"

	"github.com/Izaiaspertrelly/apple-store-cli/internal/asc"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"
)

func printOutput(data any, format string, pretty bool, headers []string, rows [][]string) error {
	return shared.PrintOutputWithRenderers(
		data,
		format,
		pretty,
		func() error { asc.RenderTable(headers, rows); return nil },
		func() error { asc.RenderMarkdown(headers, rows); return nil },
	)
}

func boolString(value bool) string { return strconv.FormatBool(value) }
