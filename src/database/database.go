package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConn struct {
	DB *sql.DB
}

// GetColumnData connects to the provided table using the name and iterates over all columns.
// The type of the data will get chosen independently
func (dbc *DBConn) GetColumnData(tableName string) (map[string]string, error) {
	m := make(map[string]string)

	getStmt := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", tableName)
	rows, err := dbc.DB.Query(getStmt)
	if err != nil {
		return m, err
	}

	var (
		column_name  string
		column_names string
		column_Type  string
	)

	defer rows.Close()
	rows.Next()

	err = rows.Scan(&column_name, &column_Type)
	if err != nil {
		return m, err
	}

	column_names = column_name

	for rows.Next() {
		err := rows.Scan(&column_name, &column_Type)
		if err != nil {
			return m, err
		}

		column_names = column_names + ", " + column_name
		m[column_name] = column_Type
	}

	return m, nil
}

// GetMaxID gets the highest ID of the table
func (dbc *DBConn) GetMaxID(tableName string) (int, error) {
	var maxID int

	rows, err := dbc.DB.Query(fmt.Sprintf("SELECT MAX(id) FROM %s", tableName))
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()

	err = rows.Scan(&maxID)
	if err != nil {
		return 0, err
	}

	return maxID, nil
}
