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
						}
					}
					if len(newNext) > 0 {
						runStages(wg, ctx, deps, errs, errMut, *next...)
					}
				}
			}
		}(stage)
	}
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

func Pipeline(ctx context.Context, stages ...Stage) error {
	deps := make(map[string]Stage)
	scanPipeline(&deps, searchRootStage(stages...)...)
	wg := sync.WaitGroup{}
	var err error
	runStages(&wg, ctx, &deps, &err, &sync.Mutex{}, stages...)
	wg.Wait()

	return err
}
