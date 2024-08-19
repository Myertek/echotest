package models

import (
	"fmt"

	db "github.com/myertek/echotest/internal/database"
)

// Define a Driver type to hold the data for an individual Driver.
type Driver struct {
	ID       int
	Code     string
	Forename string
	Surname  string
	Active   bool
	Image    string
}

// // Define a DriversModel type which wraps a sql.DB connection pool.
// type DriverModel struct {
//     DB *pgxpool.Pool
// }

// // This will insert a new driver into the database.
// func (m *DriverModel) Insert(Code string, Forename string, Surname string, Active bool, Image string) (int, error) {
//     return 0, nil
// }

// // This will return a specific snippet based on its id.
func GetDriver(id string) (Driver, error) {
	db := db.GetDB()

	qry := fmt.Sprintf(`SELECT id, code, forename, surname, active, image FROM driver WHERE id = %s`, id)
	row := db.QueryRow(qry)
	var d Driver
	err := row.Scan(&d.ID, &d.Code, &d.Forename, &d.Surname, &d.Active, &d.Image)
	if err != nil {
		return d, err
	}

	return d, nil
}

// This will return all the active drivers.
func AllActiveDrivers() ([]Driver, error) {
	db := db.GetDB()

	drivers := []Driver{}
	qry := `SELECT id, code, forename, surname, active, image FROM driver WHERE active is true order by surname`
	rows, err := db.Query(qry)
	if err != nil {
		//logger.Info("Query failed: ", err)
		return nil, err
	}

	for rows.Next() {
		var d Driver

		err := rows.Scan(&d.ID, &d.Code, &d.Forename, &d.Surname, &d.Active, &d.Image)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, d)
	}

	defer rows.Close()

	return drivers, nil
}
