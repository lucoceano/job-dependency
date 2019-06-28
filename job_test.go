package jobs

import "testing"

func TestNoJobs(t *testing.T) {
	expected := ""
	orderedJobs := OrderJobs("")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}
