.PHONY: install
install: 
	go get

.PHONY: test
test: 
	go test ./... -v

.PHONY: travis
travis: install test
