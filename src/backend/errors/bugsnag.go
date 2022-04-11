package errors

import (
	"net/http"
	"os"

	"anime-skip.com/remote-config/src/backend/env"
	"github.com/bugsnag/bugsnag-go/v2"
)

type bugsnagReporter struct{}

func newBugsnagReporter() BugReporter {
	return &bugsnagReporter{}
}

func (r *bugsnagReporter) Init() {
	apiKey := os.Getenv("BUGSNAG_API_KEY")
	if apiKey == "" {
		panic("Bugsnag requires a BUGSNAG_API_KEY environment variable")
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:              apiKey,
		ReleaseStage:        env.STAGE,
		NotifyReleaseStages: []string{"production"},
		ProjectPackages:     []string{"main", "anime-skip.com/remote-config/**"},
		AppVersion:          env.VERSION,
		AppType:             os.Getenv("BUGSNAG_APP_TYPE"),
		Hostname:            os.Getenv("BUGSNAG_HOSTNAME"),
	})
}

func (r *bugsnagReporter) ReportError(err error) {}

func (r *bugsnagReporter) Handler(handler http.Handler) http.Handler {
	return handler
}
