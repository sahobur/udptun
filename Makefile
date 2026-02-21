build:
	go build -o bin/udptun main.go

run-server:
	go build -o bin/udptun main.go
	sudo ./bin/udptun serve 

run-client:
	go build -o bin/udptun main.go
	sudo ./bin/udptun client -r 109.207.171.14 -p 9999 -s 8888 -t 10.30.1.2/24

deploy:
	scp bin/udptun root@109.207.171.14:/home/saho/udptun


docker-build:
	@echo "Building udptun in docker. This may take a while... Binary will be in ./bin/"
	docker build --output type=local,dest=./bin .