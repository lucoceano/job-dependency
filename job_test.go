package jobs

import "testing"

func TestNoJobs(t *testing.T) {
	expected := ""
	orderedJobs, _ := OrderJobs("", "")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestOneJob(t *testing.T) {
	expected := "a"
	orderedJobs, _ := OrderJobs("a", "")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestManyJobs(t *testing.T) {
	expected := "abc"
	orderedJobs, _ := OrderJobs("abc", "")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestJobsCantDependOnThemselves(t *testing.T) {
	expected := ""
	orderedJobs, err := OrderJobs("abc", "  c")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}

	if err != CantDependOnThemselves {
		t.Errorf("expected '%v' got '%v'", CantDependOnThemselves, err)
	}
}
