package pipeline

import (
	"context"
	"encoding/json"
	"sync"
)

var mutex = sync.Mutex{}

type StageRunner interface {
	Do(ctx context.Context, deps *map[string]Stage, err error) (interface{}, error)
}

const (
	RunOnSuccess = "run_on_success"
	RunOnFailure = "run_on_failure"
	RunAlways    = "run_always"
)

const (
	StageStatusPending = "pending"
	StageStatusRunning = "running"
	StageStatusSuccess = "success"
	StageStatusFailed  = "failed"
	StageStatusSkipped = "skipped"
)

type StageResult struct {
	Status string      `json:"status"`
	Error  error       `json:"error"`
	Value  interface{} `json:"value"`
}

type Stage interface {
	GetTag() string
	GetStatus() *StageResult
	Next() *[]Stage
	Prev() *[]Stage
	IsReady() bool
	Do(ctx context.Context, deps *map[string]Stage, err error) *StageResult
	AddNext(stage ...Stage) Stage
	AddPrev(stage ...Stage)
	GetCondition() string
	SkipStatus() bool
}

type Subscriber interface {
	OnBeforeRun(ctx context.Context, runner StageRunner, tag string, deps *map[string]Stage)
	OnAfterRun(ctx context.Context, runner StageRunner, tag string, deps *map[string]Stage, err error)
}

type SimpleStage struct {
	Runner      StageRunner `json:"-"`
	Tag         string      `json:"tag"`
	status      StageResult `json:"status"`
	next        []Stage     `json:"next"`
	prev        []Stage     `json:"prev"`
	condition   string      `json:"condition"`
	subscribers []Subscriber
}

func NewSimpleStage(runner StageRunner) *SimpleStage {
	return NewSimpleStageWithTag(runner, "")
}

func NewSimpleStageWithTag(runner StageRunner, tag string) *SimpleStage {
	return NewSimpleStageWithCondition(runner, tag, RunOnSuccess)
}

func NewSimpleStageWithCondition(runner StageRunner, tag, condition string) *SimpleStage {
	return &SimpleStage{
		Runner: runner,
		status: StageResult{
			Status: StageStatusPending,
		},
		Tag:       tag,
		condition: condition,
	}
}

func (s *SimpleStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":       s.Tag,
		"status":    s.status.Status,
		"next":      s.next,
		"condition": s.condition,
	})
}

func (s *SimpleStage) AddSubscriber(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

func (s *SimpleStage) SkipStatus() bool {
	wg := sync.WaitGroup{}
	if s.Next() != nil {
		go func(stages *[]Stage, wg *sync.WaitGroup) {
			for _, stage := range *stages {
				wg.Add(1)
				if stage != nil {
					stage.SkipStatus()
				}
			}
		}(s.Next(), &wg)
	}
	s.status = StageResult{
		Status: StageStatusPending,
	}
	wg.Wait()
	return true
}

func (s *SimpleStage) GetCondition() string {
	return s.condition
}

func (s *SimpleStage) GetTag() string {
	return s.Tag
}

func (s *SimpleStage) AddNext(stage ...Stage) Stage {
	for _, st := range stage {
		st.AddPrev(s)
	}
	s.next = append(s.next, stage...)
	return s
}

func (s *SimpleStage) AddPrev(stage ...Stage) {
	s.prev = append(s.prev, stage...)
}

func (s *SimpleStage) IsReady() bool {
	mutex.Lock()
	defer mutex.Unlock()
	if len(s.prev) == 0 {
		return true
	}
	isReady := s.status.Status == StageStatusPending
	for _, stage := range s.prev {
		if stage != nil && (stage.GetStatus().Status == StageStatusSuccess ||
			(stage.GetStatus().Status == StageStatusFailed && s.GetCondition() == RunOnFailure)) {
			isReady = true && isReady
		} else {
			isReady = false && isReady
		}
	}
	if isReady {
		var status = StageResult{
			Status: StageStatusRunning,
		}
		s.status = status
	}
	return isReady
}

func (s *SimpleStage) Next() *[]Stage {
	if len(s.next) == 0 {
		return nil
	}
	return &s.next
}

func (s *SimpleStage) Prev() *[]Stage {
	if len(s.prev) == 0 {
		return nil
	}
	return &s.prev
}

func (s *SimpleStage) GetStatus() *StageResult {
	return &s.status
}

func (s *SimpleStage) notificationBeforeDo(ctx context.Context, runner StageRunner, tag string, deps *map[string]Stage) {
	if s.subscribers == nil {
		return
	}
	for _, ss := range s.subscribers {
		ss.OnBeforeRun(ctx, runner, tag, deps)
	}
}

func (s *SimpleStage) notificationAfterDo(ctx context.Context, runner StageRunner, tag string, deps *map[string]Stage, err error) {
	if s.subscribers == nil {
		return
	}
	for _, ss := range s.subscribers {
		ss.OnAfterRun(ctx, runner, tag, deps, err)
	}
}

func (s *SimpleStage) Do(ctx context.Context, deps *map[string]Stage, e error) *StageResult {
	mutex.Lock()
	var status = StageResult{
		Status: StageStatusRunning,
	}
	s.status = status
	mutex.Unlock()
	s.notificationBeforeDo(ctx, s.Runner, s.Tag, deps)
	result, err := s.Runner.Do(ctx, deps, e)
	defer s.notificationAfterDo(ctx, s.Runner, s.Tag, deps, err)
	mutex.Lock()
	defer mutex.Unlock()
	if err != nil {
		status = StageResult{
			Status: StageStatusFailed,
			Error:  err,
			Value:  result,
		}
	} else {
		status = StageResult{
			Value:  result,
			Status: StageStatusSuccess,
		}
	}
	s.status = status
	return &status
}
