# go-sql-raw

The main reason that this module has been created was generic querying database without knowing table structure. 

## How to use it

```go
rows, err: = Db.Query("SELECT * FROM album")

defer rows.Close()
var data []RawStringType
for rows.Next() {
	rec := Rows2Map(rows)
	data = append(data, rec)
}
```

check also [go_sql_raw_test.go](tests/go_sql_raw_test.go)

## TODO
* type converting from DB to go in function `update`

## Credits
* module inspired by https://gist.github.com/SchumacherFM/69a167bec7dea644a20e