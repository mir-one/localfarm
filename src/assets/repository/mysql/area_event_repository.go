package mysql

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/decoder"
	"github.com/mir-one/localfarm/src/assets/repository"
	"github.com/mir-one/localfarm/src/helper/structhelper"
)

type AreaEventRepositoryMysql struct {
	DB *sql.DB
}

func NewAreaEventRepositoryMysql(db *sql.DB) repository.AreaEventRepository {
	return &AreaEventRepositoryMysql{DB: db}
}

func (f *AreaEventRepositoryMysql) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		for _, v := range events {
			stmt, err := f.DB.Prepare(`INSERT INTO AREA_EVENT (AREA_UID, VERSION, CREATED_DATE, EVENT) VALUES (?, ?, ?, ?)`)
			if err != nil {
				result <- err
			}

			latestVersion++

			e, err := json.Marshal(decoder.EventWrapper{
				EventName: structhelper.GetName(v),
				EventData: v,
			})
			if err != nil {
				result <- err
			}

			_, err = stmt.Exec(uid.Bytes(), latestVersion, time.Now(), e)
			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}
