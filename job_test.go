package jobs

import (
	"strings"
	"testing"
)

func TestNoJobs(t *testing.T) {
	expected := ""
	orderedJobs, _ := OrderJobs("", "")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestDependenciesSizeBuggerThanJobsSize(t *testing.T) {
	expected := "a"
	orderedJobs, _ := OrderJobs("a", "abc")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}
}

func TestDependenciesSizeSmallerThanJobsSize(t *testing.T) {
	expected := "abc"
	orderedJobs, _ := OrderJobs("abc", "a")

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

func TestJobsBDependsOnC(t *testing.T) {
	orderedJobs, _ := OrderJobs("abc", " c ")

	bIndex := strings.Index(orderedJobs, "b")
	cIndex := strings.Index(orderedJobs, "c")

	if bIndex <= cIndex {
		t.Errorf("expected c index < b index c = '%d' b = '%d'", cIndex, bIndex)
	}
}

func TestJobsBDependsOnCAndCDependsOnFAndDDependsOnAAndEDependsOnB(t *testing.T) {
	orderedJobs, _ := OrderJobs("abcdef", " cfab ")

	aIndex := strings.Index(orderedJobs, "a")
	cIndex := strings.Index(orderedJobs, "c")
	bIndex := strings.Index(orderedJobs, "b")
	eIndex := strings.Index(orderedJobs, "e")
	fIndex := strings.Index(orderedJobs, "f")
	dIndex := strings.Index(orderedJobs, "d")

	if cIndex <= fIndex {
		t.Errorf("expected f(%d) < c(%d)", fIndex, cIndex)
	}

	if bIndex <= cIndex {
		t.Errorf("expected c(%d) < b(%d)", cIndex, bIndex)
	}

	if eIndex <= bIndex {
		t.Errorf("expected b(%d) < e(%d)", bIndex, eIndex)
	}

	if dIndex <= aIndex {
		t.Errorf("expected a(%d) < d(%d)", aIndex, dIndex)
	}
}

func TestJobsWithoutCircularDependencyOnFCB(t *testing.T) {
	orderedJobs, _ := OrderJobs("abcdef", " c a b")

	aIndex := strings.Index(orderedJobs, "a")
	cIndex := strings.Index(orderedJobs, "c")
	bIndex := strings.Index(orderedJobs, "b")
	fIndex := strings.Index(orderedJobs, "f")
	dIndex := strings.Index(orderedJobs, "d")

	if fIndex <= bIndex {
		t.Errorf("expected b(%d) < f(%d)", fIndex, bIndex)
	}

	if bIndex <= cIndex {
		t.Errorf("expected c(%d) < b(%d)", cIndex, bIndex)
	}

	if dIndex <= aIndex {
		t.Errorf("expected a(%d) < d(%d)", aIndex, dIndex)
	}
}

func TestJobsWithCircularDependencyOnFCB(t *testing.T) {
	expected := ""
	orderedJobs, err := OrderJobs("abcdef", " cfa b")

	if expected != orderedJobs {
		t.Errorf("expected '%s' got '%s'", expected, orderedJobs)
	}

	if err != CantHaveCircularDependencies {
		t.Errorf("expected '%v' got '%v'", CantHaveCircularDependencies, err)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = OrderJobs("abcdefghijklmnopqrstuvxz", " cfab                   ")
	}
}
