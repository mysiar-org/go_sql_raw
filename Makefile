get::
#	go get -modfile=go_test.mod
	go get -modfile=go.mod
#	cp go_test.sum go.sum


tests::
	rm -rf test.db
	sqlite3 -batch test.db ""
	#go test -v ./...
	go test -modfile=go_test.mod ./... -v
