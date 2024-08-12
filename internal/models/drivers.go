package models

import (
"./internal/database/db"


)

// Define a Driver type to hold the data for an individual Driver.
type Driver struct {
    ID      int
    Code   string
    Forename string
	Surname string
	Active bool
    Image string
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
// func (m *DriverModel) Get(id int) (Driver, error) {
//     return Driver{}, nil
// }

// This will return all the active drivers.
func AllActiveDrivers() ([]Driver, error) {
	db := GetDB()
	drivers := []Driver{}
	qry := `select id, code, forename, surname, active, image from polls_driver where active is true`
	rows, err := db.Query(qry)
	if err != nil {
		//logger.Info("Query failed: ", err)
		return nil, err
	}
	
	

	for rows.Next(){
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











