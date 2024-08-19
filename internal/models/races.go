package models

import (
	"time"

	db "github.com/myertek/echotest/internal/database"
)

// Define a Driver type to hold the data for an individual Driver.
type Race struct {
	ID       int
	Round    int
	Name     string
	Location string
	Datetime time.Time
	Image    string
}

func GetRaces() ([]Race, error) {
	db := db.GetDB()

	races := []Race{}
	qry := `SELECT idRace, round, name, location, datetime, image FROM race where year(datetime)=2024 order by datetime`
	rows, err := db.Query(qry)
	if err != nil {
		//logger.Info("Query failed: ", err)
		return nil, err
	}

	for rows.Next() {
		var r Race

		err := rows.Scan(&r.ID, &r.Round, &r.Name, &r.Location, &r.Datetime, &r.Image)
		if err != nil {
			return nil, err
		}

		races = append(races, r)
	}

	defer rows.Close()

	return races, nil
}
