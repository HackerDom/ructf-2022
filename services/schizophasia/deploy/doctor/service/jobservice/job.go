package jobservice

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/usernamedt/doctor-service/models"
	"github.com/usernamedt/doctor-service/pkg/executor"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/meta_unpacker"
	"github.com/usernamedt/doctor-service/pkg/speechkit"
	"strings"
)

type JobService struct{}

func (js *JobService) Add(id string, ctx context.Context) (string, error) {
	meta, err := meta_unpacker.Unpack(id)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	logging.Infof("DECODED: %v\n", meta)

	err = executor.FinishJob(ctx, meta.Token, meta.Question, meta.UserId, speechkit.Generate())
	if err != nil {
		return "", err
	}

	s := sha1.New()
	s.Write([]byte(meta.UserId + meta.Token))

	resBytes := s.Sum(nil)

	res := hex.EncodeToString(resBytes)

	res = strings.ToUpper(res)
	logging.Infof("HASH: %s\n", res)
	return res, nil
}

func (js *JobService) Get(id string, context context.Context) (*models.Job, error) {
	job, err := models.GetJob(id)
	if err != nil {
		return nil, err
	}
	return job, nil
}
