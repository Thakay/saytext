build:
	GOOS=linux go build -ldflags '-extldflags "-static"' -o app
	docker build -t saytext .
	rm -f app