package pg_filler

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KingKeleos/golangPostgresFiller/src/database"
	"github.com/KingKeleos/golangPostgresFiller/src/generator"
)

// InsertOnce insters one single random string and integer into the database.
// Requires the name of the Database to connect to as well as the Name of the table to intert into
// Next to the connection strings, provide a string slice with the names of the columns
func InsertOnce(tableName string, db *sql.DB) error {
	dbcon := database.DBConn{DB: db}
	var b strings.Builder

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

	for key := range column_map {
		_, err := b.WriteString(key + ", ")
		if err != nil {
			return err
		}
	}

	column_names := b.String()[:len(b.String())-2]

	genString, err := generator.CreateValue(column_map)
	if err != nil {
		return err
	}

	insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, genString)

	_, err = dbcon.DB.Exec(insertStmt)
	if err != nil {
		return err
	}

	return nil
}

// Fill fills an already existing table with the structure of: 'id','text','value' in postgres.
// It fills it witha random String with the lenght of 15 characters and a random integer with 4 digits.
// With amonut, you specify the amount of inserts that you want to create.
func Fill(tableName string, amount int, db *sql.DB) error {
	dbcon := database.DBConn{DB: db}
	var b strings.Builder

	res, err := dbcon.GetMaxID(tableName)
	if err != nil {
		return nil
	}

	column_map, err := dbcon.GetColumnData(tableName)
	if err != nil {
		return err
	}

	for key := range column_map {
		_, err := b.WriteString(key + ", ")
		if err != nil {
			return err
		}
	}

	column_names := b.String()[:len(b.String())-2]

	for i := 0; i < amount; i++ {
		id := res + i + 1

		genString, err := generator.CreateValue(column_map)
		if err != nil {
			return err
		}

		insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, genString)

		_, err = dbcon.DB.Exec(insertStmt)
		if err != nil {
			return err
		}
	}

	return nil
}
