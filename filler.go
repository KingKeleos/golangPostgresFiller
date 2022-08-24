package pg_filler

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var Host, User, Password string
var Port int

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// InsertOnce insters one single random string and integer into the database.
// Requires the name of the Database to connect to as well as the Name of the table to intert into
// Next to the connection strings, provide a string slice with the names of the columns
func InsertOnce(dbname string, tableName string) {
	db := connectToDatabase(dbname)
	res := getMaxID(tableName, db)
	column_names, column_types := getColumnData(db, tableName)
	id := res + 1
	insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, createValue(column_types))

	_, err := db.Exec(insertStmt)
	if err != nil {
		fmt.Println(err)
	}
}

func createValue(column_type []string) string {
	var valueString string
	for _, a := range column_type {
		switch a {
		case "integer":
			valueString = valueString + strconv.Itoa(rand.Intn(9999-0)) + ","
		case "text":
			valueString = valueString + "'" + stringGenerator(15) + "'" + ","
		default:
		}
	}

	valueString = valueString[:len(valueString)-1]

	return valueString
}

func getColumnData(db *sql.DB, tableName string) (string, []string) {
	getStmt := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", tableName)
	rows, err := db.Query(getStmt)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	column_names = column_name

	for rows.Next() {
		err := rows.Scan(&column_name, &column_Type)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n", column_name, column_Type)

		column_names = column_names + ", " + column_name
		column_types = append(column_types, column_Type)
	}

	return column_names, column_types
}

func connectToDatabase(dbname string) *sql.DB {
	conStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", User, Password, Host, dbname)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func getMaxID(tableName string, db *sql.DB) int {
	rows, err := db.Query(fmt.Sprintf("SELECT MAX(id) FROM %s", tableName))
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var res int
	rows.Next()

	err = rows.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}

	return res
}

// Fill fills an already existing table with the structure of: 'id','text','value' in postgres.
// It fills it witha random String with the lenght of 15 characters and a random integer with 4 digits.
// With amonut, you specify the amount of inserts that you want to create.
func Fill(dbname string, tableName string, amount int) {
	db := connectToDatabase(dbname)
	res := getMaxID(tableName, db)
	column_names, column_types := getColumnData(db, tableName)

	for i := 0; i < amount; i++ {
		id := res + i + 1
		insertStmt := fmt.Sprintf("insert into %s (%s) values (%d, %s)", tableName, column_names, id, createValue(column_types))

		_, err := db.Exec(insertStmt)
		if err != nil {
			fmt.Println(err)
		}
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}

func stringGenerator(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
