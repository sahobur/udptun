# udptun - simple udp ip tunneling to own server

## Usage
```
To compile for linux:
# make docker-build

it will create binary in ./bin/udptun
```

## To run server:
```
./bin/udptun serve -p 9999 -l 0.0.0.0 -t 10.30.1.1/24
```

## To run client:
```
./bin/udptun client -r {remote server ip} -p 9999 -s 8888 -t 10.30.1.2/24
```
