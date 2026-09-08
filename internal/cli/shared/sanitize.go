package shared

import "github.com/Izaiaspertrelly/apple-store-cli/internal/asc"

// SanitizeTerminal removes characters interpreted by terminals and log viewers.
func SanitizeTerminal(input string) string {
	return asc.SanitizeTerminalText(input)
}
