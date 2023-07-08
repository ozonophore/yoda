package pipeline

import (
	"context"
	"errors"
	"sync"
)

func getErrors(stages *[]Stage) error {
	if stages == nil {
		return nil
	}
	var err error
	for _, stage := range *stages {
		if stage != nil && stage.GetStatus().Status == StageStatusFailed {
			status := stage.GetStatus()
			err = errors.Join(err, status.Error)
		}
	}
	return err
}

func runStages(wg *sync.WaitGroup, ctx context.Context, deps *map[string]Stage, errs *error, errMut *sync.Mutex, stages ...Stage) {
	for _, stage := range stages {
		if !stage.IsReady() {
			continue
		}
		wg.Add(1)
		go func(stage Stage) {
			defer wg.Done()
			result := stage.Do(ctx, deps, getErrors(stage.Prev()))
			next := stage.Next()
			if result.Status == StageStatusSuccess {
				if next != nil {
					runStages(wg, ctx, deps, errs, errMut, *next...)
				}
			} else if result.Status == StageStatusFailed {
				{
					errMut.Lock()
					defer errMut.Unlock()
					*errs = errors.Join(*errs, result.Error)
				}
				if next != nil {
					newNext := make([]Stage, 0, len(*next))
					for _, st := range *next {
						if st.GetCondition() == RunOnFailure || st.GetCondition() == RunAlways {
							newNext = append(newNext, st)
						} else {
							newNext = append(newNext, findNext(st)...)
							st.GetStatus().Status = StageStatusSkipped
						}
					}
					if len(newNext) > 0 {
						runStages(wg, ctx, deps, errs, errMut, newNext...)
					}
				}
			}
		}(stage)
	}
}

func findNext(st Stage) []Stage {
	next := st.Next()
	newNext := make([]Stage, 0, 2)
	if next != nil {
		for _, stage := range *next {
			if stage.GetCondition() == RunAlways {
				newNext = append(newNext, stage)
			} else if stage.Next() != nil {
				newNext = append(newNext, findNext(stage)...)
			}
		}
	}
	return newNext
}

func scanPipeline(stageMap *map[string]Stage, stages ...Stage) {
	sm := *stageMap
	for _, stage := range stages {
		if len(stage.GetTag()) > 0 {
			sm[stage.GetTag()] = stage
		}
		next := stage.Next()
		if next != nil {
			scanPipeline(stageMap, *next...)
		}
	}
}

func searchRootStage(stage ...Stage) []Stage {
	var rootStages []Stage
	for _, st := range stage {
		if st.Prev() == nil {
			rootStages = append(rootStages, st)
		} else {
			rootStages = append(rootStages, searchRootStage(*st.Prev()...)...)
		}
	}
	return rootStages
}

func resetStatus(stages []Stage) {
	for i := 0; i < len(stages); i++ {
		stages[i].ResetStatus()
	}
}

type Pipeline struct {
	err error
}

func (p *Pipeline) Error() error {
	return p.err
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Do(ctx context.Context, stages ...Stage) *Pipeline {
	resetStatus(stages)
	deps := make(map[string]Stage)
	scanPipeline(&deps, searchRootStage(stages...)...)
	wg := sync.WaitGroup{}
	runStages(&wg, ctx, &deps, &p.err, &sync.Mutex{}, stages...)
	wg.Wait()
	return p
}
