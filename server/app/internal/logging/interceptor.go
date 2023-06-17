package logging

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
)

type Interceptor struct {
	logger *logrus.Logger
}

func NewInterceptor(logger *logrus.Logger) *Interceptor {
	return &Interceptor{
		logger: logger,
	}
}

func (i *Interceptor) OnBeforeRun(ctx context.Context, runner pipeline.StageRunner, tag string, deps *map[string]pipeline.Stage) {
	i.logger.Debugf("OnBeforeRun: %s", tag)
}

func (i *Interceptor) OnAfterRun(ctx context.Context, runner pipeline.StageRunner, tag string, deps *map[string]pipeline.Stage, err error) {
	if err != nil {
		i.logger.Errorf("OnAfterRun: %s, error: %s", tag, err.Error())
	} else {
		i.logger.Debugf("OnAfterRun: %s, success", tag)
	}
}
