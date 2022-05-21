package models

import (
	"fmt"
	"github.com/usernamedt/doctor-service/pkg/setting"
	"time"
)

type NonExistJobError struct {
	error
}

func NewNonExistJobError(id string) NonExistJobError {
	return NonExistJobError{fmt.Errorf("non-existing job id: %s", id)}
}

type JobStatus string

const (
	Created JobStatus = "created"
	Success JobStatus = "success"
	Error   JobStatus = "error"
)

type JobExecStat struct {
	Start  time.Time
	Finish time.Time
}

type Job struct {
	ID       string      `json:"run_id"`
	MemID    string      `json:"id"`
	Status   JobStatus   `json:"status"`
	Result   string      `json:"result"`
	TimeInfo JobExecStat `json:"time_info"`
}

func GetJob(memId string) (*Job, error) {
	//job := &Job{}
	//key := e.PREFIX_JOB + memId

	// TODO IMPLEMENT GETTING THE JOB RESULTS

	return nil, NewNonExistJobError(memId)
}

func getJobLifetime() int {
	return int(setting.AppSetting.JobLifetimeMinutes.Seconds())
}
