package migrate

import (
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/assets"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/storeassets"
)

type (
	AppClipLayout = storeassets.AppClipLayout
	PreviewLayout = storeassets.PreviewLayout
)

func readAppClipLayout(dir string) (AppClipLayout, bool, error) { return storeassets.ReadAppClip(dir) }

func readPreviewLayout(dir string) ([]PreviewLayout, error) { return storeassets.ReadPreviews(dir) }

func validPosterFrame(value string) bool { return assets.ValidPreviewFrameTimeCode(value) }
