package stage

import (
	"fmt"
	"github.com/yoda/app/internal/pipeline"
	"sync"
)

type stageFunc func() pipeline.Stage

var (
	mu     sync.Mutex
	stages map[int]stageFunc
)

type StageFactory struct {
}

func NewStageFactory() *StageFactory {
	return &StageFactory{}
}

func (s *StageFactory) CreateStage(jobId int) (pipeline.Stage, error) {
	sf, ok := stages[jobId]
	if !ok {
		return nil, fmt.Errorf("job with tag(%d) is not active", jobId)
	}
	return sf(), nil
}

func Register(jobId int, stageFun stageFunc) {
	mu.Lock()
	defer mu.Unlock()
	if stages == nil {
		stages = make(map[int]stageFunc)
	}
	stages[jobId] = stageFun
}
