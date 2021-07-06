.PHONY: update-deps

update-deps:
	go get -d -u ./...
	go mod tidy