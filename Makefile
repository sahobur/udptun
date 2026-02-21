SERVER = 109.207.171.14
CLKEY = key-abcd12347890
.DEFAULT_GOAL := build

build:
	go build -o bin/udptun main.go

build-darwin:
	GOOS=darwin go build -o bin/udptun main.go

build-linux:
	GOOS=linux go build -o bin/udptun main.go

vet:
	go vet ./...

run-server:
	go build -o bin/udptun main.go
	sudo ./bin/udptun serve 

run-client:
	go build -o bin/udptun main.go
	sudo ./bin/udptun client -r $(SERVER)  -p 9999 -s 8888 -t 10.30.1.2/24 -k $(CLKEY)

deploy:
	scp bin/udptun root@$(SERVER):/home/saho/udptun
	scp config.yaml root@$(SERVER):/etc/udptun/example_config.yaml


docker-build:
	@echo "Building udptun in docker. This may take a while... Binary will be in ./bin/"
	docker build --output type=local,dest=./bin .