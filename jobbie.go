package toolz

import (
	"context"
	"sync"
)

// Jobbie is a generic helper function to run a list of jobs in parallel, limiting the
// number of concurrent jobs to the number of workers
// Args:
// - ctx: the context to run the jobs in
// - jobFunc: the function to run the jobs (takes a context and a job description, returns a result and an error)
// - jobDesc: the slice of job descriptions (the job description is the argument to the jobFunc)
// - workers: the number of workers to run the jobs in (the number of workers is the number of concurrent jobs)
// Returns:
// - results: the slice of results from the jobFunc
// - errors: the slice of errors from the jobFunc
// - err: the error if any (the error is the error from the context if the context is done)
func Jobbie[I any, O any](
	ctx context.Context,
	jobFunc func(context.Context, I) (O, error),
	jobDesc []I,
	workers int,
) ([]O, []error, error) {
	jobCount := len(jobDesc)
	// create a results slice to store the results
	results := make([]O, jobCount)
	// create a errors slice to store the errors
	errors := make([]error, jobCount)

	// define a work unit struct to feed the workers
	type workUnit struct {
		idx     int
		jobDesc I
	}

	// create a feeder channel to feed the workers, and a buffer size of the number of workers
	// so that there are always jobs immediately available for the workers.
	feeder := make(chan workUnit, workers)

	// create a wait group to wait for the workers to finish
	wg := sync.WaitGroup{}
	wg.Add(workers)

	// spin up the required number of workers
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			// the worker will run until the channel is closed
			for unit := range feeder {
				result, err := jobFunc(ctx, unit.jobDesc)
				results[unit.idx] = result
				errors[unit.idx] = err
			}
		}()
	}

	// defer the cleanup of the workers
	defer wg.Wait()
	defer close(feeder)

	// feed the workers
	for i, job := range jobDesc {
		select {
		case <-ctx.Done():
			return results, errors, ctx.Err()
		case feeder <- workUnit{idx: i, jobDesc: job}:
		}
	}

	return results, errors, nil
}
