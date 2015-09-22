package util

import "google.golang.org/appengine"

var FileNames = map[string]string{
	"py":    "app.py",
	"c":     "app.c",
	"cpp":   "app.cpp",
	"java":  "Application.java",
	"robot": "robot.json",
}

const (
	TemplateBucket = "coduno-templates"
	TestsBucket    = "coduno-tests"
	// TODO(victorbalan): Add param in the test struct to not hardcode
	// the result file name.
	JUnitResultsPath = "/run/build/test-results/TEST-Tests.xml"
)

func SubmissionBucket() string {
	if appengine.IsDevAppServer() {
		return "coduno-dev"
	}
	return "coduno"
}
