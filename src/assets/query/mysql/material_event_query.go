package mysql

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/decoder"
	"github.com/mir-one/localfarm/src/assets/query"
	"github.com/mir-one/localfarm/src/assets/storage"
)

type MaterialEventQueryMysql struct {
	DB *sql.DB
}

func NewMaterialEventQueryMysql(db *sql.DB) query.MaterialEventQuery {
	return &MaterialEventQueryMysql{DB: db}
}

func (f *MaterialEventQueryMysql) FindAllByID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		events := []storage.MaterialEvent{}

		rows, err := f.DB.Query("SELECT * FROM MATERIAL_EVENT WHERE MATERIAL_UID = ? ORDER BY VERSION ASC", uid.Bytes())
		if err != nil {
			result <- query.QueryResult{Error: err}
		}

		rowsData := struct {
			ID          int
			MaterialUID []byte
			Version     int
			CreatedDate time.Time
			Event       []byte
		}{}

		for rows.Next() {
			rows.Scan(&rowsData.ID, &rowsData.MaterialUID, &rowsData.Version, &rowsData.CreatedDate, &rowsData.Event)

			wrapper := decoder.MaterialEventWrapper{}
			err := json.Unmarshal(rowsData.Event, &wrapper)
			if err != nil {
				result <- query.QueryResult{Error: err}
			}

			materialUID, err := uuid.FromBytes(rowsData.MaterialUID)
			if err != nil {
				result <- query.QueryResult{Error: err}
			}

			events = append(events, storage.MaterialEvent{
				MaterialUID: materialUID,
				Version:     rowsData.Version,
				CreatedDate: rowsData.CreatedDate,
				Event:       wrapper.EventData,
			})
		}

		result <- query.QueryResult{Result: events}
		close(result)
	}()

	return result
}
