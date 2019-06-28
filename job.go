package jobs

import (
	"fmt"
	"strings"
)

var CantDependOnThemselves = fmt.Errorf("jobs can’t depend on themselves")

func OrderJobs(jobs string, dependencies string) (string, error) {
	if dependencies == "" {
		return jobs, nil
	}

	var jobsArray []string = strings.Split(jobs, "")
	var dependenciesArray []string = strings.Split(dependencies, "")

	for index, job := range jobsArray {
		if dependenciesArray[index] == job {
			return "", CantDependOnThemselves
		}
	}

	return jobs, nil
}
