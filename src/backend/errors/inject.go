package errors

import (
	"fmt"
	"net/http"
	"strings"

	"anime-skip.com/remote-config/src/backend/env"
)

type BugReporter interface {
	Init()
	ReportError(err error)
	Handler(handler http.Handler) http.Handler
}

func InjectBugReporter() BugReporter {
	switch strings.TrimSpace(env.BUG_REPORTER) {
	case "bugsnag":
		return newBugsnagReporter()
	case "":
		return newNoopReporter()
	default:
		panic(fmt.Errorf("Unknown bug reporter: '%s'", env.BUG_REPORTER))
	}
}
