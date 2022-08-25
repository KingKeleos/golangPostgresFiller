package pg_filler

import (
	"database/sql"
	"fmt"

	"github.com/KingKeleos/golangPostgresFiller/src/database"
	"github.com/KingKeleos/golangPostgresFiller/src/generator"
)

// InsertOnce insters one single random string and integer into the database.
// Requires the name of the Database to connect to as well as the Name of the table to intert into
// Next to the connection strings, provide a string slice with the names of the columns
func InsertOnce(dbname string, tableName string) error {
	conStr := "host=localhost port=5432 user=postgres password= dbname=tes sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	dbcon.DB = db
	if err != nil {
		fmt.Printf("Unable to connect to Database %s", dbname)
		return err
	}

	res, err := dbcon.GetMaxID(tableName)
	if err != nil {
		fmt.Printf("Unable to get MAX-ID of the table %s", tableName)
		return err
	}

	id := res + 1

	column_map, err := dbcon.GetColumnData(tableName)
	if err != nil {
		fmt.Printf("Unable to get Table-Data of the table %s", tableName)
		return err
	}

	genString, err := generator.CreateValue(column_map)
	if err != nil {
		return err
	}

	insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_map, id, genString)

	_, err = dbcon.DB.Exec(insertStmt)
	if err != nil {
		return err
	}

	return nil
}

// Fill fills an already existing table with the structure of: 'id','text','value' in postgres.
// It fills it witha random String with the lenght of 15 characters and a random integer with 4 digits.
// With amonut, you specify the amount of inserts that you want to create.
func Fill(dbname string, tableName string, amount int) error {
	conStr := "host=localhost port=5432 user=postgres password= dbname=tes sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	dbcon.DB = db
	if err != nil {
		fmt.Printf("Unable to connect to Database %s", dbname)
		return err
	}

	res, err := dbcon.GetMaxID(tableName)
	if err != nil {
		return nil
	}

	column_map, err := dbcon.GetColumnData(tableName)
	if err != nil {
		return err
	}

	for i := 0; i < amount; i++ {
		id := res + i + 1

		genString, err := generator.CreateValue(column_map)
		if err != nil {
			return err
		}

		insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_map, id, genString)

		_, err = dbcon.DB.Exec(insertStmt)
		if err != nil {
			return err
		}
	}

	return nil
}