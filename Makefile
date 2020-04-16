.PHONY: install
install: 
	go get

.PHONY: test
test: 
	go test ./... -v

.PHONY: try
try: 
	go run cmd/main.go

.PHONY: travis
travis: install try
