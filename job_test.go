package jobs

import "testing"

func TestNoJobs(t *testing.T) {
	expected := ""
	orderedJobs := OrderJobs("")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestOneJob(t *testing.T) {
	expected := "a"
	orderedJobs := OrderJobs("a")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestManyJobs(t *testing.T) {
	expected := "abc"
	orderedJobs := OrderJobs("abc")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}
