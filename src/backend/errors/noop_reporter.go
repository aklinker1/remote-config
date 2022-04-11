package errors

import "net/http"

type noopReporter struct{}

func newNoopReporter() BugReporter {
	return &noopReporter{}
}

func (r *noopReporter) Init() {}

func (r *noopReporter) ReportError(err error) {}

func (r *noopReporter) Handler(handler http.Handler) http.Handler {
	return handler
}
