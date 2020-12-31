init:
	npm install -g serverless

build:
	GOOS=linux go build -o bin/serverless src/**.go
deploy:
	make build
	sls deploy -v
	rm bin/**