package pg_filler

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
)

var Host, Port, User, Password string

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

//Fill fills an already existing table with the structure of: 'id','text','value' in postgres.
//It fills it witha random String with the lenght of 15 characters and a random integer with 4 digits.
//With amonut, you specify the amount of inserts that you want to create.
func Fill(dbname string, amount int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s  dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < amount; i++ {
		insertStmt := `insert into $1("text","value") values ($2,$3)`
		_, err = db.Exec(insertStmt, dbname, String(15), rand.Intn(9999-0))
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

func String(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
