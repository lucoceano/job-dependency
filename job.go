package jobs

import (
	"fmt"
	"strings"
)

var CantDependOnThemselves = fmt.Errorf("jobs can’t depend on themselves")

func OrderJobs(jobs string, dependencies string) (string, error) {
	if strings.Trim(dependencies, " ") == "" {
		return jobs, nil
	}

	var jobsArray []string = strings.Split(jobs, "")
	var dependenciesArray []string = strings.Split(dependencies, "")
	var orderedJobs string = ""

	for index, job := range jobsArray {
		if dependenciesArray[index] == job {
			return "", CantDependOnThemselves
		}

		if strings.Index(orderedJobs, job) >= 0 {
			continue
		}

		if strings.Index(dependencies, job) == -1 {
			orderedJobs = appendDependencies(index, jobs, job+orderedJobs, dependenciesArray)
		}
	}

	return orderedJobs, nil
}

func appendDependencies(index int, jobs string, orderedJobs string, dependenciesArray []string) string {
	dependency := dependenciesArray[index]
	if dependency != " " {
		orderedJobs = dependency + orderedJobs
		orderedJobs = appendDependencies(strings.Index(jobs, dependency), jobs, orderedJobs, dependenciesArray)
	}

	return orderedJobs
}
