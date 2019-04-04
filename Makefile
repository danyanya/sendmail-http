build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/sendmail-http
clean:
	rm ./bin/sendmail-http
