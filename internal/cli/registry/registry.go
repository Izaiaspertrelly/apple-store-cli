package registry

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"

	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/accessibility"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/account"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/actors"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/ads"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/agerating"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/agreements"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/alternativedistribution"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/analytics"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/androidiosmapping"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/app_events"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/appclips"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/apps"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/auth"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/backgroundassets"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/buildbundles"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/buildlocalizations"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/builds"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/bundleids"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/capabilities"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/categories"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/certificates"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/completion"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/devices"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/diffcmd"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/distribute"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/docs"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/encryption"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/eula"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/finance"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/gamecenter"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/iap"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/initcmd"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/insights"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/install"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/localizations"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/marketplace"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/merchantids"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/metadata"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/migrate"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/nominations"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/notarization"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/notify"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/optimize"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/passtypeids"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/performance"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/preorders"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/pricing"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/productpages"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/profiles"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/publish"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/release"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/releasenotes"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/reviews"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/routingcoverage"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/sandbox"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/schema"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/screenshots"
	searchcmd "github.com/Izaiaspertrelly/apple-store-cli/internal/cli/search"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/shared"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/signing"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/snitch"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/status"
	storekitcmd "github.com/Izaiaspertrelly/apple-store-cli/internal/cli/storekit"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/submit"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/subscriptions"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/systemstatus"
	telemetrycmd "github.com/Izaiaspertrelly/apple-store-cli/internal/cli/telemetry"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/testflight"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/users"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/validate"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/versions"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/videopreviews"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/web"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/webhooks"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/workflow"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/xcode"
	"github.com/Izaiaspertrelly/apple-store-cli/internal/cli/xcodecloud"
)

// VersionCommand returns a version subcommand.
func VersionCommand(version string) *ffcli.Command {
	return &ffcli.Command{
		Name:       "version",
		ShortUsage: "asc version",
		ShortHelp:  "Print version information and exit.",
		UsageFunc:  shared.DefaultUsageFunc,
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println(version)
			return nil
		},
	}
}

type factory struct {
	name       string
	shortHelp  string
	newCommand func() *ffcli.Command
}

// Catalog constructs root commands on demand while preserving display order.
type Catalog struct {
	factories   []factory
	rootFlagSet *flag.FlagSet
}

// NewCatalog returns the current root command catalog.
func NewCatalog(version string) *Catalog {
	catalog := &Catalog{}
	catalog.factories = []factory{
		commandFactory("auth", "Manage authentication for the App Store Connect API.", auth.AuthCommand),
		commandFactory("doctor", "Diagnose authentication configuration issues.", auth.AuthDoctorCommand),
		commandFactory("web", "Apple web-session workflows, including finance report downloads.", web.WebCommand),
		commandFactory("account", "Inspect account-level health and access signals.", account.AccountCommand),
		commandFactory("install-skills", "Install the asc skill pack globally for App Store Connect workflows.", install.InstallSkillsCommand),
		commandFactory("init", "Initialize asc helper docs in the current repo.", initcmd.InitCommand),
		commandFactory("docs", "Access embedded documentation guides and reference helpers.", docs.DocsCommand),
		commandFactory("diff", "Generate deterministic non-mutating diff plans.", diffcmd.DiffCommand),
		commandFactory("system-status", "Check Apple Developer service health.", systemstatus.Command),
		commandFactory("status", "Show a release pipeline dashboard for an app.", status.StatusCommand),
		commandFactory("insights", "Generate weekly and daily insights from App Store data sources.", insights.InsightsCommand),
		commandFactory("release-notes", "Generate and manage App Store release notes.", releasenotes.ReleaseNotesCommand),
		commandFactory("reviews", "List and manage App Store customer reviews.", reviews.ReviewsCommand),
		commandFactory("review", "Manage App Store review details, attachments, and submissions.", reviews.ReviewCommand),
		commandFactory("analytics", "Request and download analytics and sales reports.", analytics.AnalyticsCommand),
		commandFactory("ads", "Manage Apple Ads API resources.", ads.AdsCommand),
		commandFactory("optimize", "Build cross-API optimization plans.", optimize.OptimizeCommand),
		commandFactory("performance", "Access performance metrics and diagnostic logs.", performance.PerformanceCommand),
		commandFactory("finance", "Download payments and financial reports.", finance.FinanceCommand),
		commandFactory("apps", "List and manage apps in App Store Connect.", apps.AppsCommand),
		commandFactory("app-clips", "Manage App Clip experiences and invocations.", appclips.AppClipsCommand),
		commandFactory("android-ios-mapping", "Manage Android-to-iOS app mapping details.", androidiosmapping.AndroidIosMappingCommand),
		commandFactory("app-setup", "Post-create app setup automation.", apps.AppSetupCommand),
		commandFactory("app-tags", "Inspect Apple-generated App Store discoverability tags.", apps.AppTagsCommand),
		commandFactory("marketplace", "Manage marketplace resources.", marketplace.MarketplaceCommand),
		commandFactory("alternative-distribution", "Manage alternative distribution resources.", alternativedistribution.Command),
		commandFactory("webhooks", "Manage webhooks in App Store Connect.", webhooks.WebhooksCommand),
		commandFactory("nominations", "Manage featuring nominations.", nominations.NominationsCommand),
		commandFactory("bundle-ids", "Manage bundle IDs and capabilities.", bundleids.BundleIDsCommand),
		commandFactory("merchant-ids", "Manage merchant IDs and certificates.", merchantids.MerchantIDsCommand),
		commandFactory("certificates", "Manage signing certificates.", certificates.CertificatesCommand),
		commandFactory("pass-type-ids", "Manage pass type IDs.", passtypeids.PassTypeIDsCommand),
		commandFactory("profiles", "Manage provisioning profiles.", profiles.ProfilesCommand),
		commandFactory("users", "Manage users and invitations in App Store Connect.", users.UsersCommand),
		commandFactory("actors", "Lookup actors (users, API keys) by ID.", actors.ActorsCommand),
		commandFactory("devices", "Manage devices in App Store Connect.", devices.DevicesCommand),
		commandFactory("testflight", "Manage TestFlight workflows.", testflight.TestFlightCommand),
		commandFactory("builds", "Manage builds in App Store Connect.", builds.BuildsCommand),
		commandFactory("build-bundles", "Manage build bundles and App Clip data.", buildbundles.BuildBundlesCommand),
		commandFactory("publish", "High-level publish workflows for TestFlight and App Store.", publish.PublishCommand),
		commandFactory("release", "Run high-level App Store release workflows.", release.ReleaseCommand),
		commandFactory("workflow", "Run multi-step automation workflows.", workflow.WorkflowCommand),
		commandFactory("xcode", "Local Xcode build/archive/export and signing-settings helpers.", xcode.XcodeCommand),
		commandFactory("distribute", "Plan, execute, inspect, and publish iOS distribution artifacts.", distribute.DistributeCommand),
		commandFactory("versions", "Manage App Store versions.", versions.VersionsCommand),
		commandFactory("product-pages", "Manage custom product pages and product page experiments.", productpages.ProductPagesCommand),
		commandFactory("routing-coverage", "Manage routing app coverage files.", routingcoverage.RoutingCoverageCommand),
		commandFactory("eula", "Manage End User License Agreements (EULA).", eula.EULACommand),
		commandFactory("agreements", "Manage agreements in App Store Connect.", agreements.AgreementsCommand),
		commandFactory("pricing", "Manage app pricing and availability.", pricing.PricingCommand),
		commandFactory("pre-orders", "Manage app pre-orders.", preorders.PreOrdersCommand),
		commandFactory("localizations", "Manage App Store localization metadata.", localizations.LocalizationsCommand),
		commandFactory("metadata", "Manage app metadata with deterministic workflows and keyword tooling.", metadata.MetadataCommand),
		commandFactory("screenshots", "Upload and manage App Store screenshots, including local capture, framing, and matrices.", screenshots.ScreenshotsCommand),
		commandFactory("video-previews", "Manage App Store app preview videos.", videopreviews.VideoPreviewsCommand),
		commandFactory("background-assets", "Manage background assets.", backgroundassets.BackgroundAssetsCommand),
		commandFactory("build-localizations", "Manage build release notes localizations.", buildlocalizations.BuildLocalizationsCommand),
		commandFactory("sandbox", "Manage sandbox testers in App Store Connect.", sandbox.SandboxCommand),
		commandFactory("signing", "Manage signing certificates and profiles.", signing.SigningCommand),
		commandFactory("notarization", "Manage macOS notarization submissions.", notarization.NotarizationCommand),
		commandFactory("iap", "Manage in-app purchases in App Store Connect.", iap.IAPCommand),
		commandFactory("storekit", "Manage StoreKit server APIs with In-App Purchase API keys.", storekitcmd.Command),
		commandFactory("app-events", "Manage App Store in-app events.", app_events.Command),
		commandFactory("subscriptions", "Manage subscription groups and subscriptions.", subscriptions.SubscriptionsCommand),
		commandFactory("submit", "Submission lifecycle tools; use `publish appstore --submit` to ship.", submit.SubmitCommand),
		commandFactory("validate", "Canonical App Store submission readiness report.", validate.ValidateCommand),
		commandFactory("xcode-cloud", "Trigger and monitor Xcode Cloud workflows.", xcodecloud.XcodeCloudCommand),
		commandFactory("categories", "Manage App Store categories.", categories.CategoriesCommand),
		commandFactory("age-rating", "Manage App Store age rating declarations.", agerating.AgeRatingCommand),
		commandFactory("accessibility", "Manage accessibility declarations.", accessibility.AccessibilityCommand),
		commandFactory("encryption", "Manage app encryption declarations and documents.", encryption.EncryptionCommand),
		commandFactory("migrate", "Migrate metadata from/to fastlane format.", migrate.MigrateCommand),
		commandFactory("notify", "Send notifications to external services.", notify.NotifyCommand),
		commandFactory("game-center", "Manage Game Center resources in App Store Connect.", gamecenter.GameCenterCommand),
		commandFactory("capabilities", "Show CLI, API, web-only, and public-API-limited capability coverage.", capabilities.Command),
		commandFactory("schema", "Inspect App Store Connect API endpoint schemas at runtime.", schema.SchemaCommand),
		commandFactory("telemetry", "Manage CLI telemetry settings.", telemetrycmd.TelemetryCommand),
		commandFactory("search", "Search asc commands and examples for agent-oriented command discovery.", func() *ffcli.Command {
			return searchcmd.SearchCommand(catalog.All)
		}),
		commandFactory("snitch", "Report CLI friction as a GitHub issue.", func() *ffcli.Command {
			return snitch.SnitchCommand(version)
		}),
		commandFactory("version", "Print version information and exit.", func() *ffcli.Command {
			return VersionCommand(version)
		}),
		commandFactory("completion", "Print shell completion scripts.", func() *ffcli.Command {
			return completion.CompletionCommand(catalog.All, func() *flag.FlagSet {
				return catalog.rootFlagSet
			})
		}),
	}
	return catalog
}

// SetCompletionRootFlagSet supplies the flags accepted by the root command.
func (c *Catalog) SetCompletionRootFlagSet(fs *flag.FlagSet) {
	c.rootFlagSet = fs
}

func commandFactory(name, shortHelp string, newCommand func() *ffcli.Command) factory {
	return factory{name: name, shortHelp: shortHelp, newCommand: newCommand}
}

// MetadataCommands returns lightweight root entries without invoking factories.
func (c *Catalog) MetadataCommands() []*ffcli.Command {
	commands := make([]*ffcli.Command, 0, len(c.factories))
	for _, factory := range c.factories {
		commands = append(commands, &ffcli.Command{
			Name:      factory.name,
			ShortHelp: factory.shortHelp,
			UsageFunc: shared.DefaultUsageFunc,
		})
	}
	return commands
}

// CommandsFor returns root metadata with only the requested command materialized.
func (c *Catalog) CommandsFor(name string) []*ffcli.Command {
	commands := c.MetadataCommands()
	for i, factory := range c.factories {
		if strings.EqualFold(factory.name, strings.TrimSpace(name)) {
			commands[i] = materialize(factory)
			break
		}
	}
	return commands
}

// All materializes every command for full-tree callers such as search and docs.
func (c *Catalog) All() []*ffcli.Command {
	commands := make([]*ffcli.Command, 0, len(c.factories))
	for _, factory := range c.factories {
		commands = append(commands, materialize(factory))
	}
	return commands
}

func materialize(factory factory) *ffcli.Command {
	if factory.newCommand == nil {
		return nil
	}
	return factory.newCommand()
}

// Subcommands returns all root subcommands in display order.
func Subcommands(version string) []*ffcli.Command {
	return NewCatalog(version).All()
}
