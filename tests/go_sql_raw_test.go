package tests

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mysiar-org/go-sql-raw"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	const file string = "test.db"
	db, err := sql.Open("sqlite3", file)
	chkError(err)
	_, err = db.Exec(dropTable())
	chkError(err)
	_, err = db.Exec(createTable())
	chkError(err)
	_, err = db.Exec(insertData())
	chkError(err)

	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM album ORDER BY id")

	var data []go_sql_raw.RawSqlType
	for rows.Next() {
		rec := go_sql_raw.Rows2Map(rows)
		data = append(data, rec)
	}

	var expectedIds = []int64{1, 2, 3, 4, 5}
	var expectedPrices = []float64{56.99, 63.99, 17.99, 34.98, 80.99}
	var expectedArtists = []string{"John Coltrane", "John Coltrane", "Gerry Mulligan", "Sarah Vaughan", "Schizma"}
	var expectedTitle = []string{"Blue Train", "Giant Steps", "Jeru", "Sarah Vaughan", "Upadek"}

	for idx, entry := range data {
		id, _ := strconv.ParseInt(fmt.Sprintf("%v", entry["id"]), 10, 64)
		assert.Equal(t, expectedIds[idx], id)
		price, _ := strconv.ParseFloat(fmt.Sprintf("%v", entry["price"]), 64)
		assert.Equal(t, expectedPrices[idx], price)
		assert.Equal(t, expectedArtists[idx], entry["artist"])
		assert.Equal(t, expectedTitle[idx], entry["title"])
	}
}

func dropTable() string {
	return "DROP TABLE IF EXISTS album"
}

func createTable() string {
	return "CREATE TABLE album (id INT AUTO_INCREMENT NOT NULL, title VARCHAR(128) NOT NULL, artist VARCHAR(255) NOT NULL, price DECIMAL(5,2) NOT NULL, PRIMARY KEY (`id`))"
}

func insertData() string {
	return `
INSERT INTO album (id, title, artist, price)
VALUES
    (1, 'Blue Train', 'John Coltrane', 56.99),
    (2, 'Giant Steps', 'John Coltrane', 63.99),
    (3, 'Jeru', 'Gerry Mulligan', 17.99),
    (4, 'Sarah Vaughan', 'Sarah Vaughan', 34.98),
    (5, 'Upadek', 'Schizma', 80.99)
`
}

func chkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
