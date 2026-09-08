package telemetry

import (
	"os"
	"strings"
)

const (
	telemetryEphemeralEnvVar      = "ASC_TELEMETRY_EPHEMERAL"
	hostedUserWorkflowsRepository = "Izaiaspertrelly/user-workflows"
)

type RuntimeContext string

const (
	RuntimeLocal                RuntimeContext = "local"
	RuntimeEphemeral            RuntimeContext = "ephemeral"
	RuntimeHostedSandbox        RuntimeContext = "hosted_sandbox"
	RuntimeHostedGitHubWorkflow RuntimeContext = "hosted_github_workflow"
	RuntimeCI                   RuntimeContext = "ci"
)

type InvocationSource string

const (
	SourceTerminal     InvocationSource = "terminal"
	SourceClaudeCode   InvocationSource = "claude_code"
	SourceCursorAgent  InvocationSource = "cursor_agent"
	SourceCodexDesktop InvocationSource = "codex_desktop"
	SourceOpenCode     InvocationSource = "opencode"
	SourcePi           InvocationSource = "pi"
	SourceHostedAgent  InvocationSource = "hosted_agent"
)

func DetectRuntimeContext() RuntimeContext {
	switch {
	case isHostedGitHubWorkflow():
		return RuntimeHostedGitHubWorkflow
	case isHostedSandbox():
		return RuntimeHostedSandbox
	case envTruthy(telemetryEphemeralEnvVar):
		return RuntimeEphemeral
	case isKnownCIEnv():
		return RuntimeCI
	default:
		return RuntimeLocal
	}
}

func DetectInvocationSource() InvocationSource {
	switch {
	case envTruthy("PI_CODING_AGENT"):
		return SourcePi
	case envTruthy("OPENCODE"):
		return SourceOpenCode
	case os.Getenv("CLAUDE_CODE_CHILD_SESSION") == "1":
		return SourceClaudeCode
	case os.Getenv("CURSOR_AGENT") != "":
		return SourceCursorAgent
	case os.Getenv("CODEX_SHELL") == "1" && os.Getenv("CODEX_THREAD_ID") != "":
		return SourceCodexDesktop
	case isHostedSandbox() || isHostedGitHubWorkflow():
		return SourceHostedAgent
	default:
		return SourceTerminal
	}
}

func isHostedSandbox() bool {
	return strings.TrimSpace(os.Getenv("ASC_HOSTED_SANDBOX_ID")) != ""
}

func isHostedGitHubWorkflow() bool {
	return envTruthy("GITHUB_ACTIONS") && strings.EqualFold(
		strings.TrimSpace(os.Getenv("GITHUB_REPOSITORY")),
		hostedUserWorkflowsRepository,
	)
}

func isKnownCIEnv() bool {
	for _, key := range []string{
		"CI",
		"GITHUB_ACTIONS",
		"GITLAB_CI",
		"CIRCLECI",
		"BUILDKITE",
		"BITRISE_IO",
		"TF_BUILD",
		"TRAVIS",
		"APPVEYOR",
	} {
		if envTruthy(key) {
			return true
		}
	}
	return os.Getenv("TEAMCITY_VERSION") != "" || os.Getenv("JENKINS_URL") != ""
}

func envTruthy(key string) bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(key))) {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}

func shouldAttachInstallID(ctx RuntimeContext) bool {
	return ctx == RuntimeLocal
}
