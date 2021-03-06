package executor

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/usernamedt/doctor-service/pkg/libpq"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/workerpool"
	"time"
)

type JobQueueRow struct {
	Name   string
	Done   bool
	Result string
	Date   time.Time
}

func FinishJob(ctx context.Context, token, question, name, response string) error {
	res, err := workerpool.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	if res.Value() == nil {
		res.Destroy()
		return fmt.Errorf("acquired nil db resource from the pool")
	}

	db := res.Value().(*sql.DB)
	query := fmt.Sprintf("SELECT finish_job('%s', '%s', '%s', '%s')", token, question, name, response)
	print(query + "\n")
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		logging.Errorf("failed to ExecContext: %v", err)
		res.Destroy()
		return err
	}

	print("SUCCESS")

	res.Release()

	return err
}
