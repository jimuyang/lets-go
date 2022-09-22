package engine

import (
	"log"

	"github.com/jimuyang/lets-go/spider/fetcher"
)

// Job 爬虫要执行的job
type Job struct {
	URL    string
	Parser func([]byte) ([]interface{}, []Job)
	// Fetcher func(string) ([]byte, error)
}

// NilParser 空parser
func NilParser([]byte) ([]interface{}, []Job) {
	return nil, nil
}

// Run run
func Run(jobs ...Job) {
	var runJobs []Job
	runJobs = append(runJobs, jobs...)

	for len(runJobs) > 0 {
		job := runJobs[0]
		runJobs = runJobs[1:]

		body, err := fetcher.Fetch(job.URL)
		// fmt.Println(string(body))
		if err != nil {
			log.Printf("fetch error! url:%s %v", job.URL, err)
		}
		items, newJobs := job.Parser(body)
		for _, item := range items {
			log.Printf("get item: %v", item)
		}
		runJobs = append(runJobs, newJobs...)
	}
}
