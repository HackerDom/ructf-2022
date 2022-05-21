package jobservice

import (
	"context"
	"github.com/usernamedt/doctor-service/models"
	"github.com/usernamedt/doctor-service/pkg/executor"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/meta_unpacker"
	"github.com/usernamedt/doctor-service/pkg/speechkit"
)

type JobService struct{}

func (js *JobService) Add(id string, ctx context.Context) (*meta_unpacker.Meta, error) {
	meta, err := meta_unpacker.Unpack(id)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	logging.Infof("DECODED: %v\n", meta)

	err = executor.FinishJob(ctx, meta.Token, meta.Question, meta.UserId, speechkit.Generate())
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (js *JobService) Get(id string, context context.Context) (*models.Job, error) {
	job, err := models.GetJob(id)
	if err != nil {
		return nil, err
	}
	return job, nil
}
