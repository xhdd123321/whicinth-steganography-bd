BIN_FILE=whicinth-steganography-bd

hello:
	echo "Hello"

build:
	go build -o ${BIN_FILE}

run:
	./${BIN_FILE}

start:
	nohup make run > output/start_`date +%Y-%m-%d`.txt 2> output/run_`date +%Y-%m-%d`.txt &

stop:
	pidof ./${BIN_FILE} | xargs kill -9

restart: stop start

build&run: build run
