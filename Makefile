get::
#	go get -modfile=go_test.mod
	go get -modfile=go.mod
#	cp go_test.sum go.sum


tests::
	rm -rf tests/test.db
	go test -count=1 -modfile=go_test.mod ./... -v
