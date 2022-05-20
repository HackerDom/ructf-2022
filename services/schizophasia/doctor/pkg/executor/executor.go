package executor

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/usernamedt/doctor-service/pkg/libpq"
	"github.com/usernamedt/doctor-service/pkg/workerpool"
	"time"
)

type JobQueueRow struct {
	Name   string
	Done   bool
	Result string
	Date   time.Time
}

func Run(ctx context.Context, payload workerpool.JobDescriptor) (workerpool.ExecResult, error) {
	db, err := sql.Open("libpq", "host=registry port=5432 user=svcuser dbname=postgres password=svcpass")
	if err != nil {
		return workerpool.ExecResult{}, err
	}

	rows, err := db.Query("SELECT * FROM jobqueue")
	if err != nil {
		return workerpool.ExecResult{}, err
	}
	defer rows.Close()

	var result []JobQueueRow

	for rows.Next() {
		var row JobQueueRow
		if err := rows.Scan(&row.Name, &row.Done, &row.Result, &row.Date); err != nil {
			return workerpool.ExecResult{}, err
		}
		result = append(result, row)
	}
	if err = rows.Err(); err != nil {
		return workerpool.ExecResult{}, err
	}

	resBytes, err := json.Marshal(result)
	if err != nil {
		return workerpool.ExecResult{}, err
	}

	return workerpool.ExecResult{Res: resBytes, TimeInfo: workerpool.JobTimeInfo{
		AllocMemStart: time.Now(), StartContainer: time.Now(), StopContainer: time.Now(), ReadMem: time.Now(), DeallocMem: time.Now()}}, nil
}
