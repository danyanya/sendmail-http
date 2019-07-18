APP_NAME=sendmail-http


all: build

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o ./bin/$(APP_NAME)
clean:
	rm ./bin/$(APP_NAME)

docker:
	docker build -t danyanya/$(APP_NAME) .