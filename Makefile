build:
	(cd cmd/chikuchiku/ && GOOS=freebsd GOARCH=amd64 go build -tags netgo .)

.PHONY: build
