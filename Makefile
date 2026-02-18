build:
	go build -o bin/udptun main.go
run-sudo:
	go build -o bin/udptun main.go
	sudo ./udptun serve -p 9999 -l 0.0.0.0 -t 10.130.0.1/24

deploy:
	scp bin/udptun root@109.207.171.14:/usr/local/bin/udptun

