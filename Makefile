get::
#	go get -modfile=go_test.mod
	go get -modfile=go.mod
#	cp go_test.sum go.sum


tests::
	rm -f tests/test.db
	go test -count=1 -modfile=go_test.mod ./... -v

coverage::
	rm -f tests/test.db
	go test -count=1 -modfile=go_test.mod -cover -coverpkg=./... -coverprofile=coverage.out ./...
