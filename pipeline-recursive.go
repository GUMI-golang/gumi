package gumi

import (
	"sync"
)

type RecursiveJob func(pipe *Pipe) error

func WorkingPipe(pipe *Pipe, works ...RecursiveJob) error {
	return recurvieWorkingPipe(pipe, works)
}
func recurvieWorkingPipe(pipe *Pipe, works []RecursiveJob) error {
	for _, work := range works {

		if err := work(pipe); err != nil{
			return err
		}
	}

	childcount := len(pipe.Childrun)
	if childcount > 1 {
		wg := wgpool.Get().(*sync.WaitGroup)
		defer wgpool.Put(wg)
		// goroutine
		wg.Add(childcount)
		var err error
		for _, child := range pipe.Childrun {
			go func(c *Pipe) {
				if err == nil {
					err = WorkingPipe(c, works...)
				}
				wg.Done()
			}(child)
		}
		if err != nil {
			return err
		}
		wg.Wait()
	} else {
		for _, child := range pipe.Childrun {
			if err := WorkingPipe(child, works...); err != nil {
				return err
			}
		}
	}
	return nil
}
func WorkingPipeSynchronized(pipe *Pipe, works ...RecursiveJob) error {
	return recurvieWorkingPipeSynchronized(pipe, works)
}
func recurvieWorkingPipeSynchronized(pipe *Pipe, works []RecursiveJob) error {
	for _, work := range works {
		if err := work(pipe); err != nil {
			return err
		}
	}
	for _, child := range pipe.Childrun {
		if err := WorkingPipe(child, works...); err != nil {
			return err
		}
	}
	return nil
}