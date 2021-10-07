package sqlite

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/decoder"
	"github.com/mir-one/localfarm/src/assets/repository"
	"github.com/mir-one/localfarm/src/helper/structhelper"
)

type ReservoirEventRepositorySqlite struct {
	DB *sql.DB
}

func NewReservoirEventRepositorySqlite(db *sql.DB) repository.ReservoirEventRepository {
	return &ReservoirEventRepositorySqlite{DB: db}
}

func (f *ReservoirEventRepositorySqlite) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		for _, v := range events {
			latestVersion++

			stmt, err := f.DB.Prepare(`INSERT INTO RESERVOIR_EVENT
				(RESERVOIR_UID, VERSION, CREATED_DATE, EVENT)
				VALUES (?, ?, ?, ?)`)

			if err != nil {
				result <- err
			}

			e, err := json.Marshal(decoder.EventWrapper{
				EventName: structhelper.GetName(v),
				EventData: v,
			})

			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(uid, latestVersion, time.Now().Format(time.RFC3339), e)
			if err != nil {
				result <- err
			}
		}

		result <- nil
		close(result)
	}()

	return result
}
