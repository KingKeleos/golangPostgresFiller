package database_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/KingKeleos/golangPostgresFiller/src/database"
	"github.com/stretchr/testify/assert"
)

func TestGetColumnData(t *testing.T) {
	conStr := "host=localhost port=5432 user=postgres password= dbname=test sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)

	dbcon.DB = db

	test_map, err := dbcon.GetColumnData("testfill")
	assert.NoError(t, err)

	assert.NotEqual(t, len(test_map), 0)
	fmt.Println(test_map)
}

func TestGetColumnDataFalse(t *testing.T) {
	conStr := "host=localhost port=5432 user=postgres password= dbname=tes sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)

	dbcon.DB = db

	test_map, err := dbcon.GetColumnData("test")
	assert.Error(t, err)

	assert.Equal(t, len(test_map), 0)
	fmt.Println(test_map)
}

func TestGetMaxID(t *testing.T) {
	conStr := "host=localhost port=5432 user=postgres password= dbname=test sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	assert.NoError(t, err)
	dbcon.DB = db

	max, err := dbcon.GetMaxID("testfill")
	assert.NoError(t, err)

	fmt.Println(max)

	assert.True(t, max >= 0)
}

func TestGetMaxIDFalse(t *testing.T) {
	conStr := "host=localhost port=5432 user=postgres password= dbname=tes sslmode=disable"
	dbcon := database.DBConn{}

	db, err := sql.Open("postgres", conStr)
	assert.NoError(t, err)
	dbcon.DB = db

	max, err := dbcon.GetMaxID("test")
	assert.Error(t, err)

	assert.False(t, max < 0)
}
