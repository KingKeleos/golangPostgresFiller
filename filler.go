package pg_filler

import (
	"database/sql"
	"fmt"

	"github.com/KingKeleos/golangPostgresFiller/database"
	"github.com/KingKeleos/golangPostgresFiller/generator"
	_ "github.com/lib/pq"
)

var Host, User, Password string
var Port int

// InsertOnce insters one single random string and integer into the database.
// Requires the name of the Database to connect to as well as the Name of the table to intert into
// Next to the connection strings, provide a string slice with the names of the columns
func InsertOnce(dbname string, tableName string) error {
	db, err := database.ConnectToDatabase(dbname, Port, User, Password, Host)
	if err != nil {
		fmt.Printf("Unable to connect to Database %s", dbname)
		return err
	}

	res, err := database.GetMaxID(tableName, db)
	if err != nil {
		fmt.Printf("Unable to get MAX-ID of the table %s", tableName)
		return err
	}

	id := res + 1

	column_names, column_types, err := database.GetColumnData(db, tableName)
	if err != nil {
		fmt.Printf("Unable to get Table-Data of the table %s", tableName)
		return err
	}

	insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, generator.CreateValue(column_types))

	_, err = db.Exec(insertStmt)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

// Fill fills an already existing table with the structure of: 'id','text','value' in postgres.
// It fills it witha random String with the lenght of 15 characters and a random integer with 4 digits.
// With amonut, you specify the amount of inserts that you want to create.
func Fill(dbname string, tableName string, amount int) error {
	db, err := database.ConnectToDatabase(dbname, Port, User, Password, Host)
	if err != nil {
		fmt.Printf("Unable to connect to Database %s", dbname)
		return err
	}

	res, err := database.GetMaxID(tableName, db)
	if err != nil {
		return nil
	}

	column_names, column_types, err := database.GetColumnData(db, tableName)
	if err != nil {
		return err
	}

	for i := 0; i < amount; i++ {
		id := res + i + 1
		insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, generator.CreateValue(column_types))

		_, err := db.Exec(insertStmt)
		if err != nil {
			return err
		}
	}

	defer func(db *sql.DB) error {
		err := db.Close()
		if err != nil {
			return err
		}
		return nil
	}(db)

	return nil
}
