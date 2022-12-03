

.PHONY: tidy run

tidy:
	go mod tidy && go mod download && go mod vendor


run:
	go run main.go run
