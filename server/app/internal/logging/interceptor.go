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
	i.logger.WithFields(
		logrus.Fields{
			"type": "stage",
			"tag":  tag,
		}).Debugf("OnBeforeRun: %s", tag)
}

func (i *Interceptor) OnAfterRun(ctx context.Context, runner pipeline.StageRunner, tag string, deps *map[string]pipeline.Stage, err error) {
	if err != nil {
		i.logger.WithFields(
			logrus.Fields{
				"type": "stage",
				"tag":  tag,
			}).Errorf("OnAfterRun: %s, error: %s", tag, err.Error())
	} else {
		i.logger.WithFields(
			logrus.Fields{
				"type": "stage",
				"tag":  tag,
			}).Debugf("OnAfterRun: %s, success", tag)
	}
}
