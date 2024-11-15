package main

import (
	"fmt"
	"slices"
)

type Job struct {
	ID   int
	Name string
}

func main() {
	jobs := loadJobs()

	batchSize := 100
	for batch := range slices.Chunk(jobs, batchSize) {
		processBatch(batch)
	}
}

func loadJobs() []Job {
	jobs := make([]Job, 10000)
	for i := 0; i < 10000; i++ {
		jobs[i] = Job{
			ID:   i,
			Name: fmt.Sprintf("Job %d", i),
		}
	}
	return jobs
}

func processBatch(batch []Job) {
	fmt.Println("Batch:", batch)
	fmt.Println("--------------------")
}
