.PHONY: install
install: 
	go get

.PHONY: dev
dev: 
	go run main.go

.PHONY: test
test: 
	go test ./... -v

.PHONY: travis
travis: install test
