package database

import (
	"database/sql"
	"fmt"
)

// GetColumnData connects to the provided table using the name and iterates over all columns.
// The type of the data will get chosen independently
func GetColumnData(db *sql.DB, tableName string) (string, []string, error) {
	getStmt := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", tableName)
	rows, err := db.Query(getStmt)
	if err != nil {
		return "", nil, err
	}

	var (
		column_name  string
		column_names string
		column_Type  string
		column_types []string
	)

	defer rows.Close()
	rows.Next()

	err = rows.Scan(&column_name, &column_Type)
	if err != nil {
		return "", nil, err
	}

	column_names = column_name

	for rows.Next() {
		err := rows.Scan(&column_name, &column_Type)
		if err != nil {
			return "", nil, err
		}

		column_names = column_names + ", " + column_name
		column_types = append(column_types, column_Type)
	}

	return column_names, column_types, nil
}

// ConnectToDatabase creates the connection to the database with the provided Data
func ConnectToDatabase(dbname string, port int, user string, password string, host string) (*sql.DB, error) {
	conStr := fmt.Sprintf("postgresql:%d//%s:%s@%s/%s?sslmode=disable", port, user, password, host, dbname)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetMaxID gets the highest ID of the table
func GetMaxID(tableName string, db *sql.DB) (int, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT MAX(id) FROM %s", tableName))
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var res int
	rows.Next()

	err = rows.Scan(&res)
	if err != nil {
		return 0, err
	}

	return res, nil
}
