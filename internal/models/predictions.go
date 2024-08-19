package models

import (
	db "github.com/myertek/echotest/internal/database"
)

// Define a Driver type to hold the data for an individual Driver.
type Prediction struct {
	ID           int
	points_p1    int
	points_p2    int
	points_p3    int
	points_p4    int
	points_p5    int
	p1_id        int
	p2_id        int
	p3_id        int
	p4_id        int
	p5_id        int
	race_id      int
	user_id      int
	points_total int
}

func GetPredictions() ([]Prediction, error) {
	db := db.GetDB()

	predictions := []Prediction{}
	qry := `SELECT * FROM prediction`
	rows, err := db.Query(qry)
	if err != nil {
		//logger.Info("Query failed: ", err)
		return nil, err
	}

	for rows.Next() {
		var p Prediction

		err := rows.Scan(
			&p.ID, &p.points_p1, &p.points_p2,
			&p.points_p3, &p.points_p4, &p.points_p5,
			&p.p1_id, &p.p2_id, &p.p3_id, &p.p4_id, &p.p5_id,
			&p.race_id, &p.user_id, &p.points_total)

		if err != nil {
			return nil, err
		}

		predictions = append(predictions, p)
	}

	defer rows.Close()

	return predictions, nil
}
