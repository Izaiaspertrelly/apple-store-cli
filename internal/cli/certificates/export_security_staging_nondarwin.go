//go:build !windows && !darwin

package certificates

import (
	"os"

	"github.com/Izaiaspertrelly/apple-store-cli/internal/secureopen"
)

func createCertificateExportStagingFilePlatform(root *os.Root, name string, perm os.FileMode) (*os.File, error) {
	return secureopen.OpenNewFileNoFollowInRoot(root, name, perm)
}
