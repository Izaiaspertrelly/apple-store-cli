package apps

import "github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"

func normalizeAppInfoInclude(value string) ([]string, error) {
	return shared.NormalizeSelection(value, appInfoIncludeList(), "--include")
}

func appInfoIncludeList() []string {
	return []string{
		"app",
		"ageRatingDeclaration",
		"appInfoLocalizations",
		"primaryCategory",
		"primarySubcategoryOne",
		"primarySubcategoryTwo",
		"secondaryCategory",
		"secondarySubcategoryOne",
		"secondarySubcategoryTwo",
	}
}
