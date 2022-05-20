package jobservice

import (
	"context"
	"github.com/usernamedt/doctor-service/models"
	"github.com/usernamedt/doctor-service/pkg/executor"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/meta_unpacker"
	"github.com/usernamedt/doctor-service/pkg/speechkit"
	"github.com/usernamedt/doctor-service/pkg/workerpool"
)

type JobService struct{}

func (js *JobService) Add(id string, ctx context.Context) (*meta_unpacker.Meta, error) {

	meta, err := meta_unpacker.Unpack(id)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	logging.Infof("DECODED: %v\n", meta)
	speechkit.Generate()

	workerpool.Pool.AddJob(workerpool.Job{
		Descriptor: workerpool.JobDescriptor{},
		ExecFn: func(ctx context.Context, payload workerpool.JobDescriptor) (workerpool.ExecResult, error) {
			return executor.Run(ctx, payload)
		},
	})

	//job, err := models.NewJob(jobId, memId)
	//if err != nil {
	//	logging.Info(err)
	//	return nil, err
	//}

	// add job to the pool
	return meta, nil
}

func (js *JobService) Get(id string, context context.Context) (*models.Job, error) {
	job, err := models.GetJob(id)
	if err != nil {
		return nil, err
	}
	return job, nil
}
