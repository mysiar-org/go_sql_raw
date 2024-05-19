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

## TODO
* ...