package workerpool

import (
	"context"
	"database/sql"
	"github.com/jackc/puddle"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/setting"
)

var Pool *puddle.Pool

func Setup() {
	constructor := func(ctx context.Context) (res interface{}, err error) {
		db, err := sql.Open("libpq", "host=registry port=5432 user=svcuser dbname=postgres password=svcpass")
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	destructor := func(value interface{}) {
		db := value.(*sql.DB)
		if db != nil {
			logging.LoggedClose(db, "failed to close the conn")
		}
	}

	Pool = puddle.NewPool(constructor, destructor, int32(setting.AppSetting.ConnPoolSize))
}
