# go-sql-raw

The main reason that this module has been created was generic querying database without knowing table structure. 

## How to use it

```go
rows, err: = Db.Query("SELECT * FROM album")

defer rows.Close()
var data []go_sql_raw.RawSqlType
for rows.Next() {
	rec := go_sql_raw.Rows2Map(rows)
	data = append(data, rec)
}
```

check also [go_sql_raw_test.go](tests/go_sql_raw_test.go)

## TODO
* update type converting from DB to go in function `convertType`

## Credits
* module inspired by https://gist.github.com/SchumacherFM/69a167bec7dea644a20e