

.PHONY: tidy

tidy:
	go mod tidy && go mod download && go mod vendor
